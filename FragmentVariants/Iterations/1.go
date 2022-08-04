package Iterations

import (
	"github.com/Arion-Kun/GoZippy/FragmentVariants/Utilities"
	"net/http"
)

type Result1 struct{}

func (z Result1) TryParse(response *http.Response) (bool, *Utilities.ZippyFile) { // Unused for now
	return false, nil
}
