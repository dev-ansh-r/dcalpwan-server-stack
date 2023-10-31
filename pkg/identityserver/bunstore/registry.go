package store

import (
	"context"

	"github.com/uptrace/bun"
)

var models []any

func registerModels(m ...any) {
	models = append(models, m...)
}

func init() {
	registerModels(
		&Account{},
		&APIKey{},
		&Application{},
		&Attribute{},
		&Client{},
		&ContactInfo{},
		&ContactInfoValidation{},
		&EndDevice{},
		&EndDeviceLocation{},
		&EUIBlock{},
		&Gateway{},
		&GatewayAntenna{},
		&Invitation{},
		&LoginToken{},
		&Membership{},
		&Notification{},
		&NotificationReceiver{},
		&ClientAuthorization{},
		&AuthorizationCode{},
		&AccessToken{},
		&Organization{},
		&Picture{},
		&User{},
		&UserSession{},
	)
}

// clear database tables for the given models.
// This should be used with caution.
func clear(ctx context.Context, db *bun.DB, models ...any) (err error) {
	md, err := getDBMetadata(ctx, db)
	if err != nil {
		return err
	}

	if md.Type == "CockroachDB" {
		if _, err = db.ExecContext(ctx, "SET SQL_SAFE_UPDATES = FALSE"); err != nil {
			return err
		}
	}
	for _, model := range models {
		if _, err = db.NewDelete().Model(model).Where("1=1").ForceDelete().Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}

// Clear database tables for all registered models.
// This should be used with caution.
func Clear(ctx context.Context, db *bun.DB) error {
	return clear(ctx, db, models...)
}
