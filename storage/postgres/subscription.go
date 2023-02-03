package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"crud/models"
)

type SubscriptionRepo struct {
	db *pgxpool.Pool
}

func NewSubscriptionRepo(db *pgxpool.Pool) *SubscriptionRepo {
	return &SubscriptionRepo{
		db: db,
	}
}

func (f *SubscriptionRepo) Create(ctx context.Context, req *models.CreateSubscription) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
 		INSERT INTO subscription (
 			id,
 			title_ru,
			title_en,
			title_uz,
			price,
			duration_type,
			duration,
			is_active,
			free_trial,
 			updated_at
 		) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, now() )
 	`

	_, err := f.db.Exec(ctx, query,
		id,
		req.TitleRu,
		req.TitleEn,
		req.TitleUz,
		req.Price,
		req.DurationType,
		req.Duration,
		req.IsActive,
		req.FreeTrial,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *SubscriptionRepo) GetByPKey(ctx context.Context, pkey *models.SubscriptionPrimaryKey) (*models.Subscription, error) {

	var (
		id           sql.NullString
		titleRu      sql.NullString
		titleEn      sql.NullString
		titleUz      sql.NullString
		price        sql.NullFloat64
		durationType sql.NullString
		duration     sql.NullInt64
		isActive     sql.NullBool
		freeTrial    sql.NullInt64
		createdAt    sql.NullString
		updatedAt    sql.NullString
	)

	query := `
 		SELECT
			id,
			title_ru,
			title_en,
			title_uz,
			price,
			duration_type,
			duration,
			is_active,
			free_trial,
 			created_at,
 			updated_at
 		FROM
 			subscription
 		WHERE id = $1
 	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&titleRu,
			&titleEn,
			&titleUz,
			&price,
			&durationType,
			&duration,
			&isActive,
			&freeTrial,
			&createdAt,
			&updatedAt,
		)

	if err != nil {
		return nil, err
	}

	return &models.Subscription{
		Id:           id.String,
		TitleRu:      titleRu.String,
		TitleEn:      titleEn.String,
		TitleUz:      titleUz.String,
		Price:        price.Float64,
		DurationType: durationType.String,
		Duration:     int(duration.Int64),
		IsActive:     isActive.Bool,
		FreeTrial:    int(freeTrial.Int64),
		CreatedAt:    createdAt.String,
		UpdatedAt:    updatedAt.String,
	}, nil
}

func (f *SubscriptionRepo) GetList(ctx context.Context, req *models.GetListSubscriptionRequest) (*models.GetListSubscriptionResponse, error) {

	var (
		resp   = models.GetListSubscriptionResponse{}
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
		 	title_ru,
		 	title_en,
		 	title_uz,
		 	price,
		 	duration_type,
		 	duration,
		 	is_active,
		 	free_trial,
		 	created_at,
		 	updated_at
 		FROM
 			subscription
 	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {

		var (
			id           sql.NullString
			titleRu      sql.NullString
			titleEn      sql.NullString
			titleUz      sql.NullString
			price        sql.NullFloat64
			durationType sql.NullString
			duration     sql.NullInt64
			isActive     sql.NullBool
			freeTrial    sql.NullInt64
			createdAt    sql.NullString
			updatedAt    sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&titleRu,
			&titleEn,
			&titleUz,
			&price,
			&durationType,
			&duration,
			&isActive,
			&freeTrial,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Subscriptions = append(resp.Subscriptions, &models.Subscription{
			Id:           id.String,
			TitleRu:      titleRu.String,
			TitleEn:      titleEn.String,
			TitleUz:      titleUz.String,
			Price:        price.Float64,
			DurationType: durationType.String,
			Duration:     int(duration.Int64),
			IsActive:     isActive.Bool,
			FreeTrial:    int(freeTrial.Int64),
			CreatedAt:    createdAt.String,
			UpdatedAt:    updatedAt.String,
		})

	}

	return &resp, err
}

func (f *SubscriptionRepo) Update(ctx context.Context, req *models.UpdateSubscription) (int64, error) {

	var (
		query = ""
	)

	query = `
 		UPDATE
 			subscription
 		SET
		 	title_ru = $2
		 	title_en = $3
		 	title_uz = $4
		 	price = $5
		 	duration_type = $6
		 	duration = $7
		 	is_active = $8
		 	free_trial = $9
 			updated_at = now()
 		WHERE id = $1
 	`

	rowsAffected, err := f.db.Exec(ctx, query,
		req.Id,
		req.TitleRu,
		req.TitleEn,
		req.TitleUz,
		req.Price,
		req.DurationType,
		req.Duration,
		req.IsActive,
		req.FreeTrial,
	)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *SubscriptionRepo) Delete(ctx context.Context, req *models.SubscriptionPrimaryKey) error {

	_, err := f.db.Exec(ctx, "DELETE FROM subscription WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return err
}
