package database

import (
        "gorm.io/driver/sqlite"
        "gorm.io/gorm"
        "product-info/database/model"
        log "github.com/sirupsen/logrus"
)

var dbClient *gorm.DB
type product model.Product

func InitDB() *gorm.DB {

        db, err := gorm.Open(sqlite.Open("product.db"), &gorm.Config{})
        if err != nil {
                log.Println("failed to connect database")
        }

        dbClient = db
        return dbClient
}

func GetDB() *gorm.DB {
        dbClient = InitDB()
        return dbClient
}

func StartMigration() error {
        db := GetDB()
        if err := db.Table("products").AutoMigrate(&product{}); err != nil {
                log.Println("failed to create DB")
        }
        return nil
}

