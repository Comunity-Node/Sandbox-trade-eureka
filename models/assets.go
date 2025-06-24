package models

type AssetType string

const (
    Saham     AssetType = "saham"
    Reksadana AssetType = "reksadana"
    Obligasi  AssetType = "obligasi"
)

type Asset struct {
    Code  string    `json:"code"`
    Name  string    `json:"name"`
    Type  AssetType `json:"type"`
    Price float64   `json:"price"`
}

type Transaction struct {
    UserID string  `json:"user_id"`
    Code   string  `json:"code"`
    Amount float64 `json:"amount"`
    Action string  `json:"action"`
}

type Portfolio struct {
    Holdings map[string]float64 `json:"holdings"`
}
