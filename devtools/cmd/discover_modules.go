package cmd

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/imversed/imversed/devtools/cosmosgen"
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

	allModules, err := cosmosgen.DiscoverModules(cmd.Context(), appPath, "proto")
	if err != nil {
		return err
	}

	for path, modules := range allModules {
		for _, module := range modules {
			fmt.Printf("- path: %s\n  name: %s\n  go_module_path: %s\n  pkg:\n    name: %s\n    path: %s\n    gp_import_name: %s\n", path, module.Name, module.GoModulePath, module.Pkg.Name, module.Pkg.Path, module.Pkg.GoImportName)
		}
	}

	return nil
}
