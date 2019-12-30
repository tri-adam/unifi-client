// Copyright (c) 2019, Adam R. Hughes and the unifi-client contributors. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package client

import (
	"encoding/json"
	"fmt"
	"io"
)

// Error represents an API error.
type Error struct {
	rc  string
	msg string
}

// Error returns a formatted error.
func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.rc, e.msg)
}

// ReadResponse reads a JSON response.
func ReadResponse(r io.Reader) (json.RawMessage, error) {
	var u struct {
		Meta struct {
			RC  string `json:"rc"`
			Msg string `json:"msg,omitempty"`
		} `json:"meta"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.NewDecoder(r).Decode(&u); err != nil {
		return nil, err
	}

	if u.Meta.RC != "ok" {
		return u.Data, &Error{rc: u.Meta.RC, msg: u.Meta.Msg}
	}
	return u.Data, nil
}
