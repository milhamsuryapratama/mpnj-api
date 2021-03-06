package usecase

import (
	"context"
	"errors"
	"mpnj-api/domain"
	"mpnj-api/domain/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetProduk(t *testing.T) {
	mockProdukRepository := new(mocks.ProdukRepository)
	mockProduk := domain.Produk{
		IDProduk:   1,
		NamaProduk: "Produk 1",
		Kategori: domain.Kategori{
			IDKategoriProduk: 1,
			NamaKategori:     "Elektronik",
		},
		Berat:            1000,
		Diskon:           0,
		HargaJual:        20000,
		HargaModal:       10000,
		KategoriProdukID: 1,
		Keterangan:       "Oke",
		Satuan:           "pcs",
		Slug:             "produk-1",
		Status:           "aktif",
		Stok:             10,
		Terjual:          0,
		TipeProduk:       "single",
		UserID:           1,
		Wishlist:         0,
	}

	t.Run("success", func(t *testing.T) {
		mockProdukRepository.On("Get", mock.Anything).Return([]domain.Produk{mockProduk}, nil).Once()

		produkService := NewProdukUsecase(mockProdukRepository)

		result, _ := produkService.Get(context.TODO())

		assert.Equal(t, "Produk 1", result[0].NamaProduk)
		assert.Equal(t, "Elektronik", result[0].Kategori.NamaKategori)

		mockProdukRepository.AssertExpectations(t)
	})
}

func Test_CreateProduk(t *testing.T) {
	mockProdukRepository := new(mocks.ProdukRepository)
	mockProduk := domain.Produk{
		IDProduk:   1,
		NamaProduk: "Produk 1",
		Kategori: domain.Kategori{
			IDKategoriProduk: 1,
			NamaKategori:     "Elektronik",
		},
		Berat:            1000,
		Diskon:           0,
		HargaJual:        20000,
		HargaModal:       10000,
		KategoriProdukID: 1,
		Keterangan:       "Oke",
		Satuan:           "pcs",
		Slug:             "produk-1",
		Status:           "aktif",
		Stok:             10,
		Terjual:          0,
		TipeProduk:       "single",
		UserID:           1,
		Wishlist:         0,
	}

	t.Run("success", func(t *testing.T) {
		tempMockProduk := mockProduk

		mockProdukRepository.On("Create", mock.Anything, mock.AnythingOfType("*domain.Produk")).Return(nil).Once()

		produkService := NewProdukUsecase(mockProdukRepository)

		err := produkService.Create(context.TODO(), &mockProduk)

		assert.NoError(t, err)
		assert.Equal(t, "Produk 1", tempMockProduk.NamaProduk)
		assert.Equal(t, "Elektronik", tempMockProduk.Kategori.NamaKategori)

		mockProdukRepository.AssertExpectations(t)
	})
}

func Test_UpdateProduk(t *testing.T) {
	mockProdukRepository := new(mocks.ProdukRepository)
	mockProduk := domain.Produk{
		IDProduk:   1,
		NamaProduk: "Produk 1",
		Kategori: domain.Kategori{
			IDKategoriProduk: 1,
			NamaKategori:     "Elektronik",
		},
		Berat:            1000,
		Diskon:           0,
		HargaJual:        20000,
		HargaModal:       10000,
		KategoriProdukID: 1,
		Keterangan:       "Oke",
		Satuan:           "pcs",
		Slug:             "produk-1",
		Status:           "aktif",
		Stok:             10,
		Terjual:          0,
		TipeProduk:       "single",
		UserID:           1,
		Wishlist:         0,
	}

	t.Run("success", func(t *testing.T) {
		mockProdukRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Once().Return(mockProduk, nil)
		mockProdukRepository.On("Update", &mockProduk, mockProduk.IDProduk).Once().Return(mockProduk, nil)

		produkService := NewProdukUsecase(mockProdukRepository)

		result, _ := produkService.Update(&mockProduk, mockProduk.IDProduk)

		assert.Equal(t, "Produk 1", result.NamaProduk)

		mockProdukRepository.AssertExpectations(t)
	})
}

func Test_DeleteProduk(t *testing.T) {
	mockProdukRepository := new(mocks.ProdukRepository)
	mockProduk := domain.Produk{
		IDProduk:   1,
		NamaProduk: "Produk 1",
		Kategori: domain.Kategori{
			IDKategoriProduk: 1,
			NamaKategori:     "Elektronik",
		},
		Berat:            1000,
		Diskon:           0,
		HargaJual:        20000,
		HargaModal:       10000,
		KategoriProdukID: 1,
		Keterangan:       "Oke",
		Satuan:           "pcs",
		Slug:             "produk-1",
		Status:           "aktif",
		Stok:             10,
		Terjual:          0,
		TipeProduk:       "single",
		UserID:           1,
		Wishlist:         0,
	}

	t.Run("not found produk", func(t *testing.T) {
		mockProdukRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Once().Return(domain.Produk{}, errors.New("Produk not fount"))

		produkService := NewProdukUsecase(mockProdukRepository)

		_, err := produkService.GetByID(context.TODO(), mockProduk.IDProduk)

		assert.Error(t, err)
		mockProdukRepository.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		mockProdukRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Once().Return(mockProduk, nil)

		mockProdukRepository.On("Delete", mock.Anything, mock.AnythingOfType("*domain.Produk"), mock.AnythingOfType("int")).Once().Return(nil)

		produkService := NewProdukUsecase(mockProdukRepository)

		err := produkService.Delete(context.TODO(), mockProduk.IDProduk)

		assert.NoError(t, err)
		mockProdukRepository.AssertExpectations(t)
	})
}
