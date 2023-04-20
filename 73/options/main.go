package options

import (
	"context"
	"fmt"
)

// 定义结构体
type demoConfig struct {
	demoInt    int64
	demoBool   bool
	demoString string
}

type Option func(*demoConfig)

func DemoInt(in int64) Option {
	return func(dc *demoConfig) {
		dc.demoInt = in
	}
}

func DemoBool(in bool) Option {
	return func(dc *demoConfig) {
		dc.demoBool = in
	}
}

func DemoString(in string) Option {
	return func(dc *demoConfig) {
		dc.demoString = in
	}
}

// 绑定结构体的Apply方法，可变参数=》切片类型
func (dc *demoConfig) Apply(opts ...Option) {
	for _, opt := range opts {
		opt(dc)
	}
}

type demoConfigKey struct{}

// 赋值到context
func WithContext(ctx context.Context, cfg demoConfig) context.Context {
	return context.WithValue(ctx, demoConfigKey{}, cfg)
}

type createFunc func() demoConfig

// 取出context中的config
func FromContextOrCreate(ctx context.Context, create createFunc) demoConfig {
	vc, ok := ctx.Value(demoConfigKey{}).(demoConfig) //类型断言
	if !ok {
		return create()
	}
	return vc
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) DefaultDemoConfigCreater() createFunc {
	return func() demoConfig {
		return demoConfig{
			demoString: "hello",
		}
	}
}

func (s *Service) Demo(ctx context.Context) {
	cfg := FromContextOrCreate(ctx, s.DefaultDemoConfigCreater())
	fmt.Println(cfg)
}
