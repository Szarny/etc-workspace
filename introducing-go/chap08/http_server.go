package main

import (
	"io"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(
		res,
		`<!DOCTYPE html>
		<html>
			<head>
				<title>hello world</title>
			</head>
			<body>
				<h1>Hello!</h1>
			</body>
		</html>`,
	)
}

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe(":9000", nil)
}
