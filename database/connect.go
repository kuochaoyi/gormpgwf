package database

/*
ref: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL

The returned DB is safe for concurrent use by multiple goroutines and maintains
its own pool of idle connections. Thus, the Open function should be called just
once. It is rarely necessary to close a DB.
ref: https://golang.org/pkg/database/sql/#Open
*/

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBClient *dbClient

type dbClient struct {
	DB *gorm.DB
}

type pgDBInfo struct {
	// Datatype string `yaml:"datatype"`
	Hostname string `yaml:"hostname"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	// Prefix   string `yaml:"prefix"`
}

/*type pgSQL struct {
	Debug   pgDBInfo
	Test    pgDBInfo
	Release pgDBInfo
}*/

func init() {
	c := config()
	// var dsn = "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=" + c.Hostname + " user=" + c.Username + " password=" + c.Password + " dbname=" + c.Database + " port=" + c.Port + " sslmode=disable TimeZone=Asia/Taipei parseTime=true"
	log.Printf("database.init(): dns=> %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm.Open(): %s", err)
		panic("failed to connect database")
	}
	log.Printf("gorm db.Name(): %s", db.Name())

	DBClient = &dbClient{DB: db}
	log.Printf("gorm DBClient: %+v", DBClient)
}

func config() *pgDBInfo {
	pg := new(pgDBInfo)

	yamlFile, _ := ioutil.ReadFile("config/postgresql.yaml")
	err := yaml.Unmarshal(yamlFile, &pg)
	if err != nil {
		log.Fatalf("Unmarshal: %s", err)
	}
	log.Printf("database.config(): pg.Username= %s", pg.Username)
	return pg
}

/*
We are using pgx as postgres’s database/sql driver, it enables prepared statement cache by default, to disable it:
https://github.com/go-gorm/postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
	DSN: "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
	PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
*/

/*
Existing database connection
GORM allows to initialize *gorm.DB with an existing database connection
	import (
	  "database/sql"
	  "gorm.io/gorm"
	)

	sqlDB, err := sql.Open("postgres", "mydb_dsn")
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
	  Conn: sqlDB,
	}), &gorm.Config{})
*/

// https://gorm.io/docs/connecting_to_the_database.html#Connection-Pool
