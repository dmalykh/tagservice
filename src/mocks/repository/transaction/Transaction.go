// Code generated by mockery v2.12.0. DO NOT EDIT.

package mocks

import (
	context "context"
	repository "github.com/dmalykh/tagservice/tagservice/repository"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Transaction is an autogenerated mock type for the Transaction type
type Transaction struct {
	mock.Mock
}

// Category provides a mock function with given fields:
func (_m *Transaction) Category() repository.Category {
	ret := _m.Called()

	var r0 repository.Category
	if rf, ok := ret.Get(0).(func() repository.Category); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.Category)
		}
	}

	return r0
}

// Commit provides a mock function with given fields: ctx
func (_m *Transaction) Commit(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Namespace provides a mock function with given fields:
func (_m *Transaction) Namespace() repository.Namespace {
	ret := _m.Called()

	var r0 repository.Namespace
	if rf, ok := ret.Get(0).(func() repository.Namespace); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.Namespace)
		}
	}

	return r0
}

// Relation provides a mock function with given fields:
func (_m *Transaction) Relation() repository.Relation {
	ret := _m.Called()

	var r0 repository.Relation
	if rf, ok := ret.Get(0).(func() repository.Relation); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.Relation)
		}
	}

	return r0
}

// Rollback provides a mock function with given fields: ctx
func (_m *Transaction) Rollback(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Tag provides a mock function with given fields:
func (_m *Transaction) Tag() repository.Tag {
	ret := _m.Called()

	var r0 repository.Tag
	if rf, ok := ret.Get(0).(func() repository.Tag); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.Tag)
		}
	}

	return r0
}

// NewTransaction creates a new instance of Transaction. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransaction(t testing.TB) *Transaction {
	mock := &Transaction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
