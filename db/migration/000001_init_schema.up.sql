-- Users Table
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Movies Table
CREATE TABLE movies (
    movie_id SERIAL PRIMARY KEY,
    tmdb_id INT NOT NULL UNIQUE,
    imdb_id VARCHAR(255) UNIQUE,
    title VARCHAR(255) NOT NULL,
    overview TEXT,
    release_date DATE,
    poster_path VARCHAR(255),
    backdrop_path VARCHAR(255),
    tmdb_popularity DECIMAL(10, 2),
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Movie Cache Log
CREATE TABLE movie_cache_log (
    tmdb_id INT NOT NULL UNIQUE,
    last_fetched TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Genres Table
CREATE TABLE genres (
    genre_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

-- Movie Genres Junction Table
CREATE TABLE movie_genres (
    movie_id INT NOT NULL REFERENCES movies(movie_id) ON DELETE CASCADE,
    genre_id INT NOT NULL REFERENCES genres(genre_id) ON DELETE CASCADE,
    PRIMARY KEY (movie_id, genre_id)
);

-- People Table
CREATE TABLE people (
    person_id SERIAL PRIMARY KEY,
    tmdb_id INT NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    known_for_department VARCHAR(50),
    biography TEXT,
    birthday DATE,
    deathday DATE,
    gender SMALLINT,
    profile_path VARCHAR(255),
    tmdb_popularity DECIMAL(10, 2),
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Movie People Junction Table
CREATE TABLE movie_people (
    movie_id INT NOT NULL REFERENCES movies(movie_id) ON DELETE CASCADE,
    person_id INT NOT NULL REFERENCES people(person_id) ON DELETE CASCADE,
    role VARCHAR(50) NOT NULL,
    PRIMARY KEY (movie_id, person_id, role)
);

-- Comparisons Table
CREATE TABLE comparisons (
    comparison_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    base_movie_id INT NOT NULL REFERENCES movies(movie_id) ON DELETE CASCADE,
    compared_movie_id INT NOT NULL REFERENCES movies(movie_id) ON DELETE CASCADE,
    preference SMALLINT NOT NULL, -- 1 = prefers base_movie, 0 = prefers compared_movie
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Ratings Table
CREATE TABLE ratings (
    rating_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    movie_id INT NOT NULL REFERENCES movies(movie_id) ON DELETE CASCADE,
    rating_score DECIMAL(4, 2) NOT NULL,
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Streaming Services Table
CREATE TABLE streaming_services (
    service_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    logo_path VARCHAR(255)
);

-- Movie Availability Table
CREATE TABLE movie_availability (
    movie_id INT NOT NULL REFERENCES movies(movie_id) ON DELETE CASCADE,
    service_id INT NOT NULL REFERENCES streaming_services(service_id) ON DELETE CASCADE,
    PRIMARY KEY (movie_id, service_id)
);

-- User Movies (e.g., watchlist or favorites)
CREATE TABLE user_movies (
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    movie_id INT NOT NULL REFERENCES movies(movie_id) ON DELETE CASCADE,
    rating DECIMAL(4, 2),
    review TEXT,
    watchlist BOOLEAN DEFAULT FALSE,
    watched BOOLEAN DEFAULT FALSE,
    favorited BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, movie_id)
);

-- Indexes
CREATE INDEX idx_users_name ON users(name);
CREATE INDEX idx_movies_title ON movies(title);
CREATE INDEX idx_movie_genres_movie_id ON movie_genres(movie_id);
CREATE INDEX idx_movie_genres_genre_id ON movie_genres(genre_id);
CREATE INDEX idx_people_name ON people(name);
CREATE INDEX idx_comparisons_user_id ON comparisons(user_id);
CREATE INDEX idx_ratings_user_movie_id ON ratings(user_id, movie_id);
CREATE INDEX idx_movie_availability_movie_id ON movie_availability(movie_id);
CREATE INDEX idx_movie_availability_service_id ON movie_availability(service_id);
CREATE INDEX idx_user_movies_user_id ON user_movies(user_id);
