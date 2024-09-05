
-- +migrate Up
CREATE TABLE posts (
  id uuid PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  url VARCHAR(255) NOT NULL,
  status VARCHAR(255) NOT NULL,
  user_id uuid REFERENCES users(id),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);

-- +migrate Down
DROP TABLE posts;
