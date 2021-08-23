package main

import (
	"database/sql"
	"log"
)

type Redirector interface {
	GetRedirections() (map[string]string, error)
	CreateRedirection(path, url string) error
}

type dbRedirector struct {
	db *sql.DB
}

var redirector Redirector

func InitRedirector(r Redirector) {
	redirector = r
}

func (red *dbRedirector) GetRedirections() (map[string]string, error) {
	rows, err := red.db.Query("SELECT path, url FROM urlshort_schema.redirection")
	if err != nil {
		log.Println("Cannot SELECT * FROM redirection")
		return nil, err
	}
	defer rows.Close()
	sqlMap := make(map[string]string)
	for rows.Next() {
		var path, url string
		err = rows.Scan(&path, &url)
		if err != nil {
			log.Println("cannot scan rows")
			return nil, err
		}
		sqlMap[path] = url
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return sqlMap, nil
}

func (red *dbRedirector) CreateRedirection(path, url string) error {
	return nil
}