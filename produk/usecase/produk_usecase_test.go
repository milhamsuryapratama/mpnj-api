package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mpnj-api/domain"
	"mpnj-api/domain/mocks"
	"testing"
)

func Test_GetProduk(t *testing.T) {
	mockProdukRepository := new(mocks.ProdukRepository)
	mockProduk := domain.Produk{
		IDProduk: 1,
		NamaProduk: "Produk 1",
		Kategori: domain.Kategori{
			IDKategoriProduk: 1,
			NamaKategori: "Elektronik",
		},
		Berat: 1000,
		Diskon: 0,
		HargaJual: 20000,
		HargaModal: 10000,
		KategoriProdukID: 1,
		Keterangan: "Oke",
		Satuan: "pcs",
		Slug: "produk-1",
		Status: "aktif",
		Stok: 10,
		Terjual: 0,
		TipeProduk: "single",
		User: domain.User{
			IDUser: 1,
			Email: "ilham@gmail.com",
			NamaLengkap: "Ilham",
			NomorHp: "085330150827",
			Password: "ilham",
			Username: "ilham",
		},
		UserID: 1,
		Wishlist: 0,
	}

	t.Run("success", func(t *testing.T) {
		mockProdukRepository.On("Get", mock.Anything).Return([]domain.Produk{mockProduk}, nil).Once()

		produkService := NewProdukUsecase(mockProdukRepository)

		result, _ := produkService.Get(context.TODO())

		assert.Equal(t, "Produk 1", result[0].NamaProduk)
		assert.Equal(t, "Elektronik", result[0].Kategori.NamaKategori)
		assert.Equal(t, "Ilham", result[0].User.NamaLengkap)

		mockProdukRepository.AssertExpectations(t)
	})
}

func Test_CreateProduk(t *testing.T) {
	mockProdukRepository := new(mocks.ProdukRepository)
	mockProduk := domain.Produk{
		IDProduk: 1,
		NamaProduk: "Produk 1",
		Kategori: domain.Kategori{
			IDKategoriProduk: 1,
			NamaKategori: "Elektronik",
		},
		Berat: 1000,
		Diskon: 0,
		HargaJual: 20000,
		HargaModal: 10000,
		KategoriProdukID: 1,
		Keterangan: "Oke",
		Satuan: "pcs",
		Slug: "produk-1",
		Status: "aktif",
		Stok: 10,
		Terjual: 0,
		TipeProduk: "single",
		User: domain.User{
			IDUser: 1,
			Email: "ilham@gmail.com",
			NamaLengkap: "Ilham",
			NomorHp: "085330150827",
			Password: "ilham",
			Username: "ilham",
		},
		UserID: 1,
		Wishlist: 0,
	}

	t.Run("success", func(t *testing.T) {
		tempMockProduk := mockProduk

		mockProdukRepository.On("Create", mock.Anything, mock.AnythingOfType("*domain.Produk")).Return(nil).Once()

		produkService := NewProdukUsecase(mockProdukRepository)

		err := produkService.Create(context.TODO(), &mockProduk)

		assert.NoError(t, err)
		assert.Equal(t, "Produk 1", tempMockProduk.NamaProduk)
		assert.Equal(t, "Elektronik", tempMockProduk.Kategori.NamaKategori)
		assert.Equal(t, "Ilham", tempMockProduk.User.NamaLengkap)

		mockProdukRepository.AssertExpectations(t)
	})
}

//func Test_UpdateProduk(t *testing.T) {
//	mockProdukRepository := new(mocks.ProdukRepository)
//	mockProduk := domain.Produk{
//		IDProduk: 1,
//		NamaProduk: "Produk 1",
//		Kategori: domain.Kategori{
//			IDKategoriProduk: 1,
//			NamaKategori: "Elektronik",
//		},
//		Berat: 1000,
//		Diskon: 0,
//		HargaJual: 20000,
//		HargaModal: 10000,
//		KategoriProdukID: 1,
//		Keterangan: "Oke",
//		Satuan: "pcs",
//		Slug: "produk-1",
//		Status: "aktif",
//		Stok: 10,
//		Terjual: 0,
//		TipeProduk: "single",
//		User: domain.User{
//			IDUser: 1,
//			Email: "ilham@gmail.com",
//			NamaLengkap: "Ilham",
//			NomorHp: "085330150827",
//			Password: "ilham",
//			Username: "ilham",
//		},
//		UserID: 1,
//		Wishlist: 0,
//	}
//
//	var ctx context.Context
//	mockProdukRepository.On("GetByID", ctx, 1)
//	mockProdukRepository.On("Update", mockProduk, 1).Return(mockProduk, nil)
//
//	produkService := NewProdukUsecase(mockProdukRepository)
//
//	result, _ := produkService.Update(&mockProduk, 1)
//
//	assert.Equal(t, "Produk 1", result.NamaProduk)
//
//	mockProdukRepository.AssertExpectations(t)
//}