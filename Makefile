bench:
	go test -bench=. -benchtime=1000000x -memprofile memprofile.out -cpuprofile profile.out ./tests
	go tool pprof profile.out

test:
	go test -race -cover -covermode=atomic ./tests -v