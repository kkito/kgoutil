package kgoutil

import (
    "testing"
)

func TestSendMail(t *testing.T) {
    err := SendMail("test" , "test" , "smtp.saybot.com:25" , "test@test.com" , "target@target.com" , "hello" , "world" , "txt")
    if err != nil {
        t.Log("send with error")
    }
}
