package hotpepper

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/google/go-querystring/query"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/domain/model"
	"net/url"
)

// Client communicates with the HotPepper Gourmet API.
type Client struct {
	baseURL        *url.URL
	apiToken       string
	restyRequester RestyRequester
}

// GetDinnerBudgetMaster retrieves the dinner budget master data from the HotPepper Gourmet API.
func (c *Client) GetDinnerBudgetMaster(ctx context.Context) ([]model.Budget, error) {
	var result model.Response[model.GetDinnerBudgetMasterResult]

	queryParams := url.Values{}
	queryParams.Add("key", c.apiToken)
	queryParams.Add("format", "json")

	endpoint := c.baseURL.JoinPath(pathDinnerBudget, pathV1)
	req := c.restyRequester.R().
		SetContext(ctx).
		SetQueryParamsFromValues(queryParams)

	res, err := req.Get(endpoint.String())
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, fmt.Errorf(errorMessageBase, res.Error())
	}

	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		return nil, err
	}
	return result.Results.Budget, nil
}

// GetLargeServiceAreaMaster retrieves the large service area master data from the HotPepper Gourmet API.
func (c *Client) GetLargeServiceAreaMaster(ctx context.Context) ([]model.LargeServiceArea, error) {
	var result model.Response[model.GetLargeServiceAreaMasterResult]

	queryParams := url.Values{}
	queryParams.Add("key", c.apiToken)
	queryParams.Add("format", "json")

	endpoint := c.baseURL.JoinPath(pathLargeServiceArea, pathV1)
	req := c.restyRequester.R().
		SetContext(ctx).
		SetQueryParamsFromValues(queryParams)

	res, err := req.Get(endpoint.String())
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, fmt.Errorf(errorMessageBase, res.Error())
	}

	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		return nil, err
	}
	return result.Results.LargeServiceArea, nil
}

// GetServiceAreaMaster retrieves the service area master data from the HotPepper Gourmet API.
func (c *Client) GetServiceAreaMaster(ctx context.Context) ([]model.ServiceArea, error) {
	var result model.Response[model.GetServiceAreaMasterResult]

	queryParams := url.Values{}
	queryParams.Add("key", c.apiToken)
	queryParams.Add("format", "json")

	endpoint := c.baseURL.JoinPath(pathServiceArea, pathV1)
	req := c.restyRequester.R().
		SetContext(ctx).
		SetQueryParamsFromValues(queryParams)

	res, err := req.Get(endpoint.String())
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, fmt.Errorf(errorMessageBase, res.Error())
	}

	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		return nil, err
	}
	return result.Results.ServiceArea, nil
}

// GetCreditCardMaster retrieves the credit card master data from the HotPepper Gourmet API.
func (c *Client) GetCreditCardMaster(ctx context.Context) ([]model.CreditCard, error) {
	var result model.Response[model.GetCreditCardMasterResult]

	queryParams := url.Values{}
	queryParams.Add("key", c.apiToken)
	queryParams.Add("format", "json")

	endpoint := c.baseURL.JoinPath(pathCreditCard, pathV1)
	req := c.restyRequester.R().
		SetContext(ctx).
		SetQueryParamsFromValues(queryParams)

	res, err := req.Get(endpoint.String())
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, fmt.Errorf(errorMessageBase, res.Error())
	}

	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		return nil, err
	}
	return result.Results.CreditCard, nil
}

// FindGourmet retrieves gourmet data based on the provided search parameters.
func (c *Client) FindGourmet(ctx context.Context, param model.GourmetSearchParam) ([]model.Gourmet, error) {
	var result model.Response[model.GourmetSearchResult]
	queryParams, err := query.Values(param)
	if err != nil {
		return nil, err
	}
	queryParams.Add("key", c.apiToken)
	queryParams.Add("format", "json")
	queryParams.Add("type", "lite")

	endpoint := c.baseURL.JoinPath(pathGourmet, pathV1)
	req := c.restyRequester.R().
		SetContext(ctx).
		SetQueryParamsFromValues(queryParams)

	res, err := req.Get(endpoint.String())
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, fmt.Errorf(errorMessageBase, res.Error())
	}

	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		return nil, err
	}
	return result.Results.Shop, nil
}

// FindShop retrieves shop data based on the provided search parameters.
func (c *Client) FindShop(ctx context.Context, param model.ShopSearchParam) ([]model.Shop, error) {
	var result model.Response[model.ShopSearchResult]
	queryParams, err := query.Values(param)
	if err != nil {
		return nil, err
	}
	queryParams.Add("key", c.apiToken)
	queryParams.Add("format", "json")

	endpoint := c.baseURL.JoinPath(pathShop, pathV1)
	req := c.restyRequester.R().
		SetContext(ctx).
		SetQueryParamsFromValues(queryParams)

	res, err := req.Get(endpoint.String())
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, fmt.Errorf(errorMessageBase, res.Error())
	}

	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		return nil, err
	}
	return result.Results.Shop, nil
}

