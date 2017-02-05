package preact

import (
	"fmt"
	. "github.com/azer/go-style"
	"gopkg.in/olebedev/go-duktape.v2"
	"regexp"
	"strings"
)

const headers = `
require = undefined
console = { log: print, info: print, error: print, warn: print, trace: print }
`

func EvalJS(code string) (string, error) {
	ctx := duktape.New()
	ctx.PushTimers()

	code = fmt.Sprintf("%s\n%s", headers, code)

	if err := ctx.PevalString(code); err != nil {
		derr := err.(*duktape.Error)
		PrintJSError(code, derr.Message, derr.LineNumber)
		return "", err
	}

	result := ctx.GetString(-1)
	ctx.DestroyHeap()

	return result, nil
}

func PrintJSError(code, message string, lineno int) {
	lines := strings.Split(code, "\n")
	fmt.Println(Style("bold", fmt.Sprintf("\n  > JavaScript Error: %s", Style("red", message))))

	prev := lineno - 2
	for prev > 0 && (lines[prev] == "" || isBrowserifyCode(lines[prev])) {
		prev--
	}

	next := lineno
	for next < len(lines) && (lines[next] == "" || isBrowserifyCode(lines[next])) {
		next++
	}

	fmt.Println(Style("reset", fmt.Sprintf("\n      %d. %s", prev+1, lines[prev])))
	fmt.Println(Style("red", fmt.Sprintf("      %d. %s", lineno, lines[lineno-1])))
	fmt.Println(Style("reset", fmt.Sprintf("      %d. %s\n", next+1, lines[next])))
}

func isBrowserifyCode(line string) bool {
	r, _ := regexp.Compile(`^\s*\}\,\{"[^\"]+":\d+`)
	return r.MatchString(line)
}
