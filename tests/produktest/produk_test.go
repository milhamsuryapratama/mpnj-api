package produktest

import (
	"context"
	"github.com/magiconair/properties/assert"
	"mpnj-api/domain"
	"mpnj-api/usecase"
	"testing"
)

func Test_GetProduk(t *testing.T) {
	mockProdukRepository := new(ProdukRepository)
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

	var ctx context.Context
	mockProdukRepository.On("Get", ctx).Return([]domain.Produk{mockProduk}, nil)

	produkService := usecase.NewProdukUsecase(mockProdukRepository)

	result, _ := produkService.Get(ctx)

	assert.Equal(t, "Produk 1", result[0].NamaProduk)
	assert.Equal(t, "Elektronik", result[0].Kategori.NamaKategori)
	assert.Equal(t, "Ilham", result[0].User.NamaLengkap)

	mockProdukRepository.AssertExpectations(t)
}

func Test_CreateProduk(t *testing.T) {
	mockProdukRepository := new(ProdukRepository)
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

	var ctx context.Context
	mockProdukRepository.On("Create", ctx, &mockProduk).Return(mockProduk, nil)

	produkService := usecase.NewProdukUsecase(mockProdukRepository)

	result, _ := produkService.Create(ctx, &mockProduk)

	assert.Equal(t, "Produk 1", result.NamaProduk)
	assert.Equal(t, "Elektronik", result.Kategori.NamaKategori)
	assert.Equal(t, "Ilham", result.User.NamaLengkap)

	mockProdukRepository.AssertExpectations(t)
}
