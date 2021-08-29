package db

import (
	"fmt"
	"log"

	"golang-rest-crud/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Dbcon    *gorm.DB
	Errdb    error
	dbuser   string
	dbpass   string
	dbname   string
	dbaddres string
	dbport   string
	dbdebug  bool
	dbtype   string
	sslmode  string
)

func init() {

	fmt.Println("init")
	dbtype = utils.GetEnv("TYPE", "MySQL")
	if dbtype != "MySQL" {
		return
	}
	fmt.Println("DB MySQL")

	dbuser = utils.GetEnv("MYSQL_USER", "root")
	dbpass = utils.GetEnv("MYSQL_PASS", "")
	dbname = utils.GetEnv("MYSQL_DBNAME", "golang-rest-crud")
	dbaddres = utils.GetEnv("MYSQL_ADDRESS", "127.0.0.1")
	dbport = utils.GetEnv("MYSQL_PORT", "3306")
	sslmode = utils.GetEnv("MYSQL_SSLMODE", "disable")
	dbdebug = true
	if DbOpen() != nil {
		fmt.Println("Connection to database MySQL is failed")
	}
}

func DbOpen() error {
	// args := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=%s", dbuser, dbpass, dbaddres, dbport, dbname)
	args :=
		fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s",
			dbuser,   // config.Username,
			dbpass,   // config.Password,
			dbaddres, // config.Host,
			dbport,   //config.Port,
			dbname,   //config.Schema,
			"Local")

	Dbcon, Errdb = gorm.Open(mysql.Open(args), &gorm.Config{})

	if Errdb != nil {
		log.Fatal("open db Err ", Errdb)
		return Errdb
	}

	return nil
}

func GetDbCon() *gorm.DB {
	//TODO looping try connection until timeout
	// using channel timeout
	return Dbcon
}
