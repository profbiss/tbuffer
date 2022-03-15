Timed Buffer
============

Implementation of a buffer that flushes its' contents immediately when it's full or at regular intervals. This is useful if you have data that you need to buffer but don't want the items to stay in the buffer for too long.

ATTENTION! Required golang version 1.18+

Usage
-----
`go get github.com/profbiss/tbuffer`


```go
// define a flush function that will be called whenever the buffer is full or the time period has elapsed

// Initialize and use the buffer
tb := tbuffer.New(100, 10 * time.Second, func(items []int){
// flush logic
})
defer tb.Close()
tb.Put(123, 324)
tb.Put(435)
...
```


```go
// define a flush function that will be called whenever the buffer is full or the time period has elapsed

type example struct{
	val string
	...
}
// Initialize and use the buffer
tb := tbuffer.New(100, 10 * time.Second, func(items []example){
// flush logic
})
defer tb.Close()
tb.Put(example{"asd"}, example{"sadsa"})
tb.Put(example{"adsasd"})
...
```