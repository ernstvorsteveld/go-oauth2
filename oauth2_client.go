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