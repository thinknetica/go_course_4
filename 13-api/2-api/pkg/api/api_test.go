package api

import (
	"os"
	"testing"
)

var api *API

func TestMain(m *testing.M) {
	api = New()
	os.Exit(m.Run())
}
