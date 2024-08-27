
-- +migrate Up
CREATE TABLE users (
  id uuid PRIMARY KEY,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  confirmed_at TIMESTAMP,
  confirmation_token VARCHAR(255),
  reset_password_token VARCHAR(255),
  reset_password_expiry TIMESTAMP
);

-- +migrate Down
DROP TABLE users;
