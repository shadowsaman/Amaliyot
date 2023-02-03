package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"crud/models"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (f *UserRepo) Create(ctx context.Context, user *models.CreateUser) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO users (
			id,
			first_name,
			last_name,
			phone,
			email,
			updated_at
		) VALUES ( $1, $2, $3, $4, $5, now())
	`

	_, err := f.db.Exec(ctx, query,
		id,
		user.FirstName,
		user.LastName,
		user.Phone,
		user.Email,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *UserRepo) GetByPKey(ctx context.Context, pkey *models.UserPrimarKey) (*models.User, error) {

	var (
		id        sql.NullString
		firstName sql.NullString
		lastName  sql.NullString
		phone     sql.NullString
		email     sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query := `
		SELECT
			id,
			first_name,
			last_name,
			phone,
			email,
			created_at,
			updated_at
		FROM
			users
		WHERE id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&firstName,
			&lastName,
			&phone,
			&email,
			&createdAt,
			&updatedAt,
		)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:        id.String,
		FirstName: firstName.String,
		LastName:  lastName.String,
		Phone:     phone.String,
		Email:     email.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (f *UserRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error) {

	var (
		resp   = models.GetListUserResponse{}
		offset = " OFFSET 0"
		limit  = " LIMIT 5"
	)

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			first_name,
			last_name,
			phone,
			email,
			created_at,
			updated_at
		FROM
			users
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {

		var (
			id        sql.NullString
			firstName sql.NullString
			lastName  sql.NullString
			phone     sql.NullString
			email     sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&firstName,
			&lastName,
			&phone,
			&email,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &models.User{
			Id:        id.String,
			FirstName: firstName.String,
			LastName:  lastName.String,
			Phone:     phone.String,
			Email:     email.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})

	}

	return &resp, err
}

func (f *UserRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {

	var (
		query = ""
	)

	query = `
		UPDATE
			users
		SET
			first_name = $2,
			last_name = $3,
			phone = $4,
			email = $5
			updated_at = now()
		WHERE id = $1
	`

	rowsAffected, err := f.db.Exec(ctx, query,
		req.FirstName,
		req.LastName,
		req.Phone,
		req.Email,
	)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *UserRepo) Delete(ctx context.Context, req *models.UserPrimarKey) error {

	_, err := f.db.Exec(ctx, "DELETE FROM users WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return err
}
