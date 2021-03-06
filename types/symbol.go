// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type SymbolValue struct {
	Value string
}

func (v *SymbolValue) ToString() string {
	return fmt.Sprint(v.Value)
}

func (v *SymbolValue) GetType() string {
	return "symbol"
}

func (v *SymbolValue) IsValue() bool {
	return false
}

func (v *SymbolValue) ConvertTo(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert symbol to %s: does not implement yet", t)
}

func (v *SymbolValue) EqualTo(t ValueType) bool {
	if v2, ok := t.(*SymbolValue); ok {
		if v2.Value == v.Value {
			return true
		}
	}
	return false
}

func (t *SymbolValue) GetFinalValue(s *Scope) (ValueType, error) {
	val, err := s.Get(t.Value)
	if err != nil {
		return nil, err
	}
	sym, ok := val.(*SymbolValue)
	if !ok {
		return val, nil
	}
	return sym.GetFinalValue(s)
}

func NewSymbolValue(v string) *SymbolValue {
	return &SymbolValue{Value: v}
}
