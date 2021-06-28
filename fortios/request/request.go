package request

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/fortinetdev/forti-sdk-go/fortios/config"
)

// Request describes the request to FortiOS service
type Request struct {
	Config       config.Config
	Headers      *map[string][]string
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Path         string
	Params       *map[string][]string
	Data         *bytes.Buffer
}

// New creates reqeust object with http method, path, params and data,
// It will save the http request, path, etc. for the next operations
// such as sending data, getting response, etc.
// It returns the created request object to the gobal plugin client.
func New(c config.Config, method string, path string, params *map[string][]string, data *bytes.Buffer) *Request {
	var h *http.Request

	if data == nil {
		h, _ = http.NewRequest(method, "", nil)
	} else {
		h, _ = http.NewRequest(method, "", data)
	}

	r := &Request{
		Config:      c,
		Path:        path,
		HTTPRequest: h,
		Params:      params,
		Data:        data,
	}

	head := make(map[string][]string)
	r = buildHeaders(r, &head)

	return r
}

// Build Request header

// Build Request Sign/Login Info

// Send request data to FortiOS.
// If errors are encountered, it returns the error.
func (r *Request) Send() error {
	return r.Send2(15, false)
}

// Send2 request data to FortiOS with bIgnoreVdom.
// If errors are encountered, it returns the error.
func (r *Request) Send2(retries int, ignvdom bool) error {
	//Build FortiOS
	//build Sign/Login INfo

	//httpReq.URL, err = url.Parse(clientInfo.Endpoint + operation.HTTPPath)

	u := r.buildURL(ignvdom, "").String()

	var err error
	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		log.Fatal(err)
		return err
	}

	retry := 0
	for {
		//Send
		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		if errdo != nil {
			if strings.Contains(errdo.Error(), "x509: ") {
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}

			if retry > retries {
				err = fmt.Errorf("lost connection to firewall with error: %v", filterapikey(errdo.Error()))
				break
			}
			time.Sleep(time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

			retry++

		} else {
			break
		}
	}

	return err
}

func (r *Request) buildURL(ignvdom bool, vdomparam string) *url.URL {
	u := *r.HTTPRequest.URL
	u.Scheme = "https"
	u.Host = r.Config.FwTarget
	u.Path = r.Path

	q := u.Query()
	if r.Params != nil {
		for k, v := range *r.Params {
			for _, b := range v {
				q.Add(k, b)
			}
		}
	}
	if vdomparam != "" {
		q.Set("vdom", vdomparam)
	} else {
		q.Set("vdom", r.Config.Auth.Vdom)
	}

	if ignvdom {
		q.Del("vdom")
	}

	u.RawQuery = q.Encode()

	return &u
}

// Send3 request data to FortiOS with custom vdom.
// If errors are encountered, it returns the error.
func (r *Request) Send3(vdomparam string) error {
	retries := 15

	u := r.buildURL(false, vdomparam).String()
	var err error
	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		log.Fatal(err)
		return err
	}

	retry := 0
	for {
		//Send
		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		if errdo != nil {
			if strings.Contains(errdo.Error(), "x509: ") {
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}

			if retry > retries {
				err = fmt.Errorf("lost connection to firewall with error: %v", filterapikey(errdo.Error()))
				break
			}
			time.Sleep(time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

			retry++

		} else {
			break
		}
	}

	return err
}

func filterapikey(v string) string {
	re, _ := regexp.Compile("access_token=.*?\"")
	res := re.ReplaceAllString(v, "access_token=***************\"")

	return res
}

func buildHeaders(r *Request, h *map[string][]string) *Request {
	rh := &r.HTTPRequest.Header
	bearer := "Bearer " + r.Config.Auth.Token
	rh.Set("Authorization", bearer)
	rh.Set("Content-Type", "application/json")
	for k, v := range *h {
		for _, b := range v {
			rh.Add(k, b)
		}
	}

	return r
}

// SendWithSpecialParams sends request data to FortiOS with special URL paramaters.
// If errors are encountered, it returns the error.
// func (r *Request) SendWithSpecialParams(s, vdomparam string) error {
// 	r = buildURL(r, false, vdomparam)

// 	if s != "" {
// 		u += "&"
// 		u += s
// 	}

// 	var err error
// 	r.HTTPRequest.URL, err = url.Parse(u)
// 	if err != nil {
// 		log.Fatal(err)
// 		return err
// 	}

// 	retry := 0
// 	for {
// 		//Send
// 		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
// 		r.HTTPResponse = rsp
// 		if errdo != nil {
// 			if strings.Contains(errdo.Error(), "x509: ") {
// 				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
// 				break
// 			}

// 			if retry > 15 {
// 				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
// 				break
// 			}
// 			time.Sleep(time.Duration(1) * time.Second)
// 			log.Printf("Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

// 			retry++

// 		} else {
// 			break
// 		}
// 	}

// 	return err
// }
