package edesur

import (
	"cortes-programados-api/lib"
	"cortes-programados-api/models"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const basePath = "http://www.edesur.com.do/servicios/mantenimientos/"

func GetOutageAnouncement() ([]*models.Outage, error) {

	// First thing is to query the page of outages
	page, err := goquery.NewDocument(basePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	outages := make([]*models.Outage, 0)

	// Look into the page and get the List of outages
	page.Find("ul.es-site-mantenimientos.accordion li.text-center.accordion-item").Each(
		func(index int, sel *goquery.Selection) {

			// Get the Date for each of the outages
			dateStr := lib.ParseLocalTimeString(sel.Find("a span").First().Text())

			// Query the page to get the full week
			week := sel.Find("div ul li.accordion-item")
			outs := make([]*models.Outage, 0)

			// For each day of the week lets get each of the outages
			week.Each(
				func(i int, s *goquery.Selection) {
					provinces := s.Find("a.accordion-title>h4")

					// Outages are sorted by Day -> Province -> Time Frame
					// And so I'm doing the exact same thing to gather all needed info
					provinces.Each(
						func(j int, q *goquery.Selection) {

							name := strings.Trim(q.Text(), "\n ")

							if name == "" {
								return
							}

							timeFrame := s.Find("div.accordion-content>h5")

							timeFrame.Each(
								func(i int, x *goquery.Selection) {

									if strings.Contains(strings.ToLower(x.Text()), "no trabajos de mantenimiento") {
										return
									}

									frameStr := strings.TrimLeft(x.Text(), "Sectores impactados entre ")

									frameSlice := strings.Split(frameStr, " a ")
									if len(frameSlice) != 2 {
										return
									}

									zones := x.Next()

									if strings.Contains(strings.ToLower(zones.Text()), "trabajos de mantenimiento programados") {
										return
									}

									zoneStr := zones.Text()

									if zoneStr == "" {
										return
									}

									outage := &models.Outage{
										Province:      name,
										Date:          dateStr,
										StartTime:     frameSlice[0],
										EndTime:       strings.TrimRight(frameSlice[1], ":"),
										Company:       "EDESur",
										AffectedZones: strings.Split(zoneStr, ","),
									}

									outs = append(outs, outage)
								},
							)
						},
					)

				},
			)

			outages = append(outages, outs...)
		},
	)

	return outages, nil
}

// TODO -> Map Zones of this scraper to Sectors inside the provinces
// TODO -> Map Zones of this craper to Circuits.
