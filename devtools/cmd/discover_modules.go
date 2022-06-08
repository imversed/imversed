package cmd

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/ignite-hq/cli/ignite/pkg/cosmosanalysis/module"
	"github.com/ignite-hq/cli/ignite/pkg/gomodule"
	"github.com/imversed/imversed/devtools/chain/modules"
	"github.com/spf13/cobra"
)

func NewDiscoverModules(ctx context.Context) *cobra.Command {

	c := &cobra.Command{
		Use:   "discover-modules",
		Short: "discover modules",
		RunE:  DiscoverModulesHandler,
	}

	return c
}

func DiscoverModulesHandler(cmd *cobra.Command, args []string) error {
	appPath, err := filepath.Abs("")
	if err != nil {
		return err
	}

	modfile, err := gomodule.ParseAt(appPath)
	if err != nil {
		return err
	}

	deps, err := gomodule.ResolveDependencies(modfile)
	if err != nil {
		return err
	}

	appModules, err := modules.DiscoverModules(cmd.Context(), appPath, appPath, "proto")
	if err != nil {
		return err
	}

	thirdPartyModules := make(map[string][]module.Module)
	for _, dep := range deps {
		path, err := gomodule.LocatePath(cmd.Context(), appPath, dep)
		if err != nil {
			return err
		}
		thirdPartyDiscovered, err := modules.DiscoverModules(cmd.Context(), appPath, path, "")
		if err != nil {
			return err
		}
		thirdPartyModules[path] = append(thirdPartyModules[path], thirdPartyDiscovered...)
	}

	for _, module := range appModules {
		fmt.Printf("%s %s\n", module.Pkg.Name, module.Pkg.GoImportName)
	}

	for _, modules := range thirdPartyModules {
		for _, module := range modules {
			fmt.Printf("%s %s\n", module.Pkg.Name, module.Pkg.GoImportName)
		}
	}

	return nil
}
