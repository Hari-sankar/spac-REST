CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) UNIQUE NOT NULL,
    avatar_id VARCHAR(255),
    role role NOT NULL,
    FOREIGN KEY (avatar_id) REFERENCES avatars(id)
);
