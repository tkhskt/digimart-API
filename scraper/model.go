package scraper

import "time"

type Instrument struct {
	Category     string
	Name         string
	Price        int
	Condition    string
	Status       string
	URL          string
	Image        string
	RegisterDate time.Time
}
