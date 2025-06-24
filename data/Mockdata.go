package data

import "sandbox-invest/models"

var Prices = map[string]models.Asset{
    "BBCA":   {Code: "BBCA", Name: "Bank BCA", Type: models.Saham, Price: 9200},
    "RDPT":   {Code: "RDPT", Name: "RD Pendapatan Tetap", Type: models.Reksadana, Price: 1340},
    "ORI023": {Code: "ORI023", Name: "Obligasi ORI023", Type: models.Obligasi, Price: 100000},
}

var Portfolios = map[string]*models.Portfolio{}
