CREATE TABLE IF NOT EXISTS users
(
    uuid      VARCHAR(64) PRIMARY KEY,
    name      VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL
)