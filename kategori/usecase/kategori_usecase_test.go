package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"mpnj-api/domain"
	"mpnj-api/domain/mock"
	"testing"
)

func Test_GetKategori(t *testing.T) {
	mockKategori := new(mock.KategoriRepository)

	var ctx context.Context
	kategori := domain.Kategori{
		IDKategoriProduk: 1,
		NamaKategori:     "IPA",
	}

	mockKategori.On("Get", ctx).Return([]domain.Kategori{kategori}, nil)

	kategoriService := NewKategoriUseCase(mockKategori)

	result, _ := kategoriService.Get(ctx)

	mockKategori.AssertExpectations(t)

	assert.Equal(t, "IPA", result[0].NamaKategori)
}

func Test_GetKategoriByID(t *testing.T)  {
	mockKategori := new(mock.KategoriRepository)
	var ctx context.Context
	kategori := domain.Kategori{
		IDKategoriProduk: 1,
		NamaKategori: "IPA",
	}

	mockKategori.On("GetByID", ctx, kategori.IDKategoriProduk).Return(kategori, nil)

	kategoriService := NewKategoriUseCase(mockKategori)

	kat, _ := kategoriService.GetByID(ctx, 1)

	mockKategori.AssertExpectations(t)

	assert.Equal(t, 1, kat.IDKategoriProduk)
	assert.Equal(t, "IPA", kat.NamaKategori)
}