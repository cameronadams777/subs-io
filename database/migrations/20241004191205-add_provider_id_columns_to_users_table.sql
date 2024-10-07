
-- +migrate Up
ALTER TABLE users
ADD google_user_id VARCHAR(255),
ADD tiktok_user_id VARCHAR(255);

-- +migrate Down
ALTER TABLE users
DROP COLUMN google_user_id,
DROP COLUMN tiktok_user_id;
