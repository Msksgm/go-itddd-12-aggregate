BEGIN;

CREATE TABLE IF NOT EXISTS users(
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

INSERT INTO users(id, name) VALUES ('ownerId', 'ownerName');
INSERT INTO users(id, name) VALUES ('userId', 'userName');

COMMIT;