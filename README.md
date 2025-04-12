# 导航网站
本项目是使用go + gin + sqlite构建的个人导航网站，前端页面使用naive ui实现

### 已实现功能
- [x] 分组的添加、编辑、删除
- [x] 分组下链接的添加、编辑、删除
- [x] 导入chrome的书签json文件
- [x] 导入特定格式的json文件到某个组

注意：
chrome的书签json文件位置
```bash
windows 系统：
  C:/Users/Administrator/AppData/Local/Google/Chrome/User Data/Default/Bookmarks
mac 系统：
  /Users/一般是你的电脑账号名/Library/Application Support/Google/Chrome/Default/Bookmarks
```

### 运行步骤
#### 1.克隆项目
```bash
git clone git@github.com:romer12/my-nav.git
```

#### 2.安装依赖
```bash
# cd到项目目录
cd ma-mav
#安装依赖
go mod tidy
```

#### 3.运行
```bash
# 在项目目录下
go run main.go
```