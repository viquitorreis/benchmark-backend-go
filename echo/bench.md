victorreis@pop-os:~/Projetos/Trampo/bench-handlers/echo$ go-wrk -c 2048 -d 10 http://localhost:3060
Running 10s test @ http://localhost:3060
  2048 goroutine(s) running concurrently
1067513 requests in 10.004852561s, 115.04MB read
Requests/sec:           106699.52
Transfer/sec:           11.50MB
Overall Requests/sec:   96440.22
Overall Transfer/sec:   10.39MB
Fastest Request:        30µs
Avg Req Time:           19.193ms
Slowest Request:        188.127ms
Number of Errors:       0
10%:                    127µs
50%:                    715µs
75%:                    1.15ms
99%:                    1.536ms
99.9%:                  1.548ms
99.9999%:               1.549ms
99.99999%:              1.549ms
stddev:                 12.39ms