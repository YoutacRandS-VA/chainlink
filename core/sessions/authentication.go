package sessions

import (
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/v2/core/auth"
	"github.com/smartcontractkit/chainlink/v2/core/bridges"
)

// Application config constant options
type AuthenticationProviderName string

const (
	LocalAuth AuthenticationProviderName = "local"
	LDAPAuth  AuthenticationProviderName = "ldap"
)

// ErrUserSessionExpired defines the error triggered when the user session has expired
var ErrUserSessionExpired = errors.New("session missing or expired, please login again")

// ErrNotSupported defines the error where interface functionality doesn't align with a Read Only LDAP server
var ErrNotSupported = errors.New("functionality not supported with read only LDAP server")

//go:generate mockery --quiet --name LocalAdminUsersORM --output ./mocks/ --case=underscore

// LocalAdminUsersORM is the interface that defines the functionality implemented by the local
// users/sessions ORM containing local admin CLI actions. This is separate from the AuthenticationProvider,
// as local admin management (ie initial core node setup, initial admin user creation), is always
// required no matter what the pluggable AuthenticationProvider implementation is.
type LocalAdminUsersORM interface {
	ListUsers() ([]User, error)
	CreateUser(user *User) error
	FindUser(email string) (User, error)
}

//go:generate mockery --quiet --name AuthenticationProvider --output ./mocks/ --case=underscore

// AuthenticationProvider is an interface that abstracts the required application calls to a user management backend
// Currently localauth (users table DB) or LDAP server (readonly)
type AuthenticationProvider interface {
	FindUser(email string) (User, error)
	FindUserByAPIToken(apiToken string) (User, error)
	ListUsers() ([]User, error)
	AuthorizedUserWithSession(sessionID string) (User, error)
	DeleteUser(email string) error
	DeleteUserSession(sessionID string) error
	CreateSession(sr SessionRequest) (string, error)
	ClearNonCurrentSessions(sessionID string) error
	CreateUser(user *User) error
	UpdateRole(email, newRole string) (User, error)
	SetAuthToken(user *User, token *auth.Token) error
	CreateAndSetAuthToken(user *User) (*auth.Token, error)
	DeleteAuthToken(user *User) error
	SetPassword(user *User, newPassword string) error
	TestPassword(email, password string) error
	Sessions(offset, limit int) ([]Session, error)
	GetUserWebAuthn(email string) ([]WebAuthn, error)
	SaveWebAuthn(token *WebAuthn) error

	FindExternalInitiator(eia *auth.Token) (initiator *bridges.ExternalInitiator, err error)
}
