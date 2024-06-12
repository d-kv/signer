package services

import (
	"bytes"
	"context"
	"d-kv/signer/command-executor/pkg/entity"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type AppApiService struct {
	Client *http.Client
}

func NewApiService(client *http.Client) *AppApiService {
	return &AppApiService{Client: client}
}

func (s *AppApiService) SendCreateCommand(ctx context.Context, e entity.ApiEntity, token string) (*http.Response, error, bool) {
	payload, err := json.Marshal(e)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil, err, false
	}

	req, err := http.NewRequest("POST", string(e.GetURL()), bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err, false
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err, false
	}

	if resp.StatusCode != http.StatusCreated {
		if resp.StatusCode == http.StatusInternalServerError {
			return resp, errors.New(resp.Status), true
		} else {
			fmt.Println("Error: Unexpected response status code", resp.Status)
			responseBody, err1 := io.ReadAll(resp.Body)
			if err1 != nil {
				fmt.Println("Error reading response body:", err1)
				return resp, err1, false
			}
			fmt.Println("BundleIdResponse Body:", string(responseBody))
			return resp, errors.New(resp.Status), false
		}
	}
	fmt.Println("New instance created successfully!")
	return resp, err, false
}
