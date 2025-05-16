# InternalHttpUpdateResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**ShardId** | **string** |  | 
**Spec** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewInternalHttpUpdateResourceRequest

`func NewInternalHttpUpdateResourceRequest(shardId string, ) *InternalHttpUpdateResourceRequest`

NewInternalHttpUpdateResourceRequest instantiates a new InternalHttpUpdateResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalHttpUpdateResourceRequestWithDefaults

`func NewInternalHttpUpdateResourceRequestWithDefaults() *InternalHttpUpdateResourceRequest`

NewInternalHttpUpdateResourceRequestWithDefaults instantiates a new InternalHttpUpdateResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *InternalHttpUpdateResourceRequest) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *InternalHttpUpdateResourceRequest) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *InternalHttpUpdateResourceRequest) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *InternalHttpUpdateResourceRequest) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetLabels

`func (o *InternalHttpUpdateResourceRequest) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InternalHttpUpdateResourceRequest) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InternalHttpUpdateResourceRequest) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InternalHttpUpdateResourceRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetShardId

`func (o *InternalHttpUpdateResourceRequest) GetShardId() string`

GetShardId returns the ShardId field if non-nil, zero value otherwise.

### GetShardIdOk

`func (o *InternalHttpUpdateResourceRequest) GetShardIdOk() (*string, bool)`

GetShardIdOk returns a tuple with the ShardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardId

`func (o *InternalHttpUpdateResourceRequest) SetShardId(v string)`

SetShardId sets ShardId field to given value.


### GetSpec

`func (o *InternalHttpUpdateResourceRequest) GetSpec() []int32`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *InternalHttpUpdateResourceRequest) GetSpecOk() (*[]int32, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *InternalHttpUpdateResourceRequest) SetSpec(v []int32)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *InternalHttpUpdateResourceRequest) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


