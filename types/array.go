// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import (
	"fmt"
	"strings"
)

type ArrayValue struct {
	Value []ValueType
}

func (v *ArrayValue) ToString() string {
	arr := make([]string, len(v.Value))
	for i, v2 := range v.Value {
		arr[i] = v2.ToString()
	}
	return "[" + strings.Join(arr, " ") + "]"
}

func (v *ArrayValue) GetType() string {
	return "array"
}

func (v *ArrayValue) IsValue() bool {
	return false
}

func (v *ArrayValue) ConvertTo(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert list to %s: does not implement yet", t)
}

func (v *ArrayValue) EqualTo(t ValueType) bool {
	if v2, ok := t.(*ArrayValue); ok {
		if &v2.Value == &v.Value {
			return true
		}
	}
	return false
}

func NewArrayValue(v []ValueType) *ArrayValue {
	return &ArrayValue{Value: v}
}
