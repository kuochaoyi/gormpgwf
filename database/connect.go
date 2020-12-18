package database

// ref: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// We are using pgx as postgres’s database/sql driver, it enables prepared statement cache by default, to disable it:
// https://github.com/go-gorm/postgres
db, err := gorm.Open(postgres.New(postgres.Config{
DSN: "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
PreferSimpleProtocol: true, // disables implicit prepared statement usage
}), &gorm.Config{})

sqlDB, err := sql.Open("postgres", "mydb_dsn")
gormDB, err := gorm.Open(postgres.New(postgres.Config{
Conn: sqlDB,
}), &gorm.Config{})