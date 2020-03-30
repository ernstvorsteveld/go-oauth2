package oauth2

import (
	"github.com/ernstvorsteveld/go-password"
	"github.com/google/uuid"
	"net/url"
)

const (
	confidential ClientType = iota + 1
	public
)

func (d ClientType) String() string {
	return [...]string{"confidential", "public"}[d]
}

type ClientIdType uuid.UUID

type Scope string

type CodeType struct {
	code string
}

type ClientType int

type Client struct {
	clientType           ClientType
	id                   ClientIdType
	secret               password.Password
	redirect_uris        []url.URL // the URL MAY not contain a fragment and is required and MAY be more than on
	default_redirect_uri url.URL   // use when request does not contain redirect_uri
	scopes               []Scope
}

type TokenEndpointAuthMethodType int

const (
	none TokenEndpointAuthMethodType = iota
	client_secret_post
	client_secret_basic
)

type ClientMetadata struct {
	redirect_uris              []url.URL
	token_endpoint_auth_method TokenEndpointAuthMethodType
	grant_types                []GrantTypeType
	respose_types              ResponseType
}

type secondsSinds int64

type ClientInformation struct {
	client_id                ClientIdType
	client_secret            password.Password
	client_id_issued_at      secondsSinds
	client_secret_expires_at secondsSinds
}
