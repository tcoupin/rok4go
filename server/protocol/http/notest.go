// +build !test

package http

import "log"

func (r HTTPResponse) Print() {
	log.Println("HTTPResponse:")
	log.Printf("- code: %d", r.code)
	log.Printf("- headers: %s", r.header)
	log.Printf("- body: %s", r.body)
}