// FindLargeArea retrieves large area data based on the provided search parameters.
func (c *Client) FindLargeArea(ctx context.Context, param model.LargeAreaSearchParam) ([]model.LargeArea, error) {
	var result model.Response[model.LargeAreaSearchResult]
	queryParams, err := query.Values(param)
	if err != nil {
		return nil, err
	}
	queryParams.Add("key", c.apiToken)
	queryParams.Add("format", "json")

	endpoint := c.baseURL.JoinPath(pathLargeArea, pathV1)
	req := c.restyRequester.R().
		SetContext(ctx).
		SetQueryParamsFromValues(queryParams)

	res, err := req.Get(endpoint.String())
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, fmt.Errorf(errorMessageBase, res.Error())
	}

	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		return nil, err
	}
	return result.Results.LargeArea, nil
}

// FindMiddleArea retrieves middle area data based on the provided search parameters.
func (c *Client) FindMiddleArea(ctx context.Context, param model.MiddleAreaSearchParam) ([]model.MiddleArea, error) {
	var result model.Response[model.MiddleAreaSearchResult]
	queryParams, err := query.Values(param)
	if err != nil {
		return nil, err
	}
	queryParams.Add("key", c.apiToken)
	queryParams.Add("format", "json")

	endpoint := c.baseURL.JoinPath(pathMiddleArea, pathV1)
	req := c.restyRequester.R().
		SetContext(ctx).
		SetQueryParamsFromValues(queryParams)

	res, err := req.Get(endpoint.String())
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, fmt.Errorf(errorMessageBase, res.Error())
	}

	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		return nil, err
	}
	return result.Results.MiddleArea, nil
}

// FindSmallArea retrieves small area data based on the provided search parameters.
func (c *Client) FindSmallArea(ctx context.Context, param model.SmallAreaSearchParam) ([]model.SmallArea, error) {
	var result model.Response[model.SmallAreaSearchResult]
	queryParams, err := query.Values(param)
	if err != nil {
		return nil, err
	}
	queryParams.Add("key", c.apiToken)
	queryParams.Add("format", "json")

	endpoint := c.baseURL.JoinPath(pathSmallArea, pathV1)
	req := c.restyRequester.R().
		SetContext(ctx).
		SetQueryParamsFromValues(queryParams)

	res, err := req.Get(endpoint.String())
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, fmt.Errorf(errorMessageBase, res.Error())
	}

	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		return nil, err
	}
	return result.Results.SmallArea, nil
}

// FindGenre retrieves genre data based on the provided search parameters.
func (c *Client) FindGenre(ctx context.Context, param model.GenreSearchParam) ([]model.Genre, error) {
	var result model.Response[model.GenreSearchResult]
	queryParams, err := query.Values(param)
	if err != nil {
		return nil, err
	}
	queryParams.Add("key", c.apiToken)
	queryParams.Add("format", "json")

	endpoint := c.baseURL.JoinPath(pathGenre, pathV1)
	req := c.restyRequester.R().
		SetContext(ctx).
		SetQueryParamsFromValues(queryParams)

	res, err := req.Get(endpoint.String())
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, fmt.Errorf(errorMessageBase, res.Error())
	}

	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		return nil, err
	}
	return result.Results.Genre, nil
}

// NewClient returns a new Client with the given base URL and API token.
func NewClient(baseURL *url.URL, apiToken string, restyRequester RestyRequester) *Client {
	return &Client{
		baseURL:        baseURL,
		apiToken:       apiToken,
		restyRequester: restyRequester,
	}
}

// RestyRequester generates a new resty.Request
type RestyRequester interface {
	R() *resty.Request
}

// paths
const (
	pathV1               = "v1"
	pathDinnerBudget     = "budget"
	pathLargeServiceArea = "large_service_area"
	pathServiceArea      = "service_area"
	pathCreditCard       = "credit_card"
	pathGourmet          = "gourmet"
	pathShop             = "shop"
	pathLargeArea        = "large_area"
	pathMiddleArea       = "middle_area"
	pathSmallArea        = "small_area"
	pathGenre            = "genre"
)

// query string bases
const (
	queryStringBaseFindGourmet = "%s&key=%s&format=json&type=lite"
)

const errorMessageBase = "error: %v"
