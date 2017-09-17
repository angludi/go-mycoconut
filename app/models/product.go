package models

import (
	"time"

	"my-coconut.com/web/db"
)

type Produk struct {
	ID         uint32    `json:"id" gorm:"primary_key"`
	NamaProduk string    `json:"nama_produk"`
	Deskripsi  string    `json:"deskripsi"`
	Kategori []Kategori	`jzon:"kategori" gorm:"many2many:produk_kategori"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

func GetProducts() ([]Produk, error) {
	var (
		produk []Produk
		err error
	)

	tx := gorm.MysqlConn().Begin()

	if err = tx.Joins("JOIN produk_kategori pk ON pk.produk_id = produk.id").
				Joins("JOIN kategori k ON k.id = pk.kategori_id").
				Group("produk.id").
				Preload("Kategori").
				Find(&produk).
				Error; err != nil {
		tx.Rollback()
		return produk, err
	}

	tx.Commit()
	return produk, err
}

func (Produk) TableName() string {
	return "Produk"
}