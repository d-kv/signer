package handlers

import (
	"context"
	"d-kv/signer/db-common/entity"
	"d-kv/signer/db-common/usecase"
	"errors"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

type BundleIdRepoStub struct{}

func (r BundleIdRepoStub) Create(context.Context, *entity.BundleId) error {
	panic("not implemented stub function!")
}

func (r BundleIdRepoStub) FindById(_ context.Context, id string) (entity.BundleId, error) {
	if id == "bundleId" {
		return entity.BundleId{
			ID:         "bundleId",
			Identifier: "identifier",
			Name:       "name",
		}, nil
	}
	return entity.BundleId{}, errors.New("not found")
}

func (r BundleIdRepoStub) FindByIntegrationId(_ context.Context, id string) (entity.BundleId, error) {
	if id == "integrationId" {
		return entity.BundleId{
			ID:            "bundleId",
			Identifier:    "identifier",
			Name:          "name",
			IntegrationId: "integrationId",
		}, nil
	}
	return entity.BundleId{}, errors.New("not found")
}

func (r BundleIdRepoStub) Update(context.Context, *entity.BundleId) error {
	panic("not implemented stub function!")
}

func (r BundleIdRepoStub) DeleteById(context.Context, string) error {
	panic("not implemented stub function!")
}

func (r BundleIdRepoStub) FindAll(context.Context) ([]entity.BundleId, error) {
	panic("not implemented stub function!")
}

type CapabilityRepoStub struct{}

func (r CapabilityRepoStub) Create(context.Context, *entity.Capability) error {
	panic("not implemented stub function!")
}

func (r CapabilityRepoStub) FindByBundleIdId(_ context.Context, id string) ([]entity.Capability, error) {
	if id == "bundleId" {
		return []entity.Capability{
			{
				Type:       string(entity.InAppPurchase),
				BundleIdId: "bundleId",
			},
		}, nil
	}
	return []entity.Capability{}, errors.New("not found")
}

func (r CapabilityRepoStub) Update(context.Context, *entity.Capability) error {
	panic("not implemented stub function!")
}

func (r CapabilityRepoStub) DeleteById(context.Context, string) error {
	panic("not implemented stub function!")
}

func (r CapabilityRepoStub) FindAll(context.Context) ([]entity.Capability, error) {
	panic("not implemented stub function!")
}

func getHandlerWithStub() Handler {
	domainReposStub := usecase.DomainRepos{
		BundleIdRepo:   BundleIdRepoStub{},
		CapabilityRepo: CapabilityRepoStub{},
	}
	handler := Handler{DomainRepos: &domainReposStub}
	return handler
}

func TestBundleIdsHandler(t *testing.T) {
	handler := getHandlerWithStub()

	r := handler.InitRoutes()

	req, _ := http.NewRequest(http.MethodGet, "/v1/query/tenantId/integrationId/bundleIds/", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code, "Header error:", resp.Header().Get("error"))
	require.Equal(
		t,
		`{"identifier":"identifier","name":"name","capabilities":["IN_APP_PURCHASE"]}`,
		resp.Body.String(),
	)
}

func TestBundleIdsHandlerId(t *testing.T) {
	handler := getHandlerWithStub()

	r := handler.InitRoutes()

	req, _ := http.NewRequest(http.MethodGet, "/v1/query/tenantId/integrationId/bundleIds/bundleId", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code, "Header error:", resp.Header().Get("error"))
	require.Equal(
		t,
		`{"identifier":"identifier","name":"name","capabilities":["IN_APP_PURCHASE"]}`,
		resp.Body.String(),
	)
}
