package main

import (
	"context"
	"demo/73/options"
)

type config struct {
	needSwitch     bool
	needWatchLater bool
}

type Option func(*config)

func NeedSwitch(in bool) Option {
	return func(cfg *config) {
		cfg.needSwitch = in
	}
}

func (cfg *config) Apply(opts ...Option) {
	for _, opt := range opts {
		opt(cfg)
	}
}

func constructThreePoint(ctx context.Context, opts ...Option) {
	cfg := &config{}
	cfg.Apply(opts...)
	// 直接使用
}

//调用
// constructThreePoint(ctx, NeedSwitch(true))

func main() {
	svc := options.NewService()
	cfg := svc.DefaultDemoConfigCreater()()
	opts := []options.Option{
		options.DemoInt(27),
	}
	cfg.Apply(opts...)
	ctx := context.Background()
	ctx = options.WithContext(ctx, cfg) //赋值到context
	svc.Demo(ctx)
}
