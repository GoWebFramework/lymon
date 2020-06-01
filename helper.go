package lymon

import "net/url"

// FormToMAP convert r.Form into map[string]val
func FormToMAP(form url.Values) map[string]interface{} {

	result := map[string]interface{}{}

	for key, val := range form {
		if len(val) > 0 {
			result[key] = val[0] // there will be someone who will reference issue to this line
		}
	}

	return result
}
