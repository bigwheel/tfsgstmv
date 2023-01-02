package addrs

import "github.com/zclconf/go-cty/cty"

// InstanceKey represents the key of an instance within an object that
// contains multiple instances due to using "count" or "for_each" arguments
// in configuration.
//
// IntKey and StringKey are the two implementations of this type. No other
// implementations are allowed. The single instance of an object that _isn't_
// using "count" or "for_each" is represented by NoKey, which is a nil
// InstanceKey.
type InstanceKey interface {
	instanceKeySigil()
	String() string

	// Value returns the cty.Value of the appropriate type for the InstanceKey
	// value.
	Value() cty.Value
}
