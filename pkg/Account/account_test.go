package account_test

import (
	"os"
	"testing"

	Account "github.com/rohitaryal/imageGO/pkg/Account"
	types "github.com/rohitaryal/imageGO/pkg/Types"
)

func TestRefreshSession(t *testing.T) {
	account := Account.Account{
		Cookie: os.Getenv("GOOGLE_COOKIE"),
	}

	// If cookie is not defined just fail
	if account.Cookie == "" {
		t.FailNow()
	}

	err := account.RefreshSession(false)
	if err != nil {
		t.Fatalf("Failed to refresh session: %v\n", err)
	}

	// No chance for any escape :)
	if account.Token == "" || account.TokenExpiry == "" || account.User.Email == "" || account.User.Image == "" || account.User.Name == "" {
		t.FailNow()
	}
}

func TestIsAccountExpired(t *testing.T) {
	stubAccount := Account.Account{
		User: types.User{
			Name:  "Crocodile Tears",
			Email: "greencroco@deadsea.com",
			Image: "https://deadseabois.com/profile/croco/profile.png",
		},
		TokenExpiry: "2000-10-05T14:15:49.000Z",
		Token:       "ABC",
	}

	if !stubAccount.IsTokenExpired(false) {
		t.FailNow()
	}

	stubAccount.TokenExpiry = "3000-10-05T14:15:49.000Z" // Will fail on this date lulz

	if stubAccount.IsTokenExpired(false) {
		t.FailNow()
	}
}
