package kitchen

import (
	"bytes"
	"encoding/json"
	"github.com/golang/protobuf/jsonpb"
)

// SinkJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Sink. This struct is safe to replace or modify but
// should not be done so concurrently.
var SinkJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Sink) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := SinkJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Sink)(nil)

// SinkJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Sink. This struct is safe to replace or modify but
// should not be done so concurrently.
var SinkJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Sink) UnmarshalJSON(b []byte) error {
	return SinkJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Sink)(nil)

// Sink_MaterialJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Sink_Material. This struct is safe to replace or modify but
// should not be done so concurrently.
var Sink_MaterialJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Sink_Material) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := Sink_MaterialJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Sink_Material)(nil)

// Sink_MaterialJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Sink_Material. This struct is safe to replace or modify but
// should not be done so concurrently.
var Sink_MaterialJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Sink_Material) UnmarshalJSON(b []byte) error {
	return Sink_MaterialJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Sink_Material)(nil)
