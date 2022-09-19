# 爬取漏洞盒子公益SRC。

本程序可以获取到漏洞盒子所有的公益SRC厂商名称和网站用csv格式保存。

## 数据：

```text
...
"漏洞盒子","https://www.vulbox.com"
...
```

## 用法：

```bash
./VulBoxCharitySRC "Bearer e*********************************************************************************************************************************************************************************************************************************0"
2022/09/19 20:02:53 读取第1页的数据。
2022/09/19 20:02:53 读取第2页的数据。
...
2022/09/19 20:03:13 读取第84页的数据。
2022/09/19 20:03:13 查询结束！
```

## 编译：

```bash
go build .

ls -l
总用量 6680
-rw------- 1 webb users   75139  9月 19 20:03 data.csv
-rw-r--r-- 1 webb users      33  9月 19 19:10 go.mod
-rw-r--r-- 1 webb users   11357  9月 19 18:49 LICENSE.txt
-rw-r--r-- 1 webb users    2299  9月 19 20:02 main.go
-rw-r--r-- 1 webb users     891  9月 19 20:05 README.md
-rwxr-xr-x 1 webb users 6736445  9月 19 20:05 VulBoxCharitySRC
```

## 协议：

```text
                                 Apache License
                           Version 2.0, January 2004
                        http://www.apache.org/licenses/
```