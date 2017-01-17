package helpers

import (
	"GoReadNote/logger"
	"bytes"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"path"
	"strings"
)

const layoutFile = "layout.tmpl"

func getTemplate(templates ...string) *template.Template {
	templates = append(templates, layoutFile)
	files := []string{}
	for _, fname := range templates {
		files = append(files, path.Join("./templates/", fname))

	}

	funcMap := template.FuncMap{
		"Split": strings.Split,
	}

	layout := template.New(layoutFile)
	layout.Delims("<%", "%/>")
	template.Must(layout.Funcs(funcMap).ParseFiles(files...))
	logger.ALogger().Debugf("8888:%v", layout)

	return layout
}

func Render(c *gin.Context, obj map[string]interface{}, templates ...string) {
	layout := getTemplate(templates...)

	if _, ok := obj["Title"]; !ok {
		obj["Title"] = "未命名"
	}

	result := bytes.NewBufferString("")

	err := layout.Execute(result, obj)
	if err != nil {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, err.Error())

	} else {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, result.String())

	}

}
