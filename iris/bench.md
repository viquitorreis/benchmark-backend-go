victorreis@pop-os:~/Projetos/Trampo/bench-handlers/iris$ go-wrk -c 2048 -d 10 http://localhost:8080
Running 10s test @ http://localhost:8080
  2048 goroutine(s) running concurrently
199097 requests in 9.800595417s, 24.87MB read
Requests/sec:           20314.79
Transfer/sec:           2.54MB
Overall Requests/sec:   16708.48
Overall Transfer/sec:   2.09MB
Fastest Request:        108µs
Avg Req Time:           100.812ms
Slowest Request:        1.046399s
Number of Errors:       567
Error Counts:           net/http: timeout awaiting response headers=567
10%:                    170µs
50%:                    231µs
75%:                    255µs
99%:                    276µs
99.9%:                  276µs
99.9999%:               276µs
99.99999%:              276µs
stddev:                 190.458ms