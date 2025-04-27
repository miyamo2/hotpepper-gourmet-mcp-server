package handler

import (
	"context"
	"encoding/json"
	"github.com/ktr0731/go-mcp"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/app/dto"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/interfaces/mcpgen"
	"log/slog"
)

// compatibility check
var _ mcpgen.ServerToolHandler = (*ToolHandler)(nil)

// ToolHandler implements mcpgen.ServerToolHandler.
type ToolHandler struct {
	usecase ToolUseCase
}

// HandleToolGourmetSearch See: mcpgen.ServerToolHandler#HandleToolGourmetSearch
func (h *ToolHandler) HandleToolGourmetSearch(ctx context.Context, req *mcpgen.ToolGourmetSearchRequest) (*mcp.CallToolResult, error) {
	inputDTO := dto.GourmetSearchInput{
		ID:               req.ID,
		Name:             req.Name,
		NameKana:         req.NameKana,
		NameAny:          req.NameAny,
		Tel:              req.Tel,
		Address:          req.Address,
		Keyword:          req.Keyword,
		LargeServiceArea: req.LargeServiceArea,
		ServiceArea:      req.ServiceArea,
		LargeArea:        req.LargeArea,
		MiddleArea:       req.MiddleArea,
		SmallArea:        req.SmallArea,
		Latitude:         req.Latitude,
		Longitude:        req.Longitude,
		Range:            req.Range,
		Datum:            req.Datum,
		Genre:            req.Genre,
		Budget:           req.Budget,
		PartyCapacity:    req.PartyCapacity,
		KtaiCoupon:       req.KtaiCoupon,
		Wifi:             req.Wifi,
		Wedding:          req.Wedding,
		Course:           req.Course,
		FreeDrink:        req.FreeDrink,
		FreeFood:         req.FreeFood,
		PrivateRoom:      req.PrivateRoom,
		Horigotatsu:      req.Horigotatsu,
		Tatami:           req.Tatami,
		Cocktail:         req.Cocktail,
		Shochu:           req.Shochu,
		Sake:             req.Sake,
		Wine:             req.Wine,
		Card:             req.Card,
		NonSmoking:       req.NonSmoking,
		Charter:          req.Charter,
		Ktai:             req.Ktai,
		Parking:          req.Parking,
		BarrierFree:      req.BarrierFree,
		Sommelier:        req.Sommelier,
		NightView:        req.NightView,
		OpenAir:          req.OpenAir,
		Show:             req.Show,
		Equipment:        req.Equipment,
		Karaoke:          req.Karaoke,
		Band:             req.Band,
		TV:               req.TV,
		Lunch:            req.Lunch,
		Midnight:         req.Midnight,
		MidnightMeal:     req.MidnightMeal,
		English:          req.English,
		Pet:              req.Pet,
		Child:            req.Child,
		CreditCard:       req.CreditCard,
		Order:            req.Order,
		Start:            req.Start,
		Count:            req.Count,
	}
	result, err := h.usecase.GourmetSearch(ctx, inputDTO)
	if err != nil {
		slog.Error("occurred error", slog.String("error", err.Error()))
		return nil, err
	}
	return h.toJSONResult(result)
}

// HandleToolShopSearch See: mcpgen.ServerToolHandler#HandleToolShopSearch
func (h *ToolHandler) HandleToolShopSearch(ctx context.Context, req *mcpgen.ToolShopSearchRequest) (*mcp.CallToolResult, error) {
	inputDTO := dto.ShopSearchInput{
		Keyword: req.Keyword,
		Tel:     req.Tel,
		Start:   req.Start,
		Count:   req.Count,
	}
	result, err := h.usecase.ShopSearch(ctx, inputDTO)
	if err != nil {
		slog.Error("occurred error", slog.String("error", err.Error()))
		return nil, err
	}
	return h.toJSONResult(result)
}

// HandleToolLargeAreaSearch See: mcpgen.ServerToolHandler#HandleToolLargeAreaSearch
func (h *ToolHandler) HandleToolLargeAreaSearch(ctx context.Context, req *mcpgen.ToolLargeAreaSearchRequest) (*mcp.CallToolResult, error) {
	inputDTO := dto.LargeAreaSearchInput{
		LargeArea: req.LargeArea,
		Keyword:   req.Keyword,
	}
	result, err := h.usecase.LargeAreaSearch(ctx, inputDTO)
	if err != nil {
		slog.Error("occurred error", slog.String("error", err.Error()))
		return nil, err
	}
	return h.toJSONResult(result)
}

// HandleToolMiddleAreaSearch See: mcpgen.ServerToolHandler#HandleToolMiddleAreaSearch
func (h *ToolHandler) HandleToolMiddleAreaSearch(ctx context.Context, req *mcpgen.ToolMiddleAreaSearchRequest) (*mcp.CallToolResult, error) {
	inputDTO := dto.MiddleAreaSearchInput{
		MiddleArea: req.MiddleArea,
		LargeArea:  req.LargeArea,
		Keyword:    req.Keyword,
		Start:      req.Start,
		Count:      req.Count,
	}
	result, err := h.usecase.MiddleAreaSearch(ctx, inputDTO)
	if err != nil {
		slog.Error("occurred error", slog.String("error", err.Error()))
		return nil, err
	}
	return h.toJSONResult(result)
}

