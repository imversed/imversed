package chain

import (
	"context"
	"fmt"

	"github.com/ignite-hq/cli/ignite/services/chain"
	"github.com/imversed/imversed/devtools/cosmosgen"
)

type generateOptions struct {
	isGoEnabled      bool
	isVuexEnabled    bool
	isDartEnabled    bool
	isOpenAPIEnabled bool
}

func ChainGenerateTS(
	ctx context.Context,
	c chain.Chain,
	appPath string,
) error {
	conf, err := c.Config()
	if err != nil {
		return err
	}

	if err := cosmosgen.InstallDependencies(ctx, appPath); err != nil {
		return err
	}

	fmt.Println("🛠️  Building proto...")

	options := []cosmosgen.Option{
		cosmosgen.IncludeDirs(conf.Build.Proto.ThirdPartyPaths),
	}

	enableThirdPartyModuleCodegen := true

	jsClientSrcPath := "sdk/js-client/src"

	options = append(options,
		cosmosgen.WithJSUpdateGeneration(
			enableThirdPartyModuleCodegen,
			cosmosgen.TSModulePath(jsClientSrcPath),
		),
	)

	if err := cosmosgen.Generate(ctx, appPath, conf.Build.Proto.Path, options...); err != nil {
		return err
	}

	return nil
}
