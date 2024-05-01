package uri

import (
	"errors"
	"strings"
)

func ParseURI(uriStr string) (*URI, error) {
	scheme, uriAfterScheme, err := parseScheme(uriStr)
	if err != nil {
		return nil, err
	}

	authority, uriAfterAuthority, err := parseAuthority(uriAfterScheme)
	if err != nil {
		return nil, err
	}

	path, uriAfterPath, err := parsePath(uriAfterAuthority)
	if err != nil {
		return nil, err
	}

	query, uriAfterQuery, err := parseQuery(uriAfterPath)
	if err != nil {
		return nil, err
	}

	fragment, err := parseFragment(uriAfterQuery)
	if err != nil {
		return nil, err
	}

	uri := &URI{
		Scheme:    scheme,
		Authority: *authority,
		Path:      path,
		Query:     query,
		Fragment:  fragment,
	}

	return uri, nil
}

func parseScheme(uri string) (string, string, error) {
	// find index of first colon
	colonIndex := strings.Index(uri, ":")
	if colonIndex == -1 {
		return "", "", errors.New("URI has no scheme")
	}
	scheme := uri[0:colonIndex]
	remaining := uri[colonIndex+1:]
	return scheme, remaining, nil
}

// parseAuthority parses the authority part of the URI.
// If it does not start with "//" it will return nil as Authority
func parseAuthority(uriAfterScheme string) (*Authority, string, error) {
	// check if starts with "//"
	if !strings.HasPrefix(uriAfterScheme, "//") {
		return nil, uriAfterScheme, nil
	}

	var authority = &Authority{}

	// find first index of "/" after that
	// the authority part is everything before that
	//       //user:pass@example.com:1234/path
	indexOfAuthorityEnd := strings.Index(uriAfterScheme[2:], "/")
	authorityStr := uriAfterScheme[2 : indexOfAuthorityEnd+2]
	remaining := uriAfterScheme[indexOfAuthorityEnd+2:]
	//       user:pass@example.com:1234

	indexOfAt := strings.Index(authorityStr, "@")
	if indexOfAt == -1 {
		authority.UserInfo = ""
	} else {
		authority.UserInfo = authorityStr[:indexOfAt]
		authorityStr = authorityStr[indexOfAt+1:]
	}

	//  example.com:1234
	indexOfColon := strings.Index(authorityStr, ":")
	if indexOfColon == -1 {
		authority.Host = authorityStr
		authority.Port = ""
	} else {
		authority.Host = authorityStr[:indexOfColon]
		authority.Port = authorityStr[indexOfColon+1:]
	}

	return authority, remaining, nil

}

func parsePath(uriAfterAuthority string) (string, string, error) {
	indexOfQuery := strings.Index(uriAfterAuthority, "?")
	indexOfHash := strings.Index(uriAfterAuthority, "#")

	if indexOfQuery == -1 && indexOfHash == -1 { // no query, no fragment
		return uriAfterAuthority, "", nil
	}

	if indexOfQuery == -1 { // no query, but fragment
		return uriAfterAuthority[:indexOfHash], uriAfterAuthority[indexOfHash:], nil
	}

	if indexOfHash == -1 { // query, but no fragment
		return uriAfterAuthority[:indexOfQuery], uriAfterAuthority[indexOfQuery:], nil
	}

	// query and fragment
	if indexOfHash < indexOfQuery {
		return "", "", errors.New("URI has invalid fragment or query")
	} else {
		return uriAfterAuthority[:indexOfQuery], uriAfterAuthority[indexOfQuery:], nil
	}
}

func parseQuery(uriAfterPath string) (map[string]string, string, error) {
	// if no query, return empty map
	if !strings.HasPrefix(uriAfterPath, "?") {
		return map[string]string{}, uriAfterPath, nil
	}
	// TODO: parse the query into a map
	return map[string]string{}, uriAfterPath, nil
}

func parseFragment(uriAfterQuery string) (string, error) {
	// if no fragment, return empty string
	if !strings.HasPrefix(uriAfterQuery, "#") {
		return "", nil
	} else {
		return uriAfterQuery[1:], nil
	}
}
