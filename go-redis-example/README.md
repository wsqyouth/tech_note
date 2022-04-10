
待实现:参考sql工程风格、测试学习
https://cloud.tencent.com/developer/article/1517645

https://km.woa.com/group/38985/articles/show/505627


### 安装：
1. 安装
sudo yum install redis
2. 启动
systemctl start redis
3. 查看redis状态
systemctl status redis
4. 其他相关命令
```bash
systemctl stop redis
systemctl restart redis
systemctl start redis
systemctl status redis
```
5. 测试
redis-cli -h 127.0.0.1 -p 6379


### 学习记录
时间：2022年04月10日12:00
进度：他人的go进阶项目很好，fork过来测试学习
改动点：1. 使用go mod 工程风格进行修改,不局限于gopath
        2. 将测试的功能放到各自的子功能下方便后续继续完善独立功能
        3. 他人代码的代码风格、依赖注入、命令规范值得学习
