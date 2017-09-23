package gserver

import (
	Logger "github.com/gfandada/gserver/logger"
	Module "github.com/gfandada/gserver/module"
	Network "github.com/gfandada/gserver/network"
	Services "github.com/gfandada/gserver/services"
	Discovery "github.com/gfandada/gserver/services/discovery"
	GateService "github.com/gfandada/gserver/services/gateway"
	Service "github.com/gfandada/gserver/services/service"
)

// 运行模块
// @params 模块转载器...
func Run(mods ...Module.Imodule) {
	Module.Run(mods...)
}

// 运行websocket网关
// @params log:日志配置 path:网关配置文件 discpath:服务发现配置 coder:消息编码器(注意消息需要注册)
func RunWSGateway(log, path, discpath string, coder Network.Imessage) {
	Logger.Start(log)
	gate := new(GateService.WsGateway)
	gate.Config = path
	gate.Coder = coder
	Discovery.Init(discpath)
	Run(gate)
}

// 运行tcp网关
// @params log:日志配置 path:网关配置文件 discpath:服务发现配置 coder:消息编码器(注意消息需要注册)
func RunTCPGateway(log, path, discpath string, coder Network.Imessage) {
	Logger.Start(log)
	gate := new(GateService.TcpGateway)
	gate.Config = path
	gate.Coder = coder
	Discovery.Init(discpath)
	Run(gate)
}

// 运行service
// @params log:配置 path:服务配置 coder:消息编码器(注意消息需要注册)
func RunService(log, path string, coder Network.Imessage) {
	Logger.Start(log)
	service := new(Service.Service)
	service.Config = path
	service.Coder = coder
	Run(service)
}

// 注册消息handler
// 非线程安全
// @params list:消息列表
func RegisterHandler(list []*Services.MsgHandler) {
	for _, v := range list {
		Services.Register(v.MsgId, v.MsgHandler)
	}
}
