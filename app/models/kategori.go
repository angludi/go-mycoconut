package models


import (
	"my-coconut.com/web/db"
)

type Kategori struct {
	ID 	uint32	`json:"id"`
	Kategori	string 	`json:"kategori"`
}



func GetAllKategori() ([]Kategori, error) {
	var (
		kategori []Kategori
		err error
	)

	tx := gorm.MysqlConn().Begin()

	//if err = tx.Find(&produkKategori).Error; err != nil {

	if err = tx.Find(&kategori).Error; err != nil {
				tx.Rollback()
		return kategori, err
	}

	tx.Commit()
	return kategori, err
}

func GetKategoriByProductID(ProductID uint32) ([]Kategori, error) {
	var (
		kategori []Kategori
		err error
	)

	tx := gorm.MysqlConn().Begin()

	if err = tx.Joins("JOIN produk_kategori pk ON pk.kategori_id = kategori.id").
				Where("pk.produk_id=?", ProductID).
				Find(&kategori).
				Error; err != nil {
		tx.Rollback()
		return kategori, err
	}

	tx.Commit()
	return kategori, err
}


func (Kategori) TableName() string {
	return "kategori"
}