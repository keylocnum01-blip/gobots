/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package thrift

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// Default to using the shared http client. Library users are
// free to change this global client or specify one through
// THttpClientOptions.

type THttpClient struct {
	client             *http.Client
	response           *http.Response
	url                *url.URL
	urls               string
	requestBuffer      *bytes.Buffer
	header             http.Header
	nsecConnectTimeout int64
	nsecReadTimeout    int64
	results            string
	print              bool
	body               []byte
	moreCompact        bool
}

type THttpClientTransportFactory struct {
	options THttpClientOptions
	url     string
}

func (p *THttpClientTransportFactory) GetTransport(trans TTransport) (TTransport, error) {
	if trans != nil {
		t, ok := trans.(*THttpClient)
		if ok && t.url != nil {
			return NewTHttpClientWithOptions(t.url.String(), nil, p.options)
		}
	}
	return NewTHttpClientWithOptions(p.url, nil, p.options)
}

type THttpClientOptions struct {
	// If nil, DefaultHttpClient is used
	Client *http.Client
}

func NewTHttpClientTransportFactory(url string) *THttpClientTransportFactory {
	return NewTHttpClientTransportFactoryWithOptions(url, THttpClientOptions{})
}

func NewTHttpClientTransportFactoryWithOptions(url string, options THttpClientOptions) *THttpClientTransportFactory {
	return &THttpClientTransportFactory{url: url, options: options}
}

func NewTHttpClientWithOptions(urlstr string, t *http.Transport, options THttpClientOptions) (TTransport, error) {
	parsedURL, err := url.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, 0, 1024)
	httpHeader := map[string][]string{"Content-Type": {"application/x-thrift"}}
	return &THttpClient{client: &http.Client{Transport: t}, url: parsedURL, requestBuffer: bytes.NewBuffer(buf), header: httpHeader}, nil
}

func NewTHttpClientHeader(urlstr string, cl *http.Client, hed http.Header) TTransport {
	buf := make([]byte, 0, 1024)
	return &THttpClient{client: cl, requestBuffer: bytes.NewBuffer(buf), header: hed}
}

func NewTHttpClient(urlstr string) (TTransport, error) {
	return NewTHttpClientWithOptions(urlstr, nil, THttpClientOptions{})
}

// Set the HTTP Header for this specific Thrift Transport
// It is important that you first assert the TTransport as a THttpClient type
// like so:
//
// httpTrans := trans.(THttpClient)
// httpTrans.SetHeader("User-Agent","Thrift Client 1.0")

func (p *THttpClient) SetMoreCompact(value bool) {
	p.moreCompact = value
}

func ModHttpClient(urlstr string, tr *http.Transport, headers http.Header) *THttpClient {
	parsedURL, _ := url.Parse(urlstr)
	return &THttpClient{client: &http.Client{Transport: tr}, url: parsedURL, urls: urlstr, requestBuffer: bytes.NewBuffer(make([]byte, 0, 512)), header: headers}
}

//func (p *THttpClient) SetProxy(ip string) {
//	proxyStr := ip
//	proxyURL, _ := url.Parse(proxyStr)
//	t.Proxy = http.ProxyURL(proxyURL)
//	p.client = &http.Client{Transport: t}
//}

func (p *THttpClient) Printres() {
	p.print = true
}

//func (p *THttpClient) Clone() TTransport {
//	buf := make([]byte, 0, 1024)
//	p.requestBuffer = bytes.NewBuffer(buf)
//	p.client = &http.Client{Transport: t}
//	return p
//}

func (p *THttpClient) GetBody() []byte {
	return p.body
}

func (p *THttpClient) SetHeader(key string, value string) {
	p.header.Add(key, value)
}

// Get the HTTP Header represented by the supplied Header Key for this specific Thrift Transport
// It is important that you first assert the TTransport as a THttpClient type
// like so:
//
// httpTrans := trans.(THttpClient)
// hdrValue := httpTrans.GetHeader("User-Agent")
func (p *THttpClient) GetHeader(key string) string {
	return p.header.Get(key)
}

// Deletes the HTTP Header given a Header Key for this specific Thrift Transport
// It is important that you first assert the TTransport as a THttpClient type
// like so:
//
// httpTrans := trans.(THttpClient)
// httpTrans.DelHeader("User-Agent")
func (p *THttpClient) DelHeader(key string) {
	p.header.Del(key)
}

func (p *THttpClient) Open() error {
	// do nothing
	return nil
}

func (p *THttpClient) IsOpen() bool {
	return p.response != nil || p.requestBuffer != nil
}

func (p *THttpClient) closeResponse() error {
	var err error
	if p.response != nil && p.response.Body != nil {
		// The docs specify that if keepalive is enabled and the response body is not
		// read to completion the connection will never be returned to the pool and
		// reused. Errors are being ignored here because if the connection is invalid
		// and this fails for some reason, the Close() method will do any remaining
		// cleanup.
		io.Copy(ioutil.Discard, p.response.Body)

		err = p.response.Body.Close()
	}

	p.response = nil
	return err
}

