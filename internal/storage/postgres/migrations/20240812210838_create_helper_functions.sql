-- +goose Up
-- +goose StatementBegin
CREATE FUNCTION sync_updated_at_column () RETURNS TRIGGER LANGUAGE plpgsql AS $$
BEGIN
    IF NEW.deleted_at IS NULL THEN
        NEW.updated_at = CLOCK_TIMESTAMP();
    END IF;
    RETURN NEW;
END;
$$;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP FUNCTION sync_updated_at_column;

-- +goose StatementEnd