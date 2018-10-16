# queueMail
邮件队列

### 功能
> 
- 支持自定义邮件服务器配置
- 支持smtp安全连接
- 支持附件
- 支持自定义启动端口配置
- 支持多接收人

### API
|路由|请求方式|body|参数|说明|
|---:|---:|---:|---:|---:|
|/message|post| json(消息)|见下方说明|加入邮件队列|
|/upload|post| form-data|upload(多文件)|多文件上传接口(用于发送代附件邮件)|
#### body参数说明
```javascript
{
	"host": "smtp.qq.com",              //smtp服务器地址
	"port": 25,                         //smtp服务器端口 
	"username": "991154416@qq.com",     //发送邮箱用户名
	"password": "xxxx",     //发送邮箱密码
	"attachments": ["/data/filename.text"],   //附件在服务器中的路径
	"subject": "hello,subject!",              //邮件主题
	"body": "<h1>此处为标题111</h1><span>此处为正文111！</span>",    //邮件内容
	"from": {
		"email": "991154416@qq.com",            //发送者邮箱
		"name": "林文祥"                        //发送人
	},
	"to": ["linwenxiang_xm@jsptpd.com"],        //收件人邮箱
	"auth": true                              //是否是加密连接
}
```

### 部署
```shell
./queueMail --port=8080 --filepath=./attachments    //port和filepath可根据需求自定义
```
