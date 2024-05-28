package client_service

import (
	"d-kv/signer/query-handler/pkg/httpserver/entities"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type QueryService struct {
	url string
}

func NewQueryService(queryURL string) *QueryService {
	return &QueryService{url: queryURL}
}

func GetTenantIntegration(c *gin.Context) string {
	tenant := c.Param("tenantId")
	integration := c.Param("integrationId")
	return "/" + tenant + "/" + integration
}

func (s QueryService) GetBundleIds(c *gin.Context) ([]entities.BundleId, error) {
	var bundles []entities.BundleId
	res, err := http.Get(s.url + GetTenantIntegration(c) + "/bundleIds")
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil, err
	}

	err = json.Unmarshal(body, &bundles)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return bundles, err
	}
	return bundles, err
}

func (s QueryService) GetBundleIdById(c *gin.Context) (entities.BundleId, error) {
	id := c.Param("id")
	GetTenantIntegration(c)
	res, err := http.Get(s.url + GetTenantIntegration(c) + "/bundleIds/" + id)
	var bundle entities.BundleId

	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return bundle, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return bundle, err
	}

	err = json.Unmarshal(body, &bundle)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return bundle, err
	}
	return bundle, err
}

func (s QueryService) GetDevices(c *gin.Context) ([]entities.Device, error) {
	res, err := http.Get(s.url + GetTenantIntegration(c) + "/devices")
	var devices []entities.Device

	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return devices, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return devices, err
	}

	err = json.Unmarshal(body, &devices)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return devices, err
	}
	return devices, err
}

func (s QueryService) GetDeviceById(c *gin.Context) (entities.Device, error) {
	id := c.Param("id")
	res, err := http.Get(s.url + GetTenantIntegration(c) + "/devices/" + id)
	var device entities.Device

	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return device, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return device, err
	}

	err = json.Unmarshal(body, &device)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return device, err
	}
	return device, err
}

func (s QueryService) GetCertificates(c *gin.Context) ([]entities.Certificate, error) {
	res, err := http.Get(s.url + GetTenantIntegration(c) + "/certificates")
	var certificates []entities.Certificate

	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return certificates, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return certificates, err
	}

	err = json.Unmarshal(body, &certificates)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return certificates, err
	}
	return certificates, err
}

func (s QueryService) GetCertificateById(c *gin.Context) (entities.Certificate, error) {
	id := c.Param("id")
	res, err := http.Get(s.url + GetTenantIntegration(c) + "/certificates/" + id)
	var certificate entities.Certificate

	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return certificate, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return certificate, err
	}

	err = json.Unmarshal(body, &certificate)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return certificate, err
	}
	return certificate, err
}

func (s QueryService) GetProfiles(c *gin.Context) ([]entities.Profile, error) {
	res, err := http.Get(s.url + GetTenantIntegration(c) + "/profiles")
	var profiles []entities.Profile

	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return profiles, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return profiles, err
	}

	err = json.Unmarshal(body, &profiles)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return profiles, err
	}
	return profiles, err
}

func (s QueryService) GetProfileById(c *gin.Context) (entities.Profile, error) {
	id := c.Param("id")
	res, err := http.Get(s.url + GetTenantIntegration(c) + "/profiles/" + id)
	var profile entities.Profile

	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return profile, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return profile, err
	}

	err = json.Unmarshal(body, &profile)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return profile, err
	}
	return profile, err
}
