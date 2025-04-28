package model

// Response indicates the response from the HotPepper Gourmet API
type Response[T any] struct {
	Results T `json:"results"`
}

// Meta indicates the metadata of the response from the HotPepper Gourmet API
type Meta struct {
	APIVersion       string `json:"api_version"`
	ResultsAvailable int    `json:"results_available"`
	ResultsReturned  string `json:"results_returned"`
	ResultsStart     int    `json:"results_start"`
}

// GetDinnerBudgetMasterResult indicates the response from the HotPepper Gourmet API for dinner budget master
type GetDinnerBudgetMasterResult struct {
	Meta
	Budget []Budget `json:"budget"`
}

// Budget indicates the budget information from the HotPepper Gourmet API
type Budget struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// GetLargeServiceAreaMasterResult indicates the response from the HotPepper Gourmet API for large service area master
type GetLargeServiceAreaMasterResult struct {
	Meta
	LargeServiceArea []LargeServiceArea `json:"large_service_area"`
}

// LargeServiceArea indicates the large service area information from the HotPepper Gourmet API
type LargeServiceArea struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// GetServiceAreaMasterResult indicates the response from the HotPepper Gourmet API for service area master
type GetServiceAreaMasterResult struct {
	Meta
	ServiceArea []ServiceArea `json:"service_area"`
}

// ServiceArea indicates the service area information from the HotPepper Gourmet API
type ServiceArea struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// GetCreditCardMasterResult indicates the response from the HotPepper Gourmet API for credit card master
type GetCreditCardMasterResult struct {
	Meta
	CreditCard []CreditCard `json:"credit_card"`
}

// CreditCard indicates the credit card information from the HotPepper Gourmet API
type CreditCard struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// GourmetSearchParam indicates the gourmet search parameters for the HotPepper Gourmet API
type GourmetSearchParam struct {
	ID               []string `url:"id,omitempty"`
	Name             *string  `url:"name,omitempty"`
	NameKana         *string  `url:"name_kana,omitempty"`
	NameAny          *string  `url:"name_any,omitempty"`
	Tel              *string  `url:"tel,omitempty"`
	Address          *string  `url:"address,omitempty"`
	Keyword          []string `url:"keyword,omitempty"`
	LargeServiceArea *string  `url:"large_service_area,omitempty"`
	ServiceArea      []string `url:"service_area,omitempty"`
	LargeArea        []string `url:"large_area,omitempty"`
	MiddleArea       []string `url:"middle_area,omitempty"`
	SmallArea        []string `url:"small_area,omitempty"`
	Latitude         *float64 `url:"lat,omitempty"`
	Longitude        *float64 `url:"lng,omitempty"`
	Range            *uint8   `url:"range,omitempty"`
	Datum            *string  `url:"datum,omitempty"`
	Genre            []string `url:"genre,omitempty"`
	Budget           []string `url:"budget,omitempty"`
	PartyCapacity    *uint    `url:"party_capacity,omitempty"`
	KtaiCoupon       bool     `url:"ktai_coupon,int,omitempty"`
	Wifi             bool     `url:"wifi,int,omitempty"`
	Wedding          bool     `url:"wedding,int,omitempty"`
	Course           bool     `url:"course,int,omitempty"`
	FreeDrink        bool     `url:"free_drink,int,omitempty"`
	FreeFood         bool     `url:"free_food,int,omitempty"`
	PrivateRoom      bool     `url:"private_room,int,omitempty"`
	Horigotatsu      bool     `url:"horigotatsu,int,omitempty"`
	Tatami           bool     `url:"tatami,int,omitempty"`
	Cocktail         bool     `url:"cocktail,int,omitempty"`
	Shochu           bool     `url:"shochu,int,omitempty"`
	Sake             bool     `url:"sake,int,omitempty"`
	Wine             bool     `url:"wine,int,omitempty"`
	Card             bool     `url:"card,int,omitempty"`
	NonSmoking       bool     `url:"non_smoking,int,omitempty"`
	Charter          bool     `url:"charter,int,omitempty"`
	Ktai             bool     `url:"ktai,int,omitempty"`
	Parking          bool     `url:"parking,int,omitempty"`
	BarrierFree      bool     `url:"barrier_free,int,omitempty"`
	Sommelier        bool     `url:"sommelier,int,omitempty"`
	NightView        bool     `url:"night_view,int,omitempty"`
	OpenAir          bool     `url:"open_air,int,omitempty"`
	Show             bool     `url:"show,int,omitempty"`
	Equipment        bool     `url:"equipment,int,omitempty"`
	Karaoke          bool     `url:"karaoke,int,omitempty"`
	Band             bool     `url:"band,int,omitempty"`
	TV               bool     `url:"tv,int,omitempty"`
	Lunch            bool     `url:"lunch,int,omitempty"`
	Midnight         bool     `url:"midnight,int,omitempty"`
	MidnightMeal     bool     `url:"midnight_meal,int,omitempty"`
	English          bool     `url:"english,int,omitempty"`
	Pet              bool     `url:"pet,int,omitempty"`
	Child            bool     `url:"child,int,omitempty"`
	CreditCard       []string `url:"credit_card,omitempty"`
	Order            *uint8   `url:"order,omitempty"`
	Start            *uint    `url:"start,omitempty"`
	Count            *uint8   `url:"count,omitempty"`
}

