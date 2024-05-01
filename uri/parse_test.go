package uri

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseScheme(t *testing.T) {

	t.Run("ValidUri", func(t *testing.T) {
		scheme, uriAfterScheme, err := parseScheme("http://example.com")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if scheme != "http" {
			t.Errorf("Scheme is not http: %s", scheme)
		}
		if uriAfterScheme != "//example.com" {
			t.Errorf("URI after scheme is not //example.com: %s", uriAfterScheme)
		}
	})

	t.Run("EmptyScheme", func(t *testing.T) {
		_, _, err := parseScheme("example.com")
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		assert.ErrorContains(t, err, "URI has no scheme")
	})

}

func TestParseAuthority(t *testing.T) {
	t.Run("userInfo@host:port", func(t *testing.T) {
		authority, remaining, err := parseAuthority("//user:pass@example.com:1234/path")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if authority == nil {
			t.Errorf("Authority should not be nil")
		}
		if authority.UserInfo != "user:pass" {
			t.Errorf("UserInfo is not user:pass: %s", authority.UserInfo)
		}
		if authority.Host != "example.com" {
			t.Errorf("Host is not example.com: %s", authority.Host)
		}
		if authority.Port != "1234" {
			t.Errorf("Port is not 1234: %s", authority.Port)
		}
		if remaining != "/path" {
			t.Errorf("Remaining is not /path: %s", remaining)
		}
	})

	t.Run("host:port", func(t *testing.T) {
		authority, remaining, err := parseAuthority("//example.com:1234/path")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if authority == nil {
			t.Errorf("Authority should not be nil")
		}
		if authority.UserInfo != "" {
			t.Errorf("UserInfo is not empty: %s", authority.UserInfo)
		}
		if authority.Host != "example.com" {
			t.Errorf("Host is not host: %s", authority.Host)
		}
		if authority.Port != "1234" {
			t.Errorf("Port is not 1234: %s", authority.Port)
		}
		if remaining != "/path" {
			t.Errorf("Remaining is not /path: %s", remaining)
		}
	})

	t.Run("host", func(t *testing.T) {
		authority, remaining, err := parseAuthority("//example.com/path")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if authority == nil {
			t.Errorf("Authority should not be nil")
		}
		if authority.UserInfo != "" {
			t.Errorf("UserInfo is not empty: %s", authority.UserInfo)
		}
		if authority.Host != "example.com" {
			t.Errorf("Host is not host: %s", authority.Host)
		}
		if authority.Port != "" {
			t.Errorf("Port is not empty: %s", authority.Port)
		}
		if remaining != "/path" {
			t.Errorf("Remaining is not /path: %s", remaining)
		}
	})

	t.Run("userInfo@:port", func(t *testing.T) {
		authority, remaining, err := parseAuthority("//user:pass@:1234/path")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if authority == nil {
			t.Errorf("Authority should not be nil")
		}
		if authority.UserInfo != "user:pass" {
			t.Errorf("UserInfo is not user:pass: %s", authority.UserInfo)
		}
		if authority.Host != "" {
			t.Errorf("Host is not empty: %s", authority.Host)
		}
		if authority.Port != "1234" {
			t.Errorf("Port is not 1234: %s", authority.Port)
		}
		if remaining != "/path" {
			t.Errorf("Remaining is not /path: %s", remaining)
		}
	})
}

func TestParsePath(t *testing.T) {

	t.Run("path?query#fragment", func(t *testing.T) {
		path, uriAfterPath, err := parsePath("/path?query#fragment")

		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if path != "/path" {
			t.Errorf("Path is not /path: %s", path)
		}

		if uriAfterPath != "?query#fragment" {
			t.Errorf("URI after path is not ?query#fragment: %s", uriAfterPath)
		}
	})
	t.Run("path", func(t *testing.T) {
		path, uriAfterPath, err := parsePath("/path")

		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if path != "/path" {
			t.Errorf("Path is not /path: %s", path)
		}

		if uriAfterPath != "" {
			t.Errorf("URI after path is not empty: %s", uriAfterPath)
		}
	})

	t.Run("invalid: path#fragment?query", func(t *testing.T) {
		_, _, err := parsePath("/path#fragment?query")
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		assert.ErrorContains(t, err, "URI has invalid fragment or query")
	})
}
