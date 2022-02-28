## 发布流程
### 停止nginx
```shell
sudo nginx -s stop
```
### 停止程序
```shell
# 查询程序端口（port:8000）进程号
sudo netstat -nltp | grep 8000
# 杀死程序进程
sudo kill -9 进程号
```
### 替换编译后的文件，前端&后端
### 启动程序
```shell
# 启动后台守护进程，并标准输出log
nohup sudo ./main > ../../log/bridge-sms.log 2>&1 &
```
### 启动nginx
```shell
sudo nginx
```