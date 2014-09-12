package main

import (
	"strings"
)

const (
	DEFAULT_HOST           = "localhost"
	DEFAULT_SERVER_PORT    = "9999"
	DEFAULT_REPLY_PORT     = "9998"
	DEFAULT_REPLY_PORT_INT = 9998
	N_NAME                 = "n"
	SIZE_NAME              = "size"
)

var (
	DEFAULT_SERVER_ADDR        = ""
	DEFAULT_REPLY_ADDR         = ""
	DEFAULT_BYTES       []byte = []byte("7\n")
)

func init() {
	DEFAULT_SERVER_ADDR = strings.Join([]string{DEFAULT_HOST, DEFAULT_SERVER_PORT}, ":")
	DEFAULT_REPLY_ADDR = strings.Join([]string{DEFAULT_HOST, DEFAULT_REPLY_PORT}, ":")
}
