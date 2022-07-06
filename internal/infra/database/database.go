package database

import (
	"log"
	"os"

	"github.com/test_server/config"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

func NewClient() (db *db.Session, sessClose db.Session, err error) {

	var conf = config.GetConfiguration()

	settings := postgresql.ConnectionURL{
		Database: conf.DatabaseName,
		Host:     conf.DatabaseHost,
		User:     conf.DatabaseUser,
		Password: conf.DatabasePassword,
	}

	_, err = os.Stat(conf.FileStorageLocation)
	if err != nil {
		err = os.Mkdir(conf.FileStorageLocation, os.ModePerm)
	}

	sess, err := postgresql.Open(settings)

	if err != nil {
		log.Println("cannot open postgresql: ")
		return nil, nil, err
	}

	sessClose = sess

	if err = sess.Ping(); err != nil {
		log.Println("cannot ping")
		return nil, nil, err
	}

	log.Printf("Successfully connected to database: %q", sess.Name())

	return &sess, sessClose, nil
}
