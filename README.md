# Quiz CLI

Take a quiz with specified time limit

Supply questions in csv format: 'question,answer'

### Sample quiz
```sh
./quiz --limit=30 --csv=testing.csv
```

## Build

```sh
go build -ldflags "-s -w" ./cmd/quiz
```

## Help
```sh
./quiz --help
```
