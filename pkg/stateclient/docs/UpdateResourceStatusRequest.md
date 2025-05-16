# UpdateResourceStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**CurrentVersion** | Pointer to **int32** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**ShardId** | **string** |  | 
**Status** | Pointer to **map[string]interface{}** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpdateResourceStatusRequest

`func NewUpdateResourceStatusRequest(shardId string, ) *UpdateResourceStatusRequest`

NewUpdateResourceStatusRequest instantiates a new UpdateResourceStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResourceStatusRequestWithDefaults

`func NewUpdateResourceStatusRequestWithDefaults() *UpdateResourceStatusRequest`

NewUpdateResourceStatusRequestWithDefaults instantiates a new UpdateResourceStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *UpdateResourceStatusRequest) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *UpdateResourceStatusRequest) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *UpdateResourceStatusRequest) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *UpdateResourceStatusRequest) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *UpdateResourceStatusRequest) GetCurrentVersion() int32`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *UpdateResourceStatusRequest) GetCurrentVersionOk() (*int32, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *UpdateResourceStatusRequest) SetCurrentVersion(v int32)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *UpdateResourceStatusRequest) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateResourceStatusRequest) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateResourceStatusRequest) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateResourceStatusRequest) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateResourceStatusRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetShardId

`func (o *UpdateResourceStatusRequest) GetShardId() string`

GetShardId returns the ShardId field if non-nil, zero value otherwise.

### GetShardIdOk

`func (o *UpdateResourceStatusRequest) GetShardIdOk() (*string, bool)`

GetShardIdOk returns a tuple with the ShardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardId

`func (o *UpdateResourceStatusRequest) SetShardId(v string)`

SetShardId sets ShardId field to given value.


### GetStatus

`func (o *UpdateResourceStatusRequest) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateResourceStatusRequest) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateResourceStatusRequest) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateResourceStatusRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *UpdateResourceStatusRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateResourceStatusRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateResourceStatusRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateResourceStatusRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


