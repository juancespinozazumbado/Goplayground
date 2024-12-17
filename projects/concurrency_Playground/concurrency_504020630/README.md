# Go app to test sequential and concurrent execution

first of all run this to set up the project
```pws
go mod tidy
```

to use concurrency mode set it up in the main function

```go
useConcurrency := true
```
-  If there is concurrency set up run: 
```pws
go run -race main.go
```

- else run 
```pws
go run ./main.go
```