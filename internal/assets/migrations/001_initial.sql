-- +migrate Up
CREATE TABLE IF NOT EXISTS links (
  slug        text PRIMARY KEY,
  target_url  text        NOT NULL,
  created_at  timestamptz NOT NULL DEFAULT now(),
  ttl bigint
);

CREATE INDEX IF NOT EXISTS idx_links_ttl ON links (ttl);

-- +migrate Down
DROP INDEX IF EXISTS idx_links_ttl;
DROP TABLE IF EXISTS links;
