package services

import (
	"bytes"
	"command-executor/internal/entity"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const accessToken = "eyJhbGciOiJFUzI1NiIsImtpZCI6IkNITlBWMzlCNDMiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJhcHBzdG9yZWNvbm5lY3QtdjEiLCJleHAiOjE3MTMzMzM3MDUsImlhdCI6MTcxMzMzMjUwNSwiaXNzIjoiZGQ0ZmNiMjEtMGMyYy00YjIyLTllYTQtMTc5ZTBiNTFmNDgzIn0.F8b0m5gC3ib-ZUfMJx_5Ue8xk73aUAfyh1iWql5Aq99RXltc1pLxOnIZdsfQiiPxRwwoJ_vWF1_flrzB55OjrQ"

func SendCommand(e entity.ApiEntity) {
	payload, err := json.Marshal(e)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", string(e.GetURL()), bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Error: Unexpected response status code", resp.Status)
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}
		fmt.Println("Response Body:", string(responseBody))

		return
	}

	fmt.Println("New device created successfully!")
}
