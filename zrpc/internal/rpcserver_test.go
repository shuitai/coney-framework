package internal

import (
	"testing"

	"github.com/shuitai/coney-framework/core/stat"
	"github.com/stretchr/testify/assert"
)

func TestWithMetrics(t *testing.T) {
	metrics := stat.NewMetrics("foo")
	opt := WithMetrics(metrics)
	var options rpcServerOptions
	opt(&options)
	assert.Equal(t, metrics, options.metrics)
}
