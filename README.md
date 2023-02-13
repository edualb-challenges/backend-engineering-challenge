# treebabel 🌲

`treebabel` is a cli application that solves the challenge provided in the `CHALLENGE.md` file using [Segmentation Tree](https://en.wikipedia.org/wiki/Segment_tree).

## Technologies

- Go v1.20
- Docker

## Running the application
<div align="center">
    <a href="https://asciinema.org/a/BWlIWQFHSTixRpYbCVOo2HER0" target="_blank"><img src="https://asciinema.org/a/BWlIWQFHSTixRpYbCVOo2HER0.svg" width="550" /></a>
</div>

*If you don't have `make` installed in your machine run Docker by your terminal directly:*

```bash
# building docker image
$ docker build -f ./build/package/treebabel/Dockerfile . -t treebabel:latest

# running image and entering in the container
$ docker run -it treebabel:latest bash
```

*Or in case you have just Go installed in your machine, execute the program directly:*

```bash
$ go run ./cmd/treebabel/main.go --input_file ./testdata/treebabel/challenge-input.json --window_size 10
```

## Code Organization
- I am following the [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- If we have any doubts about how to right our code. We recommend to use the [Uber Style Guide](https://github.com/uber-go/guide)

## What I would like to improve:
- **Message errors**: Sometimes the message error is not so useful. I would like to put more context in the possible errors messages.
- **Debugging logs**: In this application there is no `debug` or `verbose` flag and would be useful if the application have it.
- **Concurrency**: In this application is not using concurrency to improve the processing, we could use that.
- **Graceful Shutdown**: There isn't a graceful shutdown in this application.

# treebabel Algorithm 🌲

In progress...