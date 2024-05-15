package services

import (
	"context"
	"d-kv/signer/command-executor/pkg/entity"
	usecase2 "d-kv/signer/command-executor/pkg/usecase"
	dbEntity "d-kv/signer/db-common/entity"
	dbMocks "d-kv/signer/db-common/usecase/mocks"
	"testing"
)

func TestProcessorService_SetStatusById(t *testing.T) {
	type args struct {
		ctx         context.Context
		baseCommand usecase2.DataBaseCommand
		status      dbEntity.Status
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "device test",
			args: args{
				ctx:         nil,
				baseCommand: &entity.CreateDevice{},
				status:      dbEntity.Created,
			},
			wantErr: false,
		},
		{
			name: "bundleId test",
			args: args{
				ctx:         nil,
				baseCommand: &entity.CreateBundleId{},
				status:      dbEntity.Created,
			},
			wantErr: false,
		},
		{
			name: "capability test",
			args: args{
				ctx:         nil,
				baseCommand: &entity.EnableCapabilityType{},
				status:      dbEntity.Created,
			},
			wantErr: false,
		},
		{
			name: "error test",
			args: args{
				ctx:         nil,
				baseCommand: nil,
				status:      dbEntity.Created,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := dbMocks.NewCommandRepo(t)

			queue.
				On("SetStatusByIdDeviceCommand", tt.args.ctx, uint(0), tt.args.status).Return(nil).Maybe()
			queue.
				On("SetStatusByIdBundleIdCommand", tt.args.ctx, uint(0), tt.args.status).Return(nil).Maybe()
			queue.
				On("SetStatusByIdEnableCapabilityTypeCommand", tt.args.ctx, uint(0), tt.args.status).Return(nil).Maybe()
			s := &ProcessorService{
				service: &usecase2.Service{
					Queue: queue,
				},
			}
			if err := s.SetStatusById(tt.args.ctx, &tt.args.baseCommand, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("SetStatusById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
