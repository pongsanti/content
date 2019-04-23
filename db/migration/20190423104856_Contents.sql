-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE contents (
  id serial PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  title text NOT NULL,
  subtitle text,
  detail text,
  start_at timestamp with time zone,
  end_at timestamp with time zone
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE contents;
