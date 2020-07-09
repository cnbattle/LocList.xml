package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	Name        string
	Code        string
	PID         uint
	Country     string
	CountryCode string
	State       string
	StateCode   string
	City        string
	CityCode    string
	Region      string
	RegionCode  string
}
