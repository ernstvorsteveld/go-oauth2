package oauth2

import (
	"github.com/google/uuid"
)

type TokenValueType uuid.UUID

type Token struct {
	value    TokenValueType
	validity ValiditySpecificationType
}

type Oauth2Token struct {
	access_token  Token
	refresh_token Token
	client_id     ClientIdType
}
