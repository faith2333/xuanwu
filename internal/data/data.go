package data

import (
	"fmt"
	"github/faith2333/xuanwu/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	log *log.Helper
	db  *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	data := &Data{
		log: log.NewHelper(logger),
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	var err error
	data.db, err = newDB(c)
	if err != nil {
		return nil, nil, err
	}

	return data, cleanup, nil
}

func newDB(c *conf.Data) (db *gorm.DB, err error) {
	switch c.Database.Driver {
	case "mysql":
		return gorm.Open(mysql.Open(c.Database.Source))
	case "postgres":
		return gorm.Open(postgres.Open(c.Database.Source))
	default:
		return nil, fmt.Errorf("Database Driver %s not supported ", c.Database.Driver)
	}
}
