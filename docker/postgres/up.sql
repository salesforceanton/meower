DROP TABLE IF EXISTS meows;

CREATE TABLE meows (
  id         VARCHAR(32) PRIMARY KEY,
  body       TEXT        NOT NULL,
  created_at TIMESTAMP   NOT NULL
);