package persistence

import (
	"database/sql"
	"github.com/giova333/rest-golang-postgres-docker/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) Create(user domain.User) error {
	query := `INSERT INTO "users"("uuid", "name", "last_name") VALUES($1, $2, $3)`
	statement, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(user.UUID.String(), user.Name, user.LastName)
	return err
}

func (repo *UserRepository) GetAll() ([]domain.User, error) {
	query := `SELECT * FROM "users"`
	var users domain.Users
	rows, err := repo.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user domain.User
		err = rows.Scan(&user.UUID, &user.Name, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo *UserRepository) GetOne(uuid string) (*domain.User, error) {
	query := `SELECT * FROM "users" WHERE "uuid" = $1`
	var user domain.User

	statement, err := repo.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	err = statement.QueryRow(uuid).Scan(&user.UUID, &user.Name, &user.LastName)
	return &user, err
}

func (repo *UserRepository) Delete(uuid string) error {
	query := `DELETE FROM "users" WHERE "uuid" = $1`
	statement, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(uuid)
	return err
}
