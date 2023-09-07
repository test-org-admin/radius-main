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

package backend

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	v1 "github.com/radius-project/radius/pkg/armrpc/api/v1"
	ctrl "github.com/radius-project/radius/pkg/armrpc/asyncoperation/controller"
	"github.com/radius-project/radius/pkg/armrpc/asyncoperation/worker"
	"github.com/radius-project/radius/pkg/armrpc/hostoptions"
	"github.com/radius-project/radius/pkg/ucp/aws"
	awsproxy_ctrl "github.com/radius-project/radius/pkg/ucp/backend/controller/awsproxy"
)

const (
	UCPProviderName = "ucp"
)

var (
	// ResourceTypeNames is the array that holds resource types that needs async processing.
	// We use this array to generate generic backend controller for each resource.
	ResourceTypeNames = []string{
		"AWSRESOURCE",
	}
)

// Service is a service to run AsyncReqeustProcessWorker.
type Service struct {
	worker.Service
	AWSClients *aws.Clients
}

// NewService creates new service instance to run AsyncReqeustProcessWorker.
func NewService(options hostoptions.HostOptions) *Service {
	return &Service{
		worker.Service{
			ProviderName: UCPProviderName,
			Options:      options,
		},
		nil,
	}
}

// Name returns a string containing the UCPProviderName and the text "async worker".
func (w *Service) Name() string {
	return fmt.Sprintf("%s async worker", UCPProviderName)
}

// Run starts the service and worker. It initializes the service and sets the worker options based on the configuration,
// then starts the service with the given worker options. It returns an error if the initialization fails.
func (w *Service) Run(ctx context.Context) error {
	if err := w.Init(ctx); err != nil {
		return err
	}

	// TODO: Figure out how to pass in AWS config
	awsConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}

	w.AWSClients = &aws.Clients{}
	if w.AWSClients.CloudControl == nil {
		w.AWSClients.CloudControl = cloudcontrol.NewFromConfig(awsConfig)
	}

	if w.AWSClients.CloudFormation == nil {
		w.AWSClients.CloudFormation = cloudformation.NewFromConfig(awsConfig)
	}

	opts := ctrl.Options{
		DataProvider: w.StorageProvider,
		KubeClient:   w.KubeClient,
		AWSClients:   w.AWSClients,
	}

	for _, rt := range ResourceTypeNames {
		// Register controllers

		err := w.Controllers.Register(ctx, rt, v1.OperationPutImperative, awsproxy_ctrl.NewCreateOrUpdateAWSResourceWithPost, opts)
		if err != nil {
			panic(err)
		}
	}

	workerOpts := worker.Options{}
	if w.Options.Config.WorkerServer != nil {
		if w.Options.Config.WorkerServer.MaxOperationConcurrency != nil {
			workerOpts.MaxOperationConcurrency = *w.Options.Config.WorkerServer.MaxOperationConcurrency
		}
		if w.Options.Config.WorkerServer.MaxOperationRetryCount != nil {
			workerOpts.MaxOperationRetryCount = *w.Options.Config.WorkerServer.MaxOperationRetryCount
		}
	}

	return w.Start(ctx, workerOpts)
}
