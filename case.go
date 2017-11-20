package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

// RequestResponse wraps request and response
type RequestResponse struct {
	Request  *Request
	Response *Response
}

// Case is a test file
type Case struct {
	InitJS string
	RR     []*RequestResponse
	Delay  int64
}

// NewCase return a new case
func NewCase(path string, delay int64) (*Case, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	bf := bytes.NewBuffer(b)
	js := ""
	var req *Request
	var res *Response
	c := &Case{
		Delay: delay,
		RR:    make([]*RequestResponse, 0),
	}
	for {
		l, err := bf.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			res = &Response{}
			res.JS = js
			c.RR = append(c.RR, &RequestResponse{
				Request:  req,
				Response: res,
			})
			return c, nil
		}
		if RequestBegin(l) {
			if c.InitJS == "" {
				c.InitJS = js
			} else {
				res = &Response{}
				res.JS = js
				c.RR = append(c.RR, &RequestResponse{
					Request:  req,
					Response: res,
				})
			}
			js = l
			continue
		}
		if ResponseBegin(l) {
			req = &Request{}
			req.JS = js
			js = ""
			continue
		}
		js += l
	}
}

// Run your case
func (c *Case) Run() error {
	e := func(r *Request, err error) error {
		return errors.New(fmt.Sprintf("%s on %s", err.Error(), r.Name))
	}
	if markdown {
		fmt.Printf("```\n%s```\n", c.InitJS)
	}
	_, err := VM.Run(c.InitJS)
	if err != nil {
		return err
	}
	v, err := VM.Get("url")
	if err != nil {
		return err
	}
	if !v.IsString() {
		return errors.New("Invalid url")
	}
	url, err := v.ToString()
	if err != nil {
		return err
	}
	client := NewHttpClient(url)
	for _, v := range c.RR {
		if err := v.Request.MakeStartLine(); err != nil {
			return e(v.Request, err)
		}
		if err := v.Request.Parse(); err != nil {
			return e(v.Request, err)
		}
		if err := v.Request.MakeHeader(); err != nil {
			return e(v.Request, err)
		}
		if err := v.Request.MakeQuery(); err != nil {
			return e(v.Request, err)
		}
		if v.Request.Method == "POST" || v.Request.Method == "PUT" {
			if err := v.Request.MakeBody(); err != nil {
				return e(v.Request, err)
			}
		}

		res, err := client.Do(v.Request)
		if err != nil {
			return e(v.Request, err)
		}
		if err := v.Response.CopyFrom(res); err != nil {
			return e(v.Request, err)
		}
		if err := v.Response.Parse(); err != nil {
			return e(v.Request, err)
		}
		if c.Delay != 0 {
			time.Sleep(time.Duration(c.Delay) * time.Millisecond)
		}
	}
	return nil
}
