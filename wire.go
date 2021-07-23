package main

import (
	"topics/database"
	"topics/product"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initDB(flag database.DBFlag) database.DatabaseEntry {
	wire.Build(database.ProvideDatabaseConfig, database.ProvideDatabaseEntry)
	return database.DatabaseEntry{}
}

func initProductAPI(db *gorm.DB) product.ProductAPI {
	wire.Build(product.ProvideProductRepostiory, product.ProvideProductService, product.ProvideProductAPI)
	return product.ProductAPI{}
}
