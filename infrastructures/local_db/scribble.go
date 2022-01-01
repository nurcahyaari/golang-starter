package localdb

import (
	"log"

	scribble "github.com/nanobox-io/golang-scribble"
)

type ScribleImpl struct {
	db *scribble.Driver
}

func NewScribleClient() *ScribleImpl {
	db, err := scribble.New("tmp/db", nil)
	if err != nil {
		log.Println("Error", err)
	}

	return &ScribleImpl{
		db: db,
	}
}

func (db ScribleImpl) DB() *scribble.Driver {
	return db.db
}
