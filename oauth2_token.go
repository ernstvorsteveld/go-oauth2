package oauth2

import (
	"github.com/google/uuid"
)

type TokenValueType uuid.UUID
type UserIdType uuid.UUID

type Token struct {
	value    TokenValueType
	validity ValiditySpecificationType
}

type TokenType int

const (
	Access_token TokenType = iota
	Refresh_token
)

var tokenTypes = map[string]TokenType{
	"access_token":  Access_token,
	"refresh_token": Refresh_token,
}

func (t TokenType) TokenTypeCode() string {
	return [...]string{
		"access_token",
		"refresh_token",
	}[t]
}

func TokenTypeValueOf(val string) TokenType {
	tokenType := tokenTypes[val]
	return tokenType
}

type Oauth2TokenType struct {
	token_type TokenType
	token      Token
	client_id  ClientIdType
	user_id    UserIdType
}

func NewOauth2Token(token_type string) Oauth2TokenType {
	token := Oauth2TokenType{}
	return token
}

type Oauth2Token interface {
	IsValid() bool
}
