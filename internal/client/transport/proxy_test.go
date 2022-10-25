package transport

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func tmcProxySupport(t *testing.T) {
	// define origin server URL
	originServerURL, err := url.Parse("https://playground.tmc.vmware.com")
	if err != nil {
		t.Log("invalid origin server URL")
	}

	reverseProxy := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		t.Logf("[reverse proxy server] received request at: %s\n", time.Now())

		// set req Host, URL and Request URI to forward a request to the origin server
		req.Host = originServerURL.Host
		req.URL.Host = originServerURL.Host
		req.URL.Scheme = originServerURL.Scheme
		req.RequestURI = ""

		// save the response from the origin server
		originServerResponse, err := http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(rw, err)
			return
		}

		// return response to the client
		rw.WriteHeader(http.StatusOK)
		io.Copy(rw, originServerResponse.Body)
	})

	log.Fatal(http.ListenAndServe(":8081", reverseProxy))
}

func cspProxySupport(t *testing.T) {

	// define origin server URL
	originServerURL, err := url.Parse("https://console.cloud.vmware.com")
	if err != nil {
		t.Log("invalid origin server URL")
	}

	reverseProxy := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		t.Logf("[reverse proxy server] received request at: %s\n", time.Now())

		// set req Host, URL and Request URI to forward a request to the origin server
		req.Host = originServerURL.Host
		req.URL.Host = originServerURL.Host
		req.URL.Scheme = originServerURL.Scheme
		req.RequestURI = ""

		// save the response from the origin server
		originServerResponse, err := http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(rw, err)
			return
		}

		// return response to the client
		rw.WriteHeader(http.StatusOK)
		io.Copy(rw, originServerResponse.Body)
	})

	t.Log(http.ListenAndServe(":8082", reverseProxy))
}

func TestTMC(t *testing.T) {
	t.Parallel()
	t.Run("Terraform-tmc", func(t *testing.T) {
		handler := &proxy{}
		addr := "https://playground.tmc.vmware.com"
		t.Logf("Starting proxy server on  %v", addr)
		if err := http.ListenAndServe(addr, handler); err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	})
}

func TestCSP(t *testing.T) {
	t.Parallel()
	t.Run("Terraform-csp", func(t *testing.T) {
		handler := &proxy{}
		addr := "https://console.cloud.vmware.com"
		t.Logf("Starting proxy server on  %v", addr)
		if err := http.ListenAndServe(addr, handler); err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	})
}
