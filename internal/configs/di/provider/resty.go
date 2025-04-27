package provider

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/interfaces/api/hotpepper"
)

// compatibility check
var _ hotpepper.RestyRequester = (*resty.Client)(nil)

var RestySet = wire.NewSet(
	resty.New,
	wire.Bind(new(hotpepper.RestyRequester), new(*resty.Client)))
