// Copyright (c) 2019, Adam R. Hughes and the unifi-client contributors. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package client

import (
	"net/http"
	"testing"

	"github.com/tri-adam/unifi-client/option"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name         string
		opts         []option.ClientOption
		wantEndpoint string
		wantUsername string
		wantPassword string
		wantErr      bool
	}{
		{"NoOpts", []option.ClientOption{}, defaultEndpoint, defaultUsername, defaultPassword, false},
		{"Endpoint", []option.ClientOption{
			option.WithEndpoint("https://1.2.3.4:1234"),
		}, "https://1.2.3.4:1234", defaultUsername, defaultPassword, false},
		{"EndpointInvalid", []option.ClientOption{
			option.WithEndpoint(":"),
		}, "", "", "", true},
		{"Credentials", []option.ClientOption{
			option.WithCredentials("user", "pass"),
		}, defaultEndpoint, "user", "pass", false},
		{"CredentialsInvalid", []option.ClientOption{
			option.WithCredentials("", ""),
		}, "", "", "", true},
		{"Transport", []option.ClientOption{
			option.WithTransport(http.DefaultTransport),
		}, defaultEndpoint, defaultUsername, defaultPassword, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c, err := New(tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("got error %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got, want := c.Endpoint, tt.wantEndpoint; got != want {
					t.Errorf("got endpoint %v, want %v", got, want)
				}
				if got, want := c.credentials.Username, tt.wantUsername; got != want {
					t.Errorf("got username %v, want %v", got, want)
				}
				if got, want := c.credentials.Password, tt.wantPassword; got != want {
					t.Errorf("got password %v, want %v", got, want)
				}
			}
		})
	}
}
