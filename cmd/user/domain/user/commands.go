package user

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/vardius/go-api-boilerplate/pkg/commandbus"
	"github.com/vardius/go-api-boilerplate/pkg/errors"
	"github.com/vardius/go-api-boilerplate/pkg/executioncontext"
)

// ChangeEmailAddress command
type ChangeEmailAddress struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

// OnChangeEmailAddress creates command handler
func OnChangeEmailAddress(repository Repository, db *sql.DB) commandbus.CommandHandler {
	fn := func(ctx context.Context, c *ChangeEmailAddress, out chan<- error) {
		// this goroutine runs independently to request's goroutine,
		// there for recover middlewears will not recover
		// recover from panic to prevent crash
		defer func() {
			if r := recover(); r != nil {
				out <- errors.Newf(errors.INTERNAL, "[CommandHandler] Recovered in %v", r)
			}
		}()

		var totalUsers int32

		row := db.QueryRowContext(ctx, `SELECT COUNT(distinctId) FROM users WHERE emailAddress = ?`, c.Email)
		err := row.Scan(&totalUsers)
		if err != nil {
			out <- errors.Wrap(err, errors.INTERNAL, "Could not ensure that email is not taken")
			return
		}

		if totalUsers != 0 {
			out <- errors.Wrap(err, errors.INVALID, "User with given email already registered")
			return
		}

		u := repository.Get(c.ID)
		err = u.ChangeEmailAddress(c.Email)
		if err != nil {
			out <- errors.Wrap(err, errors.INTERNAL, "Error when changing email address")
			return
		}

		out <- repository.Save(executioncontext.ContextWithFlag(context.Background(), executioncontext.LIVE), u)
	}

	return commandbus.CommandHandler(fn)
}

// RegisterWithEmail command
type RegisterWithEmail struct {
	Email string `json:"email"`
}

// OnRegisterWithEmail creates command handler
func OnRegisterWithEmail(repository Repository, db *sql.DB) commandbus.CommandHandler {
	fn := func(ctx context.Context, c *RegisterWithEmail, out chan<- error) {
		// this goroutine runs independently to request's goroutine,
		// there for recover middlewears will not recover
		// recover from panic to prevent crash
		defer func() {
			if r := recover(); r != nil {
				out <- errors.Newf(errors.INTERNAL, "[CommandHandler] Recovered in %v", r)
			}
		}()

		var totalUsers int32

		row := db.QueryRowContext(ctx, `SELECT COUNT(distinctId) FROM users WHERE emailAddress = ?`, c.Email)
		err := row.Scan(&totalUsers)
		if err != nil {
			out <- errors.Wrap(err, errors.INTERNAL, "Could not ensure that email is not taken")
			return
		}

		if totalUsers != 0 {
			out <- errors.Wrap(err, errors.INVALID, "User with given email already registered")
			return
		}

		id, err := uuid.NewRandom()
		if err != nil {
			out <- errors.Wrap(err, errors.INTERNAL, "Could not generate new id")
			return
		}

		u := New()
		err = u.RegisterWithEmail(id, c.Email)
		if err != nil {
			out <- errors.Wrap(err, errors.INTERNAL, "Error when registering new user")
			return
		}

		out <- repository.Save(executioncontext.ContextWithFlag(context.Background(), executioncontext.LIVE), u)
	}

	return commandbus.CommandHandler(fn)
}

// RegisterWithFacebook command
type RegisterWithFacebook struct {
	Email      string `json:"email"`
	FacebookID string `json:"facebookId"`
}

// OnRegisterWithFacebook creates command handler
func OnRegisterWithFacebook(repository Repository, db *sql.DB) commandbus.CommandHandler {
	fn := func(ctx context.Context, c *RegisterWithFacebook, out chan<- error) {
		// this goroutine runs independently to request's goroutine,
		// there for recover middlewears will not recover
		// recover from panic to prevent crash
		defer func() {
			if r := recover(); r != nil {
				out <- errors.Newf(errors.INTERNAL, "[CommandHandler] Recovered in %v", r)
			}
		}()

		var id, emailAddress, facebookID string

		row := db.QueryRowContext(ctx, `SELECT id, emailAddress, facebookId FROM users WHERE emailAddress = ? OR facebookId = ?`, c.Email, c.FacebookID)
		err := row.Scan(&id, &emailAddress, &facebookID)
		if err != nil {
			out <- errors.Wrap(err, errors.INTERNAL, "Could not ensure that user is not already registered")
			return
		}

		if facebookID == c.FacebookID {
			out <- errors.Wrap(err, errors.INVALID, "User facebook account already connected")
			return
		}

		var u *User
		if emailAddress == c.Email {
			u = repository.Get(uuid.MustParse(id))
			err = u.ConnectWithFacebook(c.FacebookID)
			if err != nil {
				out <- errors.Wrap(err, errors.INTERNAL, "Error when trying to connect facebook account")
				return
			}
		} else {
			id, err := uuid.NewRandom()
			if err != nil {
				out <- errors.Wrap(err, errors.INTERNAL, "Could not generate new id")
				return
			}

			u = New()
			err = u.RegisterWithFacebook(id, c.Email, c.FacebookID)
			if err != nil {
				out <- errors.Wrap(err, errors.INTERNAL, "Error when registering new user")
				return
			}
		}

		out <- repository.Save(executioncontext.ContextWithFlag(context.Background(), executioncontext.LIVE), u)
	}

	return commandbus.CommandHandler(fn)
}

// RegisterWithGoogle command
type RegisterWithGoogle struct {
	Email    string `json:"email"`
	GoogleID string `json:"googleId"`
}

// OnRegisterWithGoogle creates command handler
func OnRegisterWithGoogle(repository Repository, db *sql.DB) commandbus.CommandHandler {
	fn := func(ctx context.Context, c *RegisterWithGoogle, out chan<- error) {
		// this goroutine runs independently to request's goroutine,
		// there for recover middlewears will not recover
		// recover from panic to prevent crash
		defer func() {
			if r := recover(); r != nil {
				out <- errors.Newf(errors.INTERNAL, "[CommandHandler] Recovered in %v", r)
			}
		}()

		var id, emailAddress, googleID string

		row := db.QueryRowContext(ctx, `SELECT id, emailAddress, googleId FROM users WHERE emailAddress = ? OR googleId = ?`, c.Email, c.GoogleID)
		err := row.Scan(&id, &emailAddress, &googleID)
		if err != nil {
			out <- errors.Wrap(err, errors.INTERNAL, "Could not ensure that user is not already registered")
			return
		}

		if googleID == c.GoogleID {
			out <- errors.Wrap(err, errors.INVALID, "User google account already connected")
			return
		}

		var u *User
		if emailAddress == c.Email {
			u = repository.Get(uuid.MustParse(id))
			err = u.ConnectWithGoogle(c.GoogleID)
			if err != nil {
				out <- errors.Wrap(err, errors.INTERNAL, "Error when trying to connect google account")
				return
			}
		} else {
			id, err := uuid.NewRandom()
			if err != nil {
				out <- errors.Wrap(err, errors.INTERNAL, "Could not generate new id")
				return
			}

			u = New()
			err = u.RegisterWithGoogle(id, c.Email, c.GoogleID)
			if err != nil {
				out <- errors.Wrap(err, errors.INTERNAL, "Error when registering new user")
				return
			}
		}

		out <- repository.Save(executioncontext.ContextWithFlag(context.Background(), executioncontext.LIVE), u)
	}

	return commandbus.CommandHandler(fn)
}
