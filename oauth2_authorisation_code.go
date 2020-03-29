package oauth2

import (
	"fmt"
	"github.com/ernstvorsteveld/go-random"
	"log"
	"time"
)

func newValiditySpecificationType(duration time.Duration) ValiditySpecificationType {
	v := ValiditySpecificationType{
		issued_at:      time.Now(),
		valid_until:    time.Now().Add(duration),
		valid_duration: duration,
	}
	return v
}

var defaultDuration, _ = time.ParseDuration("10s")

func fmtDuration(d time.Duration) string {
	return fmt.Sprintf("%v", d)
}

func NewAuthorisationCodeType(validPeriod *string) AuthorisationCodeType {
	duration, err := duration(validPeriod)
	if err != nil {
		log.Fatal(err)
	}
	var ac AuthorisationCodeType = AuthorisationCodeType{
		code:     randomstring.String(15),
		validity: newValiditySpecificationType(duration),
	}
	return ac
}

func duration(durationString *string) (time.Duration, error) {
	var duration = defaultDuration
	if durationString == nil {
		return duration, nil
	}
	duration, err := time.ParseDuration(*durationString)
	if err != nil {
		return defaultDuration, fmt.Errorf("Parsing %s failed. It is not a duration.", *durationString)
	}

	return duration, nil
}

type AuthorisationCode interface {
	GetCode() string
	IsValid() bool
	String() string
}

func (c AuthorisationCodeType) GetCode() string {
	return c.code
}

func (c AuthorisationCodeType) IsValid() bool {
	return time.Now().Before(c.validity.valid_until)
}

func (c AuthorisationCodeType) String() string {
	return fmt.Sprintf("AuthorisationCode[%s] issued at[%s], duration [%s] and expires [%s].",
		c.code,
		c.validity.issued_at.Format(time.RFC3339),
		fmtDuration(c.validity.valid_duration),
		c.validity.valid_until.Format(time.RFC3339))
}
