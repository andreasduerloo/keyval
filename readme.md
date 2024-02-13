# keyval

Keyval is a minimal, multithreaded, in-memory key-value store exposed over HTTP. It came about because of two reasons:

1. I needed a very basic tool to allow for CRUD-operations on simple data over HTTP for a different project, ideally something that would containerize nicely.
2. I was going through Adam Woodbeck's excellent [Network Programming with Go](https://nostarch.com/networkprogrammingwithgo), and I wanted to practice as I went along.

Go is the ideal language for this project. It is incredibly simple to set up this kind of web service using nothing but the standard library, and this has only become more true with the recent changes to the net/http.ServeMux routing patterns in the 1.22 update. On top of that, Go's goroutines allowed me to make this program multithreaded without any headaches. By making one goroutine owner of the data and having other goroutines interact with it through blocking channels, I can avoid mind-bending bugs due to concurrent reads and writes to the same data while still allowing multiple clients to connect at the same time.

I hope to add the following features:
- TLS support (should be trivial in Go, although you could just as easily put it behind something like [caddy](https://github.com/caddyserver/caddy) and enjoy automatic TLS)
- Data persistence by periodically (ticker?) writing to a file, and reading that file on startup
- If I ever get really ambitious, support for an active-passive or active-active setup

All in all it has been a great exercise, and I have learnt a lot about writing HTTP servers and using concurrency.

## How it works

Keyval stores data as key-value pairs in tables. The table names, keys, and values are all strings. That means you could also use it as some sort of a primitve document database. (You could store JSON or base64-encoded *anything* as the values, I guess.) Under the hood, it's all just a map of maps, with the inner maps being `map[string]string` and the outer one a `map[string]map[string]string`.

The data is exposed through HTTP, allowing for the following methods:
- GET (e.g. GET http://host/table/key) to, predictably, get the value for a given key in a given table. Returns 404 if the table does not exist or if the key does not exist within the table. Returns 200 and the value otherwise.
- PUT (e.g. PUT http://host/table/key/value, value is optional) to create tables, create values in new or existing tables, or update existing values. If anything in the path does not exist, keyval will create it rather than throw an error. (Should) always return 200.
- DELETE (e.g. DELETE http://host/table/key, key is optional) to delete entire tables or keys and their associated values. Again, if anything in the path does not exist, keyval won't complain. You could call it laziness, I call it idempotency. (Should) always return 200.
- HEAD, but only because Go sets it up for free if you register a GET handler.

Any other method will return status 405 (method not allowed), and an incomplete path (or a path with a trailing forward slash) will return a 404.

## Configuration

Keyval is configured trough the config.json file. You have the following fields (over half of which are not implemented yet, see higher).

```
{
    "Portnumber": 4555,
    "IdleTimeout": "2m",
    "ReadHeaderTimeout": "1m",
    "Tls": false,
    "Tlspath": "",
    "Persist": false,
    "Persistpath": ""
}
```
Keyval will listen on all available unicast IP addresses, but at least you get to pick the port.

