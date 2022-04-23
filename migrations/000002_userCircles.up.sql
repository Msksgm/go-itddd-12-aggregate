BEGIN;

CREATE TABLE IF NOT EXISTS userCircles(
    user_id VARCHAR(255),
    circle_id VARCHAR(255)
);

COMMIT;