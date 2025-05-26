package repository

import (
	"context"
	"database/sql"
	"gofiber-restapi/domain"

	"github.com/doug-martin/goqu/v9"
)

type userRepository struct {
	db *goqu.Database
}

func NewUser(con *sql.DB) domain.UserRepository {
	return &userRepository{
		db: goqu.New("default", con),
	}
}

func (ur userRepository) FindByEmail(ctx context.Context, email string) (usr domain.User, err error) {
	dataset := ur.db.From("users").Where(goqu.C("email").Eq(email))
	_, err = dataset.ScanStructContext(ctx, &usr)
	return
}
