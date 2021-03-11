package zlib

import "github.com/blueslimee/go-zlib/native"

func checkClosed(c native.StreamCloser) error {
	if c.IsClosed() {
		return errIsClosed
	}
	return nil
}
