package models

type UserSubscriptionPrimarKey struct {
	Id string `json:"id"`
}

type UserIdSubscription struct {
	UserId string `json:"user_id"`
}

type GetSubscriptionByUserId struct {
	Purchased   *Subscription   `json:"purchased"`
	Unpurchased []*Subscription `json:"unpuchsed"`
}

type CreateUserSubscriptionSwagger struct {
	UserId         string `json:"user_id"`
	SubscriptionId string `json:"subscription_id"`
}

type CreateUserSubscription struct {
	UserId         string `json:"user_id"`
	SubscriptionId string `json:"subscription_id"`
}

type UserSubscriptionId struct {
	UserSubscriptionId string `json:"user_subscription_id"`
}

type UserSubscription struct {
	Id                 string `json:"id"`
	UserId             string `json:"user_id"`
	SubscriptionId     string `json:"subscription_id"`
	FreeTrialStartDate string `json:"free_trial_start_date"`
	FreeTrialEndDate   string `json:"free_trial_end_date"`
	Status             string `json:"status"`
	SendNotification   bool   `json:"send_notification"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}

type UserSubscriptionHistory struct {
	Status    string  `json:"status"`
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date"`
	TitleRu   string  `json:"title_ru"`
	TitleEn   string  `json:"title_en"`
	TitleUz   string  `json:"title_uz"`
	Price     float64 `json:"price"`
	Duration  int     `json:"duration"`
}

type UserHistory struct {
	Id                 string  `json:"id"`
	UserSubscriptionId string  `json:"user_subscription_id"`
	Status             string  `json:"status"`
	StartDate          string  `json:"start_date"`
	EndDate            string  `json:"end_date"`
	Price              float64 `json:"price"`
}

type UpdateUserSubscription struct {
	Id                 string `json:"id"`
	UserId             string `json:"user_id"`
	SubscriptionId     string `json:"subscription_id"`
	FreeTrialStartDate string `json:"free_trial_start_date"`
	FreeTrialEndDate   string `json:"free_trial_end_date"`
	Status             string `json:"status"`
	SendNotification   bool   `json:"send_notification"`
}

type UpdateUserSubscriptionSwagger struct {
	UserId             string `json:"user_id"`
	SubscriptionId     string `json:"subscription_id"`
	FreeTrialStartDate string `json:"free_trial_start_date"`
	FreeTrialEndDate   string `json:"free_trial_end_date"`
	Status             string `json:"status"`
	SendNotification   bool   `json:"send_notification"`
}

type UpdateStatusUserSubscriptionSwagger struct {
	Status string `json:"status"`
}

type GetListUserSubscriptionRequest struct {
	Limit  int32
	Offset int32
}

type GetListUserSubscriptionResponse struct {
	Count             int32               `json:"count"`
	UserSubscriptions []*UserSubscription `json:"user_subscription"`
}
