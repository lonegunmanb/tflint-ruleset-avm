package attrvalue

import (
	"reflect"

	"github.com/zclconf/go-cty/cty"
)

// NewSimpleStringRule is syntactic sugar, returning a new SimpleRule with the type values set for integer
// and the given resource type, attribute name, and expected values.
func NewSimpleIntRule(resourceType string, attributeName string, expectedValues []any) *SimpleRule {
	return &SimpleRule{
		resourceType:   resourceType,
		attributeName:  attributeName,
		expectedValues: expectedValues,
		ctyType:        cty.Number,
		reflectType:    reflect.TypeOf(0),
	}
}
