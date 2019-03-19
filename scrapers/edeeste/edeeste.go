package edeeste

import (
	"cortes-programados-api/models"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ledongthuc/pdf"
)

const basePath = "https://edeeste.com.do/site/wp-admin/admin-ajax.php?juwpfisadmin=false&action=wpfd&task=files.display&view=files&id=42&rootcat=30&page=null"
const tempFileName = "./edeeste_outages.pdf"

func getFileURL() (string, error) {

	var url string

	res, err := http.Get(basePath)
	if err != nil {
		return url, fmt.Errorf("Could not request EDEEste's site: %v", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return url, fmt.Errorf("Could not read response: %v", err)
	}

	stringBody := string(body)

	linkStart := strings.Index(stringBody, "/edeeste.com.do")
	linkEnd := strings.Index(stringBody, ".pdf")
	url = fmt.Sprintf("https:/%s.pdf", strings.Replace(stringBody[linkStart:linkEnd], "\\", "", -1))

	return url, nil
}

func downloadFile() error {

	url, err := getFileURL()
	if err != nil {
		return err
	}

	fmt.Println("The URL", url)

	file, err := os.Create(tempFileName)
	if err != nil {
		return err
	}

	defer file.Close()

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Error status while getting the file: %s on %s", res.StatusCode, url)
	}

	// Write the response body to the local file
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	return nil
}

func ReadOutageAnouncement() ([]*models.Outage, error) {

	res := make([]*models.Outage, 0)

	err := downloadFile()
	if err != nil {
		fmt.Println("Error downloading file", err)
		log.Fatal(err)
		return nil, err
	}

	f, r, err := pdf.Open(tempFileName)
	// remember close file
	defer f.Close()
	if err != nil {
		return nil, err
	}

	fmt.Println("R?", r)

	return res, nil
}
