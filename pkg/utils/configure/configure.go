package configure

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	// GlobalConfig 全局配置
	GlobalConfig *viper.Viper
	// RuleTable 配置规则表
	RuleTable []*ConfigureRule
	// ConfigPath 配置文件路径
	ConfigPath = ``
)

type ConfigureLoader struct {
	// ConfigName 配置文件名
	configName string
	// ConfigType 配置文件类型
	configType string
	// ConfigPath 配置文件路径
	configPath string

	// RegisterParam 注册参数
	registerParam []interface{}
	// Register 注册方法
	register func(v ...any)

	// LogPrefix 日志前缀
	logPrefix string
	// LogSuffix 日志后缀
	logSuffix string
	// LogFunc 日志方法
	log func(v ...any)
	// WarnFunc 警告方法
	wlog func(v ...any)
	// ErrorFunc 错误方法
	elog func(v ...any)
	// FatalFunc 致命错误方法
	flog func(v ...any)
}

type ConfigureOption struct {
	// ConfigName 配置文件名
	ConfigName string
	// ConfigType 配置文件类型
	ConfigType string
	// ConfigPath 配置文件路径
	ConfigPath string

	// RegisterParam 注册参数
	RegisterParam []interface{}
	// Register 注册方法
	Register func(v ...any)

	// LogPrefix 日志前缀
	LogPrefix string
	// LogSuffix 日志后缀
	LogSuffix string
	// LogFunc 日志方法
	LogFunc func(v ...any)
	// WarnFunc 警告方法
	WarnFunc func(v ...any)
	// ErrorFunc 错误方法
	ErrorFunc func(v ...any)
	// FatalFunc 致命错误方法
	FatalFunc func(v ...any)

	// Silent 是否静默模式
	Silent bool
}

// NewConfLoader 创建配置加载器
// option: 配置选项
// 返回配置加载器
func NewConfLoader(option *ConfigureOption) *ConfigureLoader {
	loader := new(ConfigureLoader)
	loader.configName = option.ConfigName
	loader.configType = option.ConfigType
	loader.configPath = option.ConfigPath
	loader.registerParam = option.RegisterParam
	loader.register = option.Register
	loader.logPrefix = option.LogPrefix
	loader.logSuffix = option.LogSuffix
	loader.log = option.LogFunc
	loader.wlog = option.WarnFunc
	loader.elog = option.ErrorFunc
	loader.flog = option.FatalFunc
	if option.Silent {
		loader.log = func(v ...any) {}
		loader.wlog = func(v ...any) {}
		loader.elog = func(v ...any) {}
		loader.flog = func(v ...any) {}
	}

	if option.LogFunc == nil {
		loader.log = log.Println
	}

	if option.WarnFunc == nil {
		loader.wlog = log.Println
	}

	if option.ErrorFunc == nil {
		loader.elog = log.Println
	}

	if option.FatalFunc == nil {
		loader.flog = log.Fatal
	}

	return loader
}

// Run 运行配置加载器
// 返回错误
func (c *ConfigureLoader) Run() error {
	GlobalConfig = viper.New()
	GlobalConfig.SetConfigName(c.configName)
	GlobalConfig.SetConfigType(c.configType)
	GlobalConfig.AddConfigPath(c.configPath)
	if err := GlobalConfig.ReadInConfig(); err != nil {
		return err
	}

	c.register(c.registerParam...)

	c.loadConfig()

	go func() {
		GlobalConfig.WatchConfig()
		GlobalConfig.OnConfigChange(func(e fsnotify.Event) {
			c.log(c.logPrefix, "The config has changed", c.logSuffix)
			c.loadConfig()
		})
	}()

	return nil
}

// LoadConfig 加载配置
func (c *ConfigureLoader) loadConfig() {
	for _, rule := range RuleTable {
		if err := rule.LoadMethod(rule.LoadMethodParam...); err != nil {
			c.alarm(rule.Level, rule.RuleName, err)
			rule.FailedTrigger(rule.FailedTriggerParam...)
		}
		rule.SuccessTrigger(rule.SuccessTriggerParam...)
	}
}

// Alarm 告警
// level: 告警级别
// rulename: 规则名
// err: 错误
func (c *ConfigureLoader) alarm(level int, rulename string, err error) {
	switch level {
	case LevelError:
		c.elog(rulename, " failed to load. ", err)
	case LevelWarn:
		c.wlog(rulename, " failed to load. ", err)
	case LevelFatal:
		c.flog(rulename, " failed to load. ", err)
	}
}
