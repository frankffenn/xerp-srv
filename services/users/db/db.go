package db

import (
	"context"

	mod "github.com/frankffenn/xerp-srv/services/users/mod"
	"github.com/go-xorm/xorm"
	"golang.org/x/xerrors"
)

var ErrUserNotFound = xerrors.New("user not found")

func CreateUser(ctx context.Context, sess *xorm.Session, user *mod.User) error {
	_, err := sess.InsertOne(user)
	return err
}

func GetUserInfo(ctx context.Context, sess *xorm.Session, userid uint64) (*mod.User, error) {
	var user mod.User
	found, err := sess.ID(userid).Get(&user)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, ErrUserNotFound
	}

	return &user, nil
}

func GetUserInfoFromUsername(ctx context.Context, sess *xorm.Session, username string) (*mod.User, error) {
	var user mod.User
	found, err := sess.Where("username = ?", username).Get(&user)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, ErrUserNotFound
	}

	return &user, nil
}
