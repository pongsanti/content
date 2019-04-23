-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE contents (
  id serial PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  content_type text NOT NULL DEFAULT 'news',
  title text NOT NULL,
  subtitle text,
  detail text,
  start_at timestamp with time zone,
  end_at timestamp with time zone,
  creator_id integer
);

CREATE INDEX idx_contents_content_type ON contents(content_type);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE contents;