// HandleToolSmallAreaSearch See: mcpgen.ServerToolHandler#HandleToolSmallAreaSearch
func (h *ToolHandler) HandleToolSmallAreaSearch(ctx context.Context, req *mcpgen.ToolSmallAreaSearchRequest) (*mcp.CallToolResult, error) {
	inputDTO := dto.SmallAreaSearchInput{
		SmallArea:  req.SmallArea,
		MiddleArea: req.MiddleArea,
		Keyword:    req.Keyword,
		Start:      req.Start,
		Count:      req.Count,
	}
	result, err := h.usecase.SmallAreaSearch(ctx, inputDTO)
	if err != nil {
		slog.Error("occurred error", slog.String("error", err.Error()))
		return nil, err
	}
	return h.toJSONResult(result)
}

// HandleToolGenreSearch See: mcpgen.ServerToolHandler#HandleToolGenreSearch
func (h *ToolHandler) HandleToolGenreSearch(ctx context.Context, req *mcpgen.ToolGenreSearchRequest) (*mcp.CallToolResult, error) {
	inputDTO := dto.GenreSearchInput{
		Code:    req.Code,
		Keyword: req.Code,
	}
	result, err := h.usecase.GenreSearch(ctx, inputDTO)
	if err != nil {
		slog.Error("occurred error", slog.String("error", err.Error()))
		return nil, err
	}
	return h.toJSONResult(result)
}

func (h *ToolHandler) HandleToolDinnerBudgetMasterSearch(ctx context.Context, _ *mcpgen.ToolDinnerBudgetMasterSearchRequest) (*mcp.CallToolResult, error) {
	result, err := h.usecase.GetDinnerBudgetMaster(ctx)
	if err != nil {
		slog.Error("occurred error", slog.String("error", err.Error()))
		return nil, err
	}
	return h.toJSONResult(result)
}

func (h *ToolHandler) HandleToolLargeServiceAreaMasterSearch(ctx context.Context, _ *mcpgen.ToolLargeServiceAreaMasterSearchRequest) (*mcp.CallToolResult, error) {
	result, err := h.usecase.GetLargeServiceAreaMaster(ctx)
	if err != nil {
		slog.Error("occurred error", slog.String("error", err.Error()))
		return nil, err
	}
	return h.toJSONResult(result)
}

func (h *ToolHandler) HandleToolServiceAreaMasterSearch(ctx context.Context, _ *mcpgen.ToolServiceAreaMasterSearchRequest) (*mcp.CallToolResult, error) {
	result, err := h.usecase.GetServiceAreaMaster(ctx)
	if err != nil {
		slog.Error("occurred error", slog.String("error", err.Error()))
		return nil, err
	}
	return h.toJSONResult(result)
}

func (h *ToolHandler) HandleToolCreditCardMasterSearch(ctx context.Context, _ *mcpgen.ToolCreditCardMasterSearchRequest) (*mcp.CallToolResult, error) {
	result, err := h.usecase.GetCreditCardMaster(ctx)
	if err != nil {
		slog.Error("occurred error", slog.String("error", err.Error()))
		return nil, err
	}
	return h.toJSONResult(result)
}

// toJSONResult parses result to JSON format
func (h *ToolHandler) toJSONResult(result interface{}) (*mcp.CallToolResult, error) {
	jsonResult, err := json.Marshal(result)
	if err != nil {
		slog.Error("occurred error", slog.String("error", err.Error()))
		return nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.CallToolContent{
			mcp.TextContent{
				Text: string(jsonResult),
			},
		},
	}, nil
}

// ToolUseCase is an interface for the tool use-case.
type ToolUseCase interface {
	// GourmetSearch retrieves gourmet information based on the provided search parameters.
	GourmetSearch(ctx context.Context, input dto.GourmetSearchInput) ([]dto.Gourmet, error)
	// ShopSearch retrieves shop information based on the provided search parameters.
	ShopSearch(ctx context.Context, input dto.ShopSearchInput) ([]dto.Shop, error)
	// LargeAreaSearch retrieves large area information based on the provided search parameters.
	LargeAreaSearch(ctx context.Context, input dto.LargeAreaSearchInput) ([]dto.LargeArea, error)
	// MiddleAreaSearch retrieves middle area data based on the provided search parameters.
	MiddleAreaSearch(ctx context.Context, input dto.MiddleAreaSearchInput) ([]dto.MiddleArea, error)
	// SmallAreaSearch retrieves small area data based on the provided search parameters.
	SmallAreaSearch(ctx context.Context, input dto.SmallAreaSearchInput) ([]dto.SmallArea, error)
	// GenreSearch retrieves genre data based on the provided search parameters.
	GenreSearch(ctx context.Context, input dto.GenreSearchInput) ([]dto.Genre, error)
	// GetDinnerBudgetMaster retrieves the dinner budget master data from the HotPepper Gourmet API.
	GetDinnerBudgetMaster(ctx context.Context) ([]dto.Budget, error)
	// GetLargeServiceAreaMaster retrieves the large service area master data from the HotPepper Gourmet API.
	GetLargeServiceAreaMaster(ctx context.Context) ([]dto.LargeServiceArea, error)
	// GetServiceAreaMaster retrieves the service area master data from the HotPepper Gourmet API.
	GetServiceAreaMaster(ctx context.Context) ([]dto.ServiceArea, error)
	// GetCreditCardMaster retrieves the credit card master data from the HotPepper Gourmet API.
	GetCreditCardMaster(ctx context.Context) ([]dto.CreditCard, error)
}

// NewToolHandler returns a new ToolHandler.
func NewToolHandler(usecase ToolUseCase) *ToolHandler {
	return &ToolHandler{
		usecase: usecase,
	}
}
