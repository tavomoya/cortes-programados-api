package cron

import (
	"cortes-programados-api/models"
	"cortes-programados-api/scrapers/edenorte"
	"cortes-programados-api/scrapers/edesur"
	"log"
)

func getOutagesScrapeData() ([]*models.Outage, error) {

	norte, err := edenorte.ReadOutageAnouncement()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	sur, err := edesur.GetOutageAnouncement()
	if err != nil {
		return nil, err
	}

	return append(norte, sur...), nil
}
