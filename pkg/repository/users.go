package repository

import (
	"github.com/FelipeUmpierre/measures/pkg/domain"
	"github.com/jmoiron/sqlx"
	lk "github.com/ulule/loukoum"
)

type (
	// UserRepo holds the database
	UserRepo struct {
		db *sqlx.DB
	}

	// UserSearch holds the findBy* type
	UserSearch struct {
		ID string
	}
)

// NewUsersRepository user repo
func NewUsersRepository(db *sqlx.DB) *UserRepo {
	return &UserRepo{db}
}

// Save saves the domain
func (u *UserRepo) Save(user domain.User) (domain.User, error) {
	builder := lk.Insert().Set(
		lk.Pair(`id`, user.ID),
		lk.Pair(`name`, user.Name),
	).OnConflict(`id`, lk.DoUpdate(
		lk.Pair(`name`, user.Name)
	)).Returning(`id`, `name`)

	query, args := builder.Prepare()

	stmt, err := u.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.Get(&user, args...)
	return user, err
}

// FindAll returns the result for all rows
func (u *UserRepo) FindAll() ([]domain.User, error) {
	query, args := lk.Select(`id`, `name`).From(`users`).Prepare()

	stmt, err := u.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	users := new([]domain.User)
	err = stmt.Select(users, args...)

	return users, err
}

// FindByID return the result for specific row
func (u *UserRepo) FindByID(ID string) (*domain.User, error) {
	query, args := lk.Select(`id`, `name`).From(`users`).Where(lk.Condition(`id`).Equal(ID)).Prepare()

	stmt, err := u.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := new(domain.User)
	if err = u.db.Get(user, args...); err != nil {
		return nil, err
	}

	return user, nil
}
