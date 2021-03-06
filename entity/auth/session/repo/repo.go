package repo

import (
	"github.com/egoholic/charcoal/entity/auth/session"
	"github.com/egoholic/charcoal/entity/auth/session/repo/idmap"
	"github.com/egoholic/serror"
)

type Repo struct {
	identityMap *idmap.IdentityMap
}

func New() *Repo {
	return &Repo{idmap.New()}
}

type InsertAdapter func(s *session.Session) (interface{}, error)

func (r *Repo) NewInserter(insert InsertAdapter) func(*session.Session) (interface{}, error) {
	return func(s *session.Session) (sid interface{}, err error) {
		sid, err = insert(s)
		if err != nil {
			err = serror.Decorate(err, "can not insert session")
			return
		}

		err = r.identityMap.Add(s.PK(), sid, s)
		return
	}
}

type FindByTokenAdapter func(token string) (interface{}, *session.Session, error)

func (r *Repo) NewByTokenFinder(find FindByTokenAdapter) func(string) (*session.Session, error) {
	return func(t string) (*session.Session, error) {
		sid, s, ok := r.identityMap.Get(t)
		if ok {
			return s, nil
		}

		sid, s, err := find(t)
		if err != nil {
			return nil, err
		}

		err = r.identityMap.Add(s.PK(), sid, s)
		return s, err
	}
}

type FindByAccountAdapter func(apk string) (interface{}, *session.Session, error)

func (r *Repo) NewByAccountFinder(find FindByAccountAdapter) func(string) (*session.Session, error) {
	return func(apk string) (*session.Session, error) {
		sid, s, ok := r.identityMap.Get(apk)
		if ok {
			return s, nil
		}

		sid, s, err := find(apk)
		if err != nil {
			return nil, err
		}

		err = r.identityMap.Add(s.PK(), sid, s)
		return s, err
	}
}
