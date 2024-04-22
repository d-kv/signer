package services

import (
	"bytes"
	"context"
	_ "d-kv/signer/db-common/config"
	"d-kv/signer/db-common/entity"
	_ "d-kv/signer/db-common/entity"
	"d-kv/signer/db-common/repo/command"
	_ "d-kv/signer/db-common/repo/command"
	_ "d-kv/signer/db-common/repo/domain"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const accessToken = "CHANGEME"

func SetStatusByIdAnything(ctx context.Context, baseCommand *entity.DataBaseCommand, status entity.Status, queue *command.Repo) error {
	err := error(nil)
	switch (*baseCommand).(type) {
	case *entity.CreateDevice:
		err = queue.SetStatusByIdDeviceCommand(ctx, (*baseCommand).GetId(), status)
	case *entity.CreateBundleId:
		err = queue.SetStatusByIdBundleIdCommand(ctx, (*baseCommand).GetId(), status)
	case *entity.EnableCapabilityType:
		err = queue.SetStatusByIdEnableCapabilityTypeCommand(ctx, (*baseCommand).GetId(), status)
	default:
		err = errors.New("unknown type")
	}
	return err
}

func Processing(ctx context.Context, queue *command.Repo, repo *domain.PostgresDomainRepo, operation entity.DataBaseCommand) error {
	err := SetStatusByIdAnything(ctx, &operation, entity.Processing, queue)
	if err != nil {
		return err
	}
	err = SendCommand(ctx, operation.Convert())
	if err != nil {
		err = SetStatusByIdAnything(ctx, &operation, entity.Error, queue)
		if err != nil {
			return err
		}
	}
	err = SetStatusByIdAnything(ctx, &operation, entity.Completed, queue)
	if err != nil {
		return err
	}
	return nil
}

func StartProcessor(ctx context.Context, queue *command.Repo, repo *domain.PostgresDomainRepo) {
	for {
		commands := queue.FindByStatusDeviceCommand(ctx, entity.Created)
		for _, operation := range commands {
			err := Processing(ctx, queue, repo, &operation)
			if err != nil {
				fmt.Println("Error while processing", err)
			}
			/*if not repo.DeviceRepo.FindById(ctx, operation.DeviceUdid), create, else update
			device := entity.ConvertDevice(operation)
			repo.DeviceRepo.Create(ctx, device)*/
		}
		bundleIdCommand := queue.FindByStatusBundleIdCommand(ctx, entity.Created)
		for _, operation := range bundleIdCommand {
			err := Processing(ctx, queue, repo, &operation)
			if err != nil {
				fmt.Println("Error while processing", err)
			}

			id, err := repo.IntegrationRepo.FindById(ctx, operation.IntegrationId)
			if err != nil {
				fmt.Println("Error while matching integration", err)
			}
			converted := entity.ConvertBundleId(&id, &operation)
			err = repo.BundleIdRepo.Create(ctx, converted)
			if err != nil {
				fmt.Println("Error while creating record in db")
			}
		}
		capabilityCommand := queue.FindByStatusEnableCapabilityTypeCommand(ctx, entity.Created)
		for _, operation := range capabilityCommand {
			err := Processing(ctx, queue, repo, &operation)
			if err != nil {
				fmt.Println("Error while processing", err)
			}

			/*id, err := repo.BundleIdRepo.FindById(ctx, operation.BundleId)
			if err != nil {
				fmt.Println("Error while matching integration", err)
			}
			converted := entity.ConvertBundleId(&id, &operation)
			err = repo.BundleIdRepo.Create(ctx, converted)
			if err != nil {
				fmt.Println("Error while creating record in db")
			}*/
		}
	}
}

func SendCommand(ctx context.Context, e entity.ApiEntity) error {
	payload, err := json.Marshal(e)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return err
	}

	req, err := http.NewRequest("POST", string(e.GetURL()), bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return err
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
			return err
		}
		fmt.Println("Response Body:", string(responseBody))
		fmt.Println("New instance created successfully!")
	}
	return nil
}
