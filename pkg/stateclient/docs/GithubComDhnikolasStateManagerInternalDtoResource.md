# GithubComDhnikolasStateManagerInternalDtoResource

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
**Spec** | Pointer to **[]int32** |  | [optional] 
**Status** | Pointer to **[]int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewGithubComDhnikolasStateManagerInternalDtoResource

`func NewGithubComDhnikolasStateManagerInternalDtoResource(shardId string, ) *GithubComDhnikolasStateManagerInternalDtoResource`

NewGithubComDhnikolasStateManagerInternalDtoResource instantiates a new GithubComDhnikolasStateManagerInternalDtoResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComDhnikolasStateManagerInternalDtoResourceWithDefaults

`func NewGithubComDhnikolasStateManagerInternalDtoResourceWithDefaults() *GithubComDhnikolasStateManagerInternalDtoResource`

NewGithubComDhnikolasStateManagerInternalDtoResourceWithDefaults instantiates a new GithubComDhnikolasStateManagerInternalDtoResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetCurrentVersion() int32`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetCurrentVersionOk() (*int32, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) SetCurrentVersion(v int32)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetDeletionTimestamp

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetDeletionTimestamp() string`

GetDeletionTimestamp returns the DeletionTimestamp field if non-nil, zero value otherwise.

### GetDeletionTimestampOk

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetDeletionTimestampOk() (*string, bool)`

GetDeletionTimestampOk returns a tuple with the DeletionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTimestamp

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) SetDeletionTimestamp(v string)`

SetDeletionTimestamp sets DeletionTimestamp field to given value.

### HasDeletionTimestamp

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) HasDeletionTimestamp() bool`

HasDeletionTimestamp returns a boolean if a field has been set.

### GetFinalizers

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetFinalizers() []string`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetFinalizersOk() (*[]string, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) SetFinalizers(v []string)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetKind

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabels

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetResourceGroup

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetShardId

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetShardId() string`

GetShardId returns the ShardId field if non-nil, zero value otherwise.

### GetShardIdOk

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetShardIdOk() (*string, bool)`

GetShardIdOk returns a tuple with the ShardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardId

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) SetShardId(v string)`

SetShardId sets ShardId field to given value.


### GetSpec

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetSpec() []int32`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetSpecOk() (*[]int32, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) SetSpec(v []int32)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetStatus() []int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetStatusOk() (*[]int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) SetStatus(v []int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GithubComDhnikolasStateManagerInternalDtoResource) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


