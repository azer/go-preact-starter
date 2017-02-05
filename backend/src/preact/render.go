package preact

import (
	"bytes"
	"github.com/labstack/echo"
	"html/template"
	"net/http"
)

var runtime *Runtime

func init() {
	runtime = &Runtime{}
}

func RenderContent(c echo.Context) (string, error) {
	return runtime.Render(c.Request().URL.Path)
}

func RenderPage(c echo.Context, options *RoutingOptions) error {
	if html, ok := cache.Get(c.Request().URL.Path); ok {
		return c.HTML(http.StatusOK, html)
	}

	content, err := RenderContent(c)
	if err != nil {
		return err
	}

	var doc bytes.Buffer
	err = templates.ExecuteTemplate(&doc, options.Template, struct {
		Title           string
		Content         interface{}
		DevelopmentMode bool
	}{
		options.Title,
		template.HTML(content),
		developmentMode,
	})

	if err != nil {
		log.Error("Failed to execute template. %v", err)
		return err
	}

	if !developmentMode {
		cache.Set(c.Request().URL.Path, doc.String())
	}

	return c.HTML(http.StatusOK, doc.String())
}
