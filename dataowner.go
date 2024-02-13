package main

type query struct {
	verb          Verb
	table         string
	key           string
	value         string
	returnChannel chan string
}

type Verb string

const (
	get = "GET"
	set = "SET"
	del = "DEL"
)

func dataOwner(requestChannel <-chan query) {
	// Set up the dataset
	data := make(map[string]map[string]string)

	for {
		req := <-requestChannel

		// This call blocks until it is completed
		handleRequest(req, &data)
	}
}

func handleRequest(q query, data *map[string]map[string]string) {
	switch q.verb {

	case get:
		if val, present := (*data)[q.table][q.key]; present {
			q.returnChannel <- val
		} else {
			q.returnChannel <- "" // 404
		}

	case set:
		if _, present := (*data)[q.table]; present {
			if q.value != "" {
				(*data)[q.table][q.key] = q.value
			} else {
				(*data)[q.table][q.key] = "" // Allows use as a set
			}
		} else {
			if q.value != "" {
				(*data)[q.table] = map[string]string{q.key: q.value}
			} else {
				(*data)[q.table] = make(map[string]string)
			}
		}
		q.returnChannel <- "OK" // 200

	case del:
		if _, present := (*data)[q.table]; present && q.key != "" {
			delete((*data)[q.table], q.key)
		} else {
			delete(*data, q.table)
		}
		q.returnChannel <- "OK" // 200
	}
}
