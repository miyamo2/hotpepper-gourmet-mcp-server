//go:build mcpgen

package main

import (
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

	// Create output file
	f, err := os.Create(filepath.Join(outDir, "mcp.gen.go"))
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer f.Close()

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
				Description: "グルメを検索します",
				InputSchema: struct {
					ID               []string `json:"id,omitempty"  jsonschema:"title=お店ID,description=お店に割り当てられた番号で検索します。20個まで指定可。,example=J999999999"`
					Name             *string  `json:"name,omitempty" jsonschema:"title=掲載店名,description=お店の名前で検索(部分一致)します。"`
					NameKana         *string  `json:"name_kana,omitempty" jsonschema:"title=掲載店名かな,description=お店の読みかなで検索(部分一致)します。"`
					NameAny          *string  `json:"name_any,omitempty" jsonschema:"title=掲載店名ORかな,description=お店の名前または読みかな両方をOR検索(部分一致)します。"`
					Tel              *string  `json:"tel,omitempty" jsonschema:"title=電話番号,description=お店の電話番号で検索します。半角数字(ハイフンなし),pattern=^\\d{9,11}$"`
					Address          *string  `json:"address,omitempty"       jsonschema:"title=住所,description=お店の住所で検索(部分一致)します。"`
					Keyword          []string `json:"keyword,omitempty" jsonschema:"title=キーワード,description=店名かな、店名、住所、駅名、お店ジャンルキャッチ、キャッチのフリーワード検索(部分一致)が可能です。複数指定可能。複数指定された場合、AND条件とみなします。"`
					LargeServiceArea *string  `json:"large_service_area,omitempty" jsonschema:"title=大サービスエリアコード,description=エリアに割り当てられたコード番号で検索します。指定できるコード番号は大サービスエリアマスタを参照。"`
					ServiceArea      []string `json:"service_area,omitempty" jsonschema:"title=サービスエリアコード,description=3個まで指定可。指定できるコード番号はサービスエリアマスタを参照。"`
					LargeArea        []string `json:"large_area,omitempty"         jsonschema:"title=大エリアコード,description=3個まで指定可。指定できるコードについては大エリアマスタを参照。"`
					MiddleArea       []string `json:"middle_area,omitempty"        jsonschema:"title=中エリアコード,description=5個まで指定可。指定できるコードについてはエリアマスタを参照。"`
					SmallArea        []string `json:"small_area,omitempty"         jsonschema:"title=小エリアコード,description=5個まで指定可。指定できるコードについては小エリアマスタを参照。"`
					Latitude         *float64 `json:"lat,omitempty"   jsonschema:"title=緯度,description=ある地点からの範囲内のお店の検索を行う場合の緯度です。"`
					Longitude        *float64 `json:"lng,omitempty"   jsonschema:"title=経度,description=ある地点からの範囲内のお店の検索を行う場合の経度です。"`
					Range            *uint8   `json:"range,omitempty" jsonschema:"title=検索範囲,description=ある地点からの範囲内のお店の検索を行う場合の範囲を5段階で指定できます。たとえば300m以内の検索ならrange=1を指定します 1=300m 2=500m 3=1000m (初期値)4=2000m 5=3000m,example=1" `
					Datum            *string  `json:"datum,omitempty" jsonschema:"title=測地系,description=緯度・経度の測地系を指定できます。world: 世界測地系、tokyo: 旧日本測地系。初期値は world。"`
					Genre            []string `json:"genre,omitempty"          jsonschema:"title=お店ジャンルコード,description=お店のジャンル(サブジャンル含む)で絞込むことができます。指定できるコードについてはジャンルマスタ参照"`
					Budget           []string `json:"budget,omitempty"         jsonschema:"title=検索用ディナー予算コード,description=ディナー予算で絞り込むことができます。指定できるコードについてはディナー予算マスタ参照"`
					PartyCapacity    *uint    `json:"party_capacity,omitempty" jsonschema:"title=宴会収容人数,description=宴会収容人数で絞り込むことができます。指定数より大きな収容人数のお店を検索します"`
					KtaiCoupon       bool     `json:"ktai_coupon,omitempty"  jsonschema:"title=携帯クーポン掲載,description=携帯クーポンの有無で絞り込み条件を指定します。true=絞り込む false=絞り込まない"`
					Wifi             bool     `json:"wifi,omitempty"         jsonschema:"title=WiFi 有無,description=WiFi 経由によるインターネット利用が可能なお店を絞り込みます。true=絞り込む false=絞り込まない"`
					Wedding          bool     `json:"wedding,omitempty"      jsonschema:"title=ウェディング二次会等,description=ウェディング・二次会等のお問い合わせが可能なお店を絞り込みます。true=絞り込む false=絞り込まない"`
					Course           bool     `json:"course,omitempty"       jsonschema:"title=コースあり,description=「コースあり」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					FreeDrink        bool     `json:"free_drink,omitempty"   jsonschema:"title=飲み放題,description=「飲み放題」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					FreeFood         bool     `json:"free_food,omitempty"    jsonschema:"title=食べ放題,description=「食べ放題」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					PrivateRoom      bool     `json:"private_room,omitempty" jsonschema:"title=個室あり,description=「個室あり」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Horigotatsu      bool     `json:"horigotatsu,omitempty"  jsonschema:"title=掘りごたつあり,description=「掘りごたつあり」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Tatami           bool     `json:"tatami,omitempty"       jsonschema:"title=座敷あり,description=「座敷あり」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Cocktail         bool     `json:"cocktail,omitempty"     jsonschema:"title=カクテル充実,description=「カクテル充実」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Shochu           bool     `json:"shochu,omitempty"       jsonschema:"title=焼酎充実,description=「焼酎充実」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Sake             bool     `json:"sake,omitempty"         jsonschema:"title=日本酒充実,description=「日本酒充実」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Wine             bool     `json:"wine,omitempty"         jsonschema:"title=ワイン充実,description=「ワイン充実」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Card             bool     `json:"card,omitempty"         jsonschema:"title=カード可,description=「カード可」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					NonSmoking       bool     `json:"non_smoking,omitempty"  jsonschema:"title=禁煙席,description=「禁煙席」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Charter          bool     `json:"charter,omitempty"      jsonschema:"title=貸切,description=「貸切可」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Ktai             bool     `json:"ktai,omitempty"         jsonschema:"title=携帯電話OK,description=「携帯電話OK」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Parking          bool     `json:"parking,omitempty"      jsonschema:"title=駐車場あり,description=「駐車場あり」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					BarrierFree      bool     `json:"barrier_free,omitempty" jsonschema:"title=バリアフリー,description=「バリアフリー」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Sommelier        bool     `json:"sommelier,omitempty"    jsonschema:"title=ソムリエがいる,description=「ソムリエがいる」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					NightView        bool     `json:"night_view,omitempty"   jsonschema:"title=夜景がキレイ,description=「夜景がキレイ」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					OpenAir          bool     `json:"open_air,omitempty"     jsonschema:"title=オープンエア,description=「オープンエア」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Show             bool     `json:"show,omitempty"         jsonschema:"title=ライブ・ショーあり,description=「ライブ・ショーあり」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Equipment        bool     `json:"equipment,omitempty"    jsonschema:"title=エンタメ設備,description=「エンタメ設備」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Karaoke          bool     `json:"karaoke,omitempty"      jsonschema:"title=カラオケあり,description=「カラオケあり」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Band             bool     `json:"band,omitempty"         jsonschema:"title=バンド演奏可,description=「バンド演奏可」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					TV               bool     `json:"tv,omitempty"           jsonschema:"title=TV・プロジェクター,description=「TV・プロジェクター」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Lunch            bool     `json:"lunch,omitempty"        jsonschema:"title=ランチあり,description=「ランチあり」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Midnight         bool     `json:"midnight,omitempty"     jsonschema:"title=23時以降も営業,description=「23時以降も営業」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					MidnightMeal     bool     `json:"midnight_meal,omitempty"jsonschema:"title=23時以降食事OK,description=「23時以降食事OK」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					English          bool     `json:"english,omitempty"      jsonschema:"title=英語メニューあり,description=「英語メニューあり」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Pet              bool     `json:"pet,omitempty"          jsonschema:"title=ペット可,description=「ペット可」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					Child            bool     `json:"child,omitempty"        jsonschema:"title=お子様連れOK,description=「お子様連れOK」という条件で絞り込むかどうかを指定します。true=絞り込む false=絞り込まない"`
					CreditCard       []string `json:"credit_card,omitempty" jsonschema:"title=クレジットカード,description=クレジットカードの種別ごとに絞り込むことができます。指定できるコードについてはクレジットカードマスタ参照。(2008/02/08追加)"`
					Order            *uint8   `json:"order,omitempty"       jsonschema:"title=ソート順,description=検索結果の並び順を指定します。おススメ順は定期的に更新されます。 ※ 位置検索の場合、「4:オススメ順」以外は指定に関係なく、強制的に距離順でソートされます。 1:店名かな順 2:ジャンルコード順 3:小エリアコード順 4:おススメ順 初期値はおススメ順。位置から検索を行った場合は距離順"`
					Start            *uint    `json:"start,omitempty"       jsonschema:"title=検索の開始位置,description=検索結果の何件目から出力するかを指定します。初期値は1。"`
					Count            *uint8   `json:"count,omitempty"       jsonschema:"title=1ページあたりの取得数,description=検索結果の最大出力データ数を指定します。初期値は10。最小1、最大100。"`
				}{},
			},
			{
				Name:        "shop_search",
				Description: "お店を検索します",
				InputSchema: struct {
					Keyword []string `json:"keyword,omitempty" jsonschema:"title=キーワード,description=お店の名前・読みがな・住所で検索（部分一致）します。複数指定可能"`
					Tel     *string  `json:"tel,omitempty" jsonschema:"title=電話番号,description=お店の電話番号で検索します。半角数字(ハイフンなし),pattern=^\\d{9,11}$"`
					Start   *uint    `json:"start,omitempty" jsonschema:"title=検索の開始位置,description=検索結果の何件目から出力するかを指定します。初期値は1。"`
					Count   *uint8   `json:"count,omitempty" jsonschema:"title=1ページあたりの取得数,description=検索結果の最大出力データ数を指定します。初期値は30。最小1、最大30。"`
				}{},
			},
			{
				Name:        "large_area_search",
				Description: "大エリアマスタを検索します",
				InputSchema: struct {
					LargeArea []string `json:"large_area,omitempty" jsonschema:"title=大エリアコード,description=大エリアコードで検索(完全一致)します。（3個まで指定可、4個以上指定すると4個目以降無視します）。"`
					Keyword   []string `json:"keyword,omitempty" jsonschema:"title=キーワード,description=大エリア名で検索(部分一致)します。"`
				}{},
			},
			{
				Name:        "middle_area_search",
				Description: "中エリアマスタを検索します",
				InputSchema: struct {
					MiddleArea []string `json:"middle_area,omitempty" jsonschema:"title=中エリアコード,description=中エリアコードで検索(完全一致)します。（5個まで指定可、6個以上指定すると6個目以降無視します）。,example=Y005"`
					LargeArea  []string `json:"large_area,omitempty" jsonschema:"title=大エリアコード,description=大エリアコードで検索(完全一致)します。（3個まで指定可、4個以上指定すると4個目以降無視します）。,example=Z011"`
					Keyword    []string `json:"keyword,omitempty" jsonschema:"title=キーワード,description=中エリア名で検索(部分一致)します。,example=飯田橋"`
					Start      *uint    `json:"start,omitempty" jsonschema:"title=検索の開始位置,description=検索結果の何件目から出力するかを指定します。初期値:1"`
					Count      *uint8   `json:"count,omitempty" jsonschema:"title=1ページあたりの取得数,description=検索結果の最大出力データ数を指定します。初期値：未設定。最小1、最大：未設定（すべて取得）"`
				}{},
			},
			{
				Name:        "small_area_search",
				Description: "小エリアマスタを検索します",
				InputSchema: struct {
					SmallArea  []string `json:"small_area,omitempty" jsonschema:"title=小エリアコード,description=小エリアコードで検索(完全一致)します。（5個まで指定可、6個以上指定すると6個目以降無視します）。,example=X005"`
					MiddleArea []string `json:"middle_area,omitempty" jsonschema:"title=中エリアコード,description=中エリアコードで検索(完全一致)します。（5個まで指定可、6個以上指定すると6個目以降無視します）。,example=Y005"`
					Keyword    []string `json:"keyword,omitempty" jsonschema:"title=キーワード,description=小エリア名で検索(部分一致)します。,example=銀座"`
					Start      *uint    `json:"start,omitempty" jsonschema:"title=検索の開始位置,description=検索結果の何件目から出力するかを指定します。初期値:1"`
					Count      *uint8   `json:"count,omitempty" jsonschema:"title=1ページあたりの取得数,description=検索結果の最大出力データ数を指定します。初期値：未設定。最小1、最大：未設定（すべて取得）"`
				}{},
			},
			{
				Name:        "genre_search",
				Description: "ジャンルマスタを検索します",
				InputSchema: struct {
					Code    []string `json:"code,omitempty" jsonschema:"title=ジャンルコード,description=ジャンルコードで検索(完全一致)します。（２個まで指定可、3個以上指定すると3個目以降無視します）。,example=G002"`
					Keyword []string `json:"keyword,omitempty" jsonschema:"title=キーワード,description=ジャンル名で検索(部分一致)します。,example=バー"`
				}{},
			},
			{
				Name:        "dinner_budget_master_search",
				Description: "ディナー予算マスタを検索します",
				InputSchema: struct {
					Start *uint `json:"start,omitempty" jsonschema:"title=検索の開始位置,description=検索結果の何件目から出力するかを指定します。初期値:1"`
					Count *uint `json:"count,omitempty" jsonschema:"title=1ページあたりの取得数,description=検索結果の最大出力データ数を指定します。初期値：未設定。最小: 1、最大：未設定（すべて取得）、推奨: 未設定（すべて取得）"`
				}{},
			},
			{
				Name:        "large_service_area_master_search",
				Description: "大サービスエリアマスタを検索します",
				InputSchema: struct {
					Start *uint `json:"start,omitempty" jsonschema:"title=検索の開始位置,description=検索結果の何件目から出力するかを指定します。初期値:1"`
					Count *uint `json:"count,omitempty" jsonschema:"title=1ページあたりの取得数,description=検索結果の最大出力データ数を指定します。初期値：未設定。最小: 1、最大：未設定（すべて取得）、推奨: 未設定（すべて取得）"`
				}{},
			},
			{
				Name:        "service_area_master_search",
				Description: "サービスエリアマスタを検索します",
				InputSchema: struct {
					Start *uint `json:"start,omitempty" jsonschema:"title=検索の開始位置,description=検索結果の何件目から出力するかを指定します。初期値:1"`
					Count *uint `json:"count,omitempty" jsonschema:"title=1ページあたりの取得数,description=検索結果の最大出力データ数を指定します。初期値：未設定。最小: 1、最大：未設定（すべて取得）、推奨: 未設定（すべて取得）"`
				}{},
			},
			{
				Name:        "credit_card_master_search",
				Description: "クレジットカードマスタを検索します",
				InputSchema: struct {
					Start *uint `json:"start,omitempty" jsonschema:"title=検索の開始位置,description=検索結果の何件目から出力するかを指定します。初期値:1"`
					Count *uint `json:"count,omitempty" jsonschema:"title=1ページあたりの取得数,description=検索結果の最大出力データ数を指定します。初期値：未設定。最小: 1、最大：未設定（すべて取得）、推奨: 未設定（すべて取得）"`
				}{},
			},
		},
		ResourceTemplates: []codegen.ResourceTemplate{
			{
				Name:        "dinner_budget_master",
				Description: "ディナー予算マスタ",
				URITemplate: "https://webservice.recruit.co.jp/hotpepper/budget/v1/",
				MimeType:    "application/json",
			},
			{
				Name:        "large_service_area_master",
				Description: "大サービスエリアマスタ",
				URITemplate: "https://webservice.recruit.co.jp/hotpepper/large_service_area/v1/",
				MimeType:    "application/json",
			},
			{
				Name:        "service_area_master",
				Description: "サービスエリアマスタ",
				URITemplate: "https://webservice.recruit.co.jp/hotpepper/service_area/v1/",
				MimeType:    "application/json",
			},
			{
				Name:        "credit_card_master",
				Description: "クレジットカードマスタ",
				URITemplate: "https://webservice.recruit.co.jp/hotpepper/credit_card/v1/",
				MimeType:    "application/json",
			},
		},
	}

	// Generate code
	if err := codegen.Generate(f, def, "mcpgen"); err != nil {
		log.Fatalf("failed to generate code: %v", err)
	}
}
