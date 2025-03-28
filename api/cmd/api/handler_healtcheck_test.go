package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerHealtcheck(t *testing.T) {
	app, assert := MockApp(t)
	serv := httptest.NewServer(app.routes())
	client := serv.Client()

	res, err := client.Get(serv.URL + "/v1/healthcheck")
	body, err := io.ReadAll(res.Body)
	body = bytes.TrimSpace(body)
	defer res.Body.Close()

	var response struct {
		Data struct {
			Status     string `json:"status"`
			SystemInfo struct {
				Environment string `json:"environment"`
				Version     string `json:"version"`
			} `json:"system_info"`
		}
	}
	json.Unmarshal(body, &response)

	assert.Nil(err)
	assert.Equal("available", response.Data.Status)
	assert.Equal("1.0.0", response.Data.SystemInfo.Version)
	assert.Equal(res.StatusCode, http.StatusOK)
}