// GourmetSearchResult indicates the response from the HotPepper Gourmet API for gourmet search
type GourmetSearchResult struct {
	Meta
	Shop []Gourmet `json:"shop"`
}

// Gourmet indicates the gourmet information from the HotPepper Gourmet API
type Gourmet struct {
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	Address        string       `json:"address"`
	Lat            float64      `json:"lat"`
	Lng            float64      `json:"lng"`
	Genre          GourmetGenre `json:"genre"`
	Catch          string       `json:"catch"`
	Access         string       `json:"access"`
	URLs           GourmetURL   `json:"urls"`
	Photo          GourmetPhoto `json:"photo"`
	OtherMemo      string       `json:"other_memo"`
	ShopDetailMemo string       `json:"shop_detail_memo"`
	BudgetMemo     string       `json:"budget_memo"`
}

// GourmetGenre indicates the genre information of the gourmet from the HotPepper Gourmet API
type GourmetGenre struct {
	Name  string `json:"name"`
	Catch string `json:"catch"`
}

// GourmetURL indicates the URL information of the gourmet from the HotPepper Gourmet API
type GourmetURL struct {
	PC string `json:"pc"`
}

// GourmetPhoto indicates the photo information of the gourmet from the HotPepper Gourmet API
type GourmetPhoto struct {
	PC struct {
		L string `json:"l"`
		M string `json:"m"`
		S string `json:"s"`
	} `json:"pc"`
}

// ShopSearchResult indicates the response from the HotPepper Gourmet API for shop search
type ShopSearchResult struct {
	Meta
	Shop []Shop `json:"shop"`
}

// ShopSearchParam indicates the shop search parameters for the HotPepper Gourmet API
type ShopSearchParam struct {
	Keyword []string `url:"keyword,omitempty"`
	Tel     *string  `url:"tel"`
	Start   *uint    `url:"start,omitempty"`
	Count   *uint8   `url:"count,omitempty"`
}

// Shop indicates the shop information from the HotPepper Gourmet API
type Shop struct {
	Address  string    `json:"address"`
	Desc     string    `json:"desc"`
	Genre    ShopGenre `json:"genre"`
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	NameKana string    `json:"name_kana"`
	URLs     ShopURL   `json:"urls"`
}

// ShopGenre indicates the genre information of the shop from the HotPepper Gourmet API
type ShopGenre struct {
	Name string `json:"name"`
}

// ShopURL indicates the URL information of the shop from the HotPepper Gourmet API
type ShopURL struct {
	PC string `json:"pc"`
}

