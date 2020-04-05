package oauth2

type GrantTypeType int

const (
	Authorisation_code_grant GrantTypeType = iota
	Implicit_grant
	Password_grant
	Client_credentials_grant
	Refresh_token_grant
	Jwt_bearer_grant
	Saml2_bearer_grant
)

func (r GrantTypeType) String() string {
	return [...]string{
		"authorisation_code",
		"implicit",
		"password",
		"client_credentials",
		"refresh_token",
		"urn:ietf:params:oauth:grant-type:jwt-bearer",
		"urn:ietf:params:oauth:grant-type:saml2-bearer"}[r]
}

var grantTypes = map[string]GrantTypeType{
	"authorisation_code": Authorisation_code_grant,
	"implicit":           Implicit_grant,
	"password":           Password_grant,
	"client_credentials": Client_credentials_grant,
	"refresh_token":      Refresh_token_grant,
	"urn:ietf:params:oauth:grant-type:jwt-bearer":   Jwt_bearer_grant,
	"urn:ietf:params:oauth:grant-type:saml2-bearer": Saml2_bearer_grant,
}
