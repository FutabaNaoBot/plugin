package plugin

import (
	"github.com/kohmebot/plugin/pkg/command"
	"github.com/kohmebot/plugin/pkg/version"
	zero "github.com/wdvxdr1123/ZeroBot"
	"gorm.io/gorm"
)

// NewPluginFunc 插件初始化函数,主程序通过该方法来新建插件实例
type NewPluginFunc = func() Plugin

// Plugin 所有插件需实现的接口
type Plugin interface {
	// Init 初始化插件,在Bot运行前调用
	Init(engine *zero.Engine, env Env) error
	// Name 插件名称，应具有唯一性
	Name() string
	// Description 插件描述
	Description() string
	// Commands 插件支持的命令描述
	//  example:
	//  func (p *myPlugin) Commands() command.Commands  {
	//		return command.NewCommands(
	//			command.NewCommand("查看当前时间","time"),
	//			command.NewCommand("关闭","close","c"),
	//		)
	//}
	Commands() command.Commands
	// Version 插件版本,使用x.y.z 格式
	//  example:
	//  func (p *myPlugin) Version() version.Version {
	//		return version.NewVersion(1,0,0)
	//}
	Version() version.Version
	// OnBoot engine准备就绪后调用
	OnBoot()
}

// Env 插件运行环境
type Env interface {
	// Get 获取环境变量
	Get(key string) any
	// FilePath 获取插件数据目录(不存在时会自动创建)
	FilePath() (string, error)
	// GetConf 从配置文件获取配置
	GetConf(conf any) error
	// GetDB 获取数据库连接
	GetDB() (*gorm.DB, error)
	// RangeBot 遍历所有机器人实例
	RangeBot(yield func(ctx *zero.Ctx) bool)
	// Groups 获取启用的群
	Groups() Groups
	// SuperUser 获取SuperUser
	SuperUser() Users
	// Error 提交错误(由上层框架决定如何处理这个错误)
	Error(ctx *zero.Ctx, err error)
}

// Groups 启用的群
type Groups interface {
	// IsContains 是否包含
	IsContains(groupId int64) bool
	// Rule 判断是否是开启的群
	Rule() zero.Rule
	// RangeGroup 遍历所有启用的群
	RangeGroup(yield func(group int64) bool)
}

// Users 用户集合
type Users interface {
	// IsContains 是否包含
	IsContains(userId int64) bool
	// Rule 判断是否包含
	Rule() zero.Rule
	// RangeUser 遍历所有用户
	RangeUser(yield func(user int64) bool)
}
