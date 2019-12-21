/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oscgo

import (
	"bytes"
	"encoding/json"
)

// CreateNatServiceRequest struct for CreateNatServiceRequest
type CreateNatServiceRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The allocation ID of the EIP to associate with the NAT service.<br /> If the EIP is already associated with another resource, you must first disassociate it.
	PublicIpId string `json:"PublicIpId"`
	// The ID of the Subnet in which you want to create the NAT service.
	SubnetId string `json:"SubnetId"`
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateNatServiceRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateNatServiceRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateNatServiceRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateNatServiceRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetPublicIpId returns the PublicIpId field value
func (o *CreateNatServiceRequest) GetPublicIpId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicIpId
}

// SetPublicIpId sets field value
func (o *CreateNatServiceRequest) SetPublicIpId(v string) {
	o.PublicIpId = v
}

// GetSubnetId returns the SubnetId field value
func (o *CreateNatServiceRequest) GetSubnetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetId
}

// SetSubnetId sets field value
func (o *CreateNatServiceRequest) SetSubnetId(v string) {
	o.SubnetId = v
}

type NullableCreateNatServiceRequest struct {
	Value        CreateNatServiceRequest
	ExplicitNull bool
}

func (v NullableCreateNatServiceRequest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCreateNatServiceRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
