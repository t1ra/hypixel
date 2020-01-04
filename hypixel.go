// Package hypixel provides Golang bindings to the Hypixel API.
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
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
	// APIKeyLength is the length of Hypixel API keys.
	APIKeyLength = 36
	// APIBaseURL is the base URL of the Hypixel API.
	APIBaseURL = "https://api.hypixel.net/"
)

var (
	errAPIKeyLength    = errors.New("bad api key length")
	errBadAPIArguments = errors.New("bad amount of arguments to apiRequest")
	// ErrThrottled is returned when a request fails because the API key used reached the maximum
	// number of queries per minute.
	ErrThrottled = errors.New("couldn't make request because key was throttled")
)

// Hypixel is the main struct of the API. It holds the API key for your session, which is
// set with the New() function.
type Hypixel struct {
	sync.RWMutex
	// Hypixel API keys are in the format of a UUID.
	// Example: 00000000-0000-0000-0000-000000000000
	APIKeys []struct {
		Name           string
		UsesLastMinute int
	}
}

// New creates a new instance of the Hypixel struct.
func New(apiKeys ...string) (*Hypixel, error) {

	session := &Hypixel{}

	for _, key := range apiKeys {
		if len(key) != APIKeyLength {
			return nil, errAPIKeyLength
		}

		keyData, err := session.Key(key)
		if err != nil {
			return session, err
		}

		session.APIKeys = append(session.APIKeys, struct {
			Name           string
			UsesLastMinute int
		}{key, keyData.QueriesInLastMinute})
	}

	return session, nil

}

// getKey() gets a random key from the available keys.
func (session *Hypixel) getKey() string {
	rand.Seed(time.Now().Unix())
	return session.APIKeys[rand.Intn(len(session.APIKeys))].Name
}

// apiRequest gets a random key from the available keys, then calls apiRequestCuetomKey with that
// key.
func (session *Hypixel) apiRequest(method string, fields ...string) ([]byte, error) {
	key := session.getKey()
	result, err := session.apiRequestCustomKey(method, key, fields...)

	if err != nil {
		return result, err
	}

	var dataInterface interface{}

	err = json.Unmarshal(result, &dataInterface)

	if err != nil {
		return result, err
	}

	if _, ok := dataInterface.(map[string]interface{})["throttled"]; ok {
		return result, ErrThrottled
	}

	return result, nil
}

// apiRequestCustomKey is very much like apiRequest, but uses a custom key instead of one set by
// the session. It's used by session.Key("<key>")
func (session *Hypixel) apiRequestCustomKey(method string, key string,
	fields ...string) ([]byte, error) {
	session.Lock()

	defer session.Unlock()

	method += ("?key=" + key)

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
