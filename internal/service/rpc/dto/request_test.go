package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestSia(t *testing.T) {
	req := NewRequest("test", []byte("hello world"))
	reqByte := req.Sia().Bytes()
	gotReq := new(RPCRequest).FromSiaBytes(reqByte)

	t.Log(req)
	t.Log(*gotReq)
	assert.Equal(t, req, *gotReq)
}
