package application

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/vardius/go-api-boilerplate/pkg/common/domain"
	"github.com/vardius/go-api-boilerplate/pkg/common/infrastructure/eventbus"
	"github.com/vardius/go-api-boilerplate/pkg/user/domain/user"
	"github.com/vardius/go-api-boilerplate/pkg/user/infrastructure/persistence"
	"github.com/vardius/go-api-boilerplate/pkg/user/infrastructure/proto"
)

// WhenUserWasRegisteredWithGoogle handles event
func WhenUserWasRegisteredWithGoogle(db *sql.DB, repository persistence.UserRepository) eventbus.EventHandler {
	fn := func(ctx context.Context, event domain.Event) {
		log.Printf("[EventHandler] %s", event.Payload)

		e := &user.WasRegisteredWithGoogle{}

		err := json.Unmarshal(event.Payload, e)
		if err != nil {
			log.Printf("[EventHandler] Error: %v", err)
			return
		}

		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			log.Printf("[EventHandler] Error: %v", err)
			return
		}
		defer tx.Rollback()

		u := &proto.User{
			Id:       e.ID.String(),
			Email:    e.Email,
			GoogleId: e.GoogleID,
		}

		err = repository.Add(ctx, u)
		if err != nil {
			log.Printf("[EventHandler] Error: %v", err)
			return
		}

		tx.Commit()
	}

	return eventbus.EventHandler(fn)
}
