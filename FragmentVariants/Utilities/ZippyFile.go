package Utilities

import (
	"fmt"
	"io"
	"net/http"
)

type ZippyFile struct {
	Domain      string
	Id          string
	Key         int64
	EncodedName string
}

func (wf *ZippyFile) GetLink() string {
	return fmt.Sprintf("https://%s/d/%s/%d/DOWNLOAD", wf.Domain, wf.Id, wf.Key)
}
func (wf *ZippyFile) GetEncodedLink() string {
	return fmt.Sprintf("https://%s/d/%s/%d/%s", wf.Domain, wf.Id, wf.Key, wf.EncodedName)
}

func TryMakeZippyFile(response *http.Response, iterations []ZippyIteration) (bool, *ZippyFile) {

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(response.Body)
	for i := range iterations {
		parseSuccess, filePtr := iterations[i].TryParse(response)
		if parseSuccess {
			return true, filePtr
		}
	}
	return false, nil
}
