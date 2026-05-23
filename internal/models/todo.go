package models

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type Bookings struct {
	ID     int     `json:"id"`
	Price  float64 `json:"price"`
	Booked bool    `json:"booked"`
}
