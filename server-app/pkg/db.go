package pkg

import (
	"fmt"
	"log"
	"swms-sa/pkg/env"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Conn struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func (c *Conn) GetConn() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("Database connected, host:", c.Host, "port:", c.Port)

	/* disable create deleted_at, updated_at, created_at */
	db = db.Session(&gorm.Session{DisableNestedTransaction: true})

	return db
}

func (c *Conn) CloseConn(db *gorm.DB) bool {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	log.Println("Database closed, host:", c.Host, "port:", c.Port)
	return true
}

var ConnDB = Conn{
	Host:     env.MYSQL_HOST,
	Port:     env.MYSQL_PORT,
	User:     env.MYSQL_USER,
	Password: env.MYSQL_PASSWORD,
	Database: env.MYSQL_DATABASE,
}
