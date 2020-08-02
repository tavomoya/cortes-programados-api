package edenorte

import (
	"cortes-programados-api/lib"
	"cortes-programados-api/models"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/PuerkitoBio/goquery"
)

const basePath = "https://edenorte.com.do/interrupciones-programadas/"
const tempFileName = "./edenorte_outage.xlsx"

func getLatestAnouncement() (string, error) {
	doc, err := goquery.NewDocument(basePath)
	if err != nil {
		log.Println(err)
		return "", err
	}

	lastPublication := doc.Find("article.programa-de-mantenimiento-de-redes").First()
	url, _ := lastPublication.Find("div a").First().Attr("href")

	return url, nil
}

func getFileURL() (string, error) {

	anouncement, err := getLatestAnouncement()
	if err != nil {
		log.Println(err)

		return "", err
	}

	doc, err := goquery.NewDocument(anouncement)
	if err != nil {
		log.Println(err)

		return "", err
	}

	url, _ := doc.Find("td>img[alt='xls']").Next().Attr("href")
	return url, nil
}

func downloadFile() error {

	url, err := getFileURL()
	if err != nil {
		log.Println(err)

		return err
	}

	// Create local file
	file, err := os.Create(tempFileName)
	if err != nil {
		log.Println(err)
		return err
	}

	defer file.Close()

	// Get the file from the URL
	res, err := http.Get(url)
	if err != nil {
		log.Println("Error getting file from URL =>", url)
		log.Println(err)
		return err
	}
	defer res.Body.Close()

	// Check server response
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Error status while getting the file: %s", res.Status)
	}

	// Write the response body to the local file
	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Println("Error writing response to local file =>", err)
		log.Println(err)
		return err
	}

	return nil
}

func buildOutageModel(column int, value string, outage *models.Outage) {

	if column < 0 || column > 7 {
		return
	}

	switch column {
	case 0:
		outage.Circuit = value
		break
	case 1:
		date, _ := time.Parse("01-02-06", value)
		outage.Date = date
		break
	case 2:
		outage.Province = value
		break
	case 3:
		outage.Sectors = strings.Split(value, ",")
		break
	case 4:
		floatVal, _ := strconv.ParseFloat(value, 64)
		hour := math.Round((floatVal*24)*100) / 100
		hourString := lib.GetTimeString(hour)
		outage.StartTime = hourString
		outage.StartTimeInt = hour
		break
	case 6:
		floatVal, _ := strconv.ParseFloat(value, 64)
		hour := math.Round((floatVal*24)*100) / 100
		hourString := lib.GetTimeString(hour)
		outage.EndTime = hourString
		outage.EndTimeInt = hour
		break
	case 7:
		fmt.Println("value: ", value)
		outage.AffectedZones = strings.Split(value, ",")
		break
	}

	return
}

func cleanupOutageList(outages []*models.Outage) {
	prev := &models.Outage{}
	zeroDate := time.Time{}
	for i, outage := range outages {
		if i == 0 {
			prev = outage
			continue
		}

		if outage.Date == zeroDate {
			outage.Date = prev.Date
		}

		if outage.Province == "" {
			outage.Province = prev.Province
		}

		prev = outage
	}

}

// ReadOutageAnouncement downloads the latest Outage Anouncement file from the EDENorte site
// and turns it into a models.Outage slice so it can be used later on.
func ReadOutageAnouncement() ([]*models.Outage, error) {
	res := make([]*models.Outage, 0)

	err := downloadFile()
	if err != nil {
		log.Println(err)
		log.Println(err)
		return nil, err
	}

	xlsx, err := excelize.OpenFile(tempFileName)
	if err != nil {
		log.Println(err)
		log.Println(err)
		return nil, err
	}

	rows := xlsx.GetRows("Publicación EDENORTE")
	for r, row := range rows {
		if r < 10 || row[0] == "" {
			continue
		} else {
			outage := &models.Outage{}
			outage.Company = "EDENorte"
			for i, cell := range row {
				if cell != "" {
					buildOutageModel(i, cell, outage)
				}
			}

			res = append(res, outage)
		}
	}

	cleanupOutageList(res)
	os.Remove(tempFileName)
	return res, nil
}
