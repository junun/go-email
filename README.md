# go-email

```
结合gin 处理简单的邮件发送接口，支持同步和异步发送。
```

# Build

```
cd ./src
go build -o ../gemail
```

# Deploy

```
把生成的 gemail 和 conf 目录一起copy到需要部署的机器同一目录下即可
```

# curl 模拟请求

```
curl -X POST -L -d "mail_from=testfrom&mail_to=1186158664@qq.com&content=test&subject=test123456" http://127.0.0.1:9999/email
```
