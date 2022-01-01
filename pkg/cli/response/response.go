package response

import "encoding/json"

func GenerateJsonResponse(field, value string) string {
	generated := make(map[string]string)
	generated[field] = value
	jsonResponse, err := json.Marshal(generated)
	if err != nil {
		panic(err)
	}

	return string(jsonResponse)
}
