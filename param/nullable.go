package param

import "encoding/json"

// Nullable represents an optional value that can be null in JSON.
// Used for API fields that distinguish between null, missing, and present values.
type Nullable[T any] struct {
	Value  T
	IsNull bool
}

// MarshalJSON implements json.Marshaler to serialize null or the value.
// Returns "null" if IsNull is true, otherwise marshals the Value.
func (n *Nullable[T]) MarshalJSON() ([]byte, error) {
	if n.IsNull {
		return []byte("null"), nil
	}
	return json.Marshal(n.Value)
}

// UnmarshalJSON implements json.Unmarshaler to deserialize null or a value.
// Sets IsNull to true if the JSON is "null", otherwise unmarshals into Value.
func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.IsNull = true
		return nil
	}
	n.IsNull = false
	return json.Unmarshal(data, &n.Value)
}

// Null returns a *Nullable[T] with IsNull set to true, signalling an explicit JSON null.
func Null[T any]() *Nullable[T] { return &Nullable[T]{IsNull: true} }
