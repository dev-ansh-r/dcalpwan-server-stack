package ttnpb_test

import (
	"testing"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestApplicationPublicSafe(t *testing.T) {
	a := assertions.New(t)

	a.So(((*Application)(nil)).PublicSafe(), should.BeNil)

	src := &Application{
		Ids:         &ApplicationIdentifiers{ApplicationId: "foo"},
		Name:        "Name",
		Description: "Description",
		Attributes:  map[string]string{"key": "value"},
	}
	safe := src.PublicSafe()

	a.So(safe.Name, should.BeEmpty)
	a.So(safe.Description, should.BeEmpty)
	a.So(safe.Attributes, should.BeEmpty)
}

func TestClientPublicSafe(t *testing.T) {
	a := assertions.New(t)

	a.So(((*Client)(nil)).PublicSafe(), should.BeNil)

	src := &Client{
		Ids:         &ClientIdentifiers{ClientId: "foo"},
		Name:        "Name",
		Description: "Description",
		Attributes:  map[string]string{"key": "value"},
	}
	safe := src.PublicSafe()

	a.So(safe.Name, should.NotBeEmpty)
	a.So(safe.Description, should.NotBeEmpty)
	a.So(safe.Attributes, should.BeEmpty)
}

func TestGatewayPublicSafe(t *testing.T) {
	a := assertions.New(t)

	a.So(((*Gateway)(nil)).PublicSafe(), should.BeNil)

	src := &Gateway{
		Ids:         &GatewayIdentifiers{GatewayId: "foo"},
		Name:        "Name",
		Description: "Description",
		Attributes:  map[string]string{"key": "value"},
	}
	safe := src.PublicSafe()

	a.So(safe.Name, should.NotBeEmpty)
	a.So(safe.Description, should.NotBeEmpty)
	a.So(safe.Attributes, should.BeEmpty)
}

func TestOrganizationPublicSafe(t *testing.T) {
	a := assertions.New(t)

	a.So(((*Organization)(nil)).PublicSafe(), should.BeNil)

	src := &Organization{
		Ids:         &OrganizationIdentifiers{OrganizationId: "foo"},
		Name:        "Name",
		Description: "Description",
		Attributes:  map[string]string{"key": "value"},
	}
	safe := src.PublicSafe()

	a.So(safe.Name, should.NotBeEmpty)
	a.So(safe.Description, should.BeEmpty)
	a.So(safe.Attributes, should.BeEmpty)
}

func TestUserPublicSafe(t *testing.T) {
	a := assertions.New(t)

	a.So(((*User)(nil)).PublicSafe(), should.BeNil)

	src := &User{
		Ids:         &UserIdentifiers{UserId: "foo"},
		Name:        "Name",
		Description: "Description",
		Attributes:  map[string]string{"key": "value"},
	}
	safe := src.PublicSafe()

	a.So(safe.Name, should.NotBeEmpty)
	a.So(safe.Description, should.NotBeEmpty)
	a.So(safe.Attributes, should.BeEmpty)
}
