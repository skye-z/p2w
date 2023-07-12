/*
守护进程
Modify based on https://github.com/zh-five/xdaemon
*/
package daemon

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"p2w/global"
)

const ENV_NAME = "BETAX_P2W_IDX"

var runIdx int = 0

type Daemon struct {
	Name     string
	LogFile  string //日志文件, 记录守护进程和子进程的标准输出和错误输出. 若为空则不记录
	MaxCount int    //循环重启最大次数, 若为0则无限重启
}

// 把本身程序转化为后台运行(启动一个子进程, 然后自己退出)
// logFile 若不为空,子程序的标准输出和错误输出将记入此文件
// isExit  启动子加进程后是否直接退出主程序, 若为false, 主程序返回*os.Process, 子程序返回 nil. 需自行判断处理
func Background(name string, logFile string, isExit bool) (*exec.Cmd, error) {
	//判断子进程还是父进程
	runIdx++
	envIdx, err := strconv.Atoi(os.Getenv(ENV_NAME))
	if err != nil {
		envIdx = 0
	}
	//子进程, 退出
	if runIdx <= envIdx {
		return nil, nil
	}

	global.Set("service.pid", "0")
	//设置子进程环境变量
	env := os.Environ()
	env = append(env, fmt.Sprintf("%s=%d", ENV_NAME, runIdx))
	//启动子进程
	cmd, err := startProc(os.Args, env, logFile)
	if err != nil {
		log.Println("Daemon service startup failed", err)
		return nil, err
	} else {
		log.Println("Daemon service startup success")
		global.Set("service.pid", cmd.Process.Pid)
	}

	if isExit {
		os.Exit(0)
	}

	return cmd, nil
}

func NewDaemon(name string, logFile string) *Daemon {
	return &Daemon{
		Name:     name,
		LogFile:  logFile,
		MaxCount: 3,
	}
}

// 启动后台守护进程
func (d *Daemon) Run() {
	//启动一个守护进程后退出
	Background(d.Name, d.LogFile, true)

	//守护进程启动一个子进程, 并循环监视
	var t int64
	count := 1
	for {
		//daemon 信息描述
		dInfo := fmt.Sprintf("Daemon(pid:%d; count:%d/%d;):",
			os.Getpid(), count, d.MaxCount)
		if d.MaxCount > 0 && count > d.MaxCount {
			log.Println(dInfo, "Too many restarts, service shutdown")
			os.Exit(0)
		}
		count++

		t = time.Now().Unix() //启动时间戳
		cmd, err := Background(d.Name, d.LogFile, false)
		if err != nil { //启动失败
			log.Println(dInfo, "Service startup failed;", "err:", err)
			continue
		}

		//子进程,
		if cmd == nil {
			log.Printf("%s service running, pid is %d", d.Name, os.Getpid())
			break
		}

		//父进程: 等待子进程退出
		err = cmd.Wait()
		dat := time.Now().Unix() - t //子进程运行秒数
		log.Printf("%s %s service(%d) shutdown, Running for %d second in total %v\n", dInfo, d.Name, cmd.ProcessState.Pid(), dat, err)
	}
}

func startProc(args, env []string, logFile string) (*exec.Cmd, error) {
	cmd := &exec.Cmd{
		Path:        args[0],
		Args:        args,
		Env:         env,
		SysProcAttr: NewSysProcAttr(),
	}

	if logFile != "" {
		stdout, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			log.Println(os.Getpid(), ": Error opening log file:", err)
			return nil, err
		}
		cmd.Stderr = stdout
		cmd.Stdout = stdout
	}

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
