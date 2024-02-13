package main

import "errors"

func getVal(table, key string, requestChannel chan<- query) (string, error) {
	q := query{
		verb:          get,
		table:         table,
		key:           key,
		returnChannel: make(chan string),
	}

	// Send the request
	requestChannel <- q

	// Block until we receive the response
	ret := <-q.returnChannel

	// Return the response
	if ret != "" {
		return ret, nil
	} else {
		return "", errors.New("Not found")
	}
}

func setVal(table, key, val string, requestChannel chan<- query) { // Creates table and key as needed
	q := query{
		verb:          set,
		table:         table,
		key:           key,
		value:         val,
		returnChannel: make(chan string),
	}

	// Send the request
	requestChannel <- q

	// Block until we receive the response
	<-q.returnChannel
}

func delKey(table, key string, requestChannel chan<- query) { // Ignores missing tables or keys
	q := query{
		verb:          del,
		table:         table,
		key:           key,
		returnChannel: make(chan string),
	}

	// Send the request
	requestChannel <- q

	// Block until we receive the response
	<-q.returnChannel
}

func newTable(table string, requestChannel chan<- query) {
	q := query{
		verb:          set,
		table:         table,
		returnChannel: make(chan string),
	}

	// Send the request
	requestChannel <- q

	// Block until we receive the response
	<-q.returnChannel
}

func delTable(table string, requestChannel chan<- query) { // Ignore missing tables
	q := query{
		verb:          del,
		table:         table,
		returnChannel: make(chan string),
	}

	// Send the request
	requestChannel <- q

	// Block until we receive the response
	<-q.returnChannel
}
