package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"

	"ace/utils/cert"

	"github.com/lucas-clemente/quic-go/http3"
)

const (
	URL1 = "https://127.0.0.1:8010/test"
	URL2 = "https://127.0.0.1:8010/test2"
)

func Test_aceServer(t *testing.T) {
	client := &http.Client{Transport:&http3.RoundTripper{
			TLSClientConfig: cert.GetCaTLSConfig(),
		},
	}

	r, err := client.Get(URL1)
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(r.Body)

	log.Println(r.Status)
	log.Println(string(body))

	r, err = client.Post(URL2, "", strings.NewReader("test"))
	if err != nil {
		panic(err)
	}

	body, _ = ioutil.ReadAll(r.Body)

	log.Println(r.Status)
	log.Println(string(body))
}

func Benchmark_aceServer(b *testing.B) {
	client := &http.Client{Transport:&http3.RoundTripper{
			TLSClientConfig: cert.GetCaTLSConfig(),
		},
	}

	for i := 0; i < b.N; i++ {
		_, err := client.Get(URL1)
		if err != nil {
			fmt.Errorf("get %s error", URL1)
			panic(err)
		}
	}
}
