package transformer

import (

	"my-coconut.com/web/app/models"
	"time"
)

type ProdukCollectionTransformer struct {
	Data []interface{} `json:"data"`
	Meta *models.Meta  `json:"meta"`
}

type ProdukTransformer struct {
	Data interface{}  `json:"data"`
	Meta *models.Meta `json:"meta"`
}

type ProdukData struct {
	ID         uint32    `json:"id" gorm:"primary_key"`
	NamaProduk string    `json:"nama_produk"`
	Deskripsi  string    `json:"deskripsi"`
	Kategori   []models.Kategori
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

func (res *ProdukTransformer) Transform(produk *models.Produk) {
	res.Data = assignProduk(produk)
	res.Meta = nil
}

func (res *ProdukCollectionTransformer) Transform(products []models.Produk) {
	for _, produk := range products {
		data := assignProduk(&produk)
		res.Data = append(res.Data, data)
	}
	//res.Meta = &models.Meta{Pagination: pagination}
	res.Meta = nil
}

func assignProduk(produk *models.Produk) interface{} {
	m := new(ProdukData)
	m.ID = produk.ID
	m.NamaProduk = produk.NamaProduk
	m.Deskripsi = produk.Deskripsi
	m.Kategori = produk.Kategori
	m.CreatedAt = produk.CreatedAt
	m.UpdatedAt = produk.UpdatedAt
	return m
}

