-- Drop Indexes
DROP INDEX IF EXISTS idx_movies_title;
DROP INDEX IF EXISTS idx_movie_genres_movie_id;
DROP INDEX IF EXISTS idx_movie_genres_genre_id;
DROP INDEX IF EXISTS idx_people_name;
DROP INDEX IF EXISTS idx_comparisons_user_id;
DROP INDEX IF EXISTS idx_ratings_user_movie_id;
DROP INDEX IF EXISTS idx_movie_availability_movie_id;
DROP INDEX IF EXISTS idx_movie_availability_service_id;
DROP INDEX IF EXISTS idx_user_movies_user_id;

-- Drop User Movies Table
DROP TABLE IF EXISTS user_movies;

-- Drop Movie Availability Table
DROP TABLE IF EXISTS movie_availability;

-- Drop Streaming Services Table
DROP TABLE IF EXISTS streaming_services;

-- Drop Ratings Table
DROP TABLE IF EXISTS ratings;

-- Drop Comparisons Table
DROP TABLE IF EXISTS comparisons;

-- Drop Movie People Junction Table
DROP TABLE IF EXISTS movie_people;

-- Drop People Table
DROP TABLE IF EXISTS people;

-- Drop Movie Genres Junction Table
DROP TABLE IF EXISTS movie_genres;

-- Drop Genres Table
DROP TABLE IF EXISTS genres;

-- Drop Movie Cache Log Table
DROP TABLE IF EXISTS movie_cache_log;

-- Drop Movies Table
DROP TABLE IF EXISTS movies;

-- Drop Users Table
DROP TABLE IF EXISTS users;