// LargeAreaSearchParam indicates the large area search parameters for the HotPepper Gourmet API
type LargeAreaSearchParam struct {
	LargeArea []string `url:"large_area"`
	Keyword   []string `url:"keyword,omitempty"`
}

// LargeAreaSearchResult indicates the response from the HotPepper Gourmet API for large area search
type LargeAreaSearchResult struct {
	Meta
	LargeArea []LargeArea `json:"large_area"`
}

// LargeArea indicates the large area information from the HotPepper Gourmet API
type LargeArea struct {
	Code             string           `json:"code"`
	Name             string           `json:"name"`
	ServiceArea      ServiceArea      `json:"service_area"`
	LargeServiceArea LargeServiceArea `json:"large_service_area"`
}

// MiddleAreaSearchParam indicates the middle area search parameters for the HotPepper Gourmet API
type MiddleAreaSearchParam struct {
	MiddleArea []string `url:"middle_area"`
	LargeArea  []string `url:"large_area,omitempty"`
	Keyword    []string `url:"keyword,omitempty"`
	Start      *uint    `url:"start,omitempty"`
	Count      *uint8   `url:"count,omitempty"`
}

// MiddleAreaSearchResult indicates the response from the HotPepper Gourmet API for middle area search
type MiddleAreaSearchResult struct {
	Meta
	MiddleArea []MiddleArea `json:"middle_area"`
}

type LargeAreaTiny struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// MiddleArea indicates the middle area information from the HotPepper Gourmet API
type MiddleArea struct {
	Code             string        `json:"code"`
	Name             string        `json:"name"`
	LargeArea        LargeAreaTiny `json:"large_area"`
	ServiceArea      ServiceArea   `json:"service_area"`
	LargeServiceArea LargeServiceArea
}

// SmallAreaSearchParam indicates the small area search parameters for the HotPepper Gourmet API
type SmallAreaSearchParam struct {
	SmallArea  []string `url:"small_area"`
	MiddleArea []string `url:"middle_area,omitempty"`
	Keyword    []string `url:"keyword,omitempty"`
	Start      *uint    `url:"start,omitempty"`
	Count      *uint8   `url:"count,omitempty"`
}

// SmallAreaSearchResult indicates the response from the HotPepper Gourmet API for small area search
type SmallAreaSearchResult struct {
	Meta
	SmallArea []SmallArea `json:"small_area"`
}

type MiddleAreaTiny struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// SmallArea indicates the small area information from the HotPepper Gourmet API
type SmallArea struct {
	Code             string           `json:"code"`
	Name             string           `json:"name"`
	MiddleArea       MiddleAreaTiny   `json:"middle_area"`
	LargeArea        LargeAreaTiny    `json:"large_area"`
	ServiceArea      ServiceArea      `json:"service_area"`
	LargeServiceArea LargeServiceArea `json:"large_service_area"`
}

// GenreSearchParam indicates the genre search parameters for the HotPepper Gourmet API
type GenreSearchParam struct {
	Code    []string `url:"code"`
	Keyword []string `url:"keyword,omitempty"`
}

// GenreSearchResult indicates the response from the HotPepper Gourmet API for genre search
type GenreSearchResult struct {
	Meta
	Genre []Genre `json:"genre"`
}

// Genre indicates the genre information from the HotPepper Gourmet API
type Genre struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// DinnerBudgetMasterSearchParam indicates the dinner budget master search parameters for the HotPepper Gourmet API
type DinnerBudgetMasterSearchParam struct {
	Start uint
	Count uint
}

// LargeServiceAreaMasterSearchParam indicates the large service area master search parameters for the HotPepper Gourmet API
type LargeServiceAreaMasterSearchParam struct {
	Start uint
	Count uint
}

// ServiceAreaMasterSearchParam indicates the service area master search parameters for the HotPepper Gourmet API
type ServiceAreaMasterSearchParam struct {
	Start uint
	Count uint
}

// CreditCardMasterSearchParam indicates the credit card master search parameters for the HotPepper Gourmet API
type CreditCardMasterSearchParam struct {
	Start uint
	Count uint
}
