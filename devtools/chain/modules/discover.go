package modules

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/ignite-hq/cli/ignite/pkg/cosmosanalysis/module"
)

func DiscoverModules(ctx context.Context, appPath, path, protoDir string) ([]module.Module, error) {
	var filteredModules []module.Module

	modules, err := module.Discover(ctx, appPath, path, protoDir)
	if err != nil {
		return nil, err
	}

	for _, m := range modules {
		pp := filepath.Join(path, protoDir)
		if !strings.HasPrefix(m.Pkg.Path, pp) {
			continue
		}
		filteredModules = append(filteredModules, m)
	}

	return filteredModules, nil
}
