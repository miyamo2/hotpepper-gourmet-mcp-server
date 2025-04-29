//go:build mcpgen

package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"

	"github.com/ktr0731/go-mcp/codegen"
)

func main() {
	// Create output directory
	outDir := "./internal/interfaces/mcpgen"
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	var buf bytes.Buffer

	// Server definition
	def := &codegen.ServerDefinition{
		Capabilities: codegen.ServerCapabilities{
			Tools: &codegen.ToolCapability{},
			Resources: &codegen.ResourceCapability{
				Subscribe:   true,
				ListChanged: true,
			},
			Logging: &codegen.LoggingCapability{},
		},
		Implementation: codegen.Implementation{
			Name:    "HOTPEPPER GOURMET MCP Server",
			Version: "0.1.0",
		},
		Tools: []codegen.Tool{
			{
				Name:        "gourmet_search",
				Description: "Searches for gourmet information.",
				InputSchema: struct {
					ID               []string `json:"id,omitempty" jsonschema:"title=Store ID,description=Search by the code assigned to the store. Up to 20 can be specified.,example=J999999999"`
					Name             *string  `json:"name,omitempty" jsonschema:"title=Store Name,description=Search by the store name (partial match)."`
					NameKana         *string  `json:"name_kana,omitempty" jsonschema:"title=Store Name Kana,description=Search by the store's phonetic name (partial match)."`
					NameAny          *string  `json:"name_any,omitempty" jsonschema:"title=Store Name OR Kana,description=Search across both store name and phonetic name with an OR condition (partial match)."`
					Tel              *string  `json:"tel,omitempty" jsonschema:"title=Phone Number,description=Search by the store's phone number. Digits only (no hyphens).,pattern=^\\d{9}(\\d|\\d{2})?$"`
					Address          *string  `json:"address,omitempty" jsonschema:"title=Address,description=Search by the store's address (partial match)."`
					Keyword          []string `json:"keyword,omitempty" jsonschema:"title=Keyword,description=Free-text search across store name kana, store name, address, station name, store genre catch, and catch. Multiple entries are possible. If multiple entries are given, they are treated as an AND condition."`
					LargeServiceArea *string  `json:"large_service_area,omitempty" jsonschema:"title=Large Service Area Code,description=Search by the code number assigned to the area. Refer to the Large Service Area master for possible codes."`
					ServiceArea      []string `json:"service_area,omitempty" jsonschema:"title=Service Area Code,description=Up to 3 entries allowed. Refer to the Service Area master for possible codes."`
					LargeArea        []string `json:"large_area,omitempty" jsonschema:"title=Large Area Code,description=Up to 3 entries allowed. Refer to the Large Area master for possible codes."`
					MiddleArea       []string `json:"middle_area,omitempty" jsonschema:"title=Middle Area Code,description=Up to 5 entries allowed. Refer to the Middle Area master for possible codes."`
					SmallArea        []string `json:"small_area,omitempty" jsonschema:"title=Small Area Code,description=Up to 5 entries allowed. Refer to the Small Area master for possible codes."`
					Latitude         *float64 `json:"lat,omitempty" jsonschema:"title=Latitude,description=Used if searching for stores within a certain distance from a given point."`
					Longitude        *float64 `json:"lng,omitempty" jsonschema:"title=Longitude,description=Used if searching for stores within a certain distance from a given point."`
					Range            *uint8   `json:"range,omitempty" jsonschema:"title=Search Range,description=Specifies a range in 5 levels for distance-based searching from a given point. For example, use range=1 for within 300m: 1=300m 2=500m 3=1000m (default)4=2000m 5=3000m,example=1"`
					Datum            *string  `json:"datum,omitempty" jsonschema:"title=Geodetic System,description=Specifies the geodetic system for latitude and longitude. 'world' for World Geodetic System, 'tokyo' for old Japanese Geodetic System. Default is 'world'."`
					Genre            []string `json:"genre,omitempty" jsonschema:"title=Store Genre Code,description=Can filter by the store's genre (including sub-genres). Refer to the Genre master for valid codes."`
					Budget           []string `json:"budget,omitempty" jsonschema:"title=Dinner Budget Code for Search,description=Can filter by dinner budget. Refer to the Dinner Budget master for valid codes."`
					PartyCapacity    *uint    `json:"party_capacity,omitempty" jsonschema:"title=Party Capacity,description=Can filter by party capacity. Finds stores with capacity larger than the specified number."`
					KtaiCoupon       bool     `json:"ktai_coupon,omitempty"  jsonschema:"title=Mobile Coupon Availability,description=Specifies whether to filter by mobile coupons. true=filter false=do not filter"`
					Wifi             bool     `json:"wifi,omitempty"         jsonschema:"title=WiFi Availability,description=Specifies whether to filter by stores offering internet access via WiFi. true=filter false=do not filter"`
					Wedding          bool     `json:"wedding,omitempty"      jsonschema:"title=Wedding/After-Parties,description=Specifies whether to filter by stores that accept wedding/after-party inquiries. true=filter false=do not filter"`
					Course           bool     `json:"course,omitempty"       jsonschema:"title=Courses,description=Specifies whether to filter by stores offering course meals. true=filter false=do not filter"`
					FreeDrink        bool     `json:"free_drink,omitempty"   jsonschema:"title=All-You-Can-Drink,description=Specifies whether to filter by stores offering all-you-can-drink. true=filter false=do not filter"`
					FreeFood         bool     `json:"free_food,omitempty"    jsonschema:"title=All-You-Can-Eat,description=Specifies whether to filter by stores offering all-you-can-eat. true=filter false=do not filter"`
					PrivateRoom      bool     `json:"private_room,omitempty" jsonschema:"title=Private Room,description=Specifies whether to filter by stores offering private rooms. true=filter false=do not filter"`
					Horigotatsu      bool     `json:"horigotatsu,omitempty"  jsonschema:"title=Sunken Kotatsu,description=Specifies whether to filter by stores offering sunken kotatsu seating. true=filter false=do not filter"`
					Tatami           bool     `json:"tatami,omitempty"       jsonschema:"title=Tatami,description=Specifies whether to filter by stores offering tatami seating. true=filter false=do not filter"`
					Cocktail         bool     `json:"cocktail,omitempty"     jsonschema:"title=Cocktail Selection,description=Specifies whether to filter by stores with a wide cocktail selection. true=filter false=do not filter"`
					Shochu           bool     `json:"shochu,omitempty"       jsonschema:"title=Shochu Selection,description=Specifies whether to filter by stores with a wide shochu selection. true=filter false=do not filter"`
					Sake             bool     `json:"sake,omitempty"         jsonschema:"title=Sake Selection,description=Specifies whether to filter by stores with a wide sake selection. true=filter false=do not filter"`
					Wine             bool     `json:"wine,omitempty"         jsonschema:"title=Wine Selection,description=Specifies whether to filter by stores with a wide wine selection. true=filter false=do not filter"`
					Card             bool     `json:"card,omitempty"         jsonschema:"title=Credit Cards Accepted,description=Specifies whether to filter by stores accepting credit cards. true=filter false=do not filter"`
					NonSmoking       bool     `json:"non_smoking,omitempty"  jsonschema:"title=Non-Smoking,description=Specifies whether to filter by stores offering non-smoking seats. true=filter false=do not filter"`
					Charter          bool     `json:"charter,omitempty"      jsonschema:"title=Charter,description=Specifies whether to filter by stores allowing private events. true=filter false=do not filter"`
					Ktai             bool     `json:"ktai,omitempty"         jsonschema:"title=Cell Phone OK,description=Specifies whether to filter by stores allowing cell phone use. true=filter false=do not filter"`
					Parking          bool     `json:"parking,omitempty"      jsonschema:"title=Parking Available,description=Specifies whether to filter by stores offering parking. true=filter false=do not filter"`
					BarrierFree      bool     `json:"barrier_free,omitempty" jsonschema:"title=Barrier-Free,description=Specifies whether to filter by barrier-free accessibility. true=filter false=do not filter"`
					Sommelier        bool     `json:"sommelier,omitempty"    jsonschema:"title=Sommelier,description=Specifies whether to filter by stores with a sommelier on staff. true=filter false=do not filter"`
					NightView        bool     `json:"night_view,omitempty"   jsonschema:"title=Night View,description=Specifies whether to filter by stores with a nice night view. true=filter false=do not filter"`
					OpenAir          bool     `json:"open_air,omitempty"     jsonschema:"title=Open Air,description=Specifies whether to filter by open-air spaces. true=filter false=do not filter"`
					Show             bool     `json:"show,omitempty"         jsonschema:"title=Live Shows,description=Specifies whether to filter by stores that offer live shows. true=filter false=do not filter"`
					Equipment        bool     `json:"equipment,omitempty"    jsonschema:"title=Entertainment Equipment,description=Specifies whether to filter by stores with entertainment equipment. true=filter false=do not filter"`
					Karaoke          bool     `json:"karaoke,omitempty"      jsonschema:"title=Karaoke,description=Specifies whether to filter by stores with karaoke. true=filter false=do not filter"`
					Band             bool     `json:"band,omitempty"         jsonschema:"title=Band Performances,description=Specifies whether to filter by stores allowing band performances. true=filter false=do not filter"`
					TV               bool     `json:"tv,omitempty"           jsonschema:"title=TV/Projector,description=Specifies whether to filter by stores with TV/projector. true=filter false=do not filter"`
					Lunch            bool     `json:"lunch,omitempty"        jsonschema:"title=Lunch Available,description=Specifies whether to filter by stores offering lunch. true=filter false=do not filter"`
					Midnight         bool     `json:"midnight,omitempty"     jsonschema:"title=Open After 11PM,description=Specifies whether to filter by stores open after 11PM. true=filter false=do not filter"`
					MidnightMeal     bool     `json:"midnight_meal,omitempty" jsonschema:"title=Meals After 11PM,description=Specifies whether to filter by stores offering meals after 11PM. true=filter false=do not filter"`
					English          bool     `json:"english,omitempty"      jsonschema:"title=English Menu Available,description=Specifies whether to filter by stores offering an English menu. true=filter false=do not filter"`
					Pet              bool     `json:"pet,omitempty"          jsonschema:"title=Pets Allowed,description=Specifies whether to filter by pet-friendly stores. true=filter false=do not filter"`
					Child            bool     `json:"child,omitempty"        jsonschema:"title=Children Allowed,description=Specifies whether to filter by stores that allow children. true=filter false=do not filter"`
					CreditCard       []string `json:"credit_card,omitempty" jsonschema:"title=Credit Cards,description=Can filter by specific credit card types. Refer to the Credit Card master for possible codes. (Added 2008/02/08)"`
					Order            *uint8   `json:"order,omitempty"       jsonschema:"title=Sort Order,description=Specifies how to sort the search results. The recommended order is updated periodically. * If searching by location, all orders other than 4 are ignored and distance order is applied. 1=Store name kana 2=Genre code 3=Small area code 4=Recommended (default). If searching by location, distance order is forced."`
					Start            *uint    `json:"start,omitempty"       jsonschema:"title=Start Position,description=Specifies which result index to begin output. Default is 1."`
					Count            *uint8   `json:"count,omitempty"       jsonschema:"title=Items per Page,description=Specifies how many results to output at once. Default is 10. Minimum 1, maximum 100."`
				}{},
			},
			{
				Name:        "shop_search",
				Description: "Searches for restaurant information.",
				InputSchema: struct {
					Keyword []string `json:"keyword,omitempty" jsonschema:"title=Keyword,description=Search by the store's name, phonetic reading, or address (partial match). Multiple entries allowed."`
					Tel     *string  `json:"tel,omitempty" jsonschema:"title=Phone Number,description=Search by the store's phone number (digits only, no hyphens).,pattern=^\\d{9}(\\d|\\d{2})?$"`
					Start   *uint    `json:"start,omitempty" jsonschema:"title=Start Position,description=Specifies which search result index to start from. The default is 1."`
					Count   *uint8   `json:"count,omitempty" jsonschema:"title=Items Per Page,description=Specifies the maximum number of results to return per page. The default is 30. Minimum is 1, maximum is 30."`
				}{},
			},
			{
				Name:        "large_area_search",
				Description: "Searches for large area.",
				InputSchema: struct {
					LargeArea []string `json:"large_area,omitempty" jsonschema:"title=Large Area Code,description=Search by large area code (exact match). Up to 3 can be specified. If 4 or more are specified, entries from the 4th onward will be ignored."`
					Keyword   []string `json:"keyword,omitempty"   jsonschema:"title=Keyword,description=Search by large area name (partial match)."`
				}{},
			},
			{
				Name:        "middle_area_search",
				Description: "Searches for middle area.",
				InputSchema: struct {
					MiddleArea []string `json:"middle_area,omitempty" jsonschema:"title=Middle Area Code,description=Search by middle area code (exact match). Up to 5 can be specified. If 6 or more are specified, the 6th and onward are ignored."`
					LargeArea  []string `json:"large_area,omitempty"  jsonschema:"title=Large Area Code,description=Search by large area code (exact match). Up to 3 can be specified. If 4 or more are specified, the 4th and onward are ignored."`
					Keyword    []string `json:"keyword,omitempty"     jsonschema:"title=Keyword,description=Search by middle area name (partial match).,example=飯田橋"`
					Start      *uint    `json:"start,omitempty"       jsonschema:"title=Start Position,description=Specifies which result index to begin output from. Default is 1."`
					Count      *uint8   `json:"count,omitempty"       jsonschema:"title=Items Per Page,description=Specifies the maximum number of search results per page. Default is unset. Minimum is 1, maximum is unset (retrieve all)."`
				}{},
			},
			{
				Name:        "small_area_search",
				Description: "Searches for small area.",
				InputSchema: struct {
					SmallArea  []string `json:"small_area,omitempty" jsonschema:"title=Small Area Code,description=Search by small area code (exact match). Up to 5 can be specified. If 6 or more are specified, entries from the 6th onward will be ignored."`
					MiddleArea []string `json:"middle_area,omitempty" jsonschema:"title=Middle Area Code,description=Search by middle area code (exact match). Up to 5 can be specified. If 6 or more are specified, entries from the 6th onward will be ignored."`
					Keyword    []string `json:"keyword,omitempty"    jsonschema:"title=Keyword,description=Search by small area name (partial match).,example=銀座"`
					Start      *uint    `json:"start,omitempty"      jsonschema:"title=Start Position,description=Specifies which search result index to start from. Default is 1."`
					Count      *uint8   `json:"count,omitempty"      jsonschema:"title=Items Per Page,description=Specifies the maximum number of results per page. Default is unset. Minimum is 1, maximum is unset (all results)."`
				}{},
			},
			{
				Name:        "genre_search",
				Description: "Searches for genre.",
				InputSchema: struct {
					Code    []string `json:"code,omitempty" jsonschema:"title=Genre Code,description=Search by genre code (exact match). Up to 2 can be specified. If 3 or more are specified, entries from the 3rd onward will be ignored."`
					Keyword []string `json:"keyword,omitempty" jsonschema:"title=Keyword,description=Search by genre name (partial match).,example=バー"`
				}{},
			},
			{
				Name:        "dinner_budget_master_search",
				Description: "Retrieves diner_budget_master.",
				InputSchema: struct {
					Start *uint `json:"start,omitempty" jsonschema:"title=Start Position,description=Specifies which search result index to start from. Default is 1."`
					Count *uint `json:"count,omitempty" jsonschema:"title=Items Per Page,description=Specifies the maximum number of search results to return per page. Default is unset. Minimum is 1, maximum is unset (retrieve all), recommended to keep unset (retrieve all)."`
				}{},
			},
			{
				Name:        "large_service_area_master_search",
				Description: "Retrieves large_service_area_master",
				InputSchema: struct {
					Start *uint `json:"start,omitempty" jsonschema:"title=Start Position,description=Specifies which search result index to begin output from. Default is 1."`
					Count *uint `json:"count,omitempty" jsonschema:"title=Items Per Page,description=Specifies the maximum number of search results per page. Default is unset. Minimum is 1, maximum is unset (all results), recommended to remain unset (all results)."`
				}{},
			},
			{
				Name:        "service_area_master_search",
				Description: "Retrieves service_area_master",
				InputSchema: struct {
					Start *uint `json:"start,omitempty" jsonschema:"title=Start Position,description=Specifies which search result index to start from. Default is 1."`
					Count *uint `json:"count,omitempty" jsonschema:"title=Items Per Page,description=Specifies the maximum number of search results per page. Default is unset. Minimum is 1, maximum is unset (retrieve all), recommended to remain unset (retrieve all)."`
				}{},
			},
			{
				Name:        "credit_card_master_search",
				Description: "Retrieves credit_card_master",
				InputSchema: struct {
					Start *uint `json:"start,omitempty" jsonschema:"title=Start Position,description=Specifies which search result index to start from. Default is 1."`
					Count *uint `json:"count,omitempty" jsonschema:"title=Items Per Page,description=Specifies the maximum number of search results per page. Default is unset. Minimum is 1, maximum is unset (retrieve all), recommended to remain unset (retrieve all)."`
				}{},
			},
		},
		ResourceTemplates: []codegen.ResourceTemplate{
			{
				Name:        "dinner_budget_master",
				Description: "dinner_budget_master",
				URITemplate: "https://webservice.recruit.co.jp/hotpepper/budget/v1/",
				MimeType:    "application/json",
			},
			{
				Name:        "large_service_area_master",
				Description: "large_service_area_master",
				URITemplate: "https://webservice.recruit.co.jp/hotpepper/large_service_area/v1/",
				MimeType:    "application/json",
			},
			{
				Name:        "service_area_master",
				Description: "service_area_master",
				URITemplate: "https://webservice.recruit.co.jp/hotpepper/service_area/v1/",
				MimeType:    "application/json",
			},
			{
				Name:        "credit_card_master",
				Description: "credit_card_master",
				URITemplate: "https://webservice.recruit.co.jp/hotpepper/credit_card/v1/",
				MimeType:    "application/json",
			},
		},
	}

	// Generate code
	if err := codegen.Generate(&buf, def, "mcpgen"); err != nil {
		log.Fatalf("failed to generate code: %v", err)
	}

	f, err := os.Create(filepath.Join(outDir, "mcp.gen.go"))
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer f.Close()

	data := bytes.ReplaceAll(buf.Bytes(), []byte(`"$schema":"https://json-schema.org/draft/2020-12/schema",`), []byte(""))
	_, err = f.Write(data)
	if err != nil {
		log.Fatalf("failed to write file: %v", err)
	}
}
