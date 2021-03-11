package native

/*
#cgo CFLAGS: -I${SRCDIR}/zlib/
#cgo windows LDFLAGS: ${SRCDIR}/libs/winlibz.a
#cgo !android,linux LDFLAGS: ${SRCDIR}/libs/linuxlibz.a
#cgo darwin LDFLAGS: ${SRCDIR}/libs/darwinlibz.a
#cgo android LDFLAGS: ${SRCDIR}/libs/androidlibz.a
*/
import "C"
