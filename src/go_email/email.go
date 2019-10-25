package go_email

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"go_email/src/models"
	"gopkg.in/gomail.v2"
	"sync"
)

const SendQueue  = "email_send_queue"

type EmailQueue struct {
	From        string
	To			[]string
	Subject		string
	Body 		string
}

// 生成消息体
func CreateMsg(mailFrom string, mailTo []string,subject string, body string) *gomail.Message{
	m := gomail.NewMessage()
	m.SetHeader("From","Monitor" + "<" + mailFrom + ">")
	m.SetHeader("To", mailTo...)  //发送给多个用户
	m.SetHeader("Subject", subject)  //设置邮件主题
	m.SetBody("text/html", body)  //设置邮件正文

	return m
}

// 生成带附件的消息体， 不支持非实时发送。
func CreateMsgWithAnnex(mailFrom string, mailTo []string,subject string, body string, annex string) *gomail.Message{
	m := gomail.NewMessage()
	m.SetHeader("From","Monitor" + "<" + mailFrom + ">")
	m.SetHeader("To", mailTo...)  //发送给多个用户
	m.SetHeader("Subject", subject)  //设置邮件主题
	m.SetBody("text/html", body)     //设置邮件正文
	m.Attach(annex)							// 设置附件

	return m
}


// 非实时发送队列添加
func AddSendQueue(msgInfo EmailQueue) {
	b, err := json.Marshal(msgInfo)

	if err != nil {
		fmt.Println(err)
	}

	models.PutinQueue(SendQueue, string(b))
}

// 读取下邮件队列 并发送邮件
func SendEmailByQueue() {
	var wg sync.WaitGroup

	for {
		length := models.GetQueueLength(SendQueue)
		if length == 0 {
			break
		}

		msgInfo := models.PopfromQueue(SendQueue)

		var data EmailQueue

		if err := json.Unmarshal([]byte(msgInfo), &data); err != nil {
			fmt.Println(err)
		}

		//协程执行该任务
		wg.Add(1)
		go func() {
			m := CreateMsg(data.From, data.To, data.Subject, data.Body)
			models.Gmd.DialAndSend(m)
			defer wg.Done()
		}()
	}
}


// 定时任务
func SendEmailCronTask() {
	c := cron.New()
	c.AddFunc("*/30 * * * * *", SendEmailByQueue)
	c.Start()
	select {}
}
