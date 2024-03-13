package model

type TokenInfo struct {
	Address  string
	Symbol   string
	Decimals int32
}

type ContractApi struct {
	App    string `json:"app"`
	Market string `json:"market"`
	Vault  string `json:"vault"`
	Farm   string `json:"farm"`
}

type CommonChartApi struct {
	Price  []PriceChartItem  `json:"price"`
	Volume []VolumeChartItem `json:"volume"`
}

type PriceChartItem struct {
	Datetime string  `json:"time"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
}

type VolumeChartItem struct {
	Datetime  string  `json:"time"`
	Value     float64 `json:"value"`
	OrderType string  `json:"type"`
}

type BarChartApi struct {
	Datetime string  `json:"time"`
	Value    float64 `json:"value"`
}

type OverviewApi struct {
	TotalValue  float64 `json:"total_tvl"`
	TotalVolume float64 `json:"total_volume"`
}
