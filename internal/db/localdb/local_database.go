package localdb

import (
	"log"

	scribble "github.com/nanobox-io/golang-scribble"
)

type LocalDB interface {
	Query() *scribble.Driver
}

type localDB struct {
	db *scribble.Driver
}

func Load() LocalDB {
	db, err := scribble.New("temp/db", nil)
	if err != nil {
		log.Println("Error", err)
	}

	return &localDB{
		db: db,
	}
}

func (db *localDB) Query() *scribble.Driver {
	return db.db
}
