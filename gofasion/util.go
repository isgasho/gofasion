package gofasion

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var encodingPlacer = strings.NewReplacer("\"", "", "\\u0026", "&", "\\u003c", "<", "\\u003e", ">", "\\u003d", "=")

func httpGet(targetUrl string, params url.Values) (bs []byte, err error) {
	resp, err := http.Get(targetUrl + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	return ret, err
}

func encodingParser(raw string) string {
	return encodingPlacer.Replace(raw)
}
