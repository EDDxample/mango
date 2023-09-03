go tool pprof -raw -output tmp.prof 'http://localhost:8080/debug/pprof/profile?seconds=20'

docker run --rm -v $PWD:/input flamegraph tmp.prof flamegraph.svg
docker run --rm -v ${PWD}:/input flamegraph tmp.prof flamegraph.svg