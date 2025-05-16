# CreateResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Finalizers** | Pointer to **[]string** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** |  | 
**ShardId** | **string** |  | 
**Spec** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateResourceRequest

`func NewCreateResourceRequest(name string, shardId string, ) *CreateResourceRequest`

NewCreateResourceRequest instantiates a new CreateResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResourceRequestWithDefaults

`func NewCreateResourceRequestWithDefaults() *CreateResourceRequest`

NewCreateResourceRequestWithDefaults instantiates a new CreateResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *CreateResourceRequest) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateResourceRequest) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateResourceRequest) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateResourceRequest) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetFinalizers

`func (o *CreateResourceRequest) GetFinalizers() []string`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *CreateResourceRequest) GetFinalizersOk() (*[]string, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *CreateResourceRequest) SetFinalizers(v []string)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *CreateResourceRequest) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetLabels

`func (o *CreateResourceRequest) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateResourceRequest) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateResourceRequest) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateResourceRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *CreateResourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateResourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateResourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetShardId

`func (o *CreateResourceRequest) GetShardId() string`

GetShardId returns the ShardId field if non-nil, zero value otherwise.

### GetShardIdOk

`func (o *CreateResourceRequest) GetShardIdOk() (*string, bool)`

GetShardIdOk returns a tuple with the ShardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardId

`func (o *CreateResourceRequest) SetShardId(v string)`

SetShardId sets ShardId field to given value.


### GetSpec

`func (o *CreateResourceRequest) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CreateResourceRequest) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CreateResourceRequest) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CreateResourceRequest) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


