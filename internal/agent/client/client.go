package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	api "github.com/kubev2v/migration-planner/api/v1alpha1/agent"
	client "github.com/kubev2v/migration-planner/internal/api/client/agent"
	baseclient "github.com/kubev2v/migration-planner/internal/client"
	"github.com/kubev2v/migration-planner/pkg/reqid"
)

// NewFromConfig returns a new FlightCtl API client from the given config.
func NewFromConfig(config *baseclient.Config) (*client.ClientWithResponses, error) {

	httpClient, err := baseclient.NewHTTPClientFromConfig(config)
	if err != nil {
		return nil, fmt.Errorf("NewFromConfig: creating HTTP client %w", err)
	}
	ref := client.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set(middleware.RequestIDHeader, reqid.GetReqID())
		return nil
	})
	return client.NewClientWithResponses(config.Service.Server, client.WithHTTPClient(httpClient), ref)
}

type Config = baseclient.Config
type Service = baseclient.Service

func NewDefault() *Config {
	return baseclient.NewDefault()
}

// Planner is the client interface for migration planning.
type Planner interface {
	UpdateSourceInventory(ctx context.Context, id string, params api.SourceInventoryUpdate, rcb ...client.RequestEditorFn) error
	UpdateSourceStatus(ctx context.Context, id string, params api.SourceStatusUpdate, rcb ...client.RequestEditorFn) error
}
