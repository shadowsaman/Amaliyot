package storage

import (
	"context"
	"crud/models"
)

type StorageI interface {
	CloseDB()
	User() UserRepoI
	Subscription() SubscriptionRepoI
	UserSubscription() UserSubscriptionI
}

type UserRepoI interface {
	Create(context.Context, *models.CreateUser) (string, error)
	GetByPKey(context.Context, *models.UserPrimarKey) (*models.User, error)
	GetList(context.Context, *models.GetListUserRequest) (*models.GetListUserResponse, error)
	Update(context.Context, *models.UpdateUser) (int64, error)
	Delete(context.Context, *models.UserPrimarKey) error
}

type SubscriptionRepoI interface {
	Create(context.Context, *models.CreateSubscription) (string, error)
	GetByPKey(context.Context, *models.SubscriptionPrimaryKey) (*models.Subscription, error)
	GetList(context.Context, *models.GetListSubscriptionRequest) (*models.GetListSubscriptionResponse, error)
	Update(context.Context, *models.UpdateSubscription) (int64, error)
	Delete(context.Context, *models.SubscriptionPrimaryKey) error
}

type UserSubscriptionI interface {
	Create(context.Context, *models.CreateUserSubscription) (string, error)
	GetByPKey(context.Context, *models.UserSubscriptionPrimarKey) (*models.UserSubscriptionHistory, error)
	GetList(context.Context, *models.GetListUserSubscriptionRequest) (*models.GetListUserSubscriptionResponse, error)
	GetSubscriptionsByUserId(context.Context, *models.UserIdSubscription) (*models.GetSubscriptionByUserId, error)
	HasAccess(context.Context, *models.UserIdSubscription) (string, error)
	MakeSubscriptionActive(ctx context.Context, req *models.UserSubscriptionId) (*models.UserHistory, error)
}
