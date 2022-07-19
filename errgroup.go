package isuconlib

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

type LimitGroup struct {
	eg  *errgroup.Group
	sem *semaphore.Weighted
}

func NewLimitGroup(concurrency int64) *LimitGroup {
	return &LimitGroup{
		eg:  &errgroup.Group{},
		sem: semaphore.NewWeighted(concurrency),
	}
}

func (lg *LimitGroup) Wait() error {
	err := lg.eg.Wait()
	if err != nil {
		return err
	}
	return nil
}

func (lg *LimitGroup) Go(ctx context.Context, f func() error) {
	lg.eg.Go(func() (err error) {
		defer func() {
			msg := recover()
			if msg != nil {
				err = errors.Errorf("panic: %s", msg)
			}
		}()
		if err = lg.sem.Acquire(ctx, 1); err != nil {
			return errors.Wrap(err, "semaphoreの確保に失敗しました")
		}
		defer lg.sem.Release(1)
		err = f()
		if err != nil {
			return err
		}
		return nil
	})
}
