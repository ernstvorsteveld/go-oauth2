package oauth2_interactors

import (
	gateways "github.com/ernstvorsteveld/go-oauth2-gateway"
)

type ClientInteractor struct {
	ClientGateway gateways.ClientGateway
}
