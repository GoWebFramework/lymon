package lymon

import (
	"net/http"
)

// FormToMAP convert r.Form into map[string]val
func FormToMAP(r *http.Request, maxMemory int64) map[string]interface{} {

	// default r.ParseForm() max memory are 10mb
	// https://golang.org/src/net/http/request.go#L1192
	r.ParseMultipartForm(maxMemory)

	result := map[string]interface{}{}
	for key, val := range r.Form {
		if len(val) > 0 {
			result[key] = val[0] // there will be someone who will reference issue to this line
		}
	}

	return result
}
