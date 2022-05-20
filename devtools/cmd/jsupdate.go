package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ignite-hq/cli/ignite/pkg/cliui/clispinner"
	"github.com/ignite-hq/cli/ignite/services/chain"
	"github.com/spf13/cobra"

	"github.com/ignite-hq/cli/ignite/pkg/cosmosgen"
)

const (
	flagPath          = "path"
	flagHome          = "home"
	flagProto3rdParty = "proto-all-modules"
	flagYes           = "yes"
)

func NewJSUpdate(ctx context.Context) *cobra.Command {

	c := &cobra.Command{
		Use:   "js-update",
		Short: "devtools make sense",
		RunE:  JSUpdateHandler,
	}

	return c
}

func JSUpdateHandler(cmd *cobra.Command, args []string) error {
	s := clispinner.New().SetText("JS Client update...")
	defer s.Stop()

	absPath, err := filepath.Abs("")
	if err != nil {
		return err
	}

	c, err := chain.New(absPath, chain.EnableThirdPartyModuleCodegen())
	if err != nil {
		return err
	}

	if err := generate(cmd.Context()); err != nil {
		return err
	}

	s.Stop()
	fmt.Println("JS Client updated.")

	return nil
}

func newChainWithHomeFlags(cmd *cobra.Command, chainOption ...chain.Option) (*chain.Chain, error) {

	absPath, err := filepath.Abs("")
	if err != nil {
		return nil, err
	}

	return chain.New(absPath, chainOption...)
}

func getHome(cmd *cobra.Command) (home string) {
	home, _ = cmd.Flags().GetString(flagHome)
	return
}

func flagGetPath(cmd *cobra.Command) (path string) {
	path, _ = cmd.Flags().GetString(flagPath)
	return
}

func generate(ctx context.Context, c chain.Chain) error {
	conf, err := c.Config()
	if err != nil {
		return err
	}

	if err := cosmosgen.InstallDependencies(ctx, c.app.Path); err != nil {
		return err
	}

	fmt.Println("🛠️  Building proto...")

	options := []cosmosgen.Option{
		cosmosgen.IncludeDirs(conf.Build.Proto.ThirdPartyPaths),
	}

	enableThirdPartyModuleCodegen := true

	// generate Vuex code as well if it is enabled.
	vuexPath := conf.Client.Vuex.Path
	if vuexPath == "" {
		vuexPath = "sdk/js-client/src"
	}

	storeRootPath := filepath.Join(c.app.Path, vuexPath, "generated")
	if err := os.MkdirAll(storeRootPath, 0766); err != nil {
		return err
	}

	options = append(options,
		cosmosgen.WithVuexGeneration(
			enableThirdPartyModuleCodegen,
			cosmosgen.VuexStoreModulePath(storeRootPath),
			storeRootPath,
		),
	)

	if err := cosmosgen.Generate(ctx, c.app.Path, conf.Build.Proto.Path, options...); err != nil {
		return &CannotBuildAppError{err}
	}

	c.protoBuiltAtLeastOnce = true

	return nil
}
