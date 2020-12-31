package users

import (
	"context"

	"github.com/frankffenn/xerp-srv/config"
	"github.com/frankffenn/xerp-srv/services/users/db"
	mod "github.com/frankffenn/xerp-srv/services/users/mod"
)

func CreateUser(ctx context.Context, user *mod.User) error {
	return db.CreateUser(ctx, config.Session(), user)
}

func GetUserInfo(ctx context.Context, userid uint64) (*mod.User, error) {
	return db.GetUserInfo(ctx, config.Session(), userid)
}
