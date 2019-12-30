// Copyright (c) 2019, Adam R. Hughes and the unifi-client contributors. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package client

import (
	"os"
	"path"
	"testing"
)

func TestReadResponse(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"OK", "ok.json", false},
		{"ErrLoginRequired", "err_login_required.json", true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open(path.Join("testdata", tt.path))
			if err != nil {
				t.Fatalf("failed to open: %v", err)
			}
			defer f.Close()

			if _, err := ReadResponse(f); (err != nil) != tt.wantErr {
				t.Errorf("got error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
