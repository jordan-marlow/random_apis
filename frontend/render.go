package frontend

import (
	"html/template"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func renderTemplate(c *gin.Context, name string, data gin.H) {
	tmpl, err := template.ParseFiles(
		filepath.Join("frontend", "templates", "base.html"),
		filepath.Join("frontend", "templates", name),
	)
	if err != nil {
		c.String(500, "Template parsing error: %v", err)
		return
	}
	data["Title"] = name
	if err := tmpl.ExecuteTemplate(c.Writer, "base", data); err != nil {
		c.String(500, "Template execution error: %v", err)
	}
}
