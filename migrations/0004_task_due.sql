-- +goose up
ALTER TABLE tasks ADD COLUMN due TEXT DEFAULT NULL;

-- +goose down
ALTER TABLE tasks DROP COLUMN due;
