package usecases

import (
	"context"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/app/dto"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/domain/model"
	"slices"
)

// ToolUseCase is a use-case for tools
type ToolUseCase struct {
	hotpepperClient HOTPEPPERGOURMETAPIClient
}

// GourmetSearch retrieves gourmet information based on the provided search parameters.
func (u *ToolUseCase) GourmetSearch(ctx context.Context, input dto.GourmetSearchInput) ([]dto.Gourmet, error) {
	gourmets, err := u.hotpepperClient.FindGourmet(ctx, input)
	if err != nil {
		return nil, err
	}
	return gourmets, nil
}

// ShopSearch retrieves shop information based on the provided search parameters.
func (u *ToolUseCase) ShopSearch(ctx context.Context, input dto.ShopSearchInput) ([]dto.Shop, error) {
	shops, err := u.hotpepperClient.FindShop(ctx, input)
	if err != nil {
		return nil, err
	}
	return shops, nil
}

// LargeAreaSearch retrieves large area data based on the provided search parameters.
func (u *ToolUseCase) LargeAreaSearch(ctx context.Context, input dto.LargeAreaSearchInput) ([]dto.LargeArea, error) {
	largeAreas, err := u.hotpepperClient.FindLargeArea(ctx, input)
	if err != nil {
		return nil, err
	}
	return largeAreas, nil
}

// MiddleAreaSearch retrieves middle area data based on the provided search parameters.
func (u *ToolUseCase) MiddleAreaSearch(ctx context.Context, input dto.MiddleAreaSearchInput) ([]dto.MiddleArea, error) {
	middleAreas, err := u.hotpepperClient.FindMiddleArea(ctx, input)
	if err != nil {
		return nil, err
	}
	return middleAreas, nil
}

// SmallAreaSearch retrieves small area data based on the provided search parameters.
func (u *ToolUseCase) SmallAreaSearch(ctx context.Context, input dto.SmallAreaSearchInput) ([]dto.SmallArea, error) {
	smallAreas, err := u.hotpepperClient.FindSmallArea(ctx, input)
	if err != nil {
		return nil, err
	}
	return smallAreas, nil
}

// GenreSearch retrieves genre data based on the provided search parameters.
func (u *ToolUseCase) GenreSearch(ctx context.Context, input dto.GenreSearchInput) ([]dto.Genre, error) {
	genres, err := u.hotpepperClient.FindGenre(ctx, input)
	if err != nil {
		return nil, err
	}
	return genres, nil
}

// GetDinnerBudgetMaster retrieves the dinner budget master data from the HotPepper Gourmet API.
func (u *ToolUseCase) GetDinnerBudgetMaster(ctx context.Context, input dto.DinnerBudgetMasterSearchInput) ([]dto.Budget, error) {
	dinnerBudgetMaster, err := u.hotpepperClient.GetDinnerBudgetMaster(ctx)
	if err != nil {
		return nil, err
	}
	if input.Count > 0 {
		var i int
		start := int(input.Start)
		for v := range slices.Chunk(dinnerBudgetMaster, int(input.Count)) {
			i++
			if i != start {
				continue
			}
			return v, nil
		}
		return nil, nil
	}
	return dinnerBudgetMaster, nil
}

// GetLargeServiceAreaMaster retrieves the large service area master data from the HotPepper Gourmet API.
func (u *ToolUseCase) GetLargeServiceAreaMaster(ctx context.Context, input dto.LargeServiceAreaMasterSearchInput) ([]dto.LargeServiceArea, error) {
	largeServiceAreaMaster, err := u.hotpepperClient.GetLargeServiceAreaMaster(ctx)
	if err != nil {
		return nil, err
	}
	if input.Count > 0 {
		var i int
		start := int(input.Start)
		for v := range slices.Chunk(largeServiceAreaMaster, int(input.Count)) {
			i++
			if i != start {
				continue
			}
			return v, nil
		}
		return nil, nil
	}
	return largeServiceAreaMaster, nil
}

// GetServiceAreaMaster retrieves the service area master data from the HotPepper Gourmet API.
func (u *ToolUseCase) GetServiceAreaMaster(ctx context.Context, input dto.ServiceAreaMasterSearchInput) ([]dto.ServiceArea, error) {
	serviceAreaMaster, err := u.hotpepperClient.GetServiceAreaMaster(ctx)
	if err != nil {
		return nil, err
	}
	if input.Count > 0 {
		var i int
		start := int(input.Start)
		for v := range slices.Chunk(serviceAreaMaster, int(input.Count)) {
			i++
			if i != start {
				continue
			}
			return v, nil
		}
		return nil, nil
	}
	return serviceAreaMaster, nil
}

// GetCreditCardMaster retrieves the credit card master data from the HotPepper Gourmet API.
func (u *ToolUseCase) GetCreditCardMaster(ctx context.Context, input dto.CreditCardMasterSearchInput) ([]dto.CreditCard, error) {
	creditCardMaster, err := u.hotpepperClient.GetCreditCardMaster(ctx)
	if err != nil {
		return nil, err
	}
	if input.Count > 0 {
		var i int
		start := int(input.Start)
		for v := range slices.Chunk(creditCardMaster, int(input.Count)) {
			i++
			if i != start {
				continue
			}
			return v, nil
		}
		return nil, nil
	}
	return creditCardMaster, nil
}

// NewToolUseCase retuns a new ToolUseCase
func NewToolUseCase(hotpepperClient HOTPEPPERGOURMETAPIClient) *ToolUseCase {
	return &ToolUseCase{
		hotpepperClient: hotpepperClient,
	}
}

// HOTPEPPERGOURMETAPIClient is a client interface for the HotPepper Gourmet API.
type HOTPEPPERGOURMETAPIClient interface {
	// FindGourmet retrieves gourmet data based on the provided search parameters.
	FindGourmet(ctx context.Context, param model.GourmetSearchParam) ([]model.Gourmet, error)
	// FindShop retrieves shop data based on the provided search parameters.
	FindShop(ctx context.Context, param model.ShopSearchParam) ([]model.Shop, error)
	// FindLargeArea retrieves large area data based on the provided search parameters.
	FindLargeArea(ctx context.Context, param model.LargeAreaSearchParam) ([]model.LargeArea, error)
	// FindMiddleArea retrieves middle area data based on the provided search parameters.
	FindMiddleArea(ctx context.Context, param model.MiddleAreaSearchParam) ([]model.MiddleArea, error)
	// FindSmallArea retrieves small area data based on the provided search parameters.
	FindSmallArea(ctx context.Context, param model.SmallAreaSearchParam) ([]model.SmallArea, error)
	// FindGenre retrieves genre data based on the provided search parameters.
	FindGenre(ctx context.Context, param model.GenreSearchParam) ([]model.Genre, error)
	// GetDinnerBudgetMaster retrieves the dinner budget master data from the HotPepper Gourmet API.
	GetDinnerBudgetMaster(ctx context.Context) ([]model.Budget, error)
	// GetLargeServiceAreaMaster retrieves the large service area master data from the HotPepper Gourmet API.
	GetLargeServiceAreaMaster(ctx context.Context) ([]model.LargeServiceArea, error)
	// GetServiceAreaMaster retrieves the service area master data from the HotPepper Gourmet API.
	GetServiceAreaMaster(ctx context.Context) ([]model.ServiceArea, error)
	// GetCreditCardMaster retrieves the credit card master data from the HotPepper Gourmet API.
	GetCreditCardMaster(ctx context.Context) ([]model.CreditCard, error)
}
