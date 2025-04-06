ALTER TABLE genres DROP CONSTRAINT fk_movies_genres;
ALTER TABLE user_movies DROP CONSTRAINT fk_users_user_movies;
ALTER TABLE user_movies DROP CONSTRAINT fk_movies_user_movies;
