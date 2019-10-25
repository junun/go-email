package go_email

import (
	"github.com/gin-gonic/gin"
	"go_email/src/models"
	"go_email/src/pkg/util"
	"strings"
)

func PostEmail(c *gin.Context)  {
	var to []string

	mailFrom 	:= c.PostForm("mail_from")
	mailTo 		:= c.PostForm("mail_to")

	//file, header, e1 := c.Request.FormFile("file")
	//if e1 != nil {
	//	util.JsonRespond(500, "附件上传失败", "", c)
	//}
	//
	//content, e2 := ioutil.ReadAll(file)
	//if e2 != nil {
	//	util.JsonRespond(500, "文件读取失败", "", c)
	//}
	//
	//fmt.Println(header.Filename)
	//fmt.Println(string(content))


	for _, item :=  range strings.Split(mailTo,",") {
		to = append(to, item)
	}


	subject := c.PostForm("subject")
	content := c.PostForm("content")

	m := CreateMsg(mailFrom, to, subject, content)
	err := models.Gmd.DialAndSend(m)

	if err != nil {
		util.JsonRespond(500, "邮件发送失败", "", c)
	}

	util.JsonRespond(200, "邮件发送成功", "", c)
}


func AddEmailToQueue(c *gin.Context)  {
	var to []string

	mailFrom 	:= c.PostForm("mail_from")
	mailTo 		:= c.PostForm("mail_to")

	for _, item :=  range strings.Split(mailTo,",") {
		to = append(to, item)
	}

	subject := c.PostForm("subject")
	content := c.PostForm("content")

	AddSendQueue(EmailQueue{
		From:    mailFrom,
		To:      to,
		Subject: subject,
		Body:    content,
	})

	util.JsonRespond(200, "邮件发送成功", "", c)

}