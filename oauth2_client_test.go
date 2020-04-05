package oauth2

import (
	"testing"
)

func Test_get_client_type(t *testing.T) {
	clientType := ClientTypeValueOf("public")
	expectedClientType := Public
	if expectedClientType != clientType {
		t.Errorf("The clientType %s is not equals to the expected value %s.\n", clientType, expectedClientType)
	}

	clientType = ClientTypeValueOf("confidential")
	expectedClientType = Confidential
	if expectedClientType != clientType {
		t.Errorf("The clientType %s is not equals to the expected value %s.\n", clientType, expectedClientType)
	}
}

func Test_client_type_to_string(t *testing.T) {
	clientType := Public.String()
	expectedString := "public"

	if expectedString != clientType {
		t.Errorf("The string value of clientType %s is not equals to the expected value %s.\n", clientType, expectedString)
	}
}

func Test_get_registration_error_type_as_string(t *testing.T) {
	registrationErrorType := Invalid_redirect_uri_registration.String()
	expectedValue := "invalid_redirect_uri"
	if expectedValue != registrationErrorType {
		t.Errorf("The string value of RegistrationErrorType %s is not equals to expected value %s.\n", registrationErrorType, expectedValue)
	}

	registrationErrorType = Uapproved_software_statement.String()
	expectedValue = "uapproved_software"
	if expectedValue != registrationErrorType {
		t.Errorf("The string value of RegistrationErrorType %s is not equals to expected value %s.\n", registrationErrorType, expectedValue)
	}

}
