// +build appenginevm

// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Google App Engine (Flex) specific implementation
// TODO Actually get working
package main

import (
	"google.golang.org/appengine"
)

func main() {
	appengine.Main()
}

/*
func handleWhois(ipAddr string) (*whoisResponse) {
	body, err := queryIpWhois(ipAddr)
	resp := &whoisResponse{
		Query: ipAddr,
		Body: body,
	}
	if err != nil {
		resp.Error = err.Error()
	}

	return resp
}

func handleReverseDns(ipAddr string) (*dnsResponse) {
	names, err := net.LookupAddr(ipAddr)
	resp := &dnsResponse{
		Query: ipAddr,
		Names: names,
	}
	if err != nil {
		resp.Error = err.Error()
	}

	return resp
}


func queryWhois(query, host string) (string, error) {
	request := &whois.Request{
		Query: query,
		Host: host,
	}
	if err := request.Prepare(); err != nil {
		return "", err
	}

	log.Printf("Whois requesting %s from %s", query, host)

	response, err := whoisClient.Fetch(request)
	if err != nil {
		return "", err
	}

	log.Printf("Whois response: %s", response.String())

	return response.String(), err
}


// queryIpWhois issues two whois queries, the first to find the right whois server, and the 2nd to
// that server.
func queryIpWhois(ipAddr string) (string, error) {
	response, err := queryWhois(ipAddr, "whois.iana.org")

	m, err := whois.ParseWhois(response)
	if err != nil {
		return "", err
	}

	host, found := m[WHOIS_KEY]
	if !found {
		return "", errors.New("no whois server returned")
	}

	return queryWhois(ipAddr, host)
}
*/
