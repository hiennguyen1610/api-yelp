package model

type Address struct {
	Name string		`json:"name"`
	Type string		`json:"type"`
	Location string  `json:"location"`
	Desc string `json:"desc"`
}