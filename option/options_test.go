// Copyright (c) 2019, Adam R. Hughes and the unifi-client contributors. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package option

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/tri-adam/unifi-client/internal"
)

func TestOptions(t *testing.T) {
	tests := []struct {
		name         string
		option       ClientOption
		wantSettings internal.ClientSettings
	}{
		{"WithCredentials", WithCredentials("user", "pass"), internal.ClientSettings{
			Credentials: &internal.ClientCredentials{
				Username: "user",
				Password: "pass",
			},
		}},
		{"WithEndpoint", WithEndpoint("https://endpoint:8443"), internal.ClientSettings{
			Endpoint: "https://endpoint:8443",
		}},
		{"WithTransport", WithTransport(http.DefaultTransport), internal.ClientSettings{
			Transport: http.DefaultTransport,
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := internal.ClientSettings{}
			tt.option.Apply(&s)

			if got, want := s, tt.wantSettings; !reflect.DeepEqual(got, want) {
				t.Errorf("got settings %v, want %v", got, want)
			}
		})
	}
}
