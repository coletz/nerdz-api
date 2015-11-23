package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/url"
	"strings"
)

// Gravatar returns the gravatar url of the given email
func Gravatar(email string) *url.URL {
	m := md5.New()
	io.WriteString(m, strings.ToLower(email))

	return &url.URL{
		Scheme: "https",
		Host:   "www.gravatar.com",
		Path:   "/avatar/" + fmt.Sprintf("%x", m.Sum(nil))}
}
