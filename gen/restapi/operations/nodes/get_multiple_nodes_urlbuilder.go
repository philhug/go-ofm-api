// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetMultipleNodesURL generates an URL for the get multiple nodes operation
type GetMultipleNodesURL struct {
	Db string

	Nodes string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetMultipleNodesURL) WithBasePath(bp string) *GetMultipleNodesURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetMultipleNodesURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetMultipleNodesURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/node/{db}/_multi"

	db := o.Db
	if db != "" {
		_path = strings.Replace(_path, "{db}", db, -1)
	} else {
		return nil, errors.New("Db is required on GetMultipleNodesURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/0.1"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	nodes := o.Nodes
	if nodes != "" {
		qs.Set("nodes", nodes)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetMultipleNodesURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetMultipleNodesURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetMultipleNodesURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetMultipleNodesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetMultipleNodesURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetMultipleNodesURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}