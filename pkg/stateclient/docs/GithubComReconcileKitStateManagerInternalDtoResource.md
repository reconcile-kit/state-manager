# GithubComReconcileKitStateManagerInternalDtoResource

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

### NewGithubComReconcileKitStateManagerInternalDtoResource

`func NewGithubComReconcileKitStateManagerInternalDtoResource(shardId string, ) *GithubComReconcileKitStateManagerInternalDtoResource`

NewGithubComReconcileKitStateManagerInternalDtoResource instantiates a new GithubComReconcileKitStateManagerInternalDtoResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComReconcileKitStateManagerInternalDtoResourceWithDefaults

`func NewGithubComReconcileKitStateManagerInternalDtoResourceWithDefaults() *GithubComReconcileKitStateManagerInternalDtoResource`

NewGithubComReconcileKitStateManagerInternalDtoResourceWithDefaults instantiates a new GithubComReconcileKitStateManagerInternalDtoResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetCurrentVersion() int32`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetCurrentVersionOk() (*int32, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) SetCurrentVersion(v int32)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetDeletionTimestamp

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetDeletionTimestamp() string`

GetDeletionTimestamp returns the DeletionTimestamp field if non-nil, zero value otherwise.

### GetDeletionTimestampOk

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetDeletionTimestampOk() (*string, bool)`

GetDeletionTimestampOk returns a tuple with the DeletionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTimestamp

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) SetDeletionTimestamp(v string)`

SetDeletionTimestamp sets DeletionTimestamp field to given value.

### HasDeletionTimestamp

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) HasDeletionTimestamp() bool`

HasDeletionTimestamp returns a boolean if a field has been set.

### GetFinalizers

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetFinalizers() []string`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetFinalizersOk() (*[]string, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) SetFinalizers(v []string)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetKind

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabels

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetResourceGroup

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetShardId

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetShardId() string`

GetShardId returns the ShardId field if non-nil, zero value otherwise.

### GetShardIdOk

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetShardIdOk() (*string, bool)`

GetShardIdOk returns a tuple with the ShardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardId

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) SetShardId(v string)`

SetShardId sets ShardId field to given value.


### GetSpec

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetSpec() []int32`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetSpecOk() (*[]int32, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) SetSpec(v []int32)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetStatus() []int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetStatusOk() (*[]int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) SetStatus(v []int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GithubComReconcileKitStateManagerInternalDtoResource) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


