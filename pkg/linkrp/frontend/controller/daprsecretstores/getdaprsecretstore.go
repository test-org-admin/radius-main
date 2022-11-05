// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package daprsecretstores

import (
	"context"
	"errors"
	"net/http"

	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	ctrl "github.com/project-radius/radius/pkg/armrpc/frontend/controller"
	"github.com/project-radius/radius/pkg/armrpc/rest"
	"github.com/project-radius/radius/pkg/linkrp/datamodel"
	"github.com/project-radius/radius/pkg/linkrp/datamodel/converter"
	"github.com/project-radius/radius/pkg/ucp/store"
)

var _ ctrl.Controller = (*GetDaprSecretStore)(nil)

// GetDaprSecretStore is the controller implementation to get the daprSecretStore conenctor resource.
type GetDaprSecretStore struct {
	ctrl.BaseController
}

// NewGetDaprSecretStore creates a new instance of GetDaprSecretStore.
func NewGetDaprSecretStore(opts ctrl.Options) (ctrl.Controller, error) {
	return &GetDaprSecretStore{ctrl.NewBaseController(opts)}, nil
}

func (daprSecretStore *GetDaprSecretStore) Run(ctx context.Context, w http.ResponseWriter, req *http.Request) (rest.Response, error) {
	serviceCtx := v1.ARMRequestContextFromContext(ctx)

	existingResource := &datamodel.DaprSecretStore{}
	_, err := daprSecretStore.GetResource(ctx, serviceCtx.ResourceID.String(), existingResource)
	if err != nil {
		if errors.Is(&store.ErrNotFound{}, err) {
			return rest.NewNotFoundResponse(serviceCtx.ResourceID), nil
		}
		return nil, err
	}

	versioned, _ := converter.DaprSecretStoreDataModelToVersioned(existingResource, serviceCtx.APIVersion)
	return rest.NewOKResponse(versioned), nil
}