package connectors

import (
	"errors"
	"github.com/boj/redistore"
	"github.com/gorilla/sessions"
	"github.com/markdicksonjr/nibbler-redis"
)

type RedisStoreConnector struct {
	RedisExtension *redis.Extension
	Secret         string
	MaxAgeSeconds  int

	store          *redistore.RediStore
}

func (s RedisStoreConnector) Connect() (error, sessions.Store) {

	// if no DB provided, use sqlite3 memory DB
	if s.RedisExtension == nil {
		return errors.New("no redis extension in redis session connector"), nil
	} else {
		var err error
		s.store, err = redistore.NewRediStore(10, "tcp", s.RedisExtension.Url, s.RedisExtension.Password, []byte(s.Secret))
		if err != nil {
			return err, nil
		}
	}

	if len(s.Secret) == 0 {
		return errors.New("sql connector requires secret"), nil
	}

	s.store.DefaultMaxAge = s.MaxAge()

	return nil, s.store
}

func (s RedisStoreConnector) MaxAge() int {
	if s.MaxAgeSeconds == 0 {
		return 60 * 60 * 24 * 15 // 15 days
	}
	return s.MaxAgeSeconds
}
