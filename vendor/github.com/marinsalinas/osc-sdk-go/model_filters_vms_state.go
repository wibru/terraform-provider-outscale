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

// FiltersVmsState One or more filters.
type FiltersVmsState struct {
	// The names of the Subregions of the VMs.
	SubregionNames *[]string `json:"SubregionNames,omitempty"`
	// One or more IDs of VMs.
	VmIds *[]string `json:"VmIds,omitempty"`
	// The states of the VMs (`pending` \\| `running` \\| `shutting-down` \\| `terminated` \\| `stopping` \\| `stopped`).
	VmStates *[]string `json:"VmStates,omitempty"`
}

// GetSubregionNames returns the SubregionNames field value if set, zero value otherwise.
func (o *FiltersVmsState) GetSubregionNames() []string {
	if o == nil || o.SubregionNames == nil {
		var ret []string
		return ret
	}
	return *o.SubregionNames
}

// GetSubregionNamesOk returns a tuple with the SubregionNames field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmsState) GetSubregionNamesOk() ([]string, bool) {
	if o == nil || o.SubregionNames == nil {
		var ret []string
		return ret, false
	}
	return *o.SubregionNames, true
}

// HasSubregionNames returns a boolean if a field has been set.
func (o *FiltersVmsState) HasSubregionNames() bool {
	if o != nil && o.SubregionNames != nil {
		return true
	}

	return false
}

// SetSubregionNames gets a reference to the given []string and assigns it to the SubregionNames field.
func (o *FiltersVmsState) SetSubregionNames(v []string) {
	o.SubregionNames = &v
}

// GetVmIds returns the VmIds field value if set, zero value otherwise.
func (o *FiltersVmsState) GetVmIds() []string {
	if o == nil || o.VmIds == nil {
		var ret []string
		return ret
	}
	return *o.VmIds
}

// GetVmIdsOk returns a tuple with the VmIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmsState) GetVmIdsOk() ([]string, bool) {
	if o == nil || o.VmIds == nil {
		var ret []string
		return ret, false
	}
	return *o.VmIds, true
}

// HasVmIds returns a boolean if a field has been set.
func (o *FiltersVmsState) HasVmIds() bool {
	if o != nil && o.VmIds != nil {
		return true
	}

	return false
}

// SetVmIds gets a reference to the given []string and assigns it to the VmIds field.
func (o *FiltersVmsState) SetVmIds(v []string) {
	o.VmIds = &v
}

// GetVmStates returns the VmStates field value if set, zero value otherwise.
func (o *FiltersVmsState) GetVmStates() []string {
	if o == nil || o.VmStates == nil {
		var ret []string
		return ret
	}
	return *o.VmStates
}

// GetVmStatesOk returns a tuple with the VmStates field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmsState) GetVmStatesOk() ([]string, bool) {
	if o == nil || o.VmStates == nil {
		var ret []string
		return ret, false
	}
	return *o.VmStates, true
}

// HasVmStates returns a boolean if a field has been set.
func (o *FiltersVmsState) HasVmStates() bool {
	if o != nil && o.VmStates != nil {
		return true
	}

	return false
}

// SetVmStates gets a reference to the given []string and assigns it to the VmStates field.
func (o *FiltersVmsState) SetVmStates(v []string) {
	o.VmStates = &v
}

type NullableFiltersVmsState struct {
	Value        FiltersVmsState
	ExplicitNull bool
}

func (v NullableFiltersVmsState) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFiltersVmsState) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
