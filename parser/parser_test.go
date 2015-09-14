package parser

import "testing"

func TestGetRoot(t *testing.T) {
	req, err := Parse("GET / HTTP/1.1")

	if err != nil {
		t.Error("error")
	}
	if req.Path != "/" {
		t.Error("wrong path")
	}

	if req.Method != "GET" {
		t.Error("wrong method")
	}
}

func TestPostFoo(t *testing.T) {
	req, err := Parse("POST /foo HTTP/1.1")

	if err != nil {
		t.Error("error")
	}
	if req.Path != "/foo" {
		t.Error("wrong path")
	}

	if req.Method != "POST" {
		t.Error("wrong method")
	}
}
