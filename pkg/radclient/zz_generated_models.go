// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package radclient

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// ApplicationCreateOrUpdateOptions contains the optional parameters for the Application.CreateOrUpdate method.
type ApplicationCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ApplicationCreateParameters - Parameters used to create an application.
type ApplicationCreateParameters struct {
	// REQUIRED; Any object
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// ApplicationDeleteOptions contains the optional parameters for the Application.Delete method.
type ApplicationDeleteOptions struct {
	// placeholder for future optional parameters
}

// ApplicationGetOptions contains the optional parameters for the Application.Get method.
type ApplicationGetOptions struct {
	// placeholder for future optional parameters
}

// ApplicationList - Application list.
type ApplicationList struct {
	// REQUIRED; List of applications.
	Value []*ApplicationResource `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationList.
func (a ApplicationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// ApplicationListByResourceGroupOptions contains the optional parameters for the Application.ListByResourceGroup method.
type ApplicationListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ApplicationResource - Application resource.
type ApplicationResource struct {
	TrackedResource
	// REQUIRED; Properties of the application.
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationResource.
func (a ApplicationResource) MarshalJSON() ([]byte, error) {
	objectMap := a.TrackedResource.marshalInternal()
	populate(objectMap, "properties", a.Properties)
	return json.Marshal(objectMap)
}

// ComponentCreateOrUpdateOptions contains the optional parameters for the Component.CreateOrUpdate method.
type ComponentCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ComponentCreateParameters - Parameters used to create a component.
type ComponentCreateParameters struct {
	// REQUIRED; Resource type of the component
	Kind *string `json:"kind,omitempty"`

	// REQUIRED; Properties of a component.
	Properties *ComponentProperties `json:"properties,omitempty"`
}

// ComponentDeleteOptions contains the optional parameters for the Component.Delete method.
type ComponentDeleteOptions struct {
	// placeholder for future optional parameters
}

// ComponentGetOptions contains the optional parameters for the Component.Get method.
type ComponentGetOptions struct {
	// placeholder for future optional parameters
}

// ComponentList - Component list.
type ComponentList struct {
	// REQUIRED; List of components.
	Value []*ComponentResource `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ComponentList.
func (c ComponentList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// ComponentListByApplicationOptions contains the optional parameters for the Component.ListByApplication method.
type ComponentListByApplicationOptions struct {
	// placeholder for future optional parameters
}

// ComponentProperties - Properties of a component.
type ComponentProperties struct {
	// Bindings spec of the component
	Bindings map[string]interface{} `json:"bindings,omitempty"`

	// Config of the component
	Config map[string]interface{} `json:"config,omitempty"`
	OutputResources []map[string]interface{} `json:"outputResources,omitempty"`

	// Revision of the component
	Revision *string `json:"revision,omitempty"`

	// Run spec of the component
	Run map[string]interface{} `json:"run,omitempty"`

	// Traits spec of the component
	Traits map[string]interface{} `json:"traits,omitempty"`

	// Uses spec of the component
	Uses map[string]interface{} `json:"uses,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ComponentProperties.
func (c ComponentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "bindings", c.Bindings)
	populate(objectMap, "config", c.Config)
	populate(objectMap, "outputResources", c.OutputResources)
	populate(objectMap, "revision", c.Revision)
	populate(objectMap, "run", c.Run)
	populate(objectMap, "traits", c.Traits)
	populate(objectMap, "uses", c.Uses)
	return json.Marshal(objectMap)
}

// ComponentResource - Component resource.
type ComponentResource struct {
	TrackedResource
	// REQUIRED; Resource type of the component
	Kind *string `json:"kind,omitempty"`

	// REQUIRED; Properties of the component.
	Properties *ComponentProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ComponentResource.
func (c ComponentResource) MarshalJSON() ([]byte, error) {
	objectMap := c.TrackedResource.marshalInternal()
	populate(objectMap, "kind", c.Kind)
	populate(objectMap, "properties", c.Properties)
	return json.Marshal(objectMap)
}

// DeploymentBeginCreateOrUpdateOptions contains the optional parameters for the Deployment.BeginCreateOrUpdate method.
type DeploymentBeginCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// DeploymentBeginDeleteOptions contains the optional parameters for the Deployment.BeginDelete method.
type DeploymentBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// DeploymentCreateParameters - Parameters used to create a deployment.
type DeploymentCreateParameters struct {
	// REQUIRED; Properties of a deployment.
	Properties *DeploymentProperties `json:"properties,omitempty"`
}

// DeploymentGetOptions contains the optional parameters for the Deployment.Get method.
type DeploymentGetOptions struct {
	// placeholder for future optional parameters
}

// DeploymentList - Deployment list.
type DeploymentList struct {
	// REQUIRED; List of deployments.
	Value []*DeploymentResource `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type DeploymentList.
func (d DeploymentList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// DeploymentListByApplicationOptions contains the optional parameters for the Deployment.ListByApplication method.
type DeploymentListByApplicationOptions struct {
	// placeholder for future optional parameters
}

// DeploymentProperties - Properties of a deployment.
type DeploymentProperties struct {
	// REQUIRED; List of components in the deployment.
	Components []*DeploymentPropertiesComponentsItem `json:"components,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type DeploymentProperties.
func (d DeploymentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "components", d.Components)
	return json.Marshal(objectMap)
}

type DeploymentPropertiesComponentsItem struct {
	// REQUIRED; Name of the component
	ComponentName *string `json:"componentName,omitempty"`
}

// DeploymentResource - Deployment resource.
type DeploymentResource struct {
	TrackedResource
	// REQUIRED; Properties of the deployment.
	Properties *DeploymentProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type DeploymentResource.
func (d DeploymentResource) MarshalJSON() ([]byte, error) {
	objectMap := d.TrackedResource.marshalInternal()
	populate(objectMap, "properties", d.Properties)
	return json.Marshal(objectMap)
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info map[string]interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations. (This also follows the OData
// error response format.).
// Implements the error and azcore.HTTPResponse interfaces.
type ErrorResponse struct {
	raw string
	// The error object.
	InnerError *ErrorDetail `json:"error,omitempty"`
}

// Error implements the error interface for type ErrorResponse.
// The contents of the error text are not contractual and subject to change.
func (e ErrorResponse) Error() string {
	return e.raw
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := r.marshalInternal()
	return json.Marshal(objectMap)
}

func (r Resource) marshalInternal() map[string]interface{} {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "type", r.Type)
	return objectMap
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags' and a 'location'
type TrackedResource struct {
	Resource
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := t.marshalInternal()
	return json.Marshal(objectMap)
}

func (t TrackedResource) marshalInternal() map[string]interface{} {
	objectMap := t.Resource.marshalInternal()
	populate(objectMap, "location", t.Location)
	populate(objectMap, "tags", t.Tags)
	return objectMap
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

