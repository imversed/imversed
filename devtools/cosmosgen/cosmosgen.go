package cosmosgen

import (
	"context"
	"path/filepath"

	gomodmodule "golang.org/x/mod/module"

	"github.com/ignite-hq/cli/ignite/pkg/cosmosanalysis/module"
	"github.com/ignite-hq/cli/ignite/pkg/gomodulepath"
)

var TSModuleNames = map[string]string{
	"ethermint.evm.v1":                                   "evm",
	"ethermint.feemarket.v1":                             "feemarket",
	"cosmos.auth.v1beta1":                                "auth",
	"cosmos.authz.v1beta1":                               "authz",
	"cosmos.bank.v1beta1":                                "bank",
	"cosmos.base.tendermint.v1beta1":                     "tendermint",
	"cosmos.crisis.v1beta1":                              "crisis",
	"cosmos.distribution.v1beta1":                        "distribution",
	"cosmos.evidence.v1beta1":                            "evidence",
	"cosmos.feegrant.v1beta1":                            "feegrant",
	"cosmos.gov.v1beta1":                                 "gov",
	"cosmos.mint.v1beta1":                                "mint",
	"cosmos.params.v1beta1":                              "params",
	"cosmos.slashing.v1beta1":                            "slashing",
	"cosmos.staking.v1beta1":                             "staking",
	"cosmos.tx.v1beta1":                                  "tx",
	"cosmos.upgrade.v1beta1":                             "upgrade",
	"cosmos.vesting.v1beta1":                             "vesting",
	"imversed.currency":                                  "currency",
	"imversed.nft":                                       "nft",
	"imversed.pools.v1beta1":                             "pools",
	"ibc.applications.interchain_accounts.controller.v1": "ibc-accounts-controller",
	"ibc.applications.interchain_accounts.host.v1":       "ibc-accounts-host",
	"ibc.applications.transfer.v1":                       "ibc-transfer",
	"ibc.core.channel.v1":                                "ibc-channel",
	"ibc.core.client.v1":                                 "ibc-client",
	"ibc.core.connection.v1":                             "ibc-connection",
}

// generateOptions used to configure code generation.
type generateOptions struct {
	includeDirs []string
	gomodPath   string

	jsOut               func(module.Module) string
	jsIncludeThirdParty bool
	vuexStoreRootPath   string

	specOut string

	dartOut               func(module.Module) string
	dartIncludeThirdParty bool
	dartRootPath          string
}

// TODO add WithInstall.

// ModulePathFunc defines a function type that returns a path based on a Cosmos SDK module.
type ModulePathFunc func(module.Module) string

// Option configures code generation.
type Option func(*generateOptions)

// WithJSGeneration adds JS code generation. out hook is called for each module to
// retrieve the path that should be used to place generated js code inside for a given module.
// if includeThirdPartyModules set to true, code generation will be made for the 3rd party modules
// used by the app -including the SDK- as well.
func WithJSGeneration(includeThirdPartyModules bool, out ModulePathFunc) Option {
	return func(o *generateOptions) {
		o.jsOut = out
		o.jsIncludeThirdParty = includeThirdPartyModules
	}
}

// WithVuexGeneration adds Vuex code generation. storeRootPath is used to determine the root path of generated
// Vuex stores. includeThirdPartyModules and out configures the underlying JS lib generation which is
// documented in WithJSGeneration.
func WithVuexGeneration(includeThirdPartyModules bool, out ModulePathFunc, storeRootPath string) Option {
	return func(o *generateOptions) {
		o.jsOut = out
		o.jsIncludeThirdParty = includeThirdPartyModules
		o.vuexStoreRootPath = storeRootPath
	}
}

func WithJSUpdateGeneration(includeThirdPartyModules bool, out ModulePathFunc) Option {
	return func(o *generateOptions) {
		o.jsOut = out
		o.jsIncludeThirdParty = includeThirdPartyModules
	}
}

func WithDartGeneration(includeThirdPartyModules bool, out ModulePathFunc, rootPath string) Option {
	return func(o *generateOptions) {
		o.dartOut = out
		o.dartIncludeThirdParty = includeThirdPartyModules
		o.dartRootPath = rootPath
	}
}

// WithGoGeneration adds Go code generation.
func WithGoGeneration(gomodPath string) Option {
	return func(o *generateOptions) {
		o.gomodPath = gomodPath
	}
}

// WithOpenAPIGeneration adds OpenAPI spec generation.
func WithOpenAPIGeneration(out string) Option {
	return func(o *generateOptions) {
		o.specOut = out
	}
}

// IncludeDirs configures the third party proto dirs that used by app's proto.
// relative to the projectPath.
func IncludeDirs(dirs []string) Option {
	return func(o *generateOptions) {
		o.includeDirs = dirs
	}
}

// generator generates code for sdk and sdk apps.
type generator struct {
	ctx          context.Context
	appPath      string
	protoDir     string
	o            *generateOptions
	sdkImport    string
	deps         []gomodmodule.Version
	appModules   []module.Module
	thirdModules map[string][]module.Module // app dependency-modules pair.
}

// Generate generates code from protoDir of an SDK app residing at appPath with given options.
// protoDir must be relative to the projectPath.
func Generate(ctx context.Context, appPath, protoDir string, options ...Option) error {
	g := &generator{
		ctx:          ctx,
		appPath:      appPath,
		protoDir:     protoDir,
		o:            &generateOptions{},
		thirdModules: make(map[string][]module.Module),
	}

	for _, apply := range options {
		apply(g.o)
	}

	if err := g.setup(); err != nil {
		return err
	}

	if g.o.gomodPath != "" {
		if err := g.generateGo(); err != nil {
			return err
		}
	}

	// js generation requires Go types to be existent in the source code. because
	// sdk.Msg implementations defined on the generated Go types.
	// so it needs to run after Go code gen.
	if g.o.jsOut != nil {
		if err := g.generateJS(); err != nil {
			return err
		}
	}

	if g.o.dartOut != nil {
		if err := g.generateDart(); err != nil {
			return err
		}
	}

	if g.o.specOut != "" {
		if err := generateOpenAPISpec(g); err != nil {
			return err
		}
	}

	return nil

}

// VuexStoreModulePath generates Vuex store module paths for Cosmos SDK modules.
// The root path is used as prefix for the generated paths.
func VuexStoreModulePath(rootPath string) ModulePathFunc {
	return func(m module.Module) string {
		appModulePath := gomodulepath.ExtractAppPath(m.GoModulePath)
		return filepath.Join(rootPath, appModulePath, m.Pkg.Name, "module")
	}
}

func TSModulePath(rootPath string) ModulePathFunc {
	return func(m module.Module) string {
		modulePath := TSModuleNames[m.Pkg.Name]
		if modulePath == "" {
			modulePath = m.Pkg.Name
		}
		return filepath.Join(rootPath, modulePath)
	}
}

func DiscoverModules(ctx context.Context, appPath, protoDir string) (map[string][]module.Module, error) {
	g := &generator{
		ctx:          ctx,
		appPath:      appPath,
		protoDir:     protoDir,
		o:            &generateOptions{},
		thirdModules: make(map[string][]module.Module),
	}

	if err := g.setup(); err != nil {
		return nil, err
	}

	allModules := make(map[string][]module.Module)
	for _, module := range g.appModules {
		allModules[appPath] = append(allModules[appPath], module)
	}
	for path, modules := range g.thirdModules {
		allModules[path] = append(allModules[path], modules...)
	}

	return allModules, nil
}
