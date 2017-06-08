// Package myRest help to realise some REST calls
package myRest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	mylog "github.com/patrickalin/GoMyLog"
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
func MakeNew() (rest RestHTTP) {
	//mylog.Init(mylog.ERROR)
	return &restHTTP{}
}

// GetWithHeaders get with headers
func (r *restHTTP) GetWithHeaders(url string, headers map[string][]string) (err error) {

	mylog.Trace.Printf("Rest Get URL: %s", url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return fmt.Errorf("RestError :> %s \n\t Rest URL :> %s \n\t Rest Advice : Check your url", err, url)
	}

	req.Header = headers
	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("RestError :> %s \n\t Rest URL :> %s \n\t Rest Advice : Check your internet connection or if the site is alive", err, url)
	}

	defer resp.Body.Close()

	//read Body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("RestError :> %s \n\t Rest URL :> %s \n\t Rest Advice : Error with read Body", err, url)
	}

	mylog.Trace.Printf("Body : \n %s \n\n", body)

	if body == nil {
		return fmt.Errorf("RestError :> %s \n\t Rest URL :> %s \n\t Rest Advice : Error the body is null, error in the secret key in the config.json ?", err, url)
	}

	if resp.StatusCode != http.StatusOK {
		mylog.Trace.Printf("URL Get : %s", url)
		mylog.Trace.Printf("Get response Status: %s>", resp.Status)
		mylog.Trace.Printf("Get response Headers: %s>", resp.Header)
		mylog.Trace.Printf("Get response Body: %s>", string(body))
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("RestError :> %s \n\t Rest URL :> %s \n\t Rest Advice : Error Status Post", err, url)
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

	mylog.Trace.Printf("URL Post : %s", url)
	mylog.Trace.Printf("Decode Post : %s", buffer)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(buffer))
	if err != nil {
		return fmt.Errorf("RestError :> %s \n\t Rest URL :> %s \n\t Rest Advice : Check your internet connection or if the site is alive", err, url)
	}

	defer resp.Body.Close()

	//read Body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("RestError :> %s \n\t Rest URL :> %s \n\t Rest Advice : Error with read Body", err, url)
	}

	if body == nil {
		return fmt.Errorf("RestError :> Error the body is null  %s \n\t Rest URL :> %s \n\t Rest Advice : Error with read Body", err, url)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		fmt.Println("\n URL Post :>", url)
		fmt.Printf("Decode Post :> %s \n\n", buffer)
		fmt.Println("Post response Status:>", resp.Status)
		fmt.Println("Post response Headers:>", resp.Header)
		fmt.Println("Post response Body:>", string(body))
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		fmt.Println("Post response Status:>", resp.Status)
		return fmt.Errorf("RestError :> %s \n\t Rest URL :> %s \n\t Rest Advice : Error Status Post", err, url)
	}

	r.status = resp.Status
	r.header = resp.Header
	r.body = body

	return nil
}

func (r *restHTTP) GetBody() []byte {
	return r.body
}
