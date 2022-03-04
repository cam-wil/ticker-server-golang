package main

type CompleteData struct {
	Worth    float64     `json:"worth"`
	Errors   int         `json:"errors"`
	MaxWorth float64     `json:"maxWorth"`
	MinWorth float64     `json:"minWorth"`
	Stocks   interface{} `json:"stocks"`
}

type SymbolData struct {
	DayChange string  `json:"dayChange"`
	Name      string  `json:"name"`
	Open      float64 `json:"open"`
	PrevClose float64 `json:"prevClose"`
	Price     float64 `json:"price"`
	Symbol    string  `json:"symbol"`
	Time      int32   `json:"time"`
	TodayHigh float64 `json:"todayHigh"`
	TodayLow  float64 `json:"todayLow"`
}
