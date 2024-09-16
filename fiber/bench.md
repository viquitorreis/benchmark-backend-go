victorreis@pop-os:~/Projetos/Trampo/bench-handlers/fiber$ go-wrk -c 2048 -d 10 http://localhost:3000
Running 10s test @ http://localhost:3000
  2048 goroutine(s) running concurrently
1209769 requests in 10.001616909s, 130.37MB read
Requests/sec:           120957.34
Transfer/sec:           13.03MB
Overall Requests/sec:   109073.65
Overall Transfer/sec:   11.75MB
Fastest Request:        22µs
Avg Req Time:           16.931ms
Slowest Request:        187.399ms
Number of Errors:       0
10%:                    45µs
50%:                    419µs
75%:                    1.27ms
99%:                    4.078ms
99.9%:                  4.219ms
99.9999%:               4.231ms
99.99999%:              4.231ms
stddev:                 5.935ms