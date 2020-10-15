// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	context "context"

	staffin "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"

	mock "github.com/stretchr/testify/mock"

	out "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/out"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, input
func (_m *Service) Create(ctx context.Context, input *staffin.CreateInput) (string, error) {
	ret := _m.Called(ctx, input)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *staffin.CreateInput) string); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *staffin.CreateInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, input
func (_m *Service) Delete(ctx context.Context, input *staffin.DeleteInput) error {
	ret := _m.Called(ctx, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *staffin.DeleteInput) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: ctx, opt
func (_m *Service) List(ctx context.Context, opt *domain.PageOption) (int, []*out.StaffView, error) {
	ret := _m.Called(ctx, opt)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *domain.PageOption) int); ok {
		r0 = rf(ctx, opt)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 []*out.StaffView
	if rf, ok := ret.Get(1).(func(context.Context, *domain.PageOption) []*out.StaffView); ok {
		r1 = rf(ctx, opt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*out.StaffView)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *domain.PageOption) error); ok {
		r2 = rf(ctx, opt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Read provides a mock function with given fields: ctx, input
func (_m *Service) Read(ctx context.Context, input *staffin.ReadInput) (*out.StaffView, error) {
	ret := _m.Called(ctx, input)

	var r0 *out.StaffView
	if rf, ok := ret.Get(0).(func(context.Context, *staffin.ReadInput) *out.StaffView); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*out.StaffView)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *staffin.ReadInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, input
func (_m *Service) Update(ctx context.Context, input *staffin.UpdateInput) error {
	ret := _m.Called(ctx, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *staffin.UpdateInput) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
