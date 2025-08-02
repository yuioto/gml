package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func DoJSONRequest[T any](client *http.Client, req *http.Request, out *T) error {
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("HTTP request failed")
		return fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error().Err(err).Int("status", resp.StatusCode).Msg("HTTP status not OK")
		return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(out)
}

func FetchJSON[T any](client *http.Client, url string, out *T) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal JSON")
		return err
	}
	return DoJSONRequest(client, req, out)
}
