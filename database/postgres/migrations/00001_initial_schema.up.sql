CREATE TABLE IF NOT EXISTS comments (
    id serial PRIMARY KEY,
    movie_id int NOT NULL,
    message varchar(500) NOT NULL,
    author_ip_address varchar(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS characters (
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    gender varchar(10) NOT NULL,
    height varchar(10) NOT NULL,
    movies_id integer[]
);