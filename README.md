# jira
项目依赖于github.com/andygrunwald/go-jira，所以首先要下载这个包，如果go语言设置的代理服务器有问题，需要
先查看proxy设置：go env，如果是goland的代理，是不能用的，需要重新设置
go env -w GOPROXY=https://goproxy.cn,direct
go get github.com/andygrunwald/go-jira
