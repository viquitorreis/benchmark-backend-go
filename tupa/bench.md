victorreis@pop-os:~/Projetos/Trampo/bench-handlers/tupa$ go-wrk -c 2048 -d 10 http://localhost:6969
Running 10s test @ http://localhost:6969
  2048 goroutine(s) running concurrently
535879 requests in 10.023956201s, 62.86MB read
Requests/sec:           53459.83
Transfer/sec:           6.27MB
Overall Requests/sec:   46924.80
Overall Transfer/sec:   5.50MB
Fastest Request:        69µs
Avg Req Time:           38.308ms
Slowest Request:        385.775ms
Number of Errors:       0
10%:                    680µs
50%:                    2.172ms
75%:                    2.87ms
99%:                    3.491ms
99.9%:                  3.509ms
99.9999%:               3.512ms
99.99999%:              3.512ms
stddev:                 29.002ms