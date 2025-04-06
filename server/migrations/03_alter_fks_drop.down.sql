ALTER TABLE genres ADD CONSTRAINT fk_movies_genres FOREIGN KEY (movie_id) REFERENCES movies(id);
ALTER TABLE user_movies ADD CONSTRAINT fk_users_user_movies FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE user_movies ADD CONSTRAINT fk_movies_user_movies FOREIGN KEY (movie_id) REFERENCES movies(id);
