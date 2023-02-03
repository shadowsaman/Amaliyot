CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    phone  VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE subscription (
    id  UUID PRIMARY KEY,
    title_ru VARCHAR,
    title_en VARCHAR,
    title_uz VARCHAR,
    price NUMERIC NOT NULL,
    duration_type VARCHAR NOT NULL,
    duration INTEGER NOT NULL,
    is_active BOOLEAN,
    free_trial INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);


CREATE TABLE user_subscription (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    subscription_id UUID NOT NULL REFERENCES subscription(id),
    free_trial_start_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    free_trial_end_date TIMESTAMP,
    status VARCHAR NOT NULL,
    send_notification BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE user_subscription_history (
    id UUID PRIMARY KEY,
    user_subscription_id UUID NOT NULL REFERENCES user_subscription(id),
    status VARCHAR NOT NULL,
    start_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, 
    end_date TIMESTAMP,
    price NUMERIC,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);