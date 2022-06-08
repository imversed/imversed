package chain

import (
	"context"
	"fmt"

	"github.com/ignite-hq/cli/ignite/pkg/cosmosgen"
	"github.com/ignite-hq/cli/ignite/services/chain"
)

type generateOptions struct {
	isGoEnabled      bool
	isVuexEnabled    bool
	isDartEnabled    bool
	isOpenAPIEnabled bool
}

func chainGenerateTS(
	c chain.Chain,
	ctx context.Context,
) error {
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

	jsClientSrcPath := "sdk/js-client/src"

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

	return nil
}
