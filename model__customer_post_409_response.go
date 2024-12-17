/*
Customer Creation API

API's cerate customer in the database

API version: 1.0.0
Contact: sharmavivek1709@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CustomerPost409Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerPost409Response{}

// CustomerPost409Response struct for CustomerPost409Response
type CustomerPost409Response struct {
	Token *string `json:"token,omitempty"`
}

// NewCustomerPost409Response instantiates a new CustomerPost409Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPost409Response() *CustomerPost409Response {
	this := CustomerPost409Response{}
	return &this
}

// NewCustomerPost409ResponseWithDefaults instantiates a new CustomerPost409Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPost409ResponseWithDefaults() *CustomerPost409Response {
	this := CustomerPost409Response{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CustomerPost409Response) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPost409Response) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CustomerPost409Response) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CustomerPost409Response) SetToken(v string) {
	o.Token = &v
}

func (o CustomerPost409Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerPost409Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableCustomerPost409Response struct {
	value *CustomerPost409Response
	isSet bool
}

func (v NullableCustomerPost409Response) Get() *CustomerPost409Response {
	return v.value
}

func (v *NullableCustomerPost409Response) Set(val *CustomerPost409Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPost409Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPost409Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPost409Response(val *CustomerPost409Response) *NullableCustomerPost409Response {
	return &NullableCustomerPost409Response{value: val, isSet: true}
}

func (v NullableCustomerPost409Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPost409Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

