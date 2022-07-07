package manager

import (
	"go-api-with-gin2/config"
	"go-api-with-gin2/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// infra -> sebagai database penyimpanan pengganti slice
type Infra interface {
	SqlDb() *gorm.DB
}

type infra struct {
	db *gorm.DB
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func NewInfra(config *config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &infra{db: resource}
}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	db.AutoMigrate(
		&model.Menu{},
		&model.MenuPrice{},
		&model.BillDetail{},
		&model.Bill{},
		&model.Table{},
		&model.TransType{},
		&model.Customer{},
		&model.Discount{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
