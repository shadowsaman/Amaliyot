package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"crud/models"
)

type UserSubscriptionRepo struct {
	db *pgxpool.Pool
}

func NewUserSubscriptionRepo(db *pgxpool.Pool) *UserSubscriptionRepo {
	return &UserSubscriptionRepo{
		db: db,
	}
}

func (f *UserSubscriptionRepo) Create(ctx context.Context, req *models.CreateUserSubscription) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO user_subscription (
			id,
			user_id,
			subscription_id,
			free_trial_start_date,
			free_trial_end_date,
			status,
			send_notification,
			updated_at
		) SELECT 
			$1,
			$2,
			s.id,
			now(),
			now() + s.free_trial * '1 days'::interval,
			case
			when s.free_trial > 0 then 'FREE_TRIAL'
			else 'INACTIVE'
			END AS status,
			false,
			now()
			FROM subscription as s
			WHERE s.id = $3
	`

	_, err := f.db.Exec(ctx, query,
		id,
		req.UserId,
		req.SubscriptionId,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *UserSubscriptionRepo) GetByPKey(ctx context.Context, pkey *models.UserSubscriptionPrimarKey) (*models.UserSubscriptionHistory, error) {

	var (
		status    sql.NullString
		startDate sql.NullString
		endDate   sql.NullString
		titleRu   sql.NullString
		titleEn   sql.NullString
		titleUz   sql.NullString
		price     sql.NullFloat64
		duration  sql.NullInt64
	)

	query := `
		SELECT
			ush.status,
			ush.start_date,
			ush.end_date,
			s.title_ru,
			s.title_en,
			s.title_uz,
			s.price,
			s.duration
		FROM
			user_subscription as us
		join user_subscription_history as ush on ush.user_subscription_id = us.id
		join subscription as s on s.id = us.subscription_id
		WHERE us.id = $1
		ORDER BY ush.start_date, ush.end_date DESC
		LIMIT 1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&status,
			&startDate,
			&endDate,
			&titleRu,
			&titleEn,
			&titleUz,
			&price,
			&duration,
		)

	if err != nil {
		return nil, err
	}

	return &models.UserSubscriptionHistory{
		Status:    status.String,
		StartDate: startDate.String,
		EndDate:   endDate.String,
		TitleRu:   titleRu.String,
		TitleEn:   titleEn.String,
		TitleUz:   titleUz.String,
		Price:     price.Float64,
		Duration:  int(duration.Int64),
	}, nil
}

