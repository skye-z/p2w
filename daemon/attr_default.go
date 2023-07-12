//go:build !windows && !plan9
// +build !windows,!plan9

/*
Modify based on https://github.com/zh-five/xdaemon
*/
package daemon

import "syscall"

func NewSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		Setsid: true,
	}
}
