package main

import (
	"bytes"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"io"
	"net/http"
	"os"
	"testing"
)

var (
	resStatus string
	resBody   []byte
)

// iSendRequestToWhereBodyIsJson realiza la solicitud a la API
func iSendRequestToWhereBodyIsJson(method, route, bodyreq string) error {
	url := "http://localhost:8080" + route

	fmt.Printf("Enviando solicitud %s a %s con cuerpo %s\n", method, route, bodyreq)
	var body io.Reader
	if bodyreq != "null" {
		body = bytes.NewBuffer([]byte(bodyreq))
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	resStatus = resp.Status
	resBody, _ = io.ReadAll(resp.Body)

	return nil
}

func theResponseStatusCodeIs(expectedStatus string) error {
	fmt.Printf("Esperando código de estado %s y se obtuvo %s\n", expectedStatus, resStatus) // Mensaje personalizado

	if resStatus != expectedStatus {
		return fmt.Errorf("expected status code %s but got %s", expectedStatus, resStatus)
	}
	return nil
}

func FeatureContext(s *godog.ScenarioContext) {
	s.Step(`^I send "([^"]*)" request to "([^"]*)" where body is json "([^"]*)"$`, iSendRequestToWhereBodyIsJson)
	s.Step(`^the response status code is "([^"]*)"$`, theResponseStatusCodeIs)

}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format: "progress",
		Paths:  []string{"features"},
		Output: colors.Colored(os.Stdout),
	}

	status := godog.TestSuite{
		Name:                "godogs",
		ScenarioInitializer: FeatureContext,
		Options:             &opts,
	}.Run()

	os.Exit(status)
}
