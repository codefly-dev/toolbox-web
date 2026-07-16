package web_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/codefly-dev/core/resources"
	web "github.com/codefly-dev/toolbox-web"
)

func TestManifestMatchesProductionCatalog(t *testing.T) {
	manifest, err := resources.LoadToolboxFromDir(context.Background(), ".")
	require.NoError(t, err)
	require.NoError(t, manifest.ValidateForProduction())

	server := web.New(manifest.Version)
	names := make([]string, 0, len(server.Tools()))
	for _, tool := range server.Tools() {
		names = append(names, tool.Name)
	}
	require.NoError(t, manifest.ValidateToolCatalog(names...))
}
