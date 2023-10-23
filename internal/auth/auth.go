package auth

import ("net/http"
"errors"
"strings"
)

// extracts api key from headers of http request
func GetAPIKey(headers http.Header) (string, error) {
  val := headers.Get("Authorization")
  if val == "" {
    return "", errors.New("missing authorization header")
  }

  vals := strings.Split(val, " ")

  if(len(vals) != 2) {
	return "", errors.New("malformed authorization header")
  }

  if vals[0]!= "ApiKey" {
    return "", errors.New("malformed authorization header")
  }

  return vals[1], nil

}