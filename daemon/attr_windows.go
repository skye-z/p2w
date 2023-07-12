//go:build windows
// +build windows

/*
Modify based on https://github.com/zh-five/xdaemon
*/
package daemon

import "syscall"

func NewSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		HideWindow: true,
	}
}
