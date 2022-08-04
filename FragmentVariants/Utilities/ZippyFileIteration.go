package Utilities

import (
	"net/http"
)

type ZippyIteration interface {
	TryParse(response *http.Response) (bool, *ZippyFile)
}
