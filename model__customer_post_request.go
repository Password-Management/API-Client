/*
Customer Creation API

APIs to manage customers in the database.

API version: 1.0.0
Contact: sharmavivek1709@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CustomerPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerPostRequest{}

// CustomerPostRequest struct for CustomerPostRequest
type CustomerPostRequest struct {
	Email string `json:"email"`
	Name string `json:"name"`
	Plan string `json:"plan"`
	Algorithm string `json:"algorithm"`
	Platform string `json:"platform"`
}

type _CustomerPostRequest CustomerPostRequest

// NewCustomerPostRequest instantiates a new CustomerPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPostRequest(email string, name string, plan string, algorithm string, platform string) *CustomerPostRequest {
	this := CustomerPostRequest{}
	this.Email = email
	this.Name = name
	this.Plan = plan
	this.Algorithm = algorithm
	this.Platform = platform
	return &this
}

// NewCustomerPostRequestWithDefaults instantiates a new CustomerPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPostRequestWithDefaults() *CustomerPostRequest {
	this := CustomerPostRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *CustomerPostRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CustomerPostRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CustomerPostRequest) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value
func (o *CustomerPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomerPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomerPostRequest) SetName(v string) {
	o.Name = v
}

// GetPlan returns the Plan field value
func (o *CustomerPostRequest) GetPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *CustomerPostRequest) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *CustomerPostRequest) SetPlan(v string) {
	o.Plan = v
}

// GetAlgorithm returns the Algorithm field value
func (o *CustomerPostRequest) GetAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *CustomerPostRequest) GetAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value
func (o *CustomerPostRequest) SetAlgorithm(v string) {
	o.Algorithm = v
}

// GetPlatform returns the Platform field value
func (o *CustomerPostRequest) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *CustomerPostRequest) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *CustomerPostRequest) SetPlatform(v string) {
	o.Platform = v
}

func (o CustomerPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["name"] = o.Name
	toSerialize["plan"] = o.Plan
	toSerialize["algorithm"] = o.Algorithm
	toSerialize["platform"] = o.Platform
	return toSerialize, nil
}

func (o *CustomerPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"name",
		"plan",
		"algorithm",
		"platform",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCustomerPostRequest := _CustomerPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomerPostRequest)

	if err != nil {
		return err
	}

	*o = CustomerPostRequest(varCustomerPostRequest)

	return err
}

type NullableCustomerPostRequest struct {
	value *CustomerPostRequest
	isSet bool
}

func (v NullableCustomerPostRequest) Get() *CustomerPostRequest {
	return v.value
}

func (v *NullableCustomerPostRequest) Set(val *CustomerPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPostRequest(val *CustomerPostRequest) *NullableCustomerPostRequest {
	return &NullableCustomerPostRequest{value: val, isSet: true}
}

func (v NullableCustomerPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


