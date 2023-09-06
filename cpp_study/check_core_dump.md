产生core dump的可能原因：(1). 内存访问越界；(2). 多线程程序使用了线程不安全的函数；(3). 多线程读写的数据未加锁保护；(4). 非法指针；(5). 堆栈溢出。

查看操作系统是否开启产生core文件：输入命令：ulimit -c或ulimit -a， core file size为0，说明系统关闭了core dump，可通过命令：ulimit -c unlimited来打开

---
ulimit -c
unlimited
gdb  coredump应用程序  coredump文件     调试coredump文件

通过以下命令调试coredump文件：
bt 显示线程堆栈信息
frame num  切换栈