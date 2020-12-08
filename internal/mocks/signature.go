// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import internal "github.com/sucsessyan/maroto/internal"
import mock "github.com/stretchr/testify/mock"
import props "github.com/sucsessyan/maroto/pkg/props"

// Signature is an autogenerated mock type for the Signature type
type Signature struct {
	mock.Mock
}

// AddSpaceFor provides a mock function with given fields: label, cell, textProp
func (_m *Signature) AddSpaceFor(label string, cell internal.Cell, textProp props.Text) {
	_m.Called(label, cell, textProp)
}
