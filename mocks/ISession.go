// Copyright (c) 2019 MindStand Technologies, Inc
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import go_cypherdsl "github.com/mindstand/go-cypherdsl"
import gogm "github.com/mindstand/gogm"
import mock "github.com/stretchr/testify/mock"

// ISession is an autogenerated mock type for the ISession type
type ISession struct {
	mock.Mock
}

// Begin provides a mock function with given fields:
func (_m *ISession) Begin() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *ISession) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Commit provides a mock function with given fields:
func (_m *ISession) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: deleteObj
func (_m *ISession) Delete(deleteObj interface{}) error {
	ret := _m.Called(deleteObj)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(deleteObj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUUID provides a mock function with given fields: uuid
func (_m *ISession) DeleteUUID(uuid string) error {
	ret := _m.Called(uuid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Load provides a mock function with given fields: respObj, id
func (_m *ISession) Load(respObj interface{}, id string) error {
	ret := _m.Called(respObj, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string) error); ok {
		r0 = rf(respObj, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadAll provides a mock function with given fields: respObj
func (_m *ISession) LoadAll(respObj interface{}) error {
	ret := _m.Called(respObj)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(respObj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadAllDepth provides a mock function with given fields: respObj, depth
func (_m *ISession) LoadAllDepth(respObj interface{}, depth int) error {
	ret := _m.Called(respObj, depth)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, int) error); ok {
		r0 = rf(respObj, depth)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadAllDepthFilter provides a mock function with given fields: respObj, depth, filter, params
func (_m *ISession) LoadAllDepthFilter(respObj interface{}, depth int, filter go_cypherdsl.ConditionOperator, params map[string]interface{}) error {
	ret := _m.Called(respObj, depth, filter, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, int, go_cypherdsl.ConditionOperator, map[string]interface{}) error); ok {
		r0 = rf(respObj, depth, filter, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadAllDepthFilterPagination provides a mock function with given fields: respObj, depth, filter, params, pagination
func (_m *ISession) LoadAllDepthFilterPagination(respObj interface{}, depth int, filter go_cypherdsl.ConditionOperator, params map[string]interface{}, pagination *gogm.Pagination) error {
	ret := _m.Called(respObj, depth, filter, params, pagination)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, int, go_cypherdsl.ConditionOperator, map[string]interface{}, *gogm.Pagination) error); ok {
		r0 = rf(respObj, depth, filter, params, pagination)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadAllEdgeConstraint provides a mock function with given fields: respObj, endNodeType, endNodeField, edgeConstraint, minJumps, maxJumps, depth, filter
func (_m *ISession) LoadAllEdgeConstraint(respObj interface{}, endNodeType string, endNodeField string, edgeConstraint interface{}, minJumps int, maxJumps int, depth int, filter go_cypherdsl.ConditionOperator) error {
	ret := _m.Called(respObj, endNodeType, endNodeField, edgeConstraint, minJumps, maxJumps, depth, filter)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, string, interface{}, int, int, int, go_cypherdsl.ConditionOperator) error); ok {
		r0 = rf(respObj, endNodeType, endNodeField, edgeConstraint, minJumps, maxJumps, depth, filter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadDepth provides a mock function with given fields: respObj, id, depth
func (_m *ISession) LoadDepth(respObj interface{}, id string, depth int) error {
	ret := _m.Called(respObj, id, depth)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, int) error); ok {
		r0 = rf(respObj, id, depth)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadDepthFilter provides a mock function with given fields: respObj, id, depth, filter, params
func (_m *ISession) LoadDepthFilter(respObj interface{}, id string, depth int, filter *go_cypherdsl.ConditionBuilder, params map[string]interface{}) error {
	ret := _m.Called(respObj, id, depth, filter, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, int, *go_cypherdsl.ConditionBuilder, map[string]interface{}) error); ok {
		r0 = rf(respObj, id, depth, filter, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadDepthFilterPagination provides a mock function with given fields: respObj, id, depth, filter, params, pagination
func (_m *ISession) LoadDepthFilterPagination(respObj interface{}, id string, depth int, filter go_cypherdsl.ConditionOperator, params map[string]interface{}, pagination *gogm.Pagination) error {
	ret := _m.Called(respObj, id, depth, filter, params, pagination)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, int, go_cypherdsl.ConditionOperator, map[string]interface{}, *gogm.Pagination) error); ok {
		r0 = rf(respObj, id, depth, filter, params, pagination)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PurgeDatabase provides a mock function with given fields:
func (_m *ISession) PurgeDatabase() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Query provides a mock function with given fields: query, properties, respObj
func (_m *ISession) Query(query string, properties map[string]interface{}, respObj interface{}) error {
	ret := _m.Called(query, properties, respObj)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}, interface{}) error); ok {
		r0 = rf(query, properties, respObj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryRaw provides a mock function with given fields: query, properties
func (_m *ISession) QueryRaw(query string, properties map[string]interface{}) ([][]interface{}, error) {
	ret := _m.Called(query, properties)

	var r0 [][]interface{}
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) [][]interface{}); ok {
		r0 = rf(query, properties)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, map[string]interface{}) error); ok {
		r1 = rf(query, properties)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Rollback provides a mock function with given fields:
func (_m *ISession) Rollback() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RollbackWithError provides a mock function with given fields: err
func (_m *ISession) RollbackWithError(err error) error {
	ret := _m.Called(err)

	var r0 error
	if rf, ok := ret.Get(0).(func(error) error); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: saveObj
func (_m *ISession) Save(saveObj interface{}) error {
	ret := _m.Called(saveObj)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(saveObj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveDepth provides a mock function with given fields: saveObj, depth
func (_m *ISession) SaveDepth(saveObj interface{}, depth int) error {
	ret := _m.Called(saveObj, depth)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, int) error); ok {
		r0 = rf(saveObj, depth)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
