package runtime

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRunWebhook(t *testing.T) {

	go func() {
		http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(writer, "hello")
		})
		err := http.ListenAndServe(":3000", nil)
		assert.NoError(t, err)
	}()
	result, err := RunWebhook(context.TODO(), "http://127.0.0.1:3000/hello", []byte("hello"))
	assert.NoError(t, err)

	assert.Equal(t, result, []byte("hello"))
}
