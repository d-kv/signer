package client_service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

type QueryService struct {
	url string
}

func NewQueryService(queryURL string) *QueryService {
	return &QueryService{url: queryURL}
}

func (s *QueryService) GetBundleIds(c *gin.Context) {
	res, err := http.Get(s.url + "/bundleIds")
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
	res, err := http.Get(s.url + "/bundleIds" + "/" + id)

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
	res, err := http.Get(s.url + "/devices")

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
	res, err := http.Get(s.url + "/devices/" + id)

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
	res, err := http.Get(s.url + "/certificates")
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
	res, err := http.Get(s.url + "/certificates" + "/" + id)

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
	res, err := http.Get(s.url + "/profiles")
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
	res, err := http.Get(s.url + "/profiles" + "/" + id)

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
