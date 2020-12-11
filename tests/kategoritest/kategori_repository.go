// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package kategoritest

import (
	context "context"
	domain "mpnj-api/domain"

	mock "github.com/stretchr/testify/mock"
)

// KategoriRepository is an autogenerated mock type for the KategoriRepository type
type KategoriRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, k
func (_m *KategoriRepository) Create(ctx context.Context, k *domain.Kategori) error {
	ret := _m.Called(ctx, k)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Kategori) error); ok {
		r0 = rf(ctx, k)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *KategoriRepository) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx
func (_m *KategoriRepository) Get(ctx context.Context) ([]domain.Kategori, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Kategori
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Kategori); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Kategori)
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
func (_m *KategoriRepository) GetByID(ctx context.Context, id int) (domain.Kategori, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Kategori
	if rf, ok := ret.Get(0).(func(context.Context, int) domain.Kategori); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Kategori)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, kategori, id
func (_m *KategoriRepository) Update(ctx context.Context, kategori *domain.Kategori, id int) (domain.Kategori, error) {
	ret := _m.Called(ctx, kategori, id)

	var r0 domain.Kategori
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Kategori, int) domain.Kategori); ok {
		r0 = rf(ctx, kategori, id)
	} else {
		r0 = ret.Get(0).(domain.Kategori)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Kategori, int) error); ok {
		r1 = rf(ctx, kategori, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
