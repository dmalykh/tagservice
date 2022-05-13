// Code generated by mockery v2.12.0. DO NOT EDIT.

package mocks

import (
	context "context"
	model "tagservice/server/model"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Namespace is an autogenerated mock type for the Namespace type
type Namespace struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, name
func (_m *Namespace) Create(ctx context.Context, name string) (model.Namespace, error) {
	ret := _m.Called(ctx, name)

	var r0 model.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Namespace); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(model.Namespace)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *Namespace) DeleteById(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetById provides a mock function with given fields: ctx, id
func (_m *Namespace) GetById(ctx context.Context, id uint64) (model.Namespace, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, uint64) model.Namespace); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Namespace)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *Namespace) GetByName(ctx context.Context, name string) (model.Namespace, error) {
	ret := _m.Called(ctx, name)

	var r0 model.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Namespace); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(model.Namespace)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetList provides a mock function with given fields: ctx, limit, offset
func (_m *Namespace) GetList(ctx context.Context, limit uint64, offset uint64) ([]model.Namespace, error) {
	ret := _m.Called(ctx, limit, offset)

	var r0 []model.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64) []model.Namespace); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, name
func (_m *Namespace) Update(ctx context.Context, id uint64, name string) (model.Namespace, error) {
	ret := _m.Called(ctx, id, name)

	var r0 model.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, uint64, string) model.Namespace); ok {
		r0 = rf(ctx, id, name)
	} else {
		r0 = ret.Get(0).(model.Namespace)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, string) error); ok {
		r1 = rf(ctx, id, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewNamespace creates a new instance of Namespace. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewNamespace(t testing.TB) *Namespace {
	mock := &Namespace{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
