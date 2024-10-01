// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "stori/internal/transactions/domain"

	mock "github.com/stretchr/testify/mock"

	summarydomain "stori/internal/summary/domain"
)

// CreateSummaryUC is an autogenerated mock type for the CreateSummaryUC type
type CreateSummaryUC struct {
	mock.Mock
}

// Exec provides a mock function with given fields: ctx, transactions
func (_m *CreateSummaryUC) Exec(ctx context.Context, transactions []domain.Transaction) summarydomain.Summary {
	ret := _m.Called(ctx, transactions)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 summarydomain.Summary
	if rf, ok := ret.Get(0).(func(context.Context, []domain.Transaction) summarydomain.Summary); ok {
		r0 = rf(ctx, transactions)
	} else {
		r0 = ret.Get(0).(summarydomain.Summary)
	}

	return r0
}

// NewCreateSummaryUC creates a new instance of CreateSummaryUC. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCreateSummaryUC(t interface {
	mock.TestingT
	Cleanup(func())
}) *CreateSummaryUC {
	mock := &CreateSummaryUC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
