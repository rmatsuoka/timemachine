package timemachine

import (
	"context"
	"time"
)

// TimeMachine は現在の時間を返す now を実装している。nowFunc が nil のとき now はローカルの現在の時間を返す。
// non-nil では now は nowFunc 呼び出す。つまり nowFunc が返す値によって現在の時間を操作することができる。
// context に入れると良さそう。
type TimeMachine struct {
	NowFunc func() time.Time
}

func (tm TimeMachine) Now() time.Time {
	if tm.NowFunc == nil {
		return time.Now()
	}
	return tm.NowFunc()
}

type key int

var timeMachineKey key

func NewContext(ctx context.Context, tm TimeMachine) context.Context {
	return context.WithValue(ctx, timeMachineKey, tm)
}

func FromContext(ctx context.Context) TimeMachine {
	tm, ok := ctx.Value(timeMachineKey).(TimeMachine)
	if !ok {
		return TimeMachine{}
	}
	return tm
}
