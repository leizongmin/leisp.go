// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type StringValue struct {
	Value string
}

func (v *StringValue) ToString() string {
	return "\"" + v.Value + "\""
}

func (v *StringValue) GetType() string {
	return "string"
}

func (v *StringValue) IsValue() bool {
	return true
}

func (v *StringValue) ConvertTo(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert string to %s: does not implement yet", t)
}

func (v *StringValue) EqualTo(t ValueType) bool {
	if v2, ok := t.(*StringValue); ok {
		if v2.Value == v.Value {
			return true
		}
	}
	return false
}

func NewStringValue(v string) *StringValue {
	return &StringValue{Value: v}
}
