//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/ktr0731/go-mcp"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/configs/di/provider"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/interfaces/mcpgen"
)

func GetHandler() *mcp.Handler {
	wire.Build(
		provider.RestySet,
		provider.HOTPEPPERSet,
		provider.UsecasesSet,
		provider.MCPGenSet,
		wire.NewSet(mcpgen.NewHandler),
	)
	return nil
}
