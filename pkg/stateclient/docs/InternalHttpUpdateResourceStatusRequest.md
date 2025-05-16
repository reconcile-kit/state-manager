# InternalHttpUpdateResourceStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**CurrentVersion** | Pointer to **int32** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**ShardId** | **string** |  | 
**Status** | Pointer to **[]int32** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewInternalHttpUpdateResourceStatusRequest

`func NewInternalHttpUpdateResourceStatusRequest(shardId string, ) *InternalHttpUpdateResourceStatusRequest`

NewInternalHttpUpdateResourceStatusRequest instantiates a new InternalHttpUpdateResourceStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalHttpUpdateResourceStatusRequestWithDefaults

`func NewInternalHttpUpdateResourceStatusRequestWithDefaults() *InternalHttpUpdateResourceStatusRequest`

NewInternalHttpUpdateResourceStatusRequestWithDefaults instantiates a new InternalHttpUpdateResourceStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *InternalHttpUpdateResourceStatusRequest) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *InternalHttpUpdateResourceStatusRequest) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *InternalHttpUpdateResourceStatusRequest) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *InternalHttpUpdateResourceStatusRequest) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *InternalHttpUpdateResourceStatusRequest) GetCurrentVersion() int32`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *InternalHttpUpdateResourceStatusRequest) GetCurrentVersionOk() (*int32, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *InternalHttpUpdateResourceStatusRequest) SetCurrentVersion(v int32)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *InternalHttpUpdateResourceStatusRequest) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetLabels

`func (o *InternalHttpUpdateResourceStatusRequest) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InternalHttpUpdateResourceStatusRequest) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InternalHttpUpdateResourceStatusRequest) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InternalHttpUpdateResourceStatusRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetShardId

`func (o *InternalHttpUpdateResourceStatusRequest) GetShardId() string`

GetShardId returns the ShardId field if non-nil, zero value otherwise.

### GetShardIdOk

`func (o *InternalHttpUpdateResourceStatusRequest) GetShardIdOk() (*string, bool)`

GetShardIdOk returns a tuple with the ShardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardId

`func (o *InternalHttpUpdateResourceStatusRequest) SetShardId(v string)`

SetShardId sets ShardId field to given value.


### GetStatus

`func (o *InternalHttpUpdateResourceStatusRequest) GetStatus() []int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalHttpUpdateResourceStatusRequest) GetStatusOk() (*[]int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalHttpUpdateResourceStatusRequest) SetStatus(v []int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InternalHttpUpdateResourceStatusRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *InternalHttpUpdateResourceStatusRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InternalHttpUpdateResourceStatusRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InternalHttpUpdateResourceStatusRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InternalHttpUpdateResourceStatusRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


