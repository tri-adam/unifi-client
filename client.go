// Copyright (c) 2019, Adam R. Hughes and the unifi-client contributors. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package client

import (
	"net/http"
	"net/http/cookiejar"

	"github.com/tri-adam/unifi-client/internal"
	"github.com/tri-adam/unifi-client/option"
)

const defaultEndpoint = "https://unifi:8443"
const defaultUsername = "ubnt"
const defaultPassword = "ubnt"

// Client represents a UniFi client connection.
type Client struct {
	Endpoint    string
	httpClient  *http.Client
	credentials *internal.ClientCredentials
}

// New creates a new client.
func New(opts ...option.ClientOption) (*Client, error) {
	s, err := getSettings(opts)
	if err != nil {
		return nil, err
	}
	hc, err := getHTTPClient(s)
	if err != nil {
		return nil, err
	}
	c := Client{
		Endpoint:    getEndpoint(s),
		httpClient:  hc,
		credentials: getCredentials(s),
	}
	return &c, nil
}

// getSettings populates client options based on the supplied client settings.
func getSettings(opts []option.ClientOption) (*internal.ClientSettings, error) {
	var s internal.ClientSettings
	for _, opt := range opts {
		opt.Apply(&s)
	}
	if err := s.Validate(); err != nil {
		return nil, err
	}
	return &s, nil
}

// getHTTPClient returns an HTTP client based on the supplied settings.
func getHTTPClient(s *internal.ClientSettings) (*http.Client, error) {
	cj, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	c := http.Client{
		Transport: s.Transport,
		Jar:       cj,
	}
	return &c, nil
}

// getEndpoint returns the endpoint based on the supplied settings.
func getEndpoint(s *internal.ClientSettings) string {
	if s.Endpoint != "" {
		return s.Endpoint
	}
	return defaultEndpoint
}

// getCredentials returns credentials based on the supplied settings.
func getCredentials(s *internal.ClientSettings) *internal.ClientCredentials {
	if s.Credentials != nil {
		return s.Credentials
	}
	c := internal.ClientCredentials{
		Username: defaultUsername,
		Password: defaultPassword,
	}
	return &c
}
