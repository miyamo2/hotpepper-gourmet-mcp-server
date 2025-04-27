package handler

import (
	"context"
	"encoding/json"
	"github.com/cockroachdb/errors"
	"github.com/ktr0731/go-mcp"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/app/dto"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/interfaces/mcpgen"
	"log/slog"
	"strings"
)

// compatibility check
var _ mcp.ServerResourceHandler = (*ResourceHandler)(nil)

// ResourceHandler implements mcp.ServerResourceHandler.
type ResourceHandler struct {
	usecase ResourceUsecase
}

// HandleResourcesList See: mcp.ServerResourceHandler#HandleResourcesList
func (r *ResourceHandler) HandleResourcesList(_ context.Context) (*mcp.ListResourcesResult, error) {
	resources := make([]mcp.Resource, 0, len(mcpgen.ResourceTemplateList))
	for _, resource := range mcpgen.ResourceTemplateList {
		resources = append(resources, mcp.Resource{
			URI:         resource.URITemplate,
			Name:        resource.Name,
			Description: resource.Description,
			MimeType:    resource.MimeType,
		})
	}
	return &mcp.ListResourcesResult{
		Resources: resources,
	}, nil
}

// HandleResourcesRead See: mcp.ServerResourceHandler#HandleResourcesRead
func (r *ResourceHandler) HandleResourcesRead(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
	parts := strings.Split(req.URI, "/")
	if len(parts) < 5 {
		return nil, errors.New("invalid URI")
	}
	resource := parts[4]
	switch resource {
	case "budget":
		result, err := r.usecase.GetDinnerBudgetMaster(ctx)
		if err != nil {
			slog.Error("occurred error", slog.String("error", err.Error()))
			return nil, err
		}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			slog.Error("occurred error", slog.String("error", err.Error()))
			return nil, err
		}
		return &mcp.ReadResourceResult{
			Contents: []mcp.ResourceContent{
				mcp.TextResourceContent{
					URI:      req.URI,
					MimeType: "application/json",
					Text:     string(jsonResult),
				},
			},
		}, nil
	case "large_service_area":
		result, err := r.usecase.GetLargeServiceAreaMaster(ctx)
		if err != nil {
			slog.Error("occurred error", slog.String("error", err.Error()))
			return nil, err
		}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			slog.Error("occurred error", slog.String("error", err.Error()))
			return nil, err
		}
		return &mcp.ReadResourceResult{
			Contents: []mcp.ResourceContent{
				mcp.TextResourceContent{
					URI:      req.URI,
					MimeType: "application/json",
					Text:     string(jsonResult),
				},
			},
		}, nil
	case "service_area":
		result, err := r.usecase.GetServiceAreaMaster(ctx)
		if err != nil {
			slog.Error("occurred error", slog.String("error", err.Error()))
			return nil, err
		}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			slog.Error("occurred error", slog.String("error", err.Error()))
			return nil, err
		}
		return &mcp.ReadResourceResult{
			Contents: []mcp.ResourceContent{
				mcp.TextResourceContent{
					URI:      req.URI,
					MimeType: "application/json",
					Text:     string(jsonResult),
				},
			},
		}, nil
	case "credit_card":
		result, err := r.usecase.GetCreditCardMaster(ctx)
		if err != nil {
			slog.Error("occurred error", slog.String("error", err.Error()))
			return nil, err
		}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			slog.Error("occurred error", slog.String("error", err.Error()))
			return nil, err
		}
		return &mcp.ReadResourceResult{
			Contents: []mcp.ResourceContent{
				mcp.TextResourceContent{
					URI:      req.URI,
					MimeType: "application/json",
					Text:     string(jsonResult),
				},
			},
		}, nil
	}
	return nil, errors.New("invalid resource")
}

// NewResourceHandler returns a new ResourceHandler.
func NewResourceHandler(usecase ResourceUsecase) *ResourceHandler {
	return &ResourceHandler{
		usecase: usecase,
	}
}

// ResourceUsecase is an interface for the resource use case.
type ResourceUsecase interface {
	// GetDinnerBudgetMaster retrieves the dinner budget master data from the HotPepper Gourmet API.
	GetDinnerBudgetMaster(ctx context.Context) ([]dto.Budget, error)
	// GetLargeServiceAreaMaster retrieves the large service area master data from the HotPepper Gourmet API.
	GetLargeServiceAreaMaster(ctx context.Context) ([]dto.LargeServiceArea, error)
	// GetServiceAreaMaster retrieves the service area master data from the HotPepper Gourmet API.
	GetServiceAreaMaster(ctx context.Context) ([]dto.ServiceArea, error)
	// GetCreditCardMaster retrieves the credit card master data from the HotPepper Gourmet API.
	GetCreditCardMaster(ctx context.Context) ([]dto.CreditCard, error)
}
