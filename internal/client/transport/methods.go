/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package transport

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
)

type Request interface {
	MarshalBinary() ([]byte, error)
}

type Response interface {
	UnmarshalBinary(b []byte) error
}

func (c *Client) Create(url string, request Request, response Response) error {
	return c.invokeAction(http.MethodPost, url, request, response)
}

func (c *Client) Update(url string, request Request, response Response) error {
	return c.invokeAction(http.MethodPut, url, request, response)
}

func (c *Client) Patch(url string, request Request, response Response) error {
	return c.invokeAction(http.MethodPatch, url, request, response)
}

func (c *Client) invokeAction(httpMethodType string, url string, request Request, response Response) error {
	requestURL := fmt.Sprintf("%s/%s", c.Host, strings.TrimPrefix(url, "/"))
	body, err := request.MarshalBinary()

	if err != nil {
		return errors.Wrap(err, "marshall request body")
	}

	headers := c.Headers.Clone()
	headers.Set(contentLengthKey, fmt.Sprintf("%d", len(body)))

	var resp *http.Response

	// nolint:bodyclose // response is being closed outside the switch block
	switch httpMethodType {
	case http.MethodPost:
		resp, err = c.post(requestURL, bytes.NewReader(body), headers)
		if err != nil {
			return errors.Wrap(err, "create")
		}
	case http.MethodPut:
		resp, err = c.put(requestURL, bytes.NewReader(body), headers)
		if err != nil {
			return errors.Wrap(err, "update")
		}
	case http.MethodPatch:
		resp, err = c.patch(requestURL, bytes.NewReader(body), headers)
		if err != nil {
			return errors.Wrap(err, "patch")
		}
	default:
		return errors.New("unsupported http method type invoked")
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "read %v response", httpMethodType)
	}

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("%s request failed with status : %v, response: %v", httpMethodType, resp.Status, string(respBody))
	}

	err = response.UnmarshalBinary(respBody)
	if err != nil {
		return errors.Wrap(err, "unmarshall")
	}

	return nil
}

func (c *Client) Delete(url string) error {
	requestURL := fmt.Sprintf("%s/%s", c.Host, strings.TrimPrefix(url, "/"))

	resp, err := c.delete(requestURL, c.Headers)
	if err != nil {
		return errors.Wrap(err, "delete")
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	respBody, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return clienterrors.ErrorWithHTTPCode(resp.StatusCode, errors.Errorf("delete request(%s) failed with status : %v, response: %v", url, resp.Status, string(respBody)))
	}

	return nil
}

func (c *Client) Get(url string, response Response) error {
	requestURL := fmt.Sprintf("%s/%s", c.Host, strings.TrimPrefix(url, "/"))

	resp, err := c.get(requestURL, c.Headers)
	if err != nil {
		return errors.Wrap(err, "get request")
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "read response")
	}

	if resp.StatusCode != http.StatusOK {
		return clienterrors.ErrorWithHTTPCode(resp.StatusCode, errors.Errorf("get request(%s) failed with status : %v, response: %v", url, resp.Status, string(respBody)))
	}

	err = response.UnmarshalBinary(respBody)
	if err != nil {
		return errors.Wrap(err, "unmarshall")
	}

	return nil
}
