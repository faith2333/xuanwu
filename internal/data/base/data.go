package base

import (
	"context"
	"fmt"
	"github.com/faith2333/xuanwu/internal/biz/base"
	"github.com/faith2333/xuanwu/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Data .
type Data struct {
	logger *log.Helper
	db     *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (_ *Data, err error) {
	data := &Data{
		logger: log.NewHelper(logger),
	}

	data.db, err = newDB(c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewTransaction(data *Data) base.Transaction {
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

type dataTxKey struct{}

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
	return data.db.WithContext(ctx)
}
