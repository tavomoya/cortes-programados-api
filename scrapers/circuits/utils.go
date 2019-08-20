package circuits

import (
	"cortes-programados-api/models"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getCircuits(basePath string, provinces []string) ([]*models.Circuit, error) {
	circuitsList := make([]*models.Circuit, 0)

	for _, x := range provinces {

		url := fmt.Sprintf(basePath, x)
		fmt.Println("For URL:", url)
		page, err := goquery.NewDocument(url)
		if err != nil {
			log.Println("Error getting to the page =>", err)
			log.Fatal(err)
			return nil, err
		}

		getProvinceCircuits(page, &circuitsList)
	}

	return circuitsList, nil
}

func getProvinceCircuits(page *goquery.Document, circuitList *[]*models.Circuit) {
	if circuitList == nil {
		return
	}

	page.Find("tbody>tr").Each(
		func(index int, sel *goquery.Selection) {
			circuit := &models.Circuit{}
			sel.Find("td").Each(
				func(i int, s *goquery.Selection) {
					switch i {
					case 0:
					case 1:
						circuit.Company = strings.TrimSpace(s.Text())
					case 2:
						circuit.Province = strings.TrimSpace(s.Text())
					case 3:
						circuit.Name = s.Find("b>strong>a").First().Text()
					case 4:
						circuit.CircuitType = strings.TrimSpace(s.Text())
					}
				},
			)

			// Add Affected Zones
			divId := fmt.Sprintf("Modal-%s", circuit.Name)
			selectorQuery := fmt.Sprintf("#%s div.ZonaInf", divId)
			circuit.InfluenceZones = strings.Split(strings.TrimSpace(page.Find(selectorQuery).First().Text()), ",")

			*circuitList = append(*circuitList, circuit)
		},
	)

	path, _ := page.Find("li.PagedList-skipToNext>a").First().Attr("href")
	if path != "" {

		url := fmt.Sprintf("https://cdeee.gob.do%s", path)
		p, err := goquery.NewDocument(url)
		if err != nil {
			log.Println("Error getting new page =>", err)
			log.Fatal(err)
			return
		}

		getProvinceCircuits(p, circuitList)
	}

	return
}

func edeNorteParams() (string, []string) {
	provinces := []string{"7", "10", "14", "15", "16", "17", "21", "22", "23", "27", "28", "29", "30"}
	basePath := "https://cdeee.gob.do/Circuitos/Circuitos?distribuidoraSearch=1&provinciaSearch=%s"

	return basePath, provinces
}

func edeSurParams() (string, []string) {
	provinces := []string{"1", "3", "4", "5", "9", "12", "19", "20", "24", "25", "31", "32"}
	basePath := "https://cdeee.gob.do/Circuitos/Circuitos?distribuidoraSearch=2&provinciaSearch=%s"

	return basePath, provinces
}

func edeEsteParams() (string, []string) {
	provinces := []string{"1", "2", "8", "11", "13", "18", "26", "32"}
	basePath := "https://cdeee.gob.do/Circuitos/Circuitos?distribuidoraSearch=3&provinciaSearch=%s"

	return basePath, provinces
}
