package helpers

import (
	"encoding/json"

	"github.com/sirupsen/logrus"

	"github.com/restuwahyu13/gin-rest-api/schemas"
)

func Strigify(payload interface{}) []byte {
	response, _ := json.Marshal(payload)
	return response
}

func Parse(payload []byte) schemas.SchemaResponses {
	var jsonResponse schemas.SchemaResponses
	err := json.Unmarshal(payload, &jsonResponse)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return jsonResponse
}
