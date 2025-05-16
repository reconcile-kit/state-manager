# UpdateResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**ShardId** | **string** |  | 
**Spec** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateResourceRequest

`func NewUpdateResourceRequest(shardId string, ) *UpdateResourceRequest`

NewUpdateResourceRequest instantiates a new UpdateResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResourceRequestWithDefaults

`func NewUpdateResourceRequestWithDefaults() *UpdateResourceRequest`

NewUpdateResourceRequestWithDefaults instantiates a new UpdateResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *UpdateResourceRequest) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *UpdateResourceRequest) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *UpdateResourceRequest) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *UpdateResourceRequest) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateResourceRequest) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateResourceRequest) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateResourceRequest) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateResourceRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetShardId

`func (o *UpdateResourceRequest) GetShardId() string`

GetShardId returns the ShardId field if non-nil, zero value otherwise.

### GetShardIdOk

`func (o *UpdateResourceRequest) GetShardIdOk() (*string, bool)`

GetShardIdOk returns a tuple with the ShardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardId

`func (o *UpdateResourceRequest) SetShardId(v string)`

SetShardId sets ShardId field to given value.


### GetSpec

`func (o *UpdateResourceRequest) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *UpdateResourceRequest) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *UpdateResourceRequest) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *UpdateResourceRequest) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


