package model

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

/**
https://gorm.io/zh_CN/docs/connecting_to_the_database.html
*/
type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Database

func (db *Database) Init() {
	DB = &Database{
		Self:   GetSelfDB(),
		Docker: GetDockerDB(),
	}
}

func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")

	// add gorm's logger.
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(config), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("Database connection failed.", err)
	}

	// set for db connection
	setupDB(db)

	/// auto migrate.
	db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{})

	return db
}

func setupDB(db *gorm.DB) {
	mysqlDB, _ := db.DB()
	mysqlDB.SetMaxIdleConns(10)
	mysqlDB.SetMaxOpenConns(20000)
	mysqlDB.SetConnMaxLifetime(time.Hour)
}

func (db *Database) Close() {
	mysqlDB, _ := DB.Self.DB()
	mysqlDB.Close()
	mysqlDB2, _ := DB.Docker.DB()
	mysqlDB2.Close()
}
