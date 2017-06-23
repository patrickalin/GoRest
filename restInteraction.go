// Package myRest help to realise some REST calls
package myRest

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

// RestHTTP interface of the package myRest
type RestHTTP interface {
	GetBody() []byte
	Get(url string) (err error)
	GetWithHeaders(url string, headers map[string][]string) (err error)
	PostJSON(url string, buffer []byte) (err error)
}

// MakeNew create the structure
func MakeNew() RestHTTP {
	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter)

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Info("Failed to log to file, using default stderr")
		return nil
	}
	log.Out = file
	return &restHTTP{}
}

const logFile = "callrest.log"

var log = logrus.New()

// GetWithHeaders get with headers
func (r *restHTTP) GetWithHeaders(url string, headers map[string][]string) (err error) {

	log.Infof("Get URL: %s", url)

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
		log.Errorf("Response Headers: %s", resp.Header)
		return fmt.Errorf("Response Status: %s", resp.Status)
	}

	//read Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ReadAll %v", err)
	}

	log.Debugf("Body : \n %s", body)
	if body == nil {
		return fmt.Errorf("Body empty")
	}

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

	log.Infof("URL Post : %s", url)
	log.Infof("Decode Post : %s", buffer)

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
		log.Errorf("Post response Headers: %v", resp.Header)
		return fmt.Errorf("Response Status: %s", resp.Status)
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
