# InternalHttpCreateResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Finalizers** | Pointer to **[]string** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** |  | 
**ShardId** | **string** |  | 
**Spec** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewInternalHttpCreateResourceRequest

`func NewInternalHttpCreateResourceRequest(name string, shardId string, ) *InternalHttpCreateResourceRequest`

NewInternalHttpCreateResourceRequest instantiates a new InternalHttpCreateResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalHttpCreateResourceRequestWithDefaults

`func NewInternalHttpCreateResourceRequestWithDefaults() *InternalHttpCreateResourceRequest`

NewInternalHttpCreateResourceRequestWithDefaults instantiates a new InternalHttpCreateResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *InternalHttpCreateResourceRequest) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *InternalHttpCreateResourceRequest) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *InternalHttpCreateResourceRequest) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *InternalHttpCreateResourceRequest) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetFinalizers

`func (o *InternalHttpCreateResourceRequest) GetFinalizers() []string`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *InternalHttpCreateResourceRequest) GetFinalizersOk() (*[]string, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *InternalHttpCreateResourceRequest) SetFinalizers(v []string)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *InternalHttpCreateResourceRequest) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetLabels

`func (o *InternalHttpCreateResourceRequest) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InternalHttpCreateResourceRequest) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InternalHttpCreateResourceRequest) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InternalHttpCreateResourceRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *InternalHttpCreateResourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalHttpCreateResourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalHttpCreateResourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetShardId

`func (o *InternalHttpCreateResourceRequest) GetShardId() string`

GetShardId returns the ShardId field if non-nil, zero value otherwise.

### GetShardIdOk

`func (o *InternalHttpCreateResourceRequest) GetShardIdOk() (*string, bool)`

GetShardIdOk returns a tuple with the ShardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardId

`func (o *InternalHttpCreateResourceRequest) SetShardId(v string)`

SetShardId sets ShardId field to given value.


### GetSpec

`func (o *InternalHttpCreateResourceRequest) GetSpec() []int32`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *InternalHttpCreateResourceRequest) GetSpecOk() (*[]int32, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *InternalHttpCreateResourceRequest) SetSpec(v []int32)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *InternalHttpCreateResourceRequest) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


