package services

import (
    "encoding/json"
    "fmt"
    "github.com/go-resty/resty/v2"
)

type YahooQuoteResult struct {
    Symbol              string  `json:"symbol"`
    RegularMarketPrice  float64 `json:"regularMarketPrice"`
    ShortName           string  `json:"shortName"`
    RegularMarketChange float64 `json:"regularMarketChange"`
}

type YahooQuoteResponse struct {
    QuoteResponse struct {
        Result []YahooQuoteResult `json:"result"`
    } `json:"quoteResponse"`
}

func GetYahooPrice(symbol string) (*YahooQuoteResult, error) {
    url := fmt.Sprintf("https://query1.finance.yahoo.com/v7/finance/quote?symbols=%s.JK", symbol)
    client := resty.New()
    resp, err := client.R().Get(url)
    if err != nil {
        return nil, err
    }

    var result YahooQuoteResponse
    if err := json.Unmarshal(resp.Body(), &result); err != nil {
        return nil, err
    }

    if len(result.QuoteResponse.Result) == 0 {
        return nil, fmt.Errorf("symbol not found")
    }

    return &result.QuoteResponse.Result[0], nil
}
