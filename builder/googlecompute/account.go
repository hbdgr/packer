package googlecompute

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"os"
	"strings"
)

var ExporterScopes = []string{
	"https://www.googleapis.com/auth/compute",
	"https://www.googleapis.com/auth/devstorage.full_control",
	"https://www.googleapis.com/auth/userinfo.email"}

var DriverScopes = []string{"https://www.googleapis.com/auth/compute",
	"https://www.googleapis.com/auth/devstorage.full_control"}

func ProcessAccountFile(scopes []string, text string) (*jwt.Config, error) {
	// Assume text is a JSON string
	if len(scopes) == 0 {
		// Default to driver scopes (defined in driver_gce)
		scopes = DriverScopes
	}
	conf, err := google.JWTConfigFromJSON([]byte(text), scopes...)
	if err != nil {
		// If text was not JSON, assume it is a file path instead
		if _, err := os.Stat(text); os.IsNotExist(err) {
			return nil, fmt.Errorf(
				"account_file path does not exist: %s",
				text)
		}
		data, err := ioutil.ReadFile(text)
		if err != nil {
			return nil, fmt.Errorf(
				"Error reading account_file from path '%s': %s",
				text, err)
		}
		conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/bigquery")
		if err != nil {
			return nil, fmt.Errorf("Error parsing account_file: %s", err)
		}
	}
	return conf, nil
}
