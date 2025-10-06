// Package account holds user account and session info
package account

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	Utils "github.com/rohitaryal/imageGO/internal/Utils"
	Types "github.com/rohitaryal/imageGO/pkg/Types"
)

type Account struct {
	User        Types.User
	Token       string
	TokenExpiry string
	Cookie      string
}

func (a *Account) RefreshSession(verbose bool) error {
	if verbose {
		fmt.Println("Trying to refresh session.")
	}
	session, err := fetchSession(a.Cookie, verbose)
	if err != nil {
		return err
	}
	a.User = session.User
	a.Token = session.AccessToken
	a.TokenExpiry = session.Expires

	return nil
}

func (a *Account) IsTokenExpired(verbose bool) bool {
	parsedTime, err := time.Parse(time.RFC3339, a.TokenExpiry)
	if err != nil {
		if verbose {
			fmt.Fprintf(os.Stderr, "Failed to parse time in RFC3339 format %s: %v", a.TokenExpiry, err)
		}
		return true
	}

	return parsedTime.Before(time.Now().UTC())
}

func fetchSession(cookie string, verbose bool) (*Types.SessionData, error) {
	req, err := http.NewRequest("GET", "https://labs.google/fx/api/auth/session", nil)
	if err != nil {
		if verbose {
			fmt.Fprint(os.Stderr, "Failed to make a new request\n")
		}
		return nil, err
	}

	temp := map[string]string{
		"Cookie": cookie,
	}
	headers := Utils.MergeMap(Types.DefaultHeader, temp)
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	session, err := Utils.Fetch(req, verbose)
	if err != nil {
		return nil, err
	}

	var parsedSession Types.SessionData
	err = json.Unmarshal([]byte(session), &parsedSession)
	if err != nil {
		if verbose {
			fmt.Fprintf(os.Stderr, "Failed to parse response to json: %s\n", session)
		}

		return nil, err
	}

	return &parsedSession, nil
}
