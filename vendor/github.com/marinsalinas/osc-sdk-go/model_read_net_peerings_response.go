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

// ReadNetPeeringsResponse struct for ReadNetPeeringsResponse
type ReadNetPeeringsResponse struct {
	// Information about one or more Net peering connections.
	NetPeerings     *[]NetPeering    `json:"NetPeerings,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetNetPeerings returns the NetPeerings field value if set, zero value otherwise.
func (o *ReadNetPeeringsResponse) GetNetPeerings() []NetPeering {
	if o == nil || o.NetPeerings == nil {
		var ret []NetPeering
		return ret
	}
	return *o.NetPeerings
}

// GetNetPeeringsOk returns a tuple with the NetPeerings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadNetPeeringsResponse) GetNetPeeringsOk() ([]NetPeering, bool) {
	if o == nil || o.NetPeerings == nil {
		var ret []NetPeering
		return ret, false
	}
	return *o.NetPeerings, true
}

// HasNetPeerings returns a boolean if a field has been set.
func (o *ReadNetPeeringsResponse) HasNetPeerings() bool {
	if o != nil && o.NetPeerings != nil {
		return true
	}

	return false
}

// SetNetPeerings gets a reference to the given []NetPeering and assigns it to the NetPeerings field.
func (o *ReadNetPeeringsResponse) SetNetPeerings(v []NetPeering) {
	o.NetPeerings = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadNetPeeringsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadNetPeeringsResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadNetPeeringsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadNetPeeringsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableReadNetPeeringsResponse struct {
	Value        ReadNetPeeringsResponse
	ExplicitNull bool
}

func (v NullableReadNetPeeringsResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableReadNetPeeringsResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
