CREATE SCHEMA IF NOT EXISTS auth_tokens;

CREATE TABLE IF NOT EXISTS rizat_blog.auth_tokens (
    id VARCHAR(100) PRIMARY KEY,
    user_id VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_auth_token_user_id ON rizat_blog.auth_tokens(user_id);
