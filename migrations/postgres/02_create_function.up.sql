CREATE OR REPLACE FUNCTION subscription_history_trigger() RETURNS trigger
LANGUAGE PLPGSQL
AS
$$

    BEGIN
        INSERT INTO user_subscription_history (
            id,
            user_subscription_id,
            status,
            start_date,
            end_date,
            price,
            created_at,
            updated_at
        ) 
        SELECT 
            uuid_generate_v4(),
            new.id,
            new.status,
            new.free_trial_start_date,
            new.free_trial_end_date,
            s.price,
            now(),
            now()
        from subscription as s
        where s.id = new.subscription_id;

        return new;
    END;
$$;

CREATE TRIGGER subscription_trigger
AFTER INSERT ON user_subscription
FOR EACH ROW EXECUTE PROCEDURE subscription_history_trigger();