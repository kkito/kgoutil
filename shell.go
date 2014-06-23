package kgoutil
import (
    "log"
    "fmt"
    "os/exec"
)

func Bash(cmd string) {
    log.Println(cmd)
    out , err := exec.Command("bash","-c",cmd).Output()
    if err != nil {
        log.Println(err)
        panic(err)
    }else {
        log.Println(fmt.Sprintf("%s" , out))
    }
}
