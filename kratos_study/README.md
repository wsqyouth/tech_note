1. 创建项目目录 kratos_study
2. go get 安装
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
3. kratos new helloworld
如果失败:  kratos new helloworld -r https://gitee.com/go-kratos/kratos-layout.git
4. cd helloworld
5. go mod download
6. kratos run
7. curl 'http://127.0.0.1:8000/helloworld/kratos'
输出:
"hello world"
恭喜，你的项目创建成功了!

