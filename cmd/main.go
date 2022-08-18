package main

import (
	"amadeus-xml/requests"
	"amadeus-xml/types"
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	config := types.LoadConfigFromFile(os.Getenv("CONFIG_FILE"))

	hotels := []string{
		"BKSEAAVF", "BKSEAAVB", "BKSEAAD7", "BKSEAACA", "BKSEAABU",
		"BKSEAABL", "BKSEAAB4", "BKSEAAAG", "BKSEAAAC",
	}

	// session := types.NewSession()

	envelope, action := requests.NewAvailabilityRequest(
		types.InfoSourceLeisure,
		"2022-09-10",
		"2022-09-11",
		"USD",
		"US",
		1,
		hotels,
		// session,
		nil,
		config,
	)

	data, err := xml.MarshalIndent(envelope, "  ", "  ")
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Println(string(data))
	}

	req, _ := http.NewRequest(http.MethodPost, config.URL, bytes.NewReader(data))
	req.Header.Set("content-type", "text/xml;charset=UTF-8")
	req.Header.Set("soapaction", action)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Request error: %v\n", err)
	} else {
		data, _ = ioutil.ReadAll(res.Body)
		log.Println(string(data))
	}
}
