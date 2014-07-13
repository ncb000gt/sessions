package sessions

import (
	"labix.org/v2/mgo"
	"github.com/kidstuff/mongostore"
	"github.com/gorilla/sessions"
)

type MongoStore interface {
	Store
	Options(Options)
}

func NewMongoStore(collection *mgo.Collection, maxAge int, ensureTTL bool, keyPairs ...[]byte) MongoStore {
	store := mongostore.NewMongoStore(collection, maxAge, ensureTTL, keyPairs...)
	return &mongoStore{store}
}

type mongoStore struct {
	*mongostore.MongoStore
}

func (c *mongoStore) Options(options Options) {
	c.MongoStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
