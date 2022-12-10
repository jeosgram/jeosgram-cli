package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/jeosgram/jeosgram-cli/session"
)

func Urlencode(m httpValue) []byte {
	data := make(url.Values)
	for k, v := range m {
		data.Set(k, fmt.Sprint(v))
	}
	return []byte(data.Encode())
}

func requestOauthToken(data httpValue) (*http.Response, error) {
	url := apiURL + "/oauth/token"

	body, _ := json.Marshal(data)
	//body := Urlencode(data)
	req, _ := http.NewRequest(
		"POST",
		url,
		bytes.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)

	if debug {
		raw, _ := httputil.DumpRequestOut(req, true)
		fmt.Printf("\n-------\n%q\n-------\n", raw)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if debug /*&& !isOK(res)*/ {
		raw, _ := httputil.DumpResponse(res, true)
		fmt.Printf("\n-------\n%q\n-------\n", raw)
	}

	return res, nil
}

func errorResponse(data []byte) error {
	var v struct {
		ErrorMessage string `json:"error"`
	}
	_ = json.Unmarshal(data, &v)
	return errors.New(v.ErrorMessage)
}

// ----------------------------

const (
	msgRequiredMFA  = "Multi-factor authentication required"
	msgInvalidToken = "The access token provided is invalid"
)

var (
	ErrRequiredMFA  = errors.New(msgRequiredMFA)
	ErrInvalidToken = errors.New(msgInvalidToken)
)

func (api JeosgramAPI) LoginByPassword(username, password string) (*session.Token, string, error) {
	res, err := requestOauthToken(httpValue{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"grant_type":    "password",
		"username":      username,
		"password":      password,
		"expires_in":    86400,
	})
	if err != nil {
		return nil, "", err
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)

	if isMFA(res) {
		var v struct {
			MFAToken string `json:"mfa_token"`
		}
		_ = json.Unmarshal(data, &v)
		return nil, v.MFAToken, ErrRequiredMFA
	}

	if !isOK(res) {
		return nil, "", errorResponse(data)
	}

	var token session.Token
	_ = json.Unmarshal(data, &token)

	return &token, "", nil
}

func (api JeosgramAPI) LoginByMFAOtp(mfaToken, otp string) (*session.Token, error) {
	res, err := requestOauthToken(httpValue{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"grant_type":    "urn:custom:mfa-otp",
		"mfa_token":     mfaToken,
		"otp":           otp,
	})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)

	if !isOK(res) {
		return nil, errorResponse(data)
	}

	var token session.Token
	_ = json.Unmarshal(data, &token)
	return &token, nil

}

func (api JeosgramAPI) LoginByRefreshToken(refreshToken string) (*session.Token, error) {
	res, err := requestOauthToken(httpValue{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)

	if !isOK(res) {
		return nil, errorResponse(data)
	}

	var token session.Token
	_ = json.Unmarshal(data, &token)
	return &token, nil
}
