package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//get the dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//DB blabla
type DB struct {
	db *gorm.DB
}

//NewDBPostresImpl blabla
func NewDBPostresImpl(host string, port int, user string, dbname string, password string) (*DB, error) {
	db, err := gorm.Open("postgres", fmt.Sprintf("sslmode=disable host=%s port=%d user=%s dbname=%s password=%s", host, port, user, dbname, password))
	return &DB{
		db: db,
	}, err
}

//MigrateAll blabla
func (odb *DB) MigrateAll() {
	odb.db.AutoMigrate(&UserAccount{})
}

//AddUser blabla
func (odb *DB) AddUser(username, password string) (UserAccount, error) {
	odb.db.Create(&UserAccount{
		Username: username,
		Password: password,
	})
	return odb.GetByUsername(username)
}

//Close blabla
func (odb *DB) Close() error {
	return odb.db.Close()
}

//GetByUsername blabla
func (odb *DB) GetByUsername(username string) (UserAccount, error) {
	user := UserAccount{}
	odb.db.Where("username = ?", username).First(&user)
	return user, nil
}
