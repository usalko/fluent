// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"fmt"
	"math/big"
	"net/url"
	"strings"

	"github.com/usalko/fluent/dialect/gremlin"
)

// ExValueScan is the model entity for the ExValueScan schema.
type ExValueScan struct {
	config `json:"-"`
	// ID of the fluent.
	ID string `json:"id,omitempty"`
	// Binary holds the value of the "binary" field.
	Binary *url.URL `json:"binary,omitempty"`
	// BinaryBytes holds the value of the "binary_bytes" field.
	BinaryBytes *url.URL `json:"binary_bytes,omitempty"`
	// BinaryOptional holds the value of the "binary_optional" field.
	BinaryOptional *url.URL `json:"binary_optional,omitempty"`
	// Text holds the value of the "text" field.
	Text *big.Int `json:"text,omitempty"`
	// TextOptional holds the value of the "text_optional" field.
	TextOptional *big.Int `json:"text_optional,omitempty"`
	// Base64 holds the value of the "base64" field.
	Base64 string `json:"base64,omitempty"`
	// Custom holds the value of the "custom" field.
	Custom string `json:"custom,omitempty"`
	// CustomOptional holds the value of the "custom_optional" field.
	CustomOptional string `json:"custom_optional,omitempty"`
}

// FromResponse scans the gremlin response data into ExValueScan.
func (evs *ExValueScan) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanevs struct {
		ID             string   `json:"id,omitempty"`
		Binary         *url.URL `json:"binary,omitempty"`
		BinaryBytes    *url.URL `json:"binary_bytes,omitempty"`
		BinaryOptional *url.URL `json:"binary_optional,omitempty"`
		Text           *big.Int `json:"text,omitempty"`
		TextOptional   *big.Int `json:"text_optional,omitempty"`
		Base64         string   `json:"base64,omitempty"`
		Custom         string   `json:"custom,omitempty"`
		CustomOptional string   `json:"custom_optional,omitempty"`
	}
	if err := vmap.Decode(&scanevs); err != nil {
		return err
	}
	evs.ID = scanevs.ID
	evs.Binary = scanevs.Binary
	evs.BinaryBytes = scanevs.BinaryBytes
	evs.BinaryOptional = scanevs.BinaryOptional
	evs.Text = scanevs.Text
	evs.TextOptional = scanevs.TextOptional
	evs.Base64 = scanevs.Base64
	evs.Custom = scanevs.Custom
	evs.CustomOptional = scanevs.CustomOptional
	return nil
}

// Update returns a builder for updating this ExValueScan.
// Note that you need to call ExValueScan.Unwrap() before calling this method if this ExValueScan
// was returned from a transaction, and the transaction was committed or rolled back.
func (evs *ExValueScan) Update() *ExValueScanUpdateOne {
	return NewExValueScanClient(evs.config).UpdateOne(evs)
}

// Unwrap unwraps the ExValueScan entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (evs *ExValueScan) Unwrap() *ExValueScan {
	_tx, ok := evs.config.driver.(*txDriver)
	if !ok {
		panic("fluent: ExValueScan is not a transactional entity")
	}
	evs.config.driver = _tx.drv
	return evs
}

// String implements the fmt.Stringer.
func (evs *ExValueScan) String() string {
	var builder strings.Builder
	builder.WriteString("ExValueScan(")
	builder.WriteString(fmt.Sprintf("id=%v, ", evs.ID))
	builder.WriteString("binary=")
	builder.WriteString(fmt.Sprintf("%v", evs.Binary))
	builder.WriteString(", ")
	builder.WriteString("binary_bytes=")
	builder.WriteString(fmt.Sprintf("%v", evs.BinaryBytes))
	builder.WriteString(", ")
	builder.WriteString("binary_optional=")
	builder.WriteString(fmt.Sprintf("%v", evs.BinaryOptional))
	builder.WriteString(", ")
	builder.WriteString("text=")
	builder.WriteString(fmt.Sprintf("%v", evs.Text))
	builder.WriteString(", ")
	builder.WriteString("text_optional=")
	builder.WriteString(fmt.Sprintf("%v", evs.TextOptional))
	builder.WriteString(", ")
	builder.WriteString("base64=")
	builder.WriteString(evs.Base64)
	builder.WriteString(", ")
	builder.WriteString("custom=")
	builder.WriteString(evs.Custom)
	builder.WriteString(", ")
	builder.WriteString("custom_optional=")
	builder.WriteString(evs.CustomOptional)
	builder.WriteByte(')')
	return builder.String()
}

// ExValueScans is a parsable slice of ExValueScan.
type ExValueScans []*ExValueScan

// FromResponse scans the gremlin response data into ExValueScans.
func (evs *ExValueScans) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanevs []struct {
		ID             string   `json:"id,omitempty"`
		Binary         *url.URL `json:"binary,omitempty"`
		BinaryBytes    *url.URL `json:"binary_bytes,omitempty"`
		BinaryOptional *url.URL `json:"binary_optional,omitempty"`
		Text           *big.Int `json:"text,omitempty"`
		TextOptional   *big.Int `json:"text_optional,omitempty"`
		Base64         string   `json:"base64,omitempty"`
		Custom         string   `json:"custom,omitempty"`
		CustomOptional string   `json:"custom_optional,omitempty"`
	}
	if err := vmap.Decode(&scanevs); err != nil {
		return err
	}
	for _, v := range scanevs {
		node := &ExValueScan{ID: v.ID}
		node.Binary = v.Binary
		node.BinaryBytes = v.BinaryBytes
		node.BinaryOptional = v.BinaryOptional
		node.Text = v.Text
		node.TextOptional = v.TextOptional
		node.Base64 = v.Base64
		node.Custom = v.Custom
		node.CustomOptional = v.CustomOptional
		*evs = append(*evs, node)
	}
	return nil
}
