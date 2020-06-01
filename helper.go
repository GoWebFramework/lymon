package lymon

import "net/url"

// FormToMAP convert r.Form into map[string]val
func FormToMAP(form url.Values) map[string]interface{} {

	result := map[string]interface{}{}

	for key, val := range form {
		result[key] = val
	}

	return result
}


