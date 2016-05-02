Namara
======

The official Go client for the Namara Open Data service. [namara.io](https://namara.io)

## Installation

```bash
go get github.com/namara-io/namara-go
```

## Usage

### Instantiation

You need a valid API key in order to access Namara (you can find it in your My Account details on namara.io).

```go
import "github.com/namara-io/namara-go"

nam := namara.New({YOUR_API_KEY}, false)
```

You can also optionally enable debug mode:

```go
nam := namara.New({YOUR_API_KEY}, true)
```

`SetAPIVersion` and `SetHost` are also available methods.

### Getting Data

To make a basic request to the Namara API you can call `Get` on your instantiated object and pass it the ID of the dataset you want and the ID of the version of the data set:

```go
res, err := nam.Get("5885fce0-92c4-4acb-960f-82ce5a0a4650", "en-1", nil)
if err != nil {
  //...
}
```

With a `nil` third options argument passed, this will return data with the Namara default offset (0) and limit (250) applied. To specify options, you can pass an options argument:

```go
options := namara.Options{
  Offset: "0",
  Limit: "150",
}

res, err := nam.Get("5885fce0-92c4-4acb-960f-82ce5a0a4650", "en-1", &options)
if err != nil {
  //...
}
```

The `Get` method returns an API response as a pointer to a `simplejson.Json` struct. See the [simplejson docs](https://godoc.org/github.com/bitly/go-simplejson) to find out which methods are available and how to interact with it.

### Options

All [Namara data options](https://namara.io/#/api) are supported.

**Basic options**

```go
options := namara.Options{
  Select: "market_nam,website",
  Where: "town = 'Toronto' AND nearby(geometry, 43.6, -79.4, 10km)",
  Offset: "0",
  Limit: "20",
}
```

**Aggregation options**
Only one aggregation option can be specified in a request, in the case of this example, all options are illustrated, but passing more than one in the options object will throw an error.

```go
options := namara.Options{
  Operation: "sum(p0)",
  Operation: "avg(p0)",
  Operation: "min(p0)",
  Operation: "max(p0)",
  Operation: "count(*)",
  Operation: "geocluster(p3, 10)",
  Operation: "geobounds(p3",
}
```

### Running Tests

From command line:

```bash
go test namara_test.go namara.go
```

### License

Apache License, Version 2.0
