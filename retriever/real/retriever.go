package real

import (
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) Get(url string) string {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	dumpResponse, err := httputil.DumpResponse(response, true)
	if err != nil {
		log.Println(err)
	}

	return string(dumpResponse)
}
