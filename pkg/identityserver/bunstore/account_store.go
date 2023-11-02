package store

import (
	"context"
	"time"

	"github.com/uptrace/bun"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	storeutil "go.thethings.network/lorawan-stack/v3/pkg/util/store"
)

// Account is the account model in the database.
type Account struct {
	bun.BaseModel `bun:"table:accounts,alias:acc"`

	Model
	SoftDelete

	UID string `bun:"uid,notnull"`

	AccountID   string `bun:"account_id,notnull"`
	AccountType string `bun:"account_type,notnull"` // user or organization
}

// BeforeAppendModel is a hook that modifies the model on SELECT and UPDATE queries.
func (m *Account) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	if err := m.Model.BeforeAppendModel(ctx, query); err != nil {
		return err
	}
	return nil
}

// GetOrganizationOrUserIdentifiers returns organization or user identifiers for the account.
func (m *Account) GetOrganizationOrUserIdentifiers() *ttnpb.OrganizationOrUserIdentifiers {
	if m == nil {
		return nil
	}
	switch m.AccountType {
	default:
		return nil
	case "organization":
		return (&ttnpb.OrganizationIdentifiers{OrganizationId: m.UID}).GetOrganizationOrUserIdentifiers()
	case "user":
		return (&ttnpb.UserIdentifiers{UserId: m.UID}).GetOrganizationOrUserIdentifiers()
	}
}

// EmbeddedAccount contains the account fields that are embedded into User and Organization
// when loading them through the user_accounts and organization_accounts views.
type EmbeddedAccount struct {
	ID string `bun:"id,notnull,scanonly"`

	CreatedAt time.Time  `bun:"created_at,notnull,scanonly"`
	UpdatedAt time.Time  `bun:"updated_at,notnull,scanonly"`
	DeletedAt *time.Time `bun:"deleted_at,scanonly"`

	UID string `bun:"uid,notnull,scanonly"`
}

func (s *baseStore) getAccountModel(
	ctx context.Context,
	accountType, uid string,
) (*Account, error) {
	model := &Account{}
	selectQuery := s.newSelectModel(ctx, model).
		Where("?TableAlias.account_type = ?", accountType).
		Where("?TableAlias.uid = ?", uid)

	if err := selectQuery.Scan(ctx); err != nil {
		err = storeutil.WrapDriverError(err)
		if errors.IsNotFound(err) {
			return nil, store.ErrAccountNotFound.WithAttributes(
				"account_type", accountType,
				"account_id", uid,
			)
		}
		return nil, err
	}

	return model, nil
}
