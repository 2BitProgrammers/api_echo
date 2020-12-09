// An HTTP Echo API which returns JSON info about the requested endpoints
// This is meant to be used for testing purposes only
//
package main

// This requires that you use the mux package
// go get github.com/gorilla/mux
//
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

// Global variables
const appName = "2bitprogrammers/api_echo"
const appVersion = "2018.11a"
const appListenPort = "1234"

// The data structure which is returned when the user hits the /health endpoint
//
type tHealthy struct {
	StatusCode    int      `json:"statusCode"`
	Headers       []string `json:"headers"`
	RequestMethod string   `json:"method"`
	RequestURI    string   `json:"uri"`
	Payload       string   `json:"payload"`
	ClientIP      string   `json:"clientIp"`
	ClientPort    string   `json:"clientPort"`
	UserAgent     string   `json:"userAgent"`
	Healthy       bool     `json:"healthy"`
}

// Retrieves the web browser client ip, port, and user agent
//
func getRemoteClientInfo(r *http.Request) (ip string, port string, userAgent string) {
	//ip := r.Header.Get("X-Forwarded-For")
	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Printf("[ERROR] userip: %q is not IP:port", r.RemoteAddr)
	}
	//userIp := net.ParseIP(ip)
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	userAgent = r.UserAgent()
	return ip, port, userAgent
}

// Retrieves an array of headers
//
func getHeaders(r *http.Request) []string {
	var sHeaders []string
	for k, v := range r.Header {
		s := fmt.Sprintf("%s::%s", k, v)
		sHeaders = append(sHeaders, s)
	}
	return sHeaders
}

// Retrieves the payload body
//
func getPayload(r *http.Request) string {
	sPayload, _ := ioutil.ReadAll(r.Body)
	return string(sPayload)
}

// Generic handler for all endpoints
//
func ratGetHttpEcho(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	ip, port, ua := getRemoteClientInfo(r)
	method := r.Method
	uri := r.RequestURI
	headers := getHeaders(r)
	payload := getPayload(r)
	fmt.Println(time.Now().Local(), "\t", method, "\t", "\t", uri, "\t", status, "\t", ip, "\t", ua)
	response := tHealthy{status, headers, method, uri, payload, ip, port, ua, true}
	joResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(joResponse)
}

// App - Main Entry Point
//
func main() {

	fmt.Printf("%s v%s\n", appName, appVersion)
	fmt.Println("www.2BitProgrammers.com\nCopyright (C) 2020. All Rights Reserved.\n")
	fmt.Println("Listening on:  " + appListenPort + "\n")
	http.HandleFunc("/", ratGetHttpEcho)
	log.Fatal(http.ListenAndServe(":"+appListenPort, nil))
}
