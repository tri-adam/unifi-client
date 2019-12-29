// Copyright (c) 2019, Adam R. Hughes and the unifi-client contributors. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package internal

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// ClientCredentials specifies login information for the UniFi Controller.
type ClientCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ClientSettings contains information about how the client should interact with the UniFi Controller.
type ClientSettings struct {
	Credentials *ClientCredentials
	Endpoint    string
	Transport   http.RoundTripper
}

// Validate returns an error if s is invalid.
func (s *ClientSettings) Validate() error {

	// If credentials are supplied, both a username and password must be suplied.
	if c := s.Credentials; c != nil {
		if c.Username == "" {
			return errors.New("empty username")
		}
		if c.Password == "" {
			return errors.New("empty password")
		}
	}

	// If endpoint is supplied, it must be valid.
	if _, err := url.Parse(s.Endpoint); err != nil {
		return fmt.Errorf("invalid endpoint: %w", err)
	}

	return nil
}
