package models

type SubscriptionPrimaryKey struct {
	Id string `json:"id"`
}

type CreateSubscription struct {
	TitleRu      string  `json:"title_ru"`
	TitleEn      string  `json:"title_en"`
	TitleUz      string  `json:"title_uz"`
	Price        float64 `json:"price"`
	DurationType string  `json:"duration_type"`
	Duration     int     `json:"duration"`
	IsActive     bool    `json:"is_active"`
	FreeTrial    int     `json:"free_trial"`
}

type Subscription struct {
	Id           string  `json:"id"`
	TitleRu      string  `json:"title_ru"`
	TitleEn      string  `json:"title_en"`
	TitleUz      string  `json:"title_uz"`
	Price        float64 `json:"price"`
	DurationType string  `json:"duration_type"`
	Duration     int     `json:"suration"`
	IsActive     bool    `json:"is_active"`
	FreeTrial    int     `json:"free_trial"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	DeletedAt    string  `json:"deleted_at"`
}

type UpdateSubscriptionSwagger struct {
	TitleRu      string  `json:"title_ru"`
	TitleEn      string  `json:"title_en"`
	TitleUz      string  `json:"title_uz"`
	Price        float64 `json:"price"`
	DurationType string  `json:"duration_type"`
	Duration     int     `json:"suration"`
	IsActive     bool    `json:"is_active"`
	FreeTrial    int     `json:"free_trial"`
}

type UpdateSubscription struct {
	Id           string  `json:"id"`
	TitleRu      string  `json:"title_ru"`
	TitleEn      string  `json:"title_en"`
	TitleUz      string  `json:"title_uz"`
	Price        float64 `json:"price"`
	DurationType string  `json:"duration_type"`
	Duration     int     `json:"suration"`
	IsActive     bool    `json:"is_active"`
	FreeTrial    int     `json:"free_trial"`
}

type GetListSubscriptionRequest struct {
	Limit  int32
	Offset int32
}

type GetListSubscriptionResponse struct {
	Count         int32           `json:"count"`
	Subscriptions []*Subscription `json:"Subscriptions"`
}
