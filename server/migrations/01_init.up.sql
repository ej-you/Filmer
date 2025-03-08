DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS movies CASCADE;
DROP TABLE IF EXISTS genres CASCADE;
DROP TABLE IF EXISTS user_movies CASCADE;


CREATE TABLE users (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    email VARCHAR(100) NOT NULL,
    password BYTES NOT NULL,
    created_at TIMESTAMP NOT NULL,

    CONSTRAINT users_pkey PRIMARY KEY (id ASC),
    UNIQUE INDEX uni_users_email (email ASC)
);

CREATE TABLE movies (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    kinopoisk_id BIGINT NOT NULL,
    title VARCHAR(100) NOT NULL,
    img_url VARCHAR(255) NOT NULL,
    web_url VARCHAR(255) NOT NULL,
    rating DECIMAL(2,1) NULL,
    year SMALLINT NULL,
    movie_length VARCHAR(10) NULL,
    description STRING NULL,
    type VARCHAR(20) NULL,
    staff JSONB NOT NULL,
    updated_at TIMESTAMP NOT NULL,

    CONSTRAINT movies_pkey PRIMARY KEY (id ASC),
    UNIQUE INDEX uni_movies_kinopoisk_id (kinopoisk_id ASC),

    FAMILY f1 (id, kinopoisk_id, title, img_url, rating, year, type, updated_at),
    FAMILY f2 (web_url, movie_length, description, staff)
);

CREATE TABLE genres (
    movie_id UUID NOT NULL,
    genre VARCHAR(50) NOT NULL,

    CONSTRAINT genres_pkey PRIMARY KEY (movie_id ASC, genre ASC),
    CONSTRAINT fk_movies_genres FOREIGN KEY (movie_id) REFERENCES movies(id)
);

CREATE TABLE user_movies (
    user_id UUID NOT NULL,
    movie_id UUID NOT NULL,
    status SMALLINT NOT NULL,
    stared BOOL NOT NULL,
    updated_at TIMESTAMP NOT NULL,

    CONSTRAINT user_movies_pkey PRIMARY KEY (user_id ASC, movie_id ASC),
    CONSTRAINT fk_users_user_movies FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_movies_user_movies FOREIGN KEY (movie_id) REFERENCES movies(id)
);
