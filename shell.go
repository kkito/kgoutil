package kgoutil
import (
    "log"
    "fmt"
    "os/exec"
)

func Bash(cmd string) string {
    log.Println(cmd)
    out , err := exec.Command("bash","-c",cmd).Output()
    if err != nil {
        log.Println(err)
        panic(err)
    }else {
        result := fmt.Sprintf("%s" , out)
        log.Println(result)
        return result
    }
}
