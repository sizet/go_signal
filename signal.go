// ©.
// https://github.com/sizet/go_signal

package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

// 處理信號.
// 參數 :
// sigQueue
//   接收信號的 channel.
func signalHandle(
    sigQueue chan os.Signal) {

    for ;; {
        switch <-sigQueue {
            case syscall.SIGUSR1:
                fmt.Printf("signal SIGUSR1\n")
            case syscall.SIGUSR2:
                fmt.Printf("signal SIGUSR2\n")
            case syscall.SIGQUIT:
                fmt.Printf("signal SIGQUIT\n")
            case syscall.SIGTERM:
                fmt.Printf("signal SIGTERM\n")
        }
    }
}

func main() {

    var sigQueue chan os.Signal

    // go 是使用同步模式來處理信號.

    // 程式收到信號後會把信號放入 channel.
    sigQueue = make(chan os.Signal, 8)

    fmt.Printf("pid %d\n", os.Getpid())

    go signalHandle(sigQueue)

    // 設定要處理的信號.
    signal.Notify(sigQueue, syscall.SIGUSR1, syscall.SIGUSR2)

    // 設定要忽略的信號.
    signal.Ignore(syscall.SIGQUIT, syscall.SIGTERM)

    // 收到不在要處理的信號集合和不在要忽略的信號集合時, 會執行預設處理, 預設處理是中止程式.
    // 例如 SIGINT (ctrl + c).

    select {}
}
