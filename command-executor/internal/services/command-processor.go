package services

import (
	"bytes"
	"context"
	"d-kv/signer/command-executor/pkg/entity"
	_ "d-kv/signer/db-common/config"
	dbEntity "d-kv/signer/db-common/entity"
	"d-kv/signer/db-common/repo/command"
	_ "d-kv/signer/db-common/repo/command"
	"d-kv/signer/db-common/repo/domain"
	_ "d-kv/signer/db-common/repo/domain"
	"d-kv/signer/db-common/repo/vault"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ProcessorService struct {
	Queue *command.Repo
	Repo  *domain.PostgresDomainRepo
	vault *vault.Repo
}

func NewProcessorService(queue *command.Repo, repo *domain.PostgresDomainRepo, vault *vault.Repo) *ProcessorService {
	return &ProcessorService{
		Queue: queue,
		Repo:  repo,
		vault: vault,
	}
}

func (s *ProcessorService) SetStatusById(ctx context.Context, baseCommand *entity.DataBaseCommand, status dbEntity.Status) error {
	err := error(nil)
	switch (*baseCommand).(type) {
	case *entity.CreateDevice:
		err = s.Queue.SetStatusByIdDeviceCommand(ctx, (*baseCommand).GetId(), status)
	case *entity.CreateBundleId:
		err = s.Queue.SetStatusByIdBundleIdCommand(ctx, (*baseCommand).GetId(), status)
	case *entity.EnableCapabilityType:
		err = s.Queue.SetStatusByIdEnableCapabilityTypeCommand(ctx, (*baseCommand).GetId(), status)
	default:
		err = errors.New("unknown type")
	}
	return err
}

func (s *ProcessorService) Processing(ctx context.Context, operation entity.DataBaseCommand) (*http.Response, error) {
	err := s.SetStatusById(ctx, &operation, dbEntity.Processing)
	if err != nil {
		err1 := s.SetStatusById(ctx, &operation, dbEntity.Error)
		if err1 != nil {
			return nil, err1
		}
		return nil, err
	}
	resp, err := s.SendCreateCommand(ctx, operation.Convert(), operation.GetIntegrationId())
	if err != nil {
		err1 := s.SetStatusById(ctx, &operation, dbEntity.Error)
		if err1 != nil {
			return resp, err1
		}
	}
	return resp, err
}

func (s *ProcessorService) StartProcessor(ctx context.Context) {
	for {
		commands := s.Queue.FindByStatusDeviceCommand(ctx, dbEntity.Created)
		for _, operation := range commands {
			localOperation := &entity.CreateDevice{Outer: operation}
			resp, err := s.Processing(ctx, localOperation)
			err = resp.Body.Close()
			if err != nil {
				return
			}
			if err != nil {
				fmt.Println("Error while processing", err)
				continue
			}
			err = s.WriteDevice(ctx, operation)
			if err != nil {
				fmt.Println("Error while writing to db", err)
				continue
			}
		}

		bundleIdCommand := s.Queue.FindByStatusBundleIdCommand(ctx, dbEntity.Created)
		for _, operation := range bundleIdCommand {
			localOperation := &entity.CreateBundleId{Outer: operation}
			resp, err := s.Processing(ctx, localOperation)
			responseObj := entity.BundleIdResponse{}
			decoder := json.NewDecoder(resp.Body)
			err = decoder.Decode(&responseObj)
			if err != nil {
				fmt.Println("Error while decoding api answer", err, responseObj.Id)
				continue
			}
			err = resp.Body.Close()
			if err != nil {
				fmt.Println("Error closing body")
			}
			if err != nil {
				fmt.Println("Error while processing", err)
				continue
			}
			err = s.WriteBundleId(ctx, operation, &responseObj)
			if err != nil {
				fmt.Println("Error while writing to db", err)
				continue
			}
		}

		capabilityCommand := s.Queue.FindByStatusEnableCapabilityTypeCommand(ctx, dbEntity.Created)
		for _, operation := range capabilityCommand {
			localOperation := entity.EnableCapabilityType{Outer: operation}
			resp, err := s.Processing(ctx, &localOperation)
			err = resp.Body.Close()
			if err != nil {
				fmt.Println("Error closing body")
			}
			if err != nil {
				fmt.Println("Error while processing", err)
				continue
			}
			err = s.WriteCapability(ctx, err, operation)
			if err != nil {
				fmt.Println("Error while writing to db", err)
				continue
			}
		}
	}
}

func (s *ProcessorService) SendCreateCommand(ctx context.Context, e entity.ApiEntity, IntegrationId string) (*http.Response, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", string(e.GetURL()), bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	token, err := s.GetTokenByIntegrationID(ctx, IntegrationId)
	if err != nil {
		fmt.Println("Error searching token:", err)
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Error: Unexpected response status code", resp.Status)
		responseBody, err1 := io.ReadAll(resp.Body)
		if err1 != nil {
			fmt.Println("Error reading response body:", err1)
			return resp, err1
		}
		fmt.Println("Response Body:", string(responseBody))
		return resp, errors.New(resp.Status)
	}
	fmt.Println("New instance created successfully!")
	return resp, err
}
