package helpers

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/sirupsen/logrus"

	"github.com/restuwahyu13/gin-rest-api/schemas"
)

func ParseHtml(fileName string, data map[string]string) string {
	html, errParse := template.ParseFiles("templates/" + fileName + ".html")

	if errParse != nil {
		defer fmt.Println("parser file html failed")
		logrus.Fatal(errParse.Error())
	}

	body := schemas.SchemaHtmlRequest{To: data["to"], Token: data["token"]}

	buf := new(bytes.Buffer)
	errExecute := html.Execute(buf, body)

	if errExecute != nil {
		defer fmt.Println("execute html file failed")
		logrus.Fatal(errExecute.Error())
	}

	return buf.String()
}
