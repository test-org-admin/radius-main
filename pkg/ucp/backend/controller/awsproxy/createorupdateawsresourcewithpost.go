/*
Copyright 2023 The Radius Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package awsproxy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/google/uuid"
	v1 "github.com/radius-project/radius/pkg/armrpc/api/v1"
	ctrl "github.com/radius-project/radius/pkg/armrpc/asyncoperation/controller"
	"github.com/radius-project/radius/pkg/to"
	ucp_aws "github.com/radius-project/radius/pkg/ucp/aws"
	"github.com/radius-project/radius/pkg/ucp/aws/servicecontext"
	"github.com/radius-project/radius/pkg/ucp/store"
	"github.com/radius-project/radius/pkg/ucp/ucplog"
)

var _ ctrl.Controller = (*CreateOrUpdateAWSResourceWithPost)(nil)

// CreateOrUpdateAWSResourceWithPost is the controller implementation to create/update an AWS resource.
type CreateOrUpdateAWSResourceWithPost struct {
	ctrl.BaseController
	AWSClients *ucp_aws.Clients
}

// NewCreateOrUpdateAWSResourceWithPost creates a new CreateOrUpdateAWSResourceWithPost.
func NewCreateOrUpdateAWSResourceWithPost(opts ctrl.Options) (ctrl.Controller, error) {
	c := CreateOrUpdateAWSResourceWithPost{
		ctrl.NewBaseAsyncController(opts),
		nil}
	c.AWSClients = opts.AWSClients
	return &c, nil
}

// "Run" reads the request body to get properties, checks if the resource exists, and creates or updates
// the resource accordingly, returning an async operation response.
func (p *CreateOrUpdateAWSResourceWithPost) Run(ctx context.Context, req *ctrl.Request) (ctrl.Result, error) {
	logger := ucplog.FromContextOrDiscard(ctx)
	serviceCtx := servicecontext.AWSRequestContextFromContext(ctx)
	region, errResponse := readRegionFromRequest(req.ResourceID, "")
	if errResponse != nil {
		// TODO: fix the message
		return ctrl.NewFailedResult(v1.ErrorDetails{Message: "Could not read region from request"}), nil
	}

	obj, err := p.StorageClient().Get(ctx, req.ResourceID)
	if err != nil {
		return ctrl.NewFailedResult(v1.ErrorDetails{Message: err.Error()}), err
	}

	var properties map[string]any
	if err := obj.As(&properties); err != nil {
		return ctrl.NewFailedResult(v1.ErrorDetails{Message: "Could not read properties from request"}), nil
	}

	// properties, err := readPropertiesFromBody(req)
	// if err != nil {
	// 	e := v1.ErrorResponse{
	// 		Error: v1.ErrorDetails{
	// 			Code:    v1.CodeInvalid,
	// 			Message: "failed to read request body",
	// 		},
	// 	}
	// 	return armrpc_rest.NewBadRequestARMResponse(e), nil
	// }

	cloudControlOpts := []func(*cloudcontrol.Options){CloudControlRegionOption(region)}
	cloudFormationOpts := []func(*cloudformation.Options){CloudFormationWithRegionOption(region)}

	describeTypeOutput, err := p.AWSClients.CloudFormation.DescribeType(ctx, &cloudformation.DescribeTypeInput{
		Type:     types.RegistryTypeResource,
		TypeName: to.Ptr(serviceCtx.ResourceTypeInAWSFormat()),
	}, cloudFormationOpts...)
	if err != nil {
		// return ucp_aws.HandleAWSError(err)
		return ctrl.NewFailedResult(v1.ErrorDetails{Message: err.Error()}), err
	}

	// var operation uuid.UUID
	desiredState, err := json.Marshal(obj)
	if err != nil {
		// return ucp_aws.HandleAWSError(err)
		return ctrl.NewFailedResult(v1.ErrorDetails{Message: err.Error()}), err
	}

	// existing := true
	var getResponse *cloudcontrol.GetResourceOutput = nil
	computedResourceID := ""
	responseProperties := map[string]any{}

	awsResourceIdentifier, err := getPrimaryIdentifierFromMultiIdentifiers(ctx, properties, *describeTypeOutput.Schema)
	if errors.Is(&ucp_aws.AWSMissingPropertyError{}, err) {
		// assume that if we can't get the AWS resource identifier, we need to create the resource
		// existing = false
	} else if err != nil {
		// return ucp_aws.HandleAWSError(err)
		return ctrl.NewFailedResult(v1.ErrorDetails{Message: err.Error()}), err
	} else {
		computedResourceID = computeResourceID(serviceCtx.ResourceID, awsResourceIdentifier)

		// Create and update work differently for AWS - we need to know if the resource
		// we're working on exists already.
		getResponse, err = p.AWSClients.CloudControl.GetResource(ctx, &cloudcontrol.GetResourceInput{
			TypeName:   to.Ptr(serviceCtx.ResourceTypeInAWSFormat()),
			Identifier: aws.String(awsResourceIdentifier),
		}, cloudControlOpts...)
		if ucp_aws.IsAWSResourceNotFoundError(err) {
			// existing = false
		} else if err != nil {
			return ctrl.NewFailedResult(v1.ErrorDetails{Message: err.Error()}), err
			// return ucp_aws.HandleAWSError(err)
		} else {
			err = json.Unmarshal([]byte(*getResponse.ResourceDescription.Properties), &responseProperties)
			if err != nil {
				// return ucp_aws.HandleAWSError(err)
				return ctrl.NewFailedResult(v1.ErrorDetails{Message: err.Error()}), err
			}
		}
	}

	// TODO????
	// Properties specified by users take precedence
	// for k, v := range properties {
	// 	responseProperties[k] = v
	// }

	// if existing {
	// 	logger.Info(fmt.Sprintf("Updating resource : resourceType %q resourceID %q", serviceCtx.ResourceTypeInAWSFormat(), awsResourceIdentifier))

	// 	// Generate patch
	// 	currentState := []byte(*getResponse.ResourceDescription.Properties)
	// 	resourceTypeSchema := []byte(*describeTypeOutput.Schema)
	// 	patch, err := awsoperations.GeneratePatch(currentState, desiredState, resourceTypeSchema)
	// 	if err != nil {
	// 		return ucp_aws.HandleAWSError(err)
	// 	}

	// 	// Call update only if the patch is not empty
	// 	if len(patch) > 0 {
	// 		marshaled, err := json.Marshal(&patch)
	// 		if err != nil {
	// 			return ucp_aws.HandleAWSError(err)
	// 		}

	// 		response, err := p.awsClients.CloudControl.UpdateResource(ctx, &cloudcontrol.UpdateResourceInput{
	// 			TypeName:      to.Ptr(serviceCtx.ResourceTypeInAWSFormat()),
	// 			Identifier:    aws.String(awsResourceIdentifier),
	// 			PatchDocument: aws.String(string(marshaled)),
	// 		}, cloudControlOpts...)
	// 		if err != nil {
	// 			return ucp_aws.HandleAWSError(err)
	// 		}

	// 		operation, err = uuid.Parse(*response.ProgressEvent.RequestToken)
	// 		if err != nil {
	// 			return ucp_aws.HandleAWSError(err)
	// 		}
	// 	} else {
	// 		// mark provisioning state as succeeded here
	// 		// and return 200, telling the deployment engine that the resource has already been created
	// 		responseProperties["provisioningState"] = v1.ProvisioningStateSucceeded
	// 		responseBody := map[string]any{
	// 			"id":         computedResourceID,
	// 			"name":       awsResourceIdentifier,
	// 			"type":       serviceCtx.ResourceID.Type(),
	// 			"properties": responseProperties,
	// 		}

	// 		resp := armrpc_rest.NewOKResponse(responseBody)
	// 		return resp, nil
	// 	}
	// } else {
	logger.Info(fmt.Sprintf("Creating resource : resourceType %q resourceID %q", serviceCtx.ResourceTypeInAWSFormat(), awsResourceIdentifier))
	response, err := p.AWSClients.CloudControl.CreateResource(ctx, &cloudcontrol.CreateResourceInput{
		TypeName:     to.Ptr(serviceCtx.ResourceTypeInAWSFormat()),
		DesiredState: aws.String(string(desiredState)),
	}, cloudControlOpts...)
	if err != nil {
		// return ucp_aws.HandleAWSError(err)
		return ctrl.NewFailedResult(v1.ErrorDetails{Message: err.Error()}), err

	}

	_, err = uuid.Parse(*response.ProgressEvent.RequestToken)
	if err != nil {
		// return ucp_aws.HandleAWSError(err)
		return ctrl.NewFailedResult(v1.ErrorDetails{Message: err.Error()}), err

	}

	// Get the resource identifier from the progress event response
	if response != nil && response.ProgressEvent != nil && response.ProgressEvent.Identifier != nil {
		awsResourceIdentifier = *response.ProgressEvent.Identifier
		computedResourceID = computeResourceID(serviceCtx.ResourceID, awsResourceIdentifier)
	}
	// }

	responseProperties["provisioningState"] = v1.ProvisioningStateProvisioning

	responseBody := map[string]any{
		"type":       serviceCtx.ResourceID.Type(),
		"properties": responseProperties,
	}
	if computedResourceID != "" && awsResourceIdentifier != "" {
		responseBody["id"] = computedResourceID
		responseBody["name"] = awsResourceIdentifier
	}

	update := &store.Object{
		Metadata: store.Metadata{
			ID: req.ResourceID,
		},
		Data: responseBody,
	}
	err = p.StorageClient().Save(ctx, update, store.WithETag(obj.ETag))
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, err
}
