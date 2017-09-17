package models

import (
	"time"
	"my-coconut.com/web/db"
	"go/ast"
)

type Produk struct {
	ID         uint32    `json:"id"`
	NamaProduk string    `json:"nama_produk"`
	Deskripsi  string    `json:"deskripsi"`
	Kategori	[]ProdukKategori `json:"kategori"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Ticker    `json:"updated_at"`
}

type ProdukKategori struct {
	IdProduk uint32 `json:"id_produk"`
	IdKategori uint32 `json:"id_kategori"`
	Kategori	[]Kategori `json:"kategori"`
}

type Kategori struct {
	ID 	uint32	`json:"id"`
	Kategori	string 	`json:"kategori"`
}

func GetProducts() ([]Produk, error) {
	var (
		produk []Produk
		err error
	)


}