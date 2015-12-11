package entities

import (
	"labix.org/v2/mgo/bson"
)

type Item struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Category    string        `json:"category"`
	Name        string        `json:"name"`
	UPC         string        `json:"upc"`
	Kcal        string        `json:"kcal"`
	Fat         string        `json:"fat"`
	Satfat      string        `json:"satfat"`
	Carbs       string        `json:"carbs"`
	Sugars      string        `json:"sugars"`
	Fibre       string        `json:"fibre"`
	Protein     string        `json:"protein"`
	Salt        string        `json:"salt"`
	PPKcal      string        `json:"ppkcal"`
	PPFat       string        `json:"ppfat"`
	PPSaturates string        `json:"ppsatfat"`
	PPCarbs     string        `json:"ppcarbs"`
	PPSugars    string        `json:"ppsugars"`
	PPProtein   string        `json:"ppprotein"`
	PPSalt      string        `json:"ppsalt"`
	ENName      string        `json:"enrichednutriname"`
	EN          string        `json:"enrichednutri"`
	PPEN        string        `json:"ppenrichednutri"`
}

type Items []Item
