CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) UNIQUE NOT NULL,
    avatar_id INTEGER,
    role role NOT NULL,
    FOREIGN KEY (avatar_id) REFERENCES avatars(id)
);
