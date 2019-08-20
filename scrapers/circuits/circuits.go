package circuits

import (
	"cortes-programados-api/lib"
	"log"
)

type CircuitsScraper struct {
	db *lib.DBLib
}

func New(db *lib.DBLib) *CircuitsScraper {
	return &CircuitsScraper{
		db: db,
	}
}

func (c *CircuitsScraper) ScrapeAllCircuits() error {

	basePath, provinces := edeNorteParams()
	norte, err := getCircuits(basePath, provinces)
	if err != nil {
		log.Println("there was an error getting EDENorte's circuits =>", err)
		return err
	}

	basePath, provinces = edeSurParams()
	sur, err := getCircuits(basePath, provinces)
	if err != nil {
		log.Println("there was an error getting EDENorte's circuits =>", err)
		return err
	}

	basePath, provinces = edeEsteParams()
	este, err := getCircuits(basePath, provinces)
	if err != nil {
		log.Println("there was an error getting EDENorte's circuits =>", err)
		return err
	}

	circuits := append(norte, sur...)
	circuits = append(circuits, este...)

	for _, o := range circuits {
		err := c.db.Insert(o)
		if err != nil {
			return err
		}
	}

	return nil
}
