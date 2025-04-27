package provider

import (
	"github.com/google/wire"
	"github.com/ktr0731/go-mcp"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/interfaces/handler"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/interfaces/mcpgen"
)

var (
	_ mcp.ServerResourceHandler = (*handler.ResourceHandler)(nil)
	_ mcpgen.ServerToolHandler  = (*handler.ToolHandler)(nil)
)

var MCPGenSet = wire.NewSet(
	handler.NewResourceHandler,
	wire.Bind(new(mcp.ServerResourceHandler), new(*handler.ResourceHandler)),
	handler.NewToolHandler,
	wire.Bind(new(mcpgen.ServerToolHandler), new(*handler.ToolHandler)))
