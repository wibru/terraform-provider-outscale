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

// ReadImageExportTasksResponse struct for ReadImageExportTasksResponse
type ReadImageExportTasksResponse struct {
	// Information about one or more image export tasks.
	ImageExportTasks *[]ImageExportTask `json:"ImageExportTasks,omitempty"`
	ResponseContext  *ResponseContext   `json:"ResponseContext,omitempty"`
}

// GetImageExportTasks returns the ImageExportTasks field value if set, zero value otherwise.
func (o *ReadImageExportTasksResponse) GetImageExportTasks() []ImageExportTask {
	if o == nil || o.ImageExportTasks == nil {
		var ret []ImageExportTask
		return ret
	}
	return *o.ImageExportTasks
}

// GetImageExportTasksOk returns a tuple with the ImageExportTasks field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadImageExportTasksResponse) GetImageExportTasksOk() ([]ImageExportTask, bool) {
	if o == nil || o.ImageExportTasks == nil {
		var ret []ImageExportTask
		return ret, false
	}
	return *o.ImageExportTasks, true
}

// HasImageExportTasks returns a boolean if a field has been set.
func (o *ReadImageExportTasksResponse) HasImageExportTasks() bool {
	if o != nil && o.ImageExportTasks != nil {
		return true
	}

	return false
}

// SetImageExportTasks gets a reference to the given []ImageExportTask and assigns it to the ImageExportTasks field.
func (o *ReadImageExportTasksResponse) SetImageExportTasks(v []ImageExportTask) {
	o.ImageExportTasks = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadImageExportTasksResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadImageExportTasksResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadImageExportTasksResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadImageExportTasksResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableReadImageExportTasksResponse struct {
	Value        ReadImageExportTasksResponse
	ExplicitNull bool
}

func (v NullableReadImageExportTasksResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableReadImageExportTasksResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
