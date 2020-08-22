CREATE TABLE Users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(32),
    created_at TIMESTAMP
);

CREATE TABLE Chats(
    id SERIAL PRIMARY KEY,
    name VARCHAR(64),
    users INTEGER[],
    created_at TIMESTAMP
);

CREATE TABLE Messages(
    id SERIAL PRIMARY KEY,
    chat INTEGER,
    author INTEGER,
    text TEXT,
    created_at TIMESTAMP
);