func (f *UserSubscriptionRepo) GetList(ctx context.Context, req *models.GetListUserSubscriptionRequest) (*models.GetListUserSubscriptionResponse, error) {

	var (
		resp   = models.GetListUserSubscriptionResponse{}
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
			user_id,
			subscription_id,
			free_trial_start_date,
			free_trial_end_date,
			status,
			send_notification,
			created_at,
			updated_at
		FROM
			user_subscription
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {

		var (
			id                 sql.NullString
			userId             sql.NullString
			subscriptionId     sql.NullString
			freeTrialStartDate sql.NullString
			freeTrialEndDate   sql.NullString
			status             sql.NullString
			sendNotification   sql.NullBool
			createdAt          sql.NullString
			updatedAt          sql.NullString
		)

		err := rows.Scan(
			&id,
			userId,
			&subscriptionId,
			&freeTrialStartDate,
			&freeTrialEndDate,
			&status,
			&sendNotification,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.UserSubscriptions = append(resp.UserSubscriptions, &models.UserSubscription{
			Id:                 id.String,
			UserId:             userId.String,
			SubscriptionId:     subscriptionId.String,
			FreeTrialStartDate: freeTrialEndDate.String,
			FreeTrialEndDate:   freeTrialEndDate.String,
			Status:             status.String,
			SendNotification:   sendNotification.Bool,
			CreatedAt:          createdAt.String,
			UpdatedAt:          updatedAt.String,
		})

	}

	return &resp, err
}

func (f *UserSubscriptionRepo) GetSubscriptionsByUserId(ctx context.Context, req *models.UserIdSubscription) (*models.GetSubscriptionByUserId, error) {

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
		 	s.id,
			s.title_ru,
			s.title_en,
			s.title_uz,
			s.price,
			s.duration_type,
			s.duration,
			s.is_active,
			s.free_trial,
			s.created_at,
			s.updated_at
		FROM 
			user_subscription as us
		join subscription as s on s.id = us.subscription_id
		where us.user_id = $1
	`

	err := f.db.QueryRow(ctx, query, req.UserId).
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

	resp := &models.GetSubscriptionByUserId{}

	resp.Purchased = &models.Subscription{
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
	}

	queryUnpur := `
		SELECT 
			s.id,
			s.title_ru,
			s.title_en,
			s.title_uz,
			s.price,
			s.duration_type,
			s.duration,
			s.is_active,
			s.free_trial,
			s.created_at,
			s.updated_at
		FROM 
			subscription as s
		left join user_subscription as us on us.subscription_id = s.id AND us.user_id = $1
		where us.user_id is null
	`

	rows, err := f.db.Query(ctx, queryUnpur, req.UserId)

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

		resp.Unpurchased = append(resp.Unpurchased, &models.Subscription{
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

	return resp, err

}

func (f *UserSubscriptionRepo) HasAccess(ctx context.Context, req *models.UserIdSubscription) (string, error) {

	var status sql.NullString

	query := `
		SELECT 
			status
		FROM
			user_subscription 
		WHERE user_id = $1
	`
	err := f.db.QueryRow(ctx, query, req.UserId).Scan(&status)

	if err != nil {
		return "", err
	}

	return status.String, nil
}

func (f *UserSubscriptionRepo) MakeSubscriptionActive(ctx context.Context, req *models.UserSubscriptionId) (*models.UserHistory, error) {

	query := `
		UPDATE user_subscription_history SET status = 'ACTIVE', start_date = now() where user_subscription_id = $1
	`
	_, err := f.db.Exec(ctx, query, req.UserSubscriptionId)

	if err != nil {
		return nil, err
	}

	var (
		id                 sql.NullString
		userSubscriptionId sql.NullString
		status             sql.NullString
		start_date         sql.NullString
		end_date           sql.NullString
		price              sql.NullFloat64
	)

	query1 := `
		SELECT 
			id,
			user_subscription_id,
			status,
			start_date,
			end_date,
			price
		FROM 
			user_subscription_history
		where user_subscription_id = $1
	`

	err = f.db.QueryRow(ctx, query1, req.UserSubscriptionId).
		Scan(
			&id,
			&userSubscriptionId,
			&status,
			&start_date,
			&end_date,
			&price,
		)

	if err != nil {
		return nil, err
	}

	return &models.UserHistory{
		Id:                 id.String,
		UserSubscriptionId: userSubscriptionId.String,
		Status:             status.String,
		StartDate:          start_date.String,
		EndDate:            end_date.String,
		Price:              price.Float64,
	}, nil

}

// func (f *UserSubscriptionRepo) Update(ctx context.Context, req *models.UpdateUserSubscription) (int64, error) {

// 	var (
// 		query = ""
// 	)

// 	query = `
// 		UPDATE
// 			UserSubscriptions
// 		SET
// 			first_name = $2,
// 			last_name = $3,
// 			phone = $4,
// 			email = $5
// 			updated_at = now()
// 		WHERE id = $1
// 	`

// 	rowsAffected, err := f.db.Exec(ctx, query,
// 		req.FirstName,
// 		req.LastName,
// 		req.Phone,
// 		req.Email,
// 	)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return rowsAffected.RowsAffected(), nil
// }

// func (f *UserSubscriptionRepo) Delete(ctx context.Context, req *models.UserSubscriptionPrimarKey) error {

// 	_, err := f.db.Exec(ctx, "DELETE FROM UserSubscriptions WHERE id = $1", req.Id)
// 	if err != nil {
// 		return err
// 	}

// 	return err
// }
