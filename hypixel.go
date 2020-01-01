package hypixel

/*
 * Copyright 2020 t1ra

 * Permission to use, copy, modify, and/or distribute this software for any purpose with or without
 * fee is hereby granted, provided that the above copyright notice and this permission notice appear
 * in all copies.

 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH REGARD TO THIS
 * SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE
 * AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT,
 * NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE
 * OF THIS SOFTWARE.
 */

import (
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
)

const (
	// APIKeyLength is the length of Hypixel API keys.
	APIKeyLength = 36
	// APIBaseURL is the base URL of the Hypixel API.
	APIBaseURL = "https://api.hypixel.net/"
)

var (
	errAPIKeyLength    = errors.New("bad api key length")
	errBadAPIArguments = errors.New("bad amount of arguments to APIRequest")
)

// Hypixel is the main struct of the API. It is expanded upon in other source files.
type Hypixel struct {
	sync.RWMutex
	// Hypixel API keys are in the format of a UUID.
	// Example: 00000000-0000-0000-0000-000000000000
	APIKey string
}

// New creates a new instance of the Hypixel struct.
func New(apiKey string) (*Hypixel, error) {

	session := &Hypixel{
		APIKey: apiKey,
	}

	if len(session.APIKey) != APIKeyLength {
		return nil, errAPIKeyLength
	}

	return session, nil

}

// APIRequest makes request requestType to the Hypixel API.
func (session *Hypixel) APIRequest(method string, fields ...string) ([]byte, error) {
	session.Lock()

	defer session.Unlock()

	method += ("?key=" + session.APIKey)

	if len(fields)%2 != 0 {
		return nil, errBadAPIArguments
	}

	if len(fields) > 0 {
		for i, j := 0, 1; j < len(fields); i, j = i+1, j+1 {
			method += ("&" + fields[i] + "=" + fields[j])
		}
	}

	response, err := http.Get(method)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(response.Body)
}
