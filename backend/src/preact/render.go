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

	return c.HTML(http.StatusOK, doc.String())
}
