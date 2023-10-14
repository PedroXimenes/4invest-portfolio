package assets

type Asset struct {
	PortfolioID     int64   `json:"portfolio_id"`
	Type            string  `json:"type"`
	Symbol          string  `json:"symbol"`
	Quantity        int64   `json:"quantity"`
	MaxValueLimit   float64 `json:"max_value_limit"`
	IdealPercentage float64 `json:"ideal_percentage"`
}