func (p *THttpClient) Close() error {
	if p.requestBuffer != nil {
		p.requestBuffer.Reset()
		p.requestBuffer = nil
	}
	return p.closeResponse()
}

func (p *THttpClient) Read(buf []byte) (int, error) {
	if p.response == nil {
		return 0, NewTTransportException(NOT_OPEN, "Response buffer is empty, no request.")
	}
	n, err := p.response.Body.Read(buf)
	if n > 0 && (err == nil || errors.Is(err, io.EOF)) {
		return n, nil
	}
	return n, NewTTransportExceptionFromError(err)
}

func (p *THttpClient) ReadByte() (c byte, err error) {
	if p.response == nil {
		return 0, NewTTransportException(NOT_OPEN, "Response buffer is empty, no request.")
	}
	return readByte(p.response.Body)
}

func (p *THttpClient) Write(buf []byte) (int, error) {
	//fmt.Println("Write",buf)
	if p.requestBuffer == nil {
		return 0, NewTTransportException(NOT_OPEN, "Request buffer is nil, connection may have been closed.")
	}
	return p.requestBuffer.Write(buf)
}

func (p *THttpClient) WriteByte(c byte) error {
	//fmt.Println("WriteByte", c)
	if p.requestBuffer == nil {
		return NewTTransportException(NOT_OPEN, "Request buffer is nil, connection may have been closed.")
	}
	return p.requestBuffer.WriteByte(c)
}

func (p *THttpClient) WriteString(s string) (n int, err error) {
	//fmt.Println("WriteString", s)
	if p.requestBuffer == nil {
		return 0, NewTTransportException(NOT_OPEN, "Request buffer is nil, connection may have been closed.")
	}
	return p.requestBuffer.WriteString(s)
}

func (p *THttpClient) ModFlush(ctx context.Context) ([]byte, error) {
	p.closeResponse()
	buf := p.requestBuffer
	p.requestBuffer = new(bytes.Buffer)
	hostUrl := p.url.String()
	req, err := http.NewRequest("POST", hostUrl, buf)
	if err != nil {
		return []byte{}, NewTTransportExceptionFromError(err)
	}
	req.Header = p.header
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	response, err := p.client.Do(req)
	if err != nil {
		return []byte{}, NewTTransportExceptionFromError(err)
	}
	if response.StatusCode != http.StatusOK {
		p.response = response
		p.closeResponse()
		return []byte{}, NewTTransportException(UNKNOWN_TRANSPORT_EXCEPTION, "HTTP Response code: "+strconv.Itoa(response.StatusCode))
	}
	res, _ := io.ReadAll(response.Body)
	return res, err
}

func (p *THttpClient) Flush(ctx context.Context) error {
	p.closeResponse()
	buf := p.requestBuffer
	p.requestBuffer = new(bytes.Buffer)
	hostUrl := p.url.String()
	req, err := http.NewRequest("POST", hostUrl, buf)
	if err != nil {
		return NewTTransportExceptionFromError(err)
	}
	req.Header = p.header
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	response, err := p.client.Do(req)
	if err != nil {
		return NewTTransportExceptionFromError(err)
	}
	if response.StatusCode != http.StatusOK {
		// Close the response to avoid leaking file descriptors. closeResponse does
		// more than just call Close(), so temporarily assign it and reuse the logic.
		p.response = response
		p.closeResponse()

		// TODO(pomack) log bad response
		return NewTTransportException(UNKNOWN_TRANSPORT_EXCEPTION, "HTTP Response code: "+strconv.Itoa(response.StatusCode))
	}
	p.response = response
	if p.moreCompact {
		p.body, _ = io.ReadAll(response.Body)
	}
	return nil
}

func (p *THttpClient) RemainingBytes() (num_bytes uint64) {
	len := p.response.ContentLength
	if len >= 0 {
		return uint64(len)
	}

	const maxSize = ^uint64(0)
	return maxSize // the truth is, we just don't know unless framed is used
}

// Deprecated: Use NewTHttpClientTransportFactory instead.
func NewTHttpPostClientTransportFactory(url string) *THttpClientTransportFactory {
	return NewTHttpClientTransportFactoryWithOptions(url, THttpClientOptions{})
}

// Deprecated: Use NewTHttpClientTransportFactoryWithOptions instead.
func NewTHttpPostClientTransportFactoryWithOptions(url string, options THttpClientOptions) *THttpClientTransportFactory {
	return NewTHttpClientTransportFactoryWithOptions(url, options)
}

// Deprecated: Use NewTHttpClientWithOptions instead.
func NewTHttpPostClientWithOptions(urlstr string, options THttpClientOptions) (TTransport, error) {
	return NewTHttpClientWithOptions(urlstr, nil, options)
}

// Deprecated: Use NewTHttpClient instead.
func NewTHttpPostClient(urlstr string) (TTransport, error) {
	return NewTHttpClientWithOptions(urlstr, nil, THttpClientOptions{})
}
