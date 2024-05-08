package lsp

import (
	"github.com/Sora233/DDBOT/lsp/cfg"
	"github.com/Sora233/sliceutil"
)

// TODO command需要重构成注册模式，然后把这个文件废弃

var CommandMaps = map[string]string{
	"GrantCommand":         GrantCommand,
	"WatchCommand":         WatchCommand,
	"UnwatchCommand":       UnwatchCommand,
	"ListCommand":          ListCommand,
	"EnableCommand":        EnableCommand,
	"DisableCommand":       DisableCommand,
	"HelpCommand":          HelpCommand,
	"ConfigCommand":        ConfigCommand,
	"PingCommand":          PingCommand,
	"LogCommand":           LogCommand,
	"BlockCommand":         BlockCommand,
	"SysinfoCommand":       SysinfoCommand,
	"WhosyourdaddyCommand": WhosyourdaddyCommand,
	"QuitCommand":          QuitCommand,
	"ModeCommand":          ModeCommand,
	"GroupRequestCommand":  GroupRequestCommand,
	"FriendRequestCommand": FriendRequestCommand,
	"AdminCommand":         AdminCommand,
	"SilenceCommand":       SilenceCommand,
	"NoUpdateCommand":      NoUpdateCommand,
	"AbnormalConcernCheck": AbnormalConcernCheck,
	"CleanConcern":         CleanConcern,
}

const (
	GrantCommand   = "grant"
	WatchCommand   = "watch"
	UnwatchCommand = "unwatch"
	ListCommand    = "list"
	EnableCommand  = "enable"
	DisableCommand = "disable"
	HelpCommand    = "help"
	ConfigCommand  = "config"
)

// private command
const (
	PingCommand          = "ping"
	LogCommand           = "log"
	BlockCommand         = "block"
	SysinfoCommand       = "sysinfo"
	WhosyourdaddyCommand = "whosyourdaddy"
	QuitCommand          = "quit"
	ModeCommand          = "mode"
	GroupRequestCommand  = "群邀请"
	FriendRequestCommand = "好友申请"
	AdminCommand         = "admin"
	SilenceCommand       = "silence"
	NoUpdateCommand      = "退订更新"
	AbnormalConcernCheck = "检测异常订阅"
	CleanConcern         = "清除订阅"
)

var allGroupCommand = [...]string{
	GrantCommand,
	WatchCommand, UnwatchCommand,
	ListCommand,
	EnableCommand, DisableCommand,
	ConfigCommand,
	HelpCommand, AdminCommand,
	SilenceCommand, NoUpdateCommand, CleanConcern,
}

var allPrivateOperate = [...]string{
	PingCommand, HelpCommand, LogCommand,
	BlockCommand, SysinfoCommand, ListCommand,
	WatchCommand, UnwatchCommand, DisableCommand,
	EnableCommand, GrantCommand, ConfigCommand,
	WhosyourdaddyCommand, QuitCommand, ModeCommand,
	GroupRequestCommand, FriendRequestCommand, AdminCommand,
	SilenceCommand, NoUpdateCommand, AbnormalConcernCheck,
	CleanConcern,
}

var nonOprateable = [...]string{
	EnableCommand, DisableCommand, GrantCommand,
	BlockCommand, LogCommand, PingCommand,
	WhosyourdaddyCommand, QuitCommand, ModeCommand,
	GroupRequestCommand, FriendRequestCommand, AdminCommand,
	SilenceCommand, NoUpdateCommand, AbnormalConcernCheck,
	CleanConcern,
}

func CheckValidCommand(command string) bool {
	return sliceutil.Contains(allGroupCommand, command)
}

func CheckCustomGroupCommand(command string) bool {
	return sliceutil.Contains(cfg.GetCustomGroupCommand(), command)
}

func CheckCustomPrivateCommand(command string) bool {
	return sliceutil.Contains(cfg.GetCustomPrivateCommand(), command)
}

func CheckOperateableCommand(command string) bool {
	return (sliceutil.Contains(allGroupCommand, command) || CheckCustomGroupCommand(command)) && !sliceutil.Contains(nonOprateable, command)
}

func CombineCommand(command string) string {
	if command == WatchCommand || command == UnwatchCommand {
		return WatchCommand
	}
	return command
}
