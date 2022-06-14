// Code generated by mockery v2.12.0. DO NOT EDIT.

package mocks

import (
	context "context"
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	transaction "github.com/dmalykh/tagservice/tagservice/repository/transaction"
)

// Transactioner is an autogenerated mock type for the Transactioner type
type Transactioner struct {
	mock.Mock
}

// BeginTx provides a mock function with given fields: ctx, opts
func (_m *Transactioner) BeginTx(ctx context.Context, opts ...*transaction.TxOptions) (transaction.Transaction, error) {
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

// NewTransactioner creates a new instance of Transactioner. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactioner(t testing.TB) *Transactioner {
	mock := &Transactioner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
