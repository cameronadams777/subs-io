
-- +migrate Up
ALTER TABLE users
ADD google_provider_id VARCHAR(255),
ADD tiktok_provider_id VARCHAR(255);

-- +migrate Down
ALTER TABLE users
DROP COLUMN google_provider_id,
DROP COLUMN tiktok_provider_id;
