package api

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/wannabehero/go-amadeus/types"
)

var (
	amadeusClient = &http.Client{}
)

func saveLogsFile(name string, content []byte) {
	os.Mkdir("logs", 0755)

	filename := fmt.Sprintf("%s_%s", time.Now().UTC().Format(time.RFC3339), name)

	os.WriteFile(fmt.Sprintf("logs/%s", filename), content, 0755)
}

func getMethodName(action string) string {
	components := strings.Split(action, "/")
	return components[len(components)-1]
}

type AmadeusAPIClient struct {
	config types.AmadeusConfig
}

func NewClient(config types.AmadeusConfig) *AmadeusAPIClient {
	return &AmadeusAPIClient{config}
}

func (client *AmadeusAPIClient) SendRequest(action string, envelope types.Envelope) ([]byte, error) {
	data, err := xml.MarshalIndent(envelope, "  ", "  ")
	if err != nil {
		return nil, err
	}

	go saveLogsFile(fmt.Sprintf("%s_RQ.xml", getMethodName(action)), data)

	req, _ := http.NewRequest(http.MethodPost, client.config.URL, bytes.NewReader(data))
	req.Header.Set("content-type", "text/xml;charset=UTF-8")
	req.Header.Set("soapaction", action)

	res, err := amadeusClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	go saveLogsFile(fmt.Sprintf("%s_RS.xml", getMethodName(action)), data)

	return data, nil
}
