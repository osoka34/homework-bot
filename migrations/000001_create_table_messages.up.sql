
CREATE TABLE IF NOT EXISTS messages (
    id UUID PRIMARY KEY, 
    content TEXT NOT NULL,
    sender_id BIGINT NOT NULL,
    create_at TIMESTAMP NOT NULL,
    username VARCHAR(70),
    group_id BIGINT NOT NULL,
    pattern VARCHAR(255)
);

