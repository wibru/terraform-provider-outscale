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

// RouteLight Information about the route.
type RouteLight struct {
	// The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).
	DestinationIpRange *string `json:"DestinationIpRange,omitempty"`
	// The type of route (always `static`).
	RouteType *string `json:"RouteType,omitempty"`
	// The current state of the static route (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	State *string `json:"State,omitempty"`
}

// GetDestinationIpRange returns the DestinationIpRange field value if set, zero value otherwise.
func (o *RouteLight) GetDestinationIpRange() string {
	if o == nil || o.DestinationIpRange == nil {
		var ret string
		return ret
	}
	return *o.DestinationIpRange
}

// GetDestinationIpRangeOk returns a tuple with the DestinationIpRange field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RouteLight) GetDestinationIpRangeOk() (string, bool) {
	if o == nil || o.DestinationIpRange == nil {
		var ret string
		return ret, false
	}
	return *o.DestinationIpRange, true
}

// HasDestinationIpRange returns a boolean if a field has been set.
func (o *RouteLight) HasDestinationIpRange() bool {
	if o != nil && o.DestinationIpRange != nil {
		return true
	}

	return false
}

// SetDestinationIpRange gets a reference to the given string and assigns it to the DestinationIpRange field.
func (o *RouteLight) SetDestinationIpRange(v string) {
	o.DestinationIpRange = &v
}

// GetRouteType returns the RouteType field value if set, zero value otherwise.
func (o *RouteLight) GetRouteType() string {
	if o == nil || o.RouteType == nil {
		var ret string
		return ret
	}
	return *o.RouteType
}

// GetRouteTypeOk returns a tuple with the RouteType field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RouteLight) GetRouteTypeOk() (string, bool) {
	if o == nil || o.RouteType == nil {
		var ret string
		return ret, false
	}
	return *o.RouteType, true
}

// HasRouteType returns a boolean if a field has been set.
func (o *RouteLight) HasRouteType() bool {
	if o != nil && o.RouteType != nil {
		return true
	}

	return false
}

// SetRouteType gets a reference to the given string and assigns it to the RouteType field.
func (o *RouteLight) SetRouteType(v string) {
	o.RouteType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RouteLight) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RouteLight) GetStateOk() (string, bool) {
	if o == nil || o.State == nil {
		var ret string
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RouteLight) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *RouteLight) SetState(v string) {
	o.State = &v
}

type NullableRouteLight struct {
	Value        RouteLight
	ExplicitNull bool
}

func (v NullableRouteLight) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableRouteLight) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
