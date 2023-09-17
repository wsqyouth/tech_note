### 注意
这里主要记录mysql学习笔记,宁简勿繁，把一些心得每日学习记录下来

ps -ef | grep mysql
sudo service mysql restart
sudo netstat -tlnp | grep mysql
mysql -h127.0.0.1 -ucoopers -p2019Youth -P3306  -A;

----
在Linux环境下，可以使用以下命令来检查MySQL服务是否可以：
查看MySQL服务状态：
使用systemctl命令：systemctl status mysql.service
使用service命令：service mysql status

重启MySQL服务：
使用systemctl命令：systemctl restart mysql.service
使用service命令：service mysql restart