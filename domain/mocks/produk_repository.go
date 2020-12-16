// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "mpnj-api/domain"

	mock "github.com/stretchr/testify/mock"
)

// ProdukRepository is an autogenerated mock type for the ProdukRepository type
type ProdukRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, p
func (_m *ProdukRepository) Create(ctx context.Context, p *domain.Produk) error {
	ret := _m.Called(ctx, p)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Produk) error); ok {
		r0 = rf(ctx, p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx
func (_m *ProdukRepository) Get(ctx context.Context) ([]domain.Produk, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Produk
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Produk); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Produk)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *ProdukRepository) GetByID(ctx context.Context, id int) (domain.Produk, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Produk
	if rf, ok := ret.Get(0).(func(context.Context, int) domain.Produk); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Produk)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: produk, id
func (_m *ProdukRepository) Update(produk *domain.Produk, id int) (domain.Produk, error) {
	ret := _m.Called(produk, id)

	var r0 domain.Produk
	if rf, ok := ret.Get(0).(func(*domain.Produk, int) domain.Produk); ok {
		r0 = rf(produk, id)
	} else {
		r0 = ret.Get(0).(domain.Produk)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Produk, int) error); ok {
		r1 = rf(produk, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
