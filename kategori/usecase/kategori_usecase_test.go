package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mpnj-api/domain"
	"mpnj-api/domain/mocks"
	"testing"
)

func Test_GetKategori(t *testing.T) {
	mockKategoriRepo := new(mocks.KategoriRepository)
	mockKategori := domain.Kategori{
		IDKategoriProduk: 1,
		NamaKategori: "Elektronik",
	}
	var ctx context.Context

	t.Run("success", func(t *testing.T) {
		mockKategoriRepo.On("Get", mock.Anything).Return([]domain.Kategori{mockKategori}, nil).Once()

		kategoriService := NewKategoriUseCase(mockKategoriRepo)

		result, _ := kategoriService.Get(ctx)

		assert.Equal(t, "Elektronik", result[0].NamaKategori)
		mockKategoriRepo.AssertExpectations(t)
	})
}

func Test_GetKategoriByID(t *testing.T)  {
	mockKategoriRepo := new(mocks.KategoriRepository)
	var ctx context.Context
	kategori := domain.Kategori{
		IDKategoriProduk: 1,
		NamaKategori: "IPA",
	}

	t.Run("success", func(t *testing.T) {
		mockKategoriRepo.On("GetByID", mock.Anything, mock.Anything).Return(kategori, nil)

		kategoriService := NewKategoriUseCase(mockKategoriRepo)

		kat, _ := kategoriService.GetByID(ctx, 1)

		mockKategoriRepo.AssertExpectations(t)

		assert.Equal(t, 1, kat.IDKategoriProduk)
		assert.Equal(t, "IPA", kat.NamaKategori)
	})
}

func Test_CreateKategori(t *testing.T)  {
	mockKategoriRepo := new(mocks.KategoriRepository)
	mockKategori := domain.Kategori{
		IDKategoriProduk: 2,
		NamaKategori: "Fashion",
	}
	var ctx context.Context
	t.Run("success", func(t *testing.T) {
		tempMockKategori := mockKategori
		mockKategoriRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Kategori")).Return(nil).Once()

		kategoriService := NewKategoriUseCase(mockKategoriRepo)

		err := kategoriService.Create(ctx, &mockKategori)

		assert.NoError(t, err)
		assert.Equal(t, tempMockKategori.NamaKategori, mockKategori.NamaKategori)
		mockKategoriRepo.AssertExpectations(t)
	})
}