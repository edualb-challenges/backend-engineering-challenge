# treebabel 🌲

`treebabel` is a cli application that solves the challenge provided in the `CHALLENGE.md` file using [Segmentation Tree](https://en.wikipedia.org/wiki/Segment_tree).

## Technologies

- Go v1.20
- Docker

## Running the application

[![asciicast](https://asciinema.org/a/BWlIWQFHSTixRpYbCVOo2HER0.svg)](https://asciinema.org/a/BWlIWQFHSTixRpYbCVOo2HER0)

*If you don't have `make` installed in your machine run Docker by your terminal directly:*

```bash
# building docker image
$ docker build -f ./build/package/treebabel/Dockerfile . -t treebabel:latest

# running image and entering in the container
$ docker run -it treebabel:latest bash
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