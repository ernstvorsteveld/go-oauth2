package oauth2

import (
	"fmt"
	"testing"
)

func Test_should_validate_password_with_bcrypt(t *testing.T) {
	fmt.Printf("Password: %s\n", "kanweg")
	t.Errorf("The passwords should not match.")
}
