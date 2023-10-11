package timemachine

import (
	"context"
	"time"
)

// NowFunc はタイムマシン上の現在の時間を返す関数。
type NowFunc func() time.Time

func now(f NowFunc) time.Time {
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

// Now は　NowFunc が ctx に存在しないとき、または存在しても nil のときはローカルの現在の時間を返す。
// もし存在すれば NowFunc 呼び出す。
func Now(ctx context.Context) time.Time {
	f, _ := FromContext(ctx)
	return now(f)
}

// Clock はタイムマシン上の現在の時間を返す Now を実装している。ゼロ値で使うことができる。
type Clock struct {
	// Func をセットしていると (Clock).Now はこの値を返す。
	Func NowFunc
}

// Now は Func が nil なら実際のローカルの現在の時間を返す。
// non-nil であれば、 NowFunc を呼び出す。
func (c Clock) Now() time.Time {
	return now(c.Func)
}
