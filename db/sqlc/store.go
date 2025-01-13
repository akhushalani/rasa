package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	decimal "github.com/shopspring/decimal"
)

type Store struct {
	*Queries
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return er
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}

type RegisterUserParams struct {
	Name         string
	Email        string
	PasswordHash string
}

func (store *Store) RegisterUser(ctx context.Context, arg RegisterUserParams) error {
	return store.execTx(ctx, func(q *Queries) error {
		// Create a new user
		_, err := q.CreateUser(ctx, CreateUserParams{
			Name:         pgtype.Text{String: arg.Name, Valid: true},
			Email:        arg.Email,
			PasswordHash: arg.PasswordHash,
		})
		return err
	})
}

type AddMovieWithGenresParams struct {
	TmdbID         int32
	ImdbID         string
	Title          string
	Overview       string
	ReleaseDate    string
	PosterPath     string
	BackdropPath   string
	TmdbPopularity decimal.Decimal
	Genres         []int32 // List of genre IDs to associate with the movie
}

func (store *Store) AddMovieWithGenres(ctx context.Context, arg AddMovieWithGenresParams) error {
	return store.execTx(ctx, func(q *Queries) error {
		// Add the movie
		movie, err := q.CreateMovie(ctx, CreateMovieParams{
			TmdbID:   arg.TmdbID,
			ImdbID:   pgtype.Text{String: arg.ImdbID, Valid: true},
			Title:    arg.Title,
			Overview: pgtype.Text{String: arg.Overview, Valid: true},
			ReleaseDate: func() pgtype.Date {
				t, err := time.Parse("2006-01-02", arg.ReleaseDate)
				if err != nil {
					return pgtype.Date{Valid: false}
				}
				return pgtype.Date{Time: t, Valid: true}
			}(),
			PosterPath:     pgtype.Text{String: arg.PosterPath, Valid: true},
			BackdropPath:   pgtype.Text{String: arg.BackdropPath, Valid: true},
			TmdbPopularity: arg.TmdbPopularity,
		})
		if err != nil {
			return err
		}

		// Add genres for the movie
		for _, genreID := range arg.Genres {
			_, err := q.CreateMovieGenre(ctx, CreateMovieGenreParams{
				MovieID: movie.MovieID,
				GenreID: genreID,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
}

type RateMovieParams struct {
	UserID      int32
	MovieID     int32
	RatingScore decimal.Decimal
}

func (store *Store) RateMovie(ctx context.Context, arg RateMovieParams) error {
	return store.execTx(ctx, func(q *Queries) error {
		// Add or update the user's rating for the movie
		_, err := q.CreateRating(ctx, CreateRatingParams{
			MovieID:     arg.MovieID,
			UserID:      arg.UserID,
			RatingScore: arg.RatingScore,
		})
		if err != nil {
			// If the rating already exists, update it
			_, err = q.UpdateRating(ctx, UpdateRatingParams{
				MovieID:     arg.MovieID,
				UserID:      arg.UserID,
				RatingScore: arg.RatingScore,
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
}

type AddUserMovieParams struct {
	UserID    int32
	MovieID   int32
	Rating    decimal.Decimal
	Review    string
	Watchlist bool
	Watched   bool
	Favorited bool
}

func (store *Store) AddUserMovie(ctx context.Context, arg AddUserMovieParams) error {
	return store.execTx(ctx, func(q *Queries) error {
		// Add or update user movie entry
		_, err := q.CreateUserMovie(ctx, CreateUserMovieParams{
			UserID:    arg.UserID,
			MovieID:   arg.MovieID,
			Rating:    arg.Rating,
			Review:    pgtype.Text{String: arg.Review, Valid: true},
			Watchlist: pgtype.Bool{Bool: arg.Watchlist, Valid: true},
			Watched:   pgtype.Bool{Bool: arg.Watched, Valid: true},
			Favorited: pgtype.Bool{Bool: arg.Favorited, Valid: true},
		})
		if err != nil {
			// If the entry already exists, update it
			_, err = q.UpdateUserMovie(ctx, UpdateUserMovieParams{
				UserID:    arg.UserID,
				MovieID:   arg.MovieID,
				Rating:    arg.Rating,
				Review:    pgtype.Text{String: arg.Review, Valid: true},
				Watchlist: pgtype.Bool{Bool: arg.Watchlist, Valid: true},
				Watched:   pgtype.Bool{Bool: arg.Watched, Valid: true},
				Favorited: pgtype.Bool{Bool: arg.Favorited, Valid: true},
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
}

type DeleteUserAndAssociatedDataParams struct {
	UserID int32
}

func (store *Store) DeleteUserAndAssociatedData(ctx context.Context, arg DeleteUserAndAssociatedDataParams) error {
	return store.execTx(ctx, func(q *Queries) error {
		// Delete user and cascade-delete associated data
		err := q.DeleteUser(ctx, arg.UserID)
		if err != nil {
			return err
		}
		return nil
	})
}

type AddMovieAvailabilityParams struct {
	MovieID    int32
	ServiceIDs []int32 // List of streaming service IDs
}

func (store *Store) AddMovieAvailability(ctx context.Context, arg AddMovieAvailabilityParams) error {
	return store.execTx(ctx, func(q *Queries) error {
		// Add availability for the movie on multiple streaming services
		for _, serviceID := range arg.ServiceIDs {
			_, err := q.CreateMovieAvailability(ctx, CreateMovieAvailabilityParams{
				MovieID:   arg.MovieID,
				ServiceID: serviceID,
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
}

type AddPersonToMovieParams struct {
	MovieID  int32
	PersonID int32
	Role     string
}

func (store *Store) AddPersonToMovie(ctx context.Context, arg AddPersonToMovieParams) error {
	return store.execTx(ctx, func(q *Queries) error {
		// Add a person to a movie in a specific role
		_, err := q.CreateMoviePerson(ctx, CreateMoviePersonParams{
			MovieID:  arg.MovieID,
			PersonID: arg.PersonID,
			Role:     arg.Role,
		})
		return err
	})
}

type DeleteMovieAndAssociationsParams struct {
	MovieID int32
}

func (store *Store) DeleteMovieAndAssociations(ctx context.Context, arg DeleteMovieAndAssociationsParams) error {
	return store.execTx(ctx, func(q *Queries) error {
		// Delete the movie and all associated records
		err := q.DeleteMovie(ctx, arg.MovieID)
		if err != nil {
			return err
		}
		return nil
	})
}

type AddComparisonParams struct {
	UserID          int32
	BaseMovieID     int32
	ComparedMovieID int32
	Preference      int16 // 1 = Base preferred, 0 = Compared preferred
}

func (store *Store) AddComparison(ctx context.Context, arg AddComparisonParams) error {
	return store.execTx(ctx, func(q *Queries) error {
		// Add a comparison between two movies for a user
		_, err := q.CreateComparison(ctx, CreateComparisonParams{
			UserID:          arg.UserID,
			BaseMovieID:     arg.BaseMovieID,
			ComparedMovieID: arg.ComparedMovieID,
			Preference:      arg.Preference,
		})
		return err
	})
}
