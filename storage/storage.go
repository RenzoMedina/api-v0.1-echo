package storage

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

const (
	dns = "usi1p2fhaqjrbcxd:7b3gioYphTLYaZoRbF8P@tcp(byr2cqsjyxldxj3oabah-mysql.services.clever-cloud.com)/byr2cqsjyxldxj3oabah"
)

func ConnectionDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
		if err != nil {
			log.Fatalf("Can't open database %v", err)
		}
		fmt.Println("Conecction succesfully!!")
	})
}

func DB() *gorm.DB {
	return db
}
