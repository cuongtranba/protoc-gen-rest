package kitchen

import (
	"bytes"
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"
)

// KitchenJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Kitchen. This struct is safe to replace or modify but
// should not be done so concurrently.
var KitchenJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Kitchen) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}
	if err := KitchenJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Kitchen)(nil)

// KitchenJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Kitchen. This struct is safe to replace or modify but
// should not be done so concurrently.
var KitchenJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Kitchen) UnmarshalJSON(b []byte) error {
	return KitchenJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Kitchen)(nil)

// ColorJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Color. This struct is safe to replace or modify but
// should not be done so concurrently.
var ColorJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Color) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}
	if err := ColorJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Color)(nil)

// ColorJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Color. This struct is safe to replace or modify but
// should not be done so concurrently.
var ColorJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Color) UnmarshalJSON(b []byte) error {
	return ColorJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Color)(nil)

// SauteRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of SauteRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var SauteRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *SauteRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}
	if err := SauteRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*SauteRequest)(nil)

// SauteRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of SauteRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var SauteRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *SauteRequest) UnmarshalJSON(b []byte) error {
	return SauteRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*SauteRequest)(nil)

// SauteResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of SauteResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var SauteResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *SauteResponse) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}
	if err := SauteResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*SauteResponse)(nil)

// SauteResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of SauteResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var SauteResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *SauteResponse) UnmarshalJSON(b []byte) error {
	return SauteResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*SauteResponse)(nil)

// IceRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of IceRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var IceRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *IceRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}
	if err := IceRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*IceRequest)(nil)

// IceRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of IceRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var IceRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *IceRequest) UnmarshalJSON(b []byte) error {
	return IceRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*IceRequest)(nil)

// IceResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of IceResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var IceResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *IceResponse) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}
	if err := IceResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*IceResponse)(nil)

// IceResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of IceResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var IceResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *IceResponse) UnmarshalJSON(b []byte) error {
	return IceResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*IceResponse)(nil)

// GroceryItemJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of GroceryItem. This struct is safe to replace or modify but
// should not be done so concurrently.
var GroceryItemJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *GroceryItem) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}
	if err := GroceryItemJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*GroceryItem)(nil)

// GroceryItemJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of GroceryItem. This struct is safe to replace or modify but
// should not be done so concurrently.
var GroceryItemJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *GroceryItem) UnmarshalJSON(b []byte) error {
	return GroceryItemJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*GroceryItem)(nil)

// LoadSummaryJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of LoadSummary. This struct is safe to replace or modify but
// should not be done so concurrently.
var LoadSummaryJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *LoadSummary) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}
	if err := LoadSummaryJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*LoadSummary)(nil)

// LoadSummaryJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of LoadSummary. This struct is safe to replace or modify but
// should not be done so concurrently.
var LoadSummaryJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *LoadSummary) UnmarshalJSON(b []byte) error {
	return LoadSummaryJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*LoadSummary)(nil)

// DrinkOrderJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of DrinkOrder. This struct is safe to replace or modify but
// should not be done so concurrently.
var DrinkOrderJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *DrinkOrder) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}
	if err := DrinkOrderJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*DrinkOrder)(nil)

// DrinkOrderJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of DrinkOrder. This struct is safe to replace or modify but
// should not be done so concurrently.
var DrinkOrderJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *DrinkOrder) UnmarshalJSON(b []byte) error {
	return DrinkOrderJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*DrinkOrder)(nil)

// PreparedDrinkJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of PreparedDrink. This struct is safe to replace or modify but
// should not be done so concurrently.
var PreparedDrinkJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *PreparedDrink) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}
	if err := PreparedDrinkJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*PreparedDrink)(nil)

// PreparedDrinkJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of PreparedDrink. This struct is safe to replace or modify but
// should not be done so concurrently.
var PreparedDrinkJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *PreparedDrink) UnmarshalJSON(b []byte) error {
	return PreparedDrinkJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*PreparedDrink)(nil)

// Color_RGBJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Color_RGB. This struct is safe to replace or modify but
// should not be done so concurrently.
var Color_RGBJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Color_RGB) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}
	if err := Color_RGBJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Color_RGB)(nil)

// Color_RGBJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Color_RGB. This struct is safe to replace or modify but
// should not be done so concurrently.
var Color_RGBJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Color_RGB) UnmarshalJSON(b []byte) error {
	return Color_RGBJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Color_RGB)(nil)

// Color_CMYKJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Color_CMYK. This struct is safe to replace or modify but
// should not be done so concurrently.
var Color_CMYKJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Color_CMYK) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}
	if err := Color_CMYKJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Color_CMYK)(nil)

// Color_CMYKJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Color_CMYK. This struct is safe to replace or modify but
// should not be done so concurrently.
var Color_CMYKJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Color_CMYK) UnmarshalJSON(b []byte) error {
	return Color_CMYKJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Color_CMYK)(nil)
