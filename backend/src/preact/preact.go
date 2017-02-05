package preact

import (
	"fmt"
	"github.com/azer/logger"
	"net/http"
	"os"
)

const REMOTE_CONTROL_URL = "http://localhost:9966"

var (
	log             = logger.New("preact")
	developmentMode = len(os.Getenv("DEVELOP")) > 0
	cache           = NewCache()
)

func SendRemoteCommand(cmd string) {
	endpoint := fmt.Sprintf("%s/commands/%s", REMOTE_CONTROL_URL, cmd)
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Error("Unable to send remote command.", logger.Attrs{
			"Error":    err,
			"Endpoint": endpoint,
		})
	}

	defer resp.Body.Close()

}
