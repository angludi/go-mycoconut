package models

import (
	"time"

	"my-coconut.com/web/db"
)

type Produk struct {
	ID         uint32    `json:"id"`
	NamaProduk string    `json:"nama_produk"`
	Deskripsi  string    `json:"deskripsi"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

func GetProducts() ([]Produk, error) {
	var (
		produk []Produk
		err error
	)

	tx := gorm.MysqlConn().Begin()

	if err = tx.Find(&produk).Error; err != nil {
		tx.Rollback()
		return produk, err
	}

	tx.Commit()
	return produk, err
}

func (Produk) TableName() string {
	return "Produk"
}