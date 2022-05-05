// Package dao for db.
package dao

import (
	"context"
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/NextSmartShip/common/util/env"
	ent "github.com/NextSmartShip/entgo/ent/gen"
	"github.com/pkg/errors" //nolint:gci
	log "github.com/sirupsen/logrus"

	// import mysql driver.
	_ "github.com/go-sql-driver/mysql"
)

// Open for mysql connect of ent.
func Open() (entClient *ent.Client, err error) {
	mysqlDSN := env.MysqlDSN

	dbConn, err := sql.Open("mysql", mysqlDSN+"?charset=utf8mb4&parseTime=True&loc=UTC")
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			},
		).Fatalf("failed connecting database")

		return nil, errors.Wrap(err, "")
	}

	if dbConn.Ping() != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			},
		).Fatalf("failed connecting database")

		return nil, errors.Wrap(err, "")
	}

	dbConn.SetMaxIdleConns(10) //nolint:gomnd

	drv := entsql.OpenDB("mysql", dbConn)

	entClient = ent.NewClient(ent.Driver(drv)).Debug()

	// if err := entClient.Schema.Create(context.Background()); err != nil {
	//	logrus.Fatalf("failed running database migration: %v", err)
	// }

	return entClient, nil
}

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}
