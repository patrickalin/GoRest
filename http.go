// Package http help to realise some REST calls
package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type restHTTP struct {
	status string
	header http.Header
	body   []byte
}

// HTTP interface of the package http
type HTTP interface {
	GetBody() []byte
	Get(url string) (err error)
	GetWithHeaders(url string, headers map[string][]string) (err error)
	PostJSON(url string, buffer []byte) (err error)
}

//constante for the log File
const logFile = "http.log"

var log = logrus.New()

// New create the structure
func New(l *logrus.Logger) HTTP {
	initLog(l)
	logInfo(funcName(), "New http structure", "")
	return &restHTTP{}
}

// GetWithHeaders get with headers
func (r *restHTTP) GetWithHeaders(url string, headers map[string][]string) error {

	logDebug(funcName(), "Get with header", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("New Request %s\n Error : %v\n Advice : Check your url", url, err)
	}
	req.Header = headers

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Execute request %s\n Error : %s \n Advice : Check your internet connection or if the site is alive", url, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Response Status: %s", resp.Status)
	}

	//read Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ReadAll %v", err)
	}
	if body == nil {
		return fmt.Errorf("Body empty")
	}

	logDebug(funcName(), "Response", string(body))
	r.status = resp.Status
	r.header = resp.Header
	r.body = body
	return nil
}

// Get Rest without header
func (r *restHTTP) Get(url string) error {
	return r.GetWithHeaders(url, nil)
}

// Post Rest on the API
func (r *restHTTP) PostJSON(url string, buffer []byte) error {

	logDebug(funcName(), "Post", url)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(buffer))
	if err != nil {
		return fmt.Errorf("Post %v\n Rest Advice : Check your internet connection or if the site is alive", err)
	}

	defer resp.Body.Close()

	//read Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ReadAll %v", err)
	}

	log.Debugf("Body : \n %s", body)
	if body == nil {
		return fmt.Errorf("Body empty")
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return fmt.Errorf("Response Status: %s \nURL : %s \nBuffer %v", resp.Status, url, buffer)
	}

	r.status = resp.Status
	r.header = resp.Header
	r.body = body
	return nil
}

//GetBody return body
func (r *restHTTP) GetBody() []byte {
	return r.body
}

/* Fun private ------------------------------------ */

//Init the logger
func initLog(l *logrus.Logger) {
	if l != nil {
		log = l
		logDebug(funcName(), "Use the logger pass in New", "")
		return
	}

	log = logrus.New()
	logDebug(funcName(), "Create new logger", "")
	log.Formatter = new(logrus.TextFormatter)

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY, 0666)
	checkErr(err, funcName(), "Failed to log to file, using default stderr", "")

	log.Out = file
}
