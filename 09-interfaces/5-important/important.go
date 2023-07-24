package important

import (
	"fmt"
	"time"
)

type E struct {
	code int
	msg  string
}

func (e *E) Error() string {
	return e.msg + fmt.Sprintf(": %v", e.code)
}

// builtin
type error interface {
	Error() string
}

// io
type Reader interface {
	Read(p []byte) (n int, err error)
}

// io
type Writer interface {
	Write(p []byte) (n int, err error)
}

// sort
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// fmt
type Stringer interface {
	String() string
}

// net
type Conn interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	Close() error
	LocalAddr() Addr
	RemoteAddr() Addr
	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}

// net
type Listener interface {
	Accept() (Conn, error)
	Close() error
	Addr() Addr
}

// net/http
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

// net/http
type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(int)
}

// encoding/json
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

// Для целей документирования.
// Реальные типы другие!
type Header int
type Addr int
type Request int
