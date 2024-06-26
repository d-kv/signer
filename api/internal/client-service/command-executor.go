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

// BundleID

func (s *CommandService) PostBundleId(c *gin.Context, ent *entity.CreateBundleId) (uint, error) {
	err, e := s.queue.CreateBundleIdCommand(c, *ent)
	if err != nil {
		return 0, err
	}
	return e.ID, err
}

func (s *CommandService) GetBundleIdStatusByID(c *gin.Context) (entity.Status, error) {
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

func (s *CommandService) DelBundleIdById(c *gin.Context) (uint, error) {
	//TODO implement me
	return 0, nil
}

// Capability

func (s *CommandService) PostCapability(c *gin.Context, ent *entity.EnableCapabilityType) (uint, error) {
	err, e := s.queue.CreateEnableCapabilityTypeCommand(c, *ent)
	if err != nil {
		return 0, err
	}
	return e.ID, err
}

func (s *CommandService) GetCapabilityStatusByID(c *gin.Context) (entity.Status, error) {
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

func (s *CommandService) DelCapability(c *gin.Context) (uint, error) {
	//TODO implement me
	return 0, nil
}

// Device

func (s *CommandService) PostDevice(c *gin.Context, ent *entity.CreateDevice) (uint, error) {
	err, e := s.queue.CreateDeviceCommand(c, *ent)
	if err != nil {
		return 0, err
	}
	return e.ID, err
}

func (s *CommandService) GetDeviceStatusByID(c *gin.Context) (entity.Status, error) {
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

// Profile

func (s *CommandService) PostProfile(c *gin.Context, ent *entity.CreateProfile) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CommandService) GetProfileStatusByID(c *gin.Context) (entity.Status, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CommandService) DelProfileById(c *gin.Context) (uint, error) {
	//TODO implement me
	panic("implement me")
}

// Certificate

func (s *CommandService) PostCertificate(c *gin.Context, ent *entity.CreateCertificate) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CommandService) GetCertificateStatusByID(c *gin.Context) (entity.Status, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CommandService) DelCertificateById(c *gin.Context) (uint, error) {
	//TODO implement me
	panic("implement me")
}
