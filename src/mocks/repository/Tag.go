// Code generated by mockery v2.12.0. DO NOT EDIT.

package mocks

import (
	context "context"
	model "tagservice/server/model"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Tag is an autogenerated mock type for the Tag type
type Tag struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, data
func (_m *Tag) Create(ctx context.Context, data *model.TagData) (model.Tag, error) {
	ret := _m.Called(ctx, data)

	var r0 model.Tag
	if rf, ok := ret.Get(0).(func(context.Context, *model.TagData) model.Tag); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(model.Tag)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.TagData) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *Tag) DeleteById(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByFilter provides a mock function with given fields: ctx, filter, limit, offset
func (_m *Tag) GetByFilter(ctx context.Context, filter model.TagFilter, limit uint, offset uint) ([]model.Tag, error) {
	ret := _m.Called(ctx, filter, limit, offset)

	var r0 []model.Tag
	if rf, ok := ret.Get(0).(func(context.Context, model.TagFilter, uint, uint) []model.Tag); ok {
		r0 = rf(ctx, filter, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.TagFilter, uint, uint) error); ok {
		r1 = rf(ctx, filter, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *Tag) GetById(ctx context.Context, id uint) (model.Tag, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Tag
	if rf, ok := ret.Get(0).(func(context.Context, uint) model.Tag); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Tag)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *Tag) GetByName(ctx context.Context, name string) ([]model.Tag, error) {
	ret := _m.Called(ctx, name)

	var r0 []model.Tag
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.Tag); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, data
func (_m *Tag) Update(ctx context.Context, id uint, data *model.TagData) (model.Tag, error) {
	ret := _m.Called(ctx, id, data)

	var r0 model.Tag
	if rf, ok := ret.Get(0).(func(context.Context, uint, *model.TagData) model.Tag); ok {
		r0 = rf(ctx, id, data)
	} else {
		r0 = ret.Get(0).(model.Tag)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, *model.TagData) error); ok {
		r1 = rf(ctx, id, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTag creates a new instance of Tag. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewTag(t testing.TB) *Tag {
	mock := &Tag{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
