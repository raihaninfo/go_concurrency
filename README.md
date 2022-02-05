# go_concurrency

## Without concurrency

```
time go run main.go google.com facebook.com amazon.com twitter.com github.com youtube.com

[200] https://google.com
[200] https://facebook.com
[200] https://amazon.com
[200] https://twitter.com
[200] https://github.com
[200] https://youtube.com

real	0m5.728s
user	0m0.858s
sys	0m0.309s
```

## With concurrency

```
time go run main.go google.com facebook.com amazon.com twitter.com github.com youtube.com
[200] https://github.com
[200] https://twitter.com
[200] https://google.com
[200] https://youtube.com
[200] https://facebook.com
[200] https://amazon.com

real	0m2.508s
user	0m0.594s
sys	0m0.187s

```
