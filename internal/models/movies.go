package models

type Movie struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Rating float64 `json:"rating"`
}
