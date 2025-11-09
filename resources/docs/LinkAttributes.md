# LinkAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** |  | 
**TargetUrl** | **string** |  | 
**CreatedAt** | Pointer to **string** | ISO 8601 date-time string | [optional] 
**Ttl** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewLinkAttributes

`func NewLinkAttributes(slug string, targetUrl string, ) *LinkAttributes`

NewLinkAttributes instantiates a new LinkAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkAttributesWithDefaults

`func NewLinkAttributesWithDefaults() *LinkAttributes`

NewLinkAttributesWithDefaults instantiates a new LinkAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *LinkAttributes) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *LinkAttributes) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *LinkAttributes) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTargetUrl

`func (o *LinkAttributes) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *LinkAttributes) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *LinkAttributes) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### GetCreatedAt

`func (o *LinkAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LinkAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LinkAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LinkAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetTtl

`func (o *LinkAttributes) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *LinkAttributes) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *LinkAttributes) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *LinkAttributes) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtlNil

`func (o *LinkAttributes) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *LinkAttributes) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


