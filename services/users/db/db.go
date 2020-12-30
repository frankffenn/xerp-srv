package db

import (
	"context"

	mod "github.com/frankffenn/xerp-srv/services/users/mod"
	"github.com/go-xorm/xorm"
	"golang.org/x/xerrors"
)

func CreateUser(ctx context.Context, sess *xorm.Session, user *mod.User) (*mod.User, error) {
	_, err := sess.Insert(user)
	return err
}

func GetUserInfo(ctx context.Context, sess *xorm.Session, userid uint64) (*mod.User, error) {
	var user mod.User
	found, err := sess.ID(userid).Get(&user)
	if err != nil {
		return err
	}

	if !found {
		return nil, xerrors.New("user not found")
	}

	return &user, nil
}
