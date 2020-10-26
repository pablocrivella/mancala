// Package openapi provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xXX2/bNhD/KgS3hw1QLMl20lXAHtphKPKwzUgT7CEwalo82WwlkiOpNEbg7z4cKcuW",
	"pfxbMqDABr9YPN7x/vx+d+QdzVWllQTpLM3uqM3XUDH/971ihn8UHPCDg82N0E4oSTP6jljBgaiCuDWQ",
	"JW6kEdVGaTBOgFfXIljsal6ugaCko+ut0YgKB5XXKZSpmKMZFdKdTWlEKyFFVVc0SyLqNhqCCFZg6Dai",
	"Fbs9D6pnrZgZwzYo1CXbgEGr3xsoaEa/i/chx0288Szs2kbUOmVg2HEvIj8sxQpj+HEwhp7rx+5uI2rg",
	"r1oY4DS7Dmnandp6O2/11PIz5A4d+9UYZS7AaiWtd7Cbb0Cx/9emsTFhnRFyhSa6uTnypDEwdPQHVg2c",
	"6OP+hHGn+MnK8o+CZtdDmau1BtMBzYpVO+Rso4drswfidr6NDo4dP3Zsqb6+2rGCD4NC8EPb+4Lvs27A",
	"1qV7DIEXYRdWqTbysd2XuOe4ggJZ6LXbQ6NOlbrJG6r0rGVLt9ayQUA/fpTsMtCgdyAHNr+fVSjqWXgu",
	"j2RIfjhnKLKLtgp9F0KyDus4IotkQX4mi5CQ9E8h7SIii/RgcbxbHPvFSwH4MfEfV5JDISTwxYhGFCT2",
	"ruskSqNxNJk/IbqIXjYo6HuLFSZuzRyRANwSp3zaiIRbd+x3z+cjf57iC6ZayEKhN6XIoWk/ARL0t/NL",
	"D1rhSv/JZM5KRt7NzmlEb8DY4Hg6SkYJblQaJNOCZnTilyKqmVt7kMWY+TA8lPWlQggyDP2c04z+YoA5",
	"+BAqjcUH694rvsGduZIOpFdiWpci92rxZ6vkfq71cR3wlg52yyAbD8iOe3hjZK/SB2BXx5kaQmfw3dy7",
	"Mk7SJwQCt6zSJQz03zBwr8+i9jc/nH5NuWatrw0lk/3MS/rd9TlGx/caxd5JJ2/46XJymp8UbxM4maaT",
	"s5O3SfHTyWmS8LN8Olmmy2LfuLLprhF6W235HmqKHhk+013S4DrJPXiw7dPpePy8TO9ma1tqkjNJpHJk",
	"CWRZMvlltC/+gGz+5BC6Q34gliupjcrBWrYsgYB0wvlrzmmSPIsGL3LiXDowkpXEgrkBQ3x+PC1sXVXM",
	"bFqytkORrXz+PniGz3FrYHt8J/gWHVrBQHO+AGcE3IAlbGep2xI+rtXXtiF0yJS8BpmS6M3u91wypd8S",
	"mSavSyZPomTar9fvypFC1ZJ/k5BEsNwHSBxDhlXgwBP9jgq0iqOJRrvK+PtVt4kfZrMdpHUdbmJHQ2Pu",
	"R12+7qftSnPmWpCT5cYPdCFX/j6SKxOwzXEFpzxpLnldLgQrLxyPjHOBIlbODgZlwUoLAy+8T0JyuB1+",
	"srHb8GQ7ffj51n8PNUb/2Qz9n/b/Hu1J7RHGn0b/l0xZ4YgHARGWCHnDSsFxwlooIXfAid9gCVTabf6b",
	"4zWQ/d7xinu9cmhmtSlpRtfOaZvFcRXu6KM1GPWlZlqPclXFNwjoriMzo3id+w98Le+tZHFcqpyVa2Vd",
	"lk7GkyFtvwMfz38HAAD//wg0yWhlEgAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
