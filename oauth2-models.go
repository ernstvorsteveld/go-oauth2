package oauth2

import (
	"net/url"
)

type StateType string

type ResponseType int

const (
	Code ResponseType = iota
	Token
)

func (r ResponseType) String() string {
	return [...]string{"code", "token"}[r]
}

var responseTypes = map[string]ResponseType{
	"code":  Code,
	"token": Token,
}

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

const (
	invalid_request Oauth2Error = iota
	unauthorized_client
	access_denied
	unsupported_response_type
	invalid_scope
	server_error
	temporarily_unavailable
)

type Oauth2Error int

func (d Oauth2Error) Code() string {
	return [...]string{
		"invalid_request",
		"unauthorized_client",
		"access_denied",
		"unsupported_response_type",
		"invalid_scope",
		"server_error",
		"temporarily_unavailable"}[d]
}

func (e Oauth2Error) Description() string {
	errorDescriptions := [...]string{
		"The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed.",
		"The client is not authorized to request an authorization code using this method.",
		"The resource owner or authorization server denied the request",
		"The authorization server does not support obtaining an authorization code using this method.",
		"The requested scope is invalid, unknown, or malformed.",
		"The authorization server encountered an unexpected condition that prevented it from fulfilling the request. (This error code is needed because a 500 Internal Server Error HTTP status code cannot be returned to the clientvia an HTTP redirect.)",
		"The authorization server is currently unable to handle the request due to a temporary overloading or maintenance of the server.  (This error code is needed because a 503 Service Unavailable HTTP status code cannot be returnedto the client via an HTTP redirect.)"}

	return errorDescriptions[e]
}

type ErrorDescriptionType string
type LocalStateType string

type ErrorResponseType struct {
	error_code        Oauth2Error
	error_description ErrorDescriptionType
	error_uri         url.URL
	state             LocalStateType
}

type PkceChallenge struct {
	code_challenge        string
	code_challenge_method string
}

type AuthorisationRequest struct {
	response_type ResponseType
	client_id     ClientIdType
	redirect_uri  url.URL
	scopes        []Scope
	state         StateType
	user_agent    string
	pkce          PkceChallenge
}

type AuthorisatonResponseType struct {
	code  string
	state StateType
}

type PkceVerifier struct {
	code_verifier string //Base64(sha356(code_verifier))
}

type AccessTokenRequestType struct {
	grant_type         GrantTypeType
	code               AuthorisationCodeType
	redirect_uri       url.URL
	client_id          ClientIdType
	authorisation_code AuthorisationCodeType
	pkce               PkceVerifier
}

type AccessTokenResponseType struct {
	redirect_uri       url.URL
	authorisation_code AuthorisationCodeType
	token              Oauth2Token
}
