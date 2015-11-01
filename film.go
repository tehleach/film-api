package main

//Film represents a film
type Film struct {
	Name       string `json:"Title"`
	Director   string
	Runtime    string
	Year       int `json:",string"`
	Rated      string
	Metascore  int     `json:",string"`
	IMDBRating float64 `json:"imdbRating,string"`
}
