package usecases

import (
	"context"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/app/dto"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/domain/model"
)

// ResourceUsecase is a use-case for managing resources.
type ResourceUsecase struct {
	hotpepperClient HOTPEPPERGOURMETMasterClient
}

// GetDinnerBudgetMaster retrieves the dinner budget master data from the HotPepper Gourmet API.
func (u *ResourceUsecase) GetDinnerBudgetMaster(ctx context.Context) ([]dto.Budget, error) {
	dinnerBudgetMaster, err := u.hotpepperClient.GetDinnerBudgetMaster(ctx)

	if err != nil {
		return nil, err
	}
	return dinnerBudgetMaster, nil
}

// GetLargeServiceAreaMaster retrieves the large service area master data from the HotPepper Gourmet API.
func (u *ResourceUsecase) GetLargeServiceAreaMaster(ctx context.Context) ([]dto.LargeServiceArea, error) {
	largeServiceAreaMaster, err := u.hotpepperClient.GetLargeServiceAreaMaster(ctx)
	if err != nil {
		return nil, err
	}
	return largeServiceAreaMaster, nil
}

// GetServiceAreaMaster retrieves the service area master data from the HotPepper Gourmet API.
func (u *ResourceUsecase) GetServiceAreaMaster(ctx context.Context) ([]dto.ServiceArea, error) {
	serviceAreaMaster, err := u.hotpepperClient.GetServiceAreaMaster(ctx)
	if err != nil {
		return nil, err
	}
	return serviceAreaMaster, nil
}

// GetCreditCardMaster retrieves the credit card master data from the HotPepper Gourmet API.
func (u *ResourceUsecase) GetCreditCardMaster(ctx context.Context) ([]dto.CreditCard, error) {
	creditCardMaster, err := u.hotpepperClient.GetCreditCardMaster(ctx)
	if err != nil {
		return nil, err
	}
	return creditCardMaster, nil
}

// HOTPEPPERGOURMETMasterClient is a client interface for HotPepper Gourmet Master Data.
type HOTPEPPERGOURMETMasterClient interface {
	// GetDinnerBudgetMaster retrieves the dinner budget master data from the HotPepper Gourmet API.
	GetDinnerBudgetMaster(ctx context.Context) ([]model.Budget, error)
	// GetLargeServiceAreaMaster retrieves the large service area master data from the HotPepper Gourmet API.
	GetLargeServiceAreaMaster(ctx context.Context) ([]model.LargeServiceArea, error)
	// GetServiceAreaMaster retrieves the service area master data from the HotPepper Gourmet API.
	GetServiceAreaMaster(ctx context.Context) ([]model.ServiceArea, error)
	// GetCreditCardMaster retrieves the credit card master data from the HotPepper Gourmet API.
	GetCreditCardMaster(ctx context.Context) ([]model.CreditCard, error)
}

// NewResourceUsecase returns a new ResourceUsecase.
func NewResourceUsecase(hotpepperClient HOTPEPPERGOURMETMasterClient) *ResourceUsecase {
	return &ResourceUsecase{
		hotpepperClient: hotpepperClient,
	}
}
