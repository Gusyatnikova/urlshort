package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "Nata2010"
	dbname   = "urlshort"
)

type Redirection struct {
	Path string
	URL string
}

func SetupConn() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "urlshort."+ defaultTableName
	}
	db.LogMode(true)
	return db
}

func GetRedirections(db *gorm.DB) map[string]string {
	redirections := make([]Redirection,0)
	db.Find(&redirections)
	if db.Error != nil {
		panic(db.Error)
	}
	return listToMapRedir(redirections)
}

func listToMapRedir (list []Redirection) map[string]string {
	if list == nil || len(list) == 0 {
		return nil
	}
	redirMap := make(map[string]string)
	for _, redir := range list {
		redirMap[redir.Path] = redir.URL
	}
	return redirMap
}

