// Copyright (c) 2019, Adam R. Hughes and the unifi-client contributors. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package option

import (
	"net/http"

	"github.com/tri-adam/unifi-client/internal"
)

// A ClientOption is an option for the UniFi Controller client.
type ClientOption interface {
	Apply(*internal.ClientSettings)
}

// WithCredentials returns a ClientOption that specifies the credentials to use to authenticate to
// the UniFi Controller.
func WithCredentials(username, password string) ClientOption {
	return withCredentials{
		&internal.ClientCredentials{
			Username: username,
			Password: password,
		},
	}
}

type withCredentials struct {
	*internal.ClientCredentials
}

func (w withCredentials) Apply(s *internal.ClientSettings) {
	s.Credentials = w.ClientCredentials
}

// WithEndpoint returns a ClientOption that overrides the default UniFi Controller endpoint.
func WithEndpoint(e string) ClientOption {
	return withEndpoint(e)
}

type withEndpoint string

func (w withEndpoint) Apply(s *internal.ClientSettings) {
	s.Endpoint = string(w)
}

// WithTransport returns a ClientOption that specifies the HTTP transport to use.
func WithTransport(t http.RoundTripper) ClientOption {
	return withTransport{t}
}

type withTransport struct {
	t http.RoundTripper
}

func (w withTransport) Apply(s *internal.ClientSettings) {
	s.Transport = w.t
}
