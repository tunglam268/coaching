package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST     = "127.0.0.1"
	USERNAME = "hppoc"
	PASSWORD = "password"
	DBNAME   = "fabricexplorer"
	PORT     = 5432
)

type DBInterface interface {
	OpenConnection()
	CloseConnection()
}

type DBConnection struct {
	Host     string `json:"host"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Port     int    `json:"port"`
	Dsn      string `json:"dsn"`
	DBInterface
	ExposeDB *gorm.DB
}

//Tao doi tuonng
func NewDBConnection() *DBConnection {
	dbConnection := new(DBConnection)
	dbConnection.Host = HOST
	dbConnection.UserName = USERNAME
	dbConnection.Password = PASSWORD
	dbConnection.Dbname = DBNAME
	dbConnection.Port = PORT
	dbConnection.Dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		HOST, USERNAME, PASSWORD, DBNAME, PORT,
	)

	return dbConnection
}

// receiver function
func (DB *DBConnection) OpenConnection() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  DB.Dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connect Success !")
		DB.ExposeDB = db
	}

}

// // receiver function
// func (DB *DBConnection) CloseConnection() {

// }
