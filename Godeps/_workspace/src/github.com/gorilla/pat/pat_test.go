// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pat

import (
	"net/http"
	"testing"

	"github.com/DigitalInnovation/vitamns/Godeps/_workspace/src/github.com/gorilla/mux"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
}

func testMatch(t *testing.T, meth, pat, path string, ok bool, vars map[string]string) {
	r := New()
	switch meth {
	case "DELETE":
		r.Delete(pat, myHandler)
	case "GET":
		r.Get(pat, myHandler)
	case "POST":
		r.Post(pat, myHandler)
	case "PUT":
		r.Put(pat, myHandler)
	}
	req, _ := http.NewRequest(meth, "http://localhost"+path, nil)
	m := mux.RouteMatch{}
	if r.Match(req, &m) != ok {
		if ok {
			t.Errorf("Expected request to %q to match %q", path, pat)
		} else {
			t.Errorf("Expected request to %q to not match %q", path, pat)
		}
	} else if ok && vars != nil {
		registerVars(req, m.Vars)
		q := req.URL.Query()
		for k, v := range vars {
			if q.Get(k) != v {
				t.Errorf("Variable missing: %q (value: %q)", k, q.Get(k))
			}
		}
	}
}

func TestPatMatch(t *testing.T) {
	testMatch(t, "DELETE", "/foo/{name}", "/foo/bar", true, map[string]string{":name": "bar"})
	testMatch(t, "GET", "/foo/{name}", "/foo/bar/baz", true, map[string]string{":name": "bar"})
	testMatch(t, "POST", "/foo/{name}/baz", "/foo/bar/baz", true, map[string]string{":name": "bar"})
	testMatch(t, "PUT", "/foo/{name}/baz", "/foo/bar/baz/ding", true, map[string]string{":name": "bar"})
	testMatch(t, "GET", "/foo/x{name}", "/foo/xbar", true, map[string]string{":name": "bar"})
	testMatch(t, "GET", "/foo/x{name}", "/foo/xbar/baz", true, map[string]string{":name": "bar"})
}
