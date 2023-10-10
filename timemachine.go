package timemachine

import (
	"context"
	"time"
)

// TimeMachine は現在の時間を返す Now を実装している。
type NowFunc func() time.Time

// Now は NowFunc が nil のときローカルの現在の時間を返す。
// non-nil では NowFunc 呼び出す。つまり NowFunc が返す値によって現在の時間を操作することができる。
// context に入れて使うことを想定している。
func (f NowFunc) Now() time.Time {
	if f == nil {
		return time.Now()
	}
	return f()
}

type key int

var nowFuncKey key

func NewContext(ctx context.Context, f NowFunc) context.Context {
	return context.WithValue(ctx, nowFuncKey, f)
}

func FromContext(ctx context.Context) (NowFunc, bool) {
	f, ok := ctx.Value(nowFuncKey).(NowFunc)
	return f, ok
}

func Now(ctx context.Context) time.Time {
	f, _ := FromContext(ctx)
	return f.Now()
}
