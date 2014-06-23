package kgoutil
import (
    "net/smtp"
    "strings"
    "log"
)

func SendMail(user, password, host, from , to, subject, body, mailtype string) error{
    hp := strings.Split(host, ":")
    auth := smtp.PlainAuth("", user, password, hp[0])
    var content_type string
    if mailtype == "html" {
        content_type = "Content-Type: text/"+ mailtype + "; charset=UTF-8"
    }else{
        content_type = "Content-Type: text/plain" + "; charset=UTF-8"
    }

    msg := []byte("To: " + to + "\r\nFrom: " + user + "<"+ user +">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
    send_to := strings.Split(to, ";")
    log.Println(user + " send email " + subject + " to " + to);
    err := smtp.SendMail(host, auth, from, send_to, msg)
    return err
}
