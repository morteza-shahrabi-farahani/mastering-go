# Chapter 8: Building Web Services
The http.Response structure embodies the response from an HTTP request. Both http.Client and http.Transport return http.Response values once the response headers have been received.

The http.Request structure represents an HTTP request as constructed by a client in order to be sent or received by an HTTP server.

## Creating a web server
The net/http package offers functions and data types that allow you to develop powerful web servers and clients. The http.Set() and http.Get() methods can be used to make HTTP and HTTPS requests, whereas http.ListenAndServe() is used for creating web servers given the user-specified handeler function or functions that handle incoming requests.

The simplest way to define the supported endpoints, as well as the handler function that responds to each client request, is with the use of http.HandleFunc(), which can be called multiple times.

```
func myHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
    fmt.Printf("Served: %s\n", r.Host)
}

http.HandleFunc("/time", timeHandler)
http.HandleFunc("/", myHandler)
```

The http.ListenAndServe() call begins the HTTP server using the predefined port number.

```
err := http.ListenAndServe(PORT, nil)
if err != nil {
    fmt.Println(err)
    return
}
```

## Implementing the handlers
Usually, handlers are put in a separate package.

```
func deleteHandler(w http.ResponseWriter, r *http.Request) {
    // Get telephone
    paramStr := strings.Split(r.URL.Path, "/")
    fmt.Println("Path:", paramStr)
    if len(paramStr) < 3 {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintln(w, "Not found: "+r.URL.Path)
        return
    }
```

If we do no have enough parameters, we should send an error message back to the client with the desired HTTP code, which in this case is http.StatusNotFound.

```
log.Println("Serving:", r.URL.Path, "from", r.Host)
```

This is where the HTTP server sends data to log files - this mainly happens for debugging reasons.

```
mux := http.NewServeMux()
s := &http.Server{
    Addr: PORT,
    Handler: mux,
    IdleTimeout: 10 * time.Second,
    ReadTimeout: time.Second,
    WriteTimeout: time.Second,
}

mux.Handle("/list", http.HandlerFunc(listHandler))
mux.Handle("/insert/", http.HandlerFunc(insertHandler))
mux.Handle("/insert", http.HandlerFunc(insertHandler))

```

Here, we store the parameters of the HTTP server in the http.Server structure and use our own http.NewServeMux() instead of the default one.

```
err = s.ListenAndServe()
if err != nil {
    fmt.Println(err)
    return
}
```
The ListenAndServe() method starts the HTTP server using the parameters defined previously in the http.Server structure.

\* The http package uses multiple goroutines for interacting with clients - in practice, this means that you application runs concurrently!

## Exposing metrics to Prometeus
The list of supported data types for metrics is the following:

Counter: Counters are usually used for representing cumulative values such as the number of requests served so far, the total number of errors, etc.

Gauge: Gauges are usually used for representing values that can go up or down such as the number of requests, time durations, etc.

Histogram: A histogram is used for sampling observations and creating counts and buckets. Histograms are usually used for counting request durations, response times, etc. 

Summary: A summary is like a histogram but can also calculate quantiles over sliding windows that work with times.

The runtime/metrics package makes metrics exported by the Go runtime available to the developer. If you want to collect all available metrics, you should use metrics.All().

\* You might ask, "why not use a normal Go binary instead of a Docker image?" The answer is simple: Docker images can be put in docker-compose.yml files and can be deployed using Kubernetes. The same is not true about Go binaries.

\* If you want a program with watching metrics, such program should definitely have at least two goroutines: one for running the HTTP server and another one for collecting the metrics. Usually, the HTTP server is on the goroutine that runs the main() function and the metric collection happens in a user-defined goroutine.

```
FROM golang:alpine AS builder
```

golang:alpine always contains the latest Go version as long as you update it regularly.

```
scrape_configs:
    scrape_interval: 5s
```

\* Prometheus pulls data every 5 seconds, according to the value of the scrape_interval field.

\* You should put all Docker images under the same network.

\* Prometheus and Grafana work very well together so we are going to use Grafana for the visualization part.

## Developing web clients
```
data, err := http.Get(URL)
```

In the previous statement we get the URL and get its data using http.Get(), which returns an *http.Response and an error variable.

### Using http.NewRequest() to improve the client

```
request, err := http.NewRequest(http.MethodGet, URL.string(), nil)
```

The http.NewRequest() function returns an http.Request object given a method, a URL, and an optional body. The http.MethodGet parameter defines that we want to retrieve the data using a GET HTTP method whereas URL.string() returns the string value of an http.URL variable.

```
httpData, err := c.Do(request)
```

The http.Do() function sends an HTTP request (http.Request) using an http.Client and gets an http.REsponse. So, http.Do() doest the job of http.Get() in a more detailed way. 

```
if data.StatusCode != http.StatusOK {
    fmt.Println("success")
    return
}
```
The httpData.Status holds the HTTP status code of the response. Checking the HTTP status code is considered a good practice. Therefore, if everything is OK with the server response, we continue by reading the data.
