# CreateResource200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CurrentVersion** | Pointer to **int32** |  | [optional] 
**DeletionTimestamp** | Pointer to **string** |  | [optional] 
**Finalizers** | Pointer to **[]string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**ResourceGroup** | Pointer to **string** |  | [optional] 
**ShardId** | **string** |  | 
**Spec** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreateResource200Response

`func NewCreateResource200Response(shardId string, ) *CreateResource200Response`

NewCreateResource200Response instantiates a new CreateResource200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResource200ResponseWithDefaults

`func NewCreateResource200ResponseWithDefaults() *CreateResource200Response`

NewCreateResource200ResponseWithDefaults instantiates a new CreateResource200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *CreateResource200Response) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateResource200Response) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateResource200Response) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateResource200Response) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateResource200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateResource200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateResource200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateResource200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *CreateResource200Response) GetCurrentVersion() int32`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *CreateResource200Response) GetCurrentVersionOk() (*int32, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *CreateResource200Response) SetCurrentVersion(v int32)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *CreateResource200Response) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetDeletionTimestamp

`func (o *CreateResource200Response) GetDeletionTimestamp() string`

GetDeletionTimestamp returns the DeletionTimestamp field if non-nil, zero value otherwise.

### GetDeletionTimestampOk

`func (o *CreateResource200Response) GetDeletionTimestampOk() (*string, bool)`

GetDeletionTimestampOk returns a tuple with the DeletionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTimestamp

`func (o *CreateResource200Response) SetDeletionTimestamp(v string)`

SetDeletionTimestamp sets DeletionTimestamp field to given value.

### HasDeletionTimestamp

`func (o *CreateResource200Response) HasDeletionTimestamp() bool`

HasDeletionTimestamp returns a boolean if a field has been set.

### GetFinalizers

`func (o *CreateResource200Response) GetFinalizers() []string`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *CreateResource200Response) GetFinalizersOk() (*[]string, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *CreateResource200Response) SetFinalizers(v []string)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *CreateResource200Response) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetKind

`func (o *CreateResource200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateResource200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateResource200Response) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateResource200Response) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabels

`func (o *CreateResource200Response) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateResource200Response) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateResource200Response) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateResource200Response) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *CreateResource200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateResource200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateResource200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateResource200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *CreateResource200Response) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreateResource200Response) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreateResource200Response) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CreateResource200Response) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetResourceGroup

`func (o *CreateResource200Response) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *CreateResource200Response) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *CreateResource200Response) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *CreateResource200Response) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetShardId

`func (o *CreateResource200Response) GetShardId() string`

GetShardId returns the ShardId field if non-nil, zero value otherwise.

### GetShardIdOk

`func (o *CreateResource200Response) GetShardIdOk() (*string, bool)`

GetShardIdOk returns a tuple with the ShardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardId

`func (o *CreateResource200Response) SetShardId(v string)`

SetShardId sets ShardId field to given value.


### GetSpec

`func (o *CreateResource200Response) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CreateResource200Response) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CreateResource200Response) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CreateResource200Response) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *CreateResource200Response) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateResource200Response) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateResource200Response) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateResource200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateResource200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateResource200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateResource200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateResource200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *CreateResource200Response) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateResource200Response) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateResource200Response) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateResource200Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


