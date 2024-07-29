package runtime

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDockerRun(t *testing.T) {
	result, err := RunDocker(context.TODO(), "python-add", []byte{3, 5})
	assert.NoError(t, err)
	t.Log(string(result))
	assert.Equal(t, []byte("8"), result)
}
