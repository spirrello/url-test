# Golang for testing HTTP endpoints

## Usage

Flags:
```
--url for target http endpoint

--request-count number of simultaneous requests to submit

--request-type GET, HTTP, etc.

--httpBody true or false to print out http body
```

```
go run url-test.go --request-count 5 --request-type GET --url https://yahoo.com
Testing with: https://yahoo.com
2019/07/05 14:31:41 test: 1, time spent: 0.87 seconds, result: 200 OK
2019/07/05 14:31:41 test: 3, time spent: 0.87 seconds, result: 200 OK
2019/07/05 14:31:41 test: 4, time spent: 0.87 seconds, result: 200 OK
2019/07/05 14:31:41 test: 2, time spent: 0.87 seconds, result: 200 OK
2019/07/05 14:31:41 test: 0, time spent: 1.05 seconds, result: 200 OK

```
