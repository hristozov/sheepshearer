package server

import "fmt"

func getServerHeader() string {
	return "Server: sheepshearer/0.1"
}

func getContentTypeHeader(contentType string) (result string) {
	var charset string
	if contentType == "text/html" {
		charset = "; charset=UTF-8"
	} else {
		charset = ""
	}
	return fmt.Sprintf("Content-Type: %s%s", contentType, charset)
}

func getContentLengthHeader(data []byte) (result string) {
	return fmt.Sprintf("Content-Length: %d", len(data))
}

// BuildErrorResponse builds an error response
func BuildErrorResponse(errorCode int) (result string) {
	errorStrings := map[int]string{
		400: "BAD REQUEST",
		404: "NOT FOUND",
	}

	var errorString string
	if val, ok := errorStrings[errorCode]; ok {
		errorString = "ERROR"
	} else {
		errorString = val
	}
	return fmt.Sprintf("HTTP/1.1 %d %s\n%s",
		errorCode,
		errorString,
		getServerHeader())
}

// BuildOkResponse builds a 200 response
func BuildOkResponse(data []byte, contentType string) (result string) {
	length := len(data)
	return fmt.Sprintf(
		"HTTP/1.1 200 OK\n%s\n%s\n%s\n\n%s",
		getServerHeader(),
		getContentTypeHeader(contentType),
		getContentLengthHeader(data),
		string(data[:length]),
	)
}
