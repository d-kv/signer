package services

import (
	"context"
	"d-kv/signer/command-executor/pkg/entity"
	"d-kv/signer/command-executor/pkg/usecase"
	_ "d-kv/signer/db-common/config"
	dbEntity "d-kv/signer/db-common/entity"
	_ "d-kv/signer/db-common/repo/command"
	_ "d-kv/signer/db-common/repo/domain"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ProcessorService struct {
	service *usecase.Service
}

func NewProcessorService(service *usecase.Service) *ProcessorService {
	return &ProcessorService{
		service: service,
	}
}

func (s *ProcessorService) SetStatusById(ctx context.Context, baseCommand *usecase.DataBaseCommand, status dbEntity.Status) error {
	err := error(nil)
	switch (*baseCommand).(type) {
	case *entity.CreateDevice:
		err = s.service.Queue.SetStatusByIdDeviceCommand(ctx, (*baseCommand).GetId(), status)
	case *entity.CreateBundleId:
		err = s.service.Queue.SetStatusByIdBundleIdCommand(ctx, (*baseCommand).GetId(), status)
	case *entity.EnableCapabilityType:
		err = s.service.Queue.SetStatusByIdEnableCapabilityTypeCommand(ctx, (*baseCommand).GetId(), status)
	default:
		err = errors.New("unknown type")
	}
	return err
}

func (s *ProcessorService) Processing(ctx context.Context, operation usecase.DataBaseCommand) (*http.Response, error) {
	err := s.SetStatusById(ctx, &operation, dbEntity.Processing)
	if err != nil {
		return nil, err
	}
	err, tokenInfo := s.service.Vault.FindTokenByIntegrationId(ctx, operation.GetIntegrationId())
	if err != nil {
		fmt.Println("Incorrect integration:", err)
		return nil, err
	}
	token, err := s.service.Token.GetJwtToken(tokenInfo)
	if err != nil {
		fmt.Println("Error searching token:", err)
		return nil, err
	}
	resp, err := s.service.Api.SendCreateCommand(ctx, operation.Convert(), token)
	if err != nil {
	}
	return resp, err
}

func (s *ProcessorService) StartProcessor(ctx context.Context) {
	for {
		commands := s.service.Queue.FindByStatusDeviceCommand(ctx, dbEntity.Created)
		for _, operation := range commands {
			localOperation := &entity.CreateDevice{Outer: operation}
			resp, err := s.Processing(ctx, localOperation)
			if err != nil {
				fmt.Println("Error while processing", err)
				err = s.service.Queue.SetStatusByIdDeviceCommand(ctx, operation.ID, dbEntity.Error)
				if err != nil {
					fmt.Println("Error while status change", err)
				}
				continue
			}
			err = resp.Body.Close()
			if err != nil {
				fmt.Println("Error while closing response body", err)
				err = s.service.Queue.SetStatusByIdDeviceCommand(ctx, operation.ID, dbEntity.Error)
				if err != nil {
					fmt.Println("Error while status change", err)
				}
				continue
			}
			err = s.service.Db.WriteDevice(ctx, operation)
			if err != nil {
				fmt.Println("Error while writing to db", err)
				err = s.service.Queue.SetStatusByIdDeviceCommand(ctx, operation.ID, dbEntity.Error)
				if err != nil {
					fmt.Println("Error while status change", err)
				}
				continue
			}
			err = s.service.Queue.SetStatusByIdDeviceCommand(ctx, operation.ID, dbEntity.Completed)
			if err != nil {
				fmt.Println("Error while changing status", err)
				continue
			}
		}

		bundleIdCommand := s.service.Queue.FindByStatusBundleIdCommand(ctx, dbEntity.Created)
		for _, operation := range bundleIdCommand {
			localOperation := &entity.CreateBundleId{Outer: operation}
			resp, err := s.Processing(ctx, localOperation)
			if err != nil {
				fmt.Println("Error while processing", err)
				err = s.service.Queue.SetStatusByIdBundleIdCommand(ctx, operation.ID, dbEntity.Error)
				if err != nil {
					fmt.Println("Error while status change", err)
				}
				continue
			}
			responseObj := entity.BundleIdResponse{}
			decoder := json.NewDecoder(resp.Body)
			err = decoder.Decode(&responseObj)
			if err != nil {
				fmt.Println("Error while decoding api answer", err, responseObj.Id)
				err = s.service.Queue.SetStatusByIdBundleIdCommand(ctx, operation.ID, dbEntity.Error)
				if err != nil {
					fmt.Println("Error while status change", err)
				}
				continue
			}
			err = resp.Body.Close()
			if err != nil {
				fmt.Println("Error closing body", err)
				err = s.service.Queue.SetStatusByIdBundleIdCommand(ctx, operation.ID, dbEntity.Error)
				if err != nil {
					fmt.Println("Error while status change", err)
				}
			}
			err = s.service.Db.WriteBundleId(ctx, operation, &responseObj)
			if err != nil {
				fmt.Println("Error while writing to db", err)
				err = s.service.Queue.SetStatusByIdBundleIdCommand(ctx, operation.ID, dbEntity.Error)
				if err != nil {
					fmt.Println("Error while status change", err)
				}
				continue
			}
			err = s.service.Queue.SetStatusByIdBundleIdCommand(ctx, operation.ID, dbEntity.Completed)
			if err != nil {
				fmt.Println("Error while changing status", err)
				continue
			}
		}

		capabilityCommand := s.service.Queue.FindByStatusEnableCapabilityTypeCommand(ctx, dbEntity.Created)
		for _, operation := range capabilityCommand {
			localOperation := entity.EnableCapabilityType{Outer: operation}
			resp, err := s.Processing(ctx, &localOperation)
			err = resp.Body.Close()
			if err != nil {
				fmt.Println("Error closing body")
				err = s.service.Queue.SetStatusByIdEnableCapabilityTypeCommand(ctx, operation.ID, dbEntity.Error)
				if err != nil {
					fmt.Println("Error while status change", err)
				}
				continue
			}
			if err != nil {
				fmt.Println("Error while processing", err)
				err = s.service.Queue.SetStatusByIdEnableCapabilityTypeCommand(ctx, operation.ID, dbEntity.Error)
				if err != nil {
					fmt.Println("Error while status change", err)
				}
				continue
			}
			err = s.service.Db.WriteCapability(ctx, err, operation)
			if err != nil {
				fmt.Println("Error while writing to db", err)
				err = s.service.Queue.SetStatusByIdEnableCapabilityTypeCommand(ctx, operation.ID, dbEntity.Error)
				if err != nil {
					fmt.Println("Error while status change", err)
				}
				continue
			}
			err = s.service.Queue.SetStatusByIdEnableCapabilityTypeCommand(ctx, operation.ID, dbEntity.Completed)
			if err != nil {
				fmt.Println("Error while changing status", err)
				continue
			}
		}
	}
}
