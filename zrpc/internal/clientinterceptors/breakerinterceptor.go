package clientinterceptors

import (
	"context"
	"path"

	"github.com/shuitai/coney-framework/core/breaker"
	"github.com/shuitai/coney-framework/zrpc/internal/codes"
	"google.golang.org/grpc"
)

func BreakerInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	breakerName := path.Join(cc.Target(), method)
	return breaker.DoWithAcceptable(breakerName, func() error {
		return invoker(ctx, method, req, reply, cc, opts...)
	}, codes.Acceptable)
}
