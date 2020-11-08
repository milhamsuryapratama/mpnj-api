package tests

import (
	"context"
	"mpnj-api/domain"
	"mpnj-api/usecase"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

type KategoriRepository struct {
	mock.Mock
}

// GetKategori ...
func (k *KategoriRepository) Get(ctx context.Context) ([]domain.Kategori, error) {
	args := k.Called()
	result := args.Get(0)

	return result.([]domain.Kategori), args.Error(1)
}

func (k *KategoriRepository) Create(ctx context.Context, kategori *domain.Kategori) (domain.Kategori, error) {
	args := k.Called()
	result := args.Get(0)

	return result.(domain.Kategori), args.Error(1)
}

func (k *KategoriRepository) Delete(ctx context.Context, id int) error {
	ret := k.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (k *KategoriRepository) GetByID(ctx context.Context, id int) (domain.Kategori, error) {
	args := k.Called()
	result := args.Get(0)

	return result.(domain.Kategori), args.Error(1)
}

func (k *KategoriRepository) Update(kategori *domain.Kategori, id int) (domain.Kategori, error) {
	args := k.Called()
	result := args.Get(0)

	return result.(domain.Kategori), args.Error(1)
}

func Test_GetKategori(t *testing.T) {
	mockKategori := new(KategoriRepository)

	var ctx context.Context
	kategori := domain.Kategori{
		IDKategoriProduk: 1,
		NamaKategori:     "IPA",
	}

	mockKategori.On("Get").Return([]domain.Kategori{kategori}, nil)

	kategoriService := usecase.NewKategoriUseCase(mockKategori)

	result, _ := kategoriService.Get(ctx)

	mockKategori.AssertExpectations(t)

	assert.Equal(t, "IPA", result[0].NamaKategori)
}

func Test_CreateKategori(t *testing.T) {
	mockKategori := new(KategoriRepository)

	var ctx context.Context
	kategori := domain.Kategori{
		IDKategoriProduk: 1,
		NamaKategori:     "IPA",
	}

	mockKategori.On("Create").Return(kategori, nil)

	kategoriService := usecase.NewKategoriUseCase(mockKategori)

	kat, _ := kategoriService.Create(ctx, &kategori)

	mockKategori.AssertExpectations(t)

	// assert.Equal(t, nil, err.Error())
	assert.Equal(t, "IPA", kat.NamaKategori)
}

func Test_GetKategoriByID(t *testing.T) {
	mockKategori := new(KategoriRepository)

	var ctx context.Context
	kategori := domain.Kategori{
		IDKategoriProduk: 1,
		NamaKategori:     "IPA",
	}

	mockKategori.On("GetByID").Return(kategori, nil)

	kategoriService := usecase.NewKategoriUseCase(mockKategori)

	kat, _ := kategoriService.GetByID(ctx, 1)

	mockKategori.AssertExpectations(t)

	assert.Equal(t, 1, kat.IDKategoriProduk)
	assert.Equal(t, "IPA", kat.NamaKategori)
}

func Test_UpdateKategeori(t *testing.T) {
	mockKategori := new(KategoriRepository)

	var ctx context.Context
	kategori := domain.Kategori{
		IDKategoriProduk: 1,
		NamaKategori:     "IPA",
	}

	// mockKategori.On("Create").Return(kategori, nil)

	kategoriService := usecase.NewKategoriUseCase(mockKategori)

	// kat, _ := kategoriService.Create(ctx, &kategori)

	mockKategori.On("GetByID").Return(kategori, nil)

	get, _ := kategoriService.GetByID(ctx, 1)

	kategoriUpdate := domain.Kategori{
		IDKategoriProduk: get.IDKategoriProduk,
		NamaKategori:     "IPS",
	}

	mockKategori.On("Update").Return(kategoriUpdate, nil)

	update, _ := kategoriService.Update(&kategoriUpdate, 1)

	mockKategori.AssertExpectations(t)

	assert.Equal(t, "IPS", update.NamaKategori)
}

func Test_DeleteKategori(t *testing.T) {
	mockKategori := new(KategoriRepository)

	var ctx context.Context

	mockKategori.On("Delete").Return(nil)

	kategoriService := usecase.NewKategoriUseCase(mockKategori)

	err := kategoriService.Delete(ctx, 1)

	mockKategori.AssertExpectations(t)

	assert.NoError(t, err)
}
