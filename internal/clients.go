package internal

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var Clients *clients

type clients struct {
	config    Config
	context   context.Context
	pgSQL     *bun.DB
	pgSQLOnce sync.Once
}

// init /
func init() {
	Clients = initWithConfig(prodConfig)
}

// initWithConfig /
func initWithConfig(c Config) *clients {
	return &clients{
		context: context.Background(),
		config:  c,
	}
}

// Close release all resources
func (c *clients) Close() {
	if c.pgSQL != nil {
		_ = c.pgSQL.Close()
	}
}

// Config /
func (c *clients) Config() Config {
	return c.config
}

// PostgreSQL /
func (c *clients) PostgreSQL() *bun.DB {
	c.pgSQLOnce.Do(func() {
		connector := pgdriver.NewConnector(
			pgdriver.WithNetwork("tcp"),
			pgdriver.WithAddr(fmt.Sprintf("%s:%s", c.Config().Postgres.Host, c.Config().Postgres.Port)),
			pgdriver.WithInsecure(true),
		)
		connector.Config().User = c.Config().Postgres.User
		connector.Config().Password = c.Config().Postgres.Password
		connector.Config().Database = c.Config().Postgres.DbName

		c.pgSQL = bun.NewDB(sql.OpenDB(connector), pgdialect.New())
	})

	return c.pgSQL
}
