package services

import (
	_ "command-executor/internal/config"
	"command-executor/internal/entity"
	_ "command-executor/internal/entity"
	"command-executor/internal/repo/command"
	_ "command-executor/internal/repo/command"
	_ "command-executor/internal/repo/domain"
	"context"
	"fmt"
)

func StartProcessor(ctx context.Context, repo *command.Repo) {
	for {
		deviceCommand := repo.FindByStatusDeviceCommand(ctx, entity.Created)
		for _, operation := range deviceCommand {
			err := repo.SetStatusByIdDeviceCommand(ctx, operation.ID, entity.Processing)
			if err != nil {
				fmt.Println("Error while changing status of Device operation: ", operation.ID)
			}
			fmt.Println(operation) // transport to the service layer
		}
		bundleIdCommand := repo.FindByStatusBundleIdCommand(ctx, entity.Created)
		for _, operation := range bundleIdCommand {
			err := repo.SetStatusByIdBundleIdCommand(ctx, operation.ID, entity.Processing)
			if err != nil {
				fmt.Println("Error while changing status of BundleId operation: ", operation.ID)
			}
			fmt.Println(operation)
		}
		capabilityCommand := repo.FindByStatusEnableCapabilityTypeCommand(ctx, entity.Created)
		for _, operation := range capabilityCommand {
			err := repo.SetStatusByIdEnableCapabilityTypeCommand(ctx, operation.ID, entity.Processing)
			if err != nil {
				fmt.Println("Error while changing status of capability operation: ", operation.ID)
			}
			fmt.Println(operation)
		}
	} // обработка команд по отдельности, подумать над обработкой команд а порядке очереди по времени
}
