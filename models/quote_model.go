package models

type (
	QuoteRequest struct {
		Profile        int    `json:"profile"`
		SourceCurrency string `json:"sourceCurrency"`
		TargetCurrency string `json:"targetCurrency"`
		SourceAmount   int    `json:"sourceAmount"`
	}

	QuoteResponse struct {
		Id             string  `json:"id"`
		SourceCurrency string  `json:"sourceCurrency"`
		TargetCurrency string  `json:"targetCurrency"`
		SourceAmount   float32 `json:"sourceAmount"`
		PayOut         string  `json:"payOut"`
		Rate           float32 `json:"rate"`
		CreatedTime    string  `json:"createdTime"`
		RateType       string  `json:"rateType"`
		Status         string  `json:"status"`
	}
)
