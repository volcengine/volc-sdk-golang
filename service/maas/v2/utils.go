package v2

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type (
	reqIdCtxKeyType struct{}
)

var (
	reqIdCtxKey = reqIdCtxKeyType{}
)

const (
	reqIdHeaderKey = "x-tt-logid"
)

func ctxWithReqId(ctx context.Context, reqId string) context.Context {
	return context.WithValue(ctx, reqIdCtxKey, reqId)
}

func reqIdFromCtx(ctx context.Context) string {
	if val := ctx.Value(reqIdCtxKey); val != nil {
		if reqId, ok := val.(string); ok {
			return reqId
		}
	}
	return ""
}

func getContext(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	reqId := genReqId()
	return ctxWithReqId(ctx, reqId)
}

func genReqId() string {
	const length = 34

	sb := strings.Builder{}
	sb.Grow(length)
	sb.WriteString(time.Now().Format("20060102150405"))
	sb.WriteString(fmt.Sprintf("%020X", rand.Uint64()))
	return sb.String()
}
