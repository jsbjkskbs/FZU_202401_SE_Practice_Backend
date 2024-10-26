package configure

import (
	"errors"
	"log"
)

var (
	example_param_a = 1
	example_param_b = `success`
	example_param_c = `failed`
)

// 你应该这样注册规则表

// 创建这个函数并在这个函数内添加规则
func Register(v ...any) {
	RuleTable = []*ConfigureRule{
		{
			RuleName: "example",
			Level:    LevelWarn,

			LoadMethodParam: []interface{}{
				&example_param_a,
			},
			LoadMethod: func(v ...any) error {
				if v[0] == nil {
					return errors.New("Nothing") // 触发失败
				}
				return nil // 触发成功
			},

			SuccessTriggerParam: []interface{}{
				&example_param_b,
			},
			SuccessTrigger: func(v ...any) {
				for _, item := range v {
					log.Print(*(item.(*string)))
				}
			},

			FailedTriggerParam: []interface{}{
				&example_param_c,
			},
			FailedTrigger: func(v ...any) {
				log.Print(v)
			},
		},
	}
}
