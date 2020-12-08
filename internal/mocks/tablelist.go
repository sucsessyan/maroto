// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import internal "github.com/sucsessyan/maroto/internal"
import mock "github.com/stretchr/testify/mock"
import props "github.com/sucsessyan/maroto/pkg/props"

// TableList is an autogenerated mock type for the TableList type
type TableList struct {
	mock.Mock
}

// BindGrid provides a mock function with given fields: part
func (_m *TableList) BindGrid(part internal.MarotoGridPart) {
	_m.Called(part)
}

// Create provides a mock function with given fields: header, contents, prop
func (_m *TableList) Create(header []string, contents [][]string, prop ...props.TableList) {
	_va := make([]interface{}, len(prop))
	for _i := range prop {
		_va[_i] = prop[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, header, contents)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}