package handlers

import (
	"github.com/go-openapi/runtime"
	ofmdb "github.com/philhug/go-ofm-api"
	"net/http"
)

func BasicAuthAuth(user string, pass string) (interface{}, error) {
	return user, nil
	//return nil, errors.NotImplemented("basic auth  (BasicAuth) has still not yet been implemented")
}

func (authorizer) Authorize(*http.Request, interface{}) error {
	return nil
}

type authorizer struct {
	rt *ofmdb.Runtime
}

func NewAuthorizer(rt *ofmdb.Runtime) runtime.Authorizer {
	return authorizer{rt: rt}
}
