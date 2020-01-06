# golang-lottery
golang lottery

---
### 项目目录
- dao: 面向数据库
- services：面向数据服务

---
### xorm-cmd
```
# 进入到 xorm-cmd项目目录
$ cd /Users/a123/Go/pkg/mod/xorm.io/cmd/xorm@v0.0.0-20191108140657-006dbf24bb9b

# xorm reverse
$ xorm reverse mysql "root:123456@tcp(127.0.0.1:3306)/lottery?charset=utf8" templates/goxorm /Users/a123/Docker/golang-lottery/models/
```

---
## TODO:
- dao 完成 其他table
