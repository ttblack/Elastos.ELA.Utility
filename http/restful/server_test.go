package restful

import (
	"bytes"
	"fmt"
	"github.com/elastos/Elastos.ELA.Utility/http/util"
	"github.com/stretchr/testify/assert"
	"net/http"
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestRestfulServer(t *testing.T) {
	s := NewServer(&Config{})

	methods := []string{
		"GET",
		"PUT",
		"PATCH",
		"DELETE",
	}

	// Invalid handler
	err := s.RegisterAction(http.MethodGet, "/api/test", nil)
	assert.EqualError(t, err, fmt.Sprintf("handler %T not match %s method", nil, http.MethodGet))

	// Invalid method
	err = s.RegisterAction("INVALID", "/api/test", nil)
	assert.EqualError(t, err, fmt.Sprintf("handler %T not match %s method", nil, "INVALID"))

	// Invalid url
	err = s.RegisterAction(http.MethodGet, "api/test/", func(util.Params) (interface{}, error) {
		return nil, nil
	})
	assert.EqualError(t, err, `resource url must begin with "/"`)

	err = s.RegisterAction(http.MethodGet, "/api/test/", func(util.Params) (interface{}, error) {
		return nil, nil
	})
	assert.EqualError(t, err, `resource url must not end with "/"`)

	// Invalid post url
	err = s.RegisterAction(http.MethodPost, "/api/test/:param", func(data []byte) (interface{}, error) {
		return nil, nil
	})
	assert.EqualError(t, err, "post method do not allow path parameters")

	// Register methods
	for _, method := range methods {
		err = s.RegisterAction(method, "/api/test", func(params util.Params) (interface{}, error) {
			t.Logf("%s request from path %s", method, "/api/test")
			return nil, nil
		})
		assert.NoError(t, err)
	}

	err = s.RegisterAction(http.MethodPost, "/api/test", func(data []byte) (interface{}, error) {
		t.Logf("POST request from path %s", "/api/test")
		return nil, nil
	})

	err = s.RegisterAction(http.MethodGet, "/api/test/:param", func(params util.Params) (interface{}, error) {
		t.Logf("GET request from path %s", "/api/test/:param")
		return nil, nil
	})
	assert.NoError(t, err)
}

func TestServer_ServeHTTP(t *testing.T) {
	s := NewServer(&Config{
		ServePort: 20336,
	})

	methods := []string{
		"GET",
		"PUT",
		"PATCH",
		"DELETE",
	}

	// Register methods
	for i, method := range methods {
		me := methods[i]
		err := s.RegisterAction(method, "/api/test", func(params util.Params) (interface{}, error) {
			t.Logf("%s request from path %s", me, "/api/test")
			return nil, nil
		})
		assert.NoError(t, err)
	}

	// Register methods with parameters
	for i, method := range methods {
		me := methods[i]
		err := s.RegisterAction(method, "/api/test/param/:param", func(params util.Params) (interface{}, error) {
			t.Logf("%s request from path %s with param %v", me, "/api/test/param", params)
			param := params["param"]
			if param != "param" {
				t.Errorf("expected param = param get %s", param)
			}
			return nil, nil
		})
		assert.NoError(t, err)
	}

	// Register methods with multiple parameters
	for i, method := range methods {
		me := methods[i]
		err := s.RegisterAction(method, "/api/test/params/:a/:b/:c", func(params util.Params) (interface{}, error) {
			t.Logf("%s request from path %s with params %v", me, "/api/test/params", params)
			a := params["a"]
			if a != "a" {
				t.Errorf("expected a = a get %s", a)
			}
			b := params["b"]
			if b != "b" {
				t.Errorf("expected b = b get %s", b)
			}
			c := params["c"]
			if c != "c" {
				t.Errorf("expected c = c get %s", c)
			}
			return nil, nil
		})
		assert.NoError(t, err)
	}

	// Register methods with separated multiple parameters
	for i, method := range methods {
		me := methods[i]
		err := s.RegisterAction(method, "/api/test/param/:a/param/:b", func(params util.Params) (interface{}, error) {
			t.Logf("%s request from path %s with params %v", me, "/api/test/param/:a/param/:b", params)
			a := params["a"]
			if a != "a" {
				t.Errorf("expected a = a get %s", a)
			}
			b := params["b"]
			if b != "b" {
				t.Errorf("expected b = b get %s", b)
			}
			return nil, nil
		})
		assert.NoError(t, err)
	}

	// Register post method.
	err := s.RegisterAction(http.MethodPost, "/api/test", func(data []byte) (interface{}, error) {
		t.Logf("POST request from path %s, data %s", "/api/test", data)
		if !bytes.Equal(data, []byte("data")) {
			t.Errorf("expected data = data get %s", data)
		}
		return nil, nil
	})
	assert.NoError(t, err)

	go s.Start()

	// request no parameters.
	for _, method := range methods {
		req, err := http.NewRequest(method, "http://127.0.0.1:20336/api/test", nil)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		_, err = http.DefaultClient.Do(req)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	}

	// request single parameter.
	for _, method := range methods {
		req, err := http.NewRequest(method, "http://127.0.0.1:20336/api/test/param/param", nil)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		_, err = http.DefaultClient.Do(req)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	}

	// request multiple parameters.
	for _, method := range methods {
		req, err := http.NewRequest(method, "http://127.0.0.1:20336/api/test/params/a/b/c", nil)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		_, err = http.DefaultClient.Do(req)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	}

	// request separated multiple parameters.
	for _, method := range methods {
		req, err := http.NewRequest(method, "http://127.0.0.1:20336/api/test/param/a/param/b", nil)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		_, err = http.DefaultClient.Do(req)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	}

	// post request
	_, err = http.Post("http://127.0.0.1:20336/api/test", "text/plain",
		bytes.NewReader([]byte("data")))
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	// not found
	resp, err := http.Get( "http://127.0.0.1:20336/api/notfound")
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Error("expecting not found get %", resp.Status)
	}

	select {
	case <-time.After(time.Second):
		s.Stop()
	}
}

func TestNewResource(t *testing.T) {
	url := "/api/test/:param/:height"

	// get URL root.
	index := strings.Index(url, ":")
	t.Log(index)
	root := url[:index-1]

	t.Log(root)

	var params []string
	matches := regexp.MustCompile(`:(\w+)`).FindAllStringSubmatch(url, -1)
	t.Log(matches)
	for _, v := range matches {
		params = append(params, v[1])
	}
	t.Log(params)

	url = "/api/test/param/height"
	match := regexp.MustCompile(`/api/test/(\w+)/(\w+)`).FindStringSubmatch(url)
	t.Log(match)
	params = []string{}
	for _, v := range match[1:] {
		params = append(params, v)
	}
	t.Log(params)

}
