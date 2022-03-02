CREATE TABLE IF NOT EXISTS comments (
    id serial PRIMARY KEY,
    movie_id int NOT NULL,
    message varchar(500) NOT NULL,
    author_ip_address varchar(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);