// Copyright (c) 2019, Adam R. Hughes and the unifi-client contributors. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package internal

import (
	"net/http"
	"testing"
)

func TestDialSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		s       ClientSettings
		wantErr bool
	}{
		{"ZeroValue", ClientSettings{}, false},
		{"Credentials", ClientSettings{
			Credentials: &ClientCredentials{Username: "user", Password: "pass"},
		}, false},
		{"CredentialsNoPassword", ClientSettings{
			Credentials: &ClientCredentials{Username: "user"},
		}, true},
		{"CredentialsNoUsername", ClientSettings{
			Credentials: &ClientCredentials{Password: "pass"},
		}, true},
		{"Endpoint", ClientSettings{
			Endpoint: "https://1.2.3.4:8765",
		}, false},
		{"EndpointBad", ClientSettings{
			Endpoint: ":",
		}, true},
		{"Transport", ClientSettings{
			Transport: http.DefaultTransport,
		}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("got error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
