package provider

import (
	"github.com/google/wire"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/app/usecases"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/interfaces/handler"
)

// compatibility check
var (
	_ handler.ToolUseCase     = (*usecases.ToolUseCase)(nil)
	_ handler.ResourceUsecase = (*usecases.ResourceUsecase)(nil)
)

var UsecasesSet = wire.NewSet(
	usecases.NewToolUseCase,
	wire.Bind(new(handler.ToolUseCase), new(*usecases.ToolUseCase)),
	usecases.NewResourceUsecase,
	wire.Bind(new(handler.ResourceUsecase), new(*usecases.ResourceUsecase)))
