package preact

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Runtime struct {
	CachedSourceCode string
}

func (runtime *Runtime) CleanCache() {
	if runtime.CachedSourceCode == "" {
		return
	}

	log.Info("Clearing JS runtime cache...")

	runtime.CachedSourceCode = ""
}

func (runtime *Runtime) CheckForErrors() error {
	_, err := runtime.Render("/")
	return err
}

func (runtime *Runtime) Render(route string) (string, error) {
	body, err := runtime.SourceCode()
	if err != nil {
		return "", err
	}

	html, err := EvalJS(fmt.Sprintf(`
  var render;
  %s

  render("%s");

  function send (renderFn) {
    render = renderFn
  }`,
		body,
		route,
	))

	if err != nil {
		runtime.CleanCache()
		return "", err
	}

	return html, nil

}

func (runtime *Runtime) SourceCode() (string, error) {
	if runtime.CachedSourceCode != "" {
		return runtime.CachedSourceCode, nil
	}

	code, err := Browserify()
	if err != nil {
		return "", err
	}

	runtime.CachedSourceCode = code

	return runtime.CachedSourceCode, nil
}

func Browserify() (string, error) {
	var stderr bytes.Buffer
	cmd := exec.Command("make", "browserify-for-backend")
	cmd.Stderr = &stderr

	source, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("%s: %s", err, stderr.String())
	}

	return string(source), nil
}
