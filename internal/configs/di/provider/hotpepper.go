package provider

import (
	"github.com/google/wire"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/app/usecases"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/interfaces/api/hotpepper"
	"net/url"
	"os"
)

var (
	_ usecases.HOTPEPPERGOURMETMasterClient = (*hotpepper.Client)(nil)
	_ usecases.HOTPEPPERGOURMETAPIClient    = (*hotpepper.Client)(nil)
)

func NewHOTPEPPERGOURMETClient(restyRequester hotpepper.RestyRequester) *hotpepper.Client {
	baseURL, err := url.Parse("https://webservice.recruit.co.jp/hotpepper")
	if err != nil {
		panic(err)
	}
	return hotpepper.NewClient(baseURL, os.Getenv("HOTPEPPER_GOURMET_API_KEY"), restyRequester)
}

var HOTPEPPERSet = wire.NewSet(
	NewHOTPEPPERGOURMETClient,
	wire.Bind(new(usecases.HOTPEPPERGOURMETMasterClient), new(*hotpepper.Client)),
	wire.Bind(new(usecases.HOTPEPPERGOURMETAPIClient), new(*hotpepper.Client)))
