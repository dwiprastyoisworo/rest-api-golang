package repositories

import (
	"context"
	"database/sql"
	"rest-api-golang/src/entity"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r UserRepository) Get(ctx context.Context, db *sql.DB) ([]entity.Users, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return []entity.Users{}, err
	}
	defer conn.Close()
	rows, err := conn.QueryContext(ctx, `select id, "name" , address from users u`)
	if err != nil {
		return []entity.Users{}, err
	}
	defer rows.Close()
	var users []entity.Users
	for rows.Next() {
		var user entity.Users
		errRowScan := rows.Scan(&user.ID, &user.Name, &user.Address)
		if errRowScan != nil {
			return []entity.Users{}, errRowScan
		}
		users = append(users, user)
	}
	return users, nil
}

func (r UserRepository) Insert(ctx context.Context, db *sql.DB, user entity.UserRequest) (int, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return 0, nil
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, `insert into users (id, "name", address) values ($1, $2, $3)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, user.ID, user.Name, user.Address)
	if err != nil {
		return 0, err
	}
	return 1, nil
}
