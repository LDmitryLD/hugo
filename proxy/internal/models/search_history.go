package models

type SearchHistory struct {
	ID    int
	Query string
}

type AddressSQL struct {
	Lat float64
	Lng float64
}

type HistorySearchAddress struct {
	ID              int
	SearchHistoryID int
	AddressID       int
}
