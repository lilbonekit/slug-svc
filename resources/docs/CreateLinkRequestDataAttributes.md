# CreateLinkRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetUrl** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewCreateLinkRequestDataAttributes

`func NewCreateLinkRequestDataAttributes() *CreateLinkRequestDataAttributes`

NewCreateLinkRequestDataAttributes instantiates a new CreateLinkRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLinkRequestDataAttributesWithDefaults

`func NewCreateLinkRequestDataAttributesWithDefaults() *CreateLinkRequestDataAttributes`

NewCreateLinkRequestDataAttributesWithDefaults instantiates a new CreateLinkRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetUrl

`func (o *CreateLinkRequestDataAttributes) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *CreateLinkRequestDataAttributes) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *CreateLinkRequestDataAttributes) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *CreateLinkRequestDataAttributes) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetSlug

`func (o *CreateLinkRequestDataAttributes) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CreateLinkRequestDataAttributes) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CreateLinkRequestDataAttributes) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CreateLinkRequestDataAttributes) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTtl

`func (o *CreateLinkRequestDataAttributes) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CreateLinkRequestDataAttributes) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CreateLinkRequestDataAttributes) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CreateLinkRequestDataAttributes) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtlNil

`func (o *CreateLinkRequestDataAttributes) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *CreateLinkRequestDataAttributes) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


