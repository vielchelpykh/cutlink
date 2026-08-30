CREATE TABLE links (
    id SERIAL PRIMARY KEY,
    full_link VARCHAR(200) NOT NULL,
    short_link VARCHAR(200) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    pressed INTEGER NOT NULL,
    UNIQUE(full_link, short_link)
);