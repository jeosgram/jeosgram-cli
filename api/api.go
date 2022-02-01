package api

import (
	"bytes"
	"constraints"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"

	"gitlab.com/jeosgram-go/jeosgram-cli/session"
)

const debug = false

const Version = "v0.0.1"

//const accessToken = "AQABAZR192EDWdDSkLuVe_0cecwewl7QppDRwQ"

//const apiURL = "http://api.jeosgram.io"

const apiURL = "http://localhost:8080"

const (
	userAgent = "Jeosgram-CLI/" + Version

	clientID     = "jeosgram-cli"
	clientSecret = "jeosgram-cli"

	minApiDelay = 400 * time.Millisecond
)

func isOK(res *http.Response) bool {
	return res.StatusCode == http.StatusOK
}

func isMFA(res *http.Response) bool {
	return res.StatusCode == http.StatusForbidden // 403
}

func isUnauthorized(res *http.Response) bool {
	return res.StatusCode == http.StatusUnauthorized // 401
}

func hasBadToken(body []byte) bool {
	const msg = "invalid_token"
	return bytes.Index(body, []byte(msg)) != -1
}

// -------------------

type httpValue = map[string]any

type JeosgramAPI struct {
	token string
}

func NewJeosgramAPI(token string) *JeosgramAPI {
	return &JeosgramAPI{
		token: "Bearer " + token,
	}
}

func (api JeosgramAPI) post(uri string, data any) (*http.Response, error) {
	body, _ := json.Marshal(data)
	return api.raw("POST", uri, body)
}

func (api JeosgramAPI) get(uri string) (*http.Response, error) {
	return api.raw("GET", uri, nil)
}

func (api JeosgramAPI) raw(method, uri string, body []byte) (*http.Response, error) {
	start := time.Now()

	url := apiURL + uri
	req, _ := http.NewRequest(
		method,
		url,
		bytes.NewReader(body),
	)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Authorization", api.token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json") // voy a suponer q solo se envia json
	}

	if debug {
		raw, _ := httputil.DumpRequestOut(req, true)
		fmt.Printf("\n-------\n%q\n-------\n", raw)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if debug /*&& !isOK(res)*/ {
		isStream := res.Header.Get("Transfer-Encoding") == "chunked"
		raw, _ := httputil.DumpResponse(res, isStream)
		fmt.Printf("\n-------\n%q\n-------\n", raw)
	}

	if isUnauthorized(res) /*&& hasBadToken(data)*/ {
		// refrescar token y reintentar la peticion

		conf, _ := session.ReadConfig()
		token, err := api.LoginByRefreshToken(conf.RefreshToken)
		if err != nil {
			// msg requiere login
			panic(err)
		}
		session.SaveTokens(token)

		// reintento la peticion
		return NewJeosgramAPI(token.AccessToken).raw(method, uri, body)
	}

	duration := time.Since(start)
	delay := Max(minApiDelay-duration, 0)
	if debug {
		fmt.Printf("duration=%v delay=%v\n", duration, delay)
	}

	if delay > 0 {
		time.Sleep(delay)
	}

	return res, nil
}

// mover a std.Max ...
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
