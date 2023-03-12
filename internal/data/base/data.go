package base

import (
	"context"
	"fmt"
	"github.com/faith2333/xuanwu/internal/biz"
	"github.com/faith2333/xuanwu/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewTransaction)

// Data Implemented biz.Transaction interface
type Data struct {
	log *log.Helper
	db  *gorm.DB
}

// NewData Constructor function for Data struct.
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	data := &Data{
		log: log.NewHelper(logger),
	}

	cleanup := func() {
		data.log.Info("closing the data resources")
	}

	var err error
	data.db, err = newDB(c)
	if err != nil {
		return nil, nil, err
	}

	return data, cleanup, nil
}

type dataTxKey struct{}

func NewTransaction(data *Data) biz.Transaction {
	return data
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

// ExecTx Execute fn with data.db, then set tx to context
func (data *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, dataTxKey{}, tx)
		return fn(ctx)
	})
}

// DB Get tx from context, and create it if not exists.
func (data *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(dataTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return data.db
}
