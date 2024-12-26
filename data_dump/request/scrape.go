package request

type Scraper interface {
	// Scrape scrapes data returning some data in a run format
	Scrape() (interface{}, error)
}
