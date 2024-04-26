package client_service

import (
	"d-kv/signer/db-common/entity"
	"d-kv/signer/db-common/repo/command"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CommandService struct {
	queue *command.Repo
}

func NewCommandService(db *command.Repo) *CommandService {
	return &CommandService{queue: db}
}

func (s CommandService) PostBundleId(c *gin.Context, ent *entity.CreateBundleId) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (s CommandService) GetBundleIdStatusByID(c *gin.Context) (entity.Status, error) {
	strId := c.Param("id")
	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		return "", err
	}
	err, status := s.queue.GetStatusByIdBundleIdCommand(c, uint(id))
	if err != nil {
		return "", err
	}
	return status, err
}

func (s CommandService) DelBundleIdById(c *gin.Context) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (s CommandService) PostCapability(c *gin.Context, ent *entity.EnableCapabilityType) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (s CommandService) GetCapabilityStatusByID(c *gin.Context) (entity.Status, error) {
	strId := c.Param("id")
	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		return "", err
	}
	err, status := s.queue.GetStatusByIdBundleIdCommand(c, uint(id))
	if err != nil {
		return "", err
	}
	return status, err
}

func (s CommandService) DelCapability(c *gin.Context) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (s CommandService) PostDevice(c *gin.Context, ent *entity.CreateDevice) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (s CommandService) GetDeviceStatusByID(c *gin.Context) (entity.Status, error) {
	strId := c.Param("id")
	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		return "", err
	}
	err, status := s.queue.GetStatusByIdBundleIdCommand(c, uint(id))
	if err != nil {
		return "", err
	}
	return status, err
}
