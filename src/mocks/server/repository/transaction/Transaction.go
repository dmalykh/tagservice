// Code generated by mockery v2.12.0. DO NOT EDIT.

package mocks

import (
	context "context"
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	transaction "tagservice/server/repository/transaction"
)

// Transaction is an autogenerated mock type for the Transaction type
type Transaction struct {
	mock.Mock
}

// BeginTx provides a mock function with given fields: ctx, opts
func (_m *Transaction) BeginTx(ctx context.Context, opts ...*transaction.TxOptions) (transaction.Transaction, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 transaction.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, ...*transaction.TxOptions) transaction.Transaction); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(transaction.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...*transaction.TxOptions) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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

// NewTransaction creates a new instance of Transaction. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransaction(t testing.TB) *Transaction {
	mock := &Transaction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}