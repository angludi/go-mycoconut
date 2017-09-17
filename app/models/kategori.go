package models


import (
	"my-coconut.com/web/db"
)

type Kategori struct {
	ID 	uint32	`json:"id"`
	Kategori	string 	`json:"kategori"`
}


type ProdukKategori struct {
	ID 	uint32 `json:"id"`
	IdProduk uint32 `json:"id_produk" gorm:"foreign_key:id_produk"`
	IdKategori uint32 `json:"id_kategori"`
}


func GetAllKategori() ([]ProdukKategori, error) {
	var (
		produkKategori []ProdukKategori
		err error
	)

	tx := gorm.MysqlConn().Begin()

	//if err = tx.Find(&produkKategori).Error; err != nil {

	if err = tx.Joins("JOIN kategori on kategori.id = produk_kategori.id_kategori").Preload("Kategori").Find(&produkKategori).Error; err != nil {
				tx.Rollback()
		return produkKategori, err
	}

	tx.Commit()
	return produkKategori, err
}

func (ProdukKategori) TableName() string {
	return "produk_kategori"
}

func (Kategori) TableName() string {
	return "kategori"
}