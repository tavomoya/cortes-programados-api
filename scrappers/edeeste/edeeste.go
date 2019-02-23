package edeeste

import (
	"cortes-programados-api/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ledongthuc/pdf"
)

const basePath = "https://edeeste.com.do/site/wp-admin/admin-ajax.php?juwpfisadmin=false&action=wpfd&task=files.display&view=files&id=42&rootcat=30&page=null"
const tempFileName = "./edeeste_outages.pdf"

func getFileURL() (string, error) {

	var url string
	category := DataCategory{}

	res, err := http.Get(basePath)
	if err != nil {
		return url, fmt.Errorf("Could not request EDEEste's site: %v", err)
	}

	defer res.Body.Close()

	// by, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	return url, fmt.Errorf("Could not read response: %v", err)
	// }

	// byString := fmt.Sprintf(`%s`, string(by))

	// buf := new(bytes.Buffer)
	// enc := json.NewEncoder(buf)
	// enc.SetEscapeHTML(false)
	// err = enc.Encode(by)
	// if err != nil {
	// 	return url, fmt.Errorf("Could not encode response body: %v", err)
	// }

	// payload := fmt.Sprintf(`%s`, string(buf.String()))

	err = json.NewDecoder(res.Body).Decode(&category)
	if err != nil {
		return url, fmt.Errorf("Could not parse response body: %v", err)
	}

	// b, err := json.Marshal(buf.String())
	// if err != nil {
	// 	return url, fmt.Errorf("Could not marshal response body: %v", err)
	// }

	// fmt.Println(":) ", string(b))

	// err = json.Unmarshal(b, &category)
	// if err != nil {
	// 	return url, fmt.Errorf("Could not parse response body: %v", err)
	// }

	fmt.Println("Cat: ", category)

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
