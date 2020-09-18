package gofinance

type DataQuotation struct {
	By            string     `json:"by"`
	ValidKey      bool       `json:"valid_key"`
	Results       *Quotation `json:"results"`
	ExecutionTime float64    `json:"execution_time"`
	FromCache     bool       `json:"from_cache"`
}

type Quotation struct {
	Currencies       *Currencies `json:"currencies"`
	Stocks           *Stocks     `json:"stocks"`
	AvailableSources []string    `json:"available_sources"`
	BitCoin          *BitCoin    `json:"bitcoin"`
}

type Currencies struct {
	Source string `json:"source"`
	USD    *Currencie
	EUR    *Currencie
	GBP    *Currencie
	ARS    *Currencie
	BTC    *Currencie
}

type Currencie struct {
	Name      string  `json:"name"`
	Buy       float64 `json:"buy"`
	Sell      float64 `json:"sell"`
	Variation float64 `json:"variation"`
}

type Stocks struct {
	IBOVESPA *Stock
	NASDAQ   *Stock
	CAC      *Stock
	NIKKEI   *Stock
}

type Stock struct {
	Name      string  `json:"name"`
	Location  string  `json:"location"`
	Variation float64 `json:"variation"`
}

type BitCoin struct {
	BlockChainInfo *Exchange `json:"blockchain_info"`
	CoinBase       *Exchange `json:"coinbase"`
	FoxBit         *Exchange `json:"foxbit"`
	MercadoBitCoin *Exchange `json:"mercadobitcoin"`
	OmniTrade      *Exchange `json:"omnitrade"`
	Xdex           *Exchange `json:"xdex"`
}

type Exchange struct {
	Name      string   `json:"name"`
	Format    []string `json:"format"`
	Last      float64  `json:"last"`
	Buy       float64  `json:"buy"`
	Sell      float64  `json:"sell"`
	Variation float64  `json:"variation"`
}

type DataTaxe struct {
	By            string  `json:"by"`
	ValidKey      bool    `json:"valid_key"`
	Results       []*Taxe `json:"results"`
	ExecutionTime float64 `json:"execution_time"`
	FromCache     bool    `json:"from_cache"`
}

type Taxe struct {
	Date        string  `json:"date"`
	CDI         float64 `json:"cdi"`
	Selic       float64 `json:"selic"`
	DailyFactor float64 `json:"daily_factor"`
}
