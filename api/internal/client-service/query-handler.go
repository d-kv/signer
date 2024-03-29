package client_service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

type QueryService struct {
}

func NewQueryService() *QueryService {
	return &QueryService{}
}

func (s *QueryService) GetBundleIds(c *gin.Context) {
	res, err := http.Get("http://localhost:8081/v1/query/bundleIds")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	fmt.Print(string(body))
}

func (s *QueryService) GetBundleIdById(c *gin.Context) {
	id := c.Param("id")
	res, err := http.Get("http://localhost:8081/v1/query/bundleIds" + "/" + id)

	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	fmt.Print(string(body))
}

func (s *QueryService) GetDevices(c *gin.Context) {
	res, err := http.Get("http://localhost:8081/v1/query/devices")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	fmt.Print(string(body))
}

func (s *QueryService) GetDeviceById(c *gin.Context) {
	id := c.Param("id")
	res, err := http.Get("http://localhost:8081/v1/query/devices" + "/" + id)

	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	fmt.Print(string(body))
}

func (s *QueryService) GetCertificates(c *gin.Context) {
	res, err := http.Get("http://localhost:8081/v1/query/certificates")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	fmt.Print(string(body))
}

func (s *QueryService) GetCertificateById(c *gin.Context) {
	id := c.Param("id")
	res, err := http.Get("http://localhost:8081/v1/query/certificates" + "/" + id)

	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	fmt.Print(string(body))
}

func (s *QueryService) GetProfiles(c *gin.Context) {
	res, err := http.Get("http://localhost:8081/v1/query/profiles")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	fmt.Print(string(body))
}

func (s *QueryService) GetProfileById(c *gin.Context) {
	id := c.Param("id")
	res, err := http.Get("http://localhost:8081/v1/query/profiles" + "/" + id)

	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	fmt.Print(string(body))
}
