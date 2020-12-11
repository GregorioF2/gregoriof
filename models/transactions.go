package models

type Transaction struct {
	Type string 						`json:"type"`
	Ammount float64					`json:"ammount"`
	Id int									`json:"id"`
	EffectiveDate string		`json:"effectiveDate"`
}