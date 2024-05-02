package services

import (
	"bytes"
	"context"
	repoWriter "d-kv/signer/command-executor/internal/repo"
	"d-kv/signer/command-executor/pkg/entity"
	_ "d-kv/signer/db-common/config"
	dbEntity "d-kv/signer/db-common/entity"
	"d-kv/signer/db-common/repo/command"
	_ "d-kv/signer/db-common/repo/command"
	"d-kv/signer/db-common/repo/domain"
	_ "d-kv/signer/db-common/repo/domain"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const accessToken = "CHANGEME"

func SetStatusById(ctx context.Context, baseCommand *entity.DataBaseCommand, status dbEntity.Status, queue *command.Repo) error {
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

func Processing(ctx context.Context, queue *command.Repo, operation entity.DataBaseCommand) error {
	err := SetStatusById(ctx, &operation, dbEntity.Processing, queue)
	if err != nil {
		err = SetStatusById(ctx, &operation, dbEntity.Error, queue)
		return err
	}
	err = SendCommand(ctx, operation.Convert())
	if err != nil {
		err = SetStatusById(ctx, &operation, dbEntity.Error, queue)
	}
	return err
}

func StartProcessor(ctx context.Context, queue *command.Repo, repo *domain.PostgresDomainRepo) {
	for {
		commands := queue.FindByStatusDeviceCommand(ctx, dbEntity.Created)
		for _, operation := range commands {
			localOperation := &entity.CreateDevice{Outer: operation}
			err := Processing(ctx, queue, localOperation)
			if err != nil {
				fmt.Println("Error while processing", err)
				continue
			}
			err = repoWriter.WriteDevice(ctx, queue, repo, operation)
			if err != nil {
				fmt.Println("Error while writing to db", err)
				continue
			}
		}

		bundleIdCommand := queue.FindByStatusBundleIdCommand(ctx, dbEntity.Created)
		for _, operation := range bundleIdCommand {
			localOperation := &entity.CreateBundleId{Outer: operation}
			err := Processing(ctx, queue, localOperation)
			if err != nil {
				fmt.Println("Error while processing", err)
				continue
			}
			err = repoWriter.WriteBundleId(ctx, queue, repo, operation)
			if err != nil {
				fmt.Println("Error while writing to db", err)
				continue
			}
		}

		capabilityCommand := queue.FindByStatusEnableCapabilityTypeCommand(ctx, dbEntity.Created)
		for _, operation := range capabilityCommand {
			localOperation := entity.EnableCapabilityType{Outer: operation}
			err := Processing(ctx, queue, &localOperation)
			if err != nil {
				fmt.Println("Error while processing", err)
				continue
			}
			err = repoWriter.WriteCapability(ctx, queue, repo, err, operation)
			if err != nil {
				fmt.Println("Error while writing to db", err)
				continue
			}
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
		err = Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Error: Unexpected response status code", resp.Status)
		responseBody, err1 := io.ReadAll(resp.Body)
		if err1 != nil {
			fmt.Println("Error reading response body:", err1)
			return err1
		}
		fmt.Println("Response Body:", string(responseBody))
	}
	fmt.Println("New instance created successfully!")
	return nil
}
