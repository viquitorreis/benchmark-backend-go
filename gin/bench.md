victorreis@pop-os:~/Projetos/Trampo/bench-handlers/gin$ go-wrk -c 2048 -d 10 http://localhost:8080
Running 10s test @ http://localhost:8080
  2048 goroutine(s) running concurrently
543348 requests in 10.015041841s, 68.92MB read
Requests/sec:           54253.19
Transfer/sec:           6.88MB
Overall Requests/sec:   47580.81
Overall Transfer/sec:   6.04MB
Fastest Request:        39µs
Avg Req Time:           37.748ms
Slowest Request:        397.983ms
Number of Errors:       0
10%:                    60µs
50%:                    78µs
75%:                    87µs
99%:                    95µs
99.9%:                  96µs
99.9999%:               96µs
99.99999%:              96µs
stddev:                 35.058ms