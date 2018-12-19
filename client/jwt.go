// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Daikon-BE": jwt Resource Client
//
// Command:
// $ goagen
// --design=github.com/Kumi-comb/Daikon-BE/design
// --out=$(GOPATH)\src\github.com\Kumi-comb\Daikon-BE
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// SigninJWTPayload is the jwt signin action payload.
type SigninJWTPayload struct {
	Password       *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
	UserScreenName *string `form:"userScreenName,omitempty" json:"userScreenName,omitempty" yaml:"userScreenName,omitempty" xml:"userScreenName,omitempty"`
}

// SigninJWTPath computes a request path to the signin action of jwt.
func SigninJWTPath() string {

	return fmt.Sprintf("/jwt/signin")
}

// ID/Passペアで認証を行う
func (c *Client) SigninJWT(ctx context.Context, path string, payload *SigninJWTPayload, contentType string) (*http.Response, error) {
	req, err := c.NewSigninJWTRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSigninJWTRequest create the request corresponding to the signin action endpoint of the jwt resource.
func (c *Client) NewSigninJWTRequest(ctx context.Context, path string, payload *SigninJWTPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
