# 导航网站
本项目是使用go + gin + sqlite构建的个人导航网站，前端页面使用naive ui实现

### 已实现功能
- [x] 分组的添加、编辑、删除
- [x] 分组下链接的添加、编辑、删除
- [x] 导入chrome的书签json文件
- [x] 导入特定格式的json文件到某个组
- [x] 搜索

注意：
chrome的书签json文件位置
```bash
windows 系统：
  C:/Users/Administrator/AppData/Local/Google/Chrome/User Data/Default/Bookmarks
mac 系统：
  /Users/一般是你的电脑账号名/Library/Application Support/Google/Chrome/Default/Bookmarks
```
找到该文件后，重命名为Bookmarks.json，这个时候就能上传json并导入数据了

### 运行步骤
#### 1.克隆项目
```bash
git clone git@github.com:romer12/my-nav.git
```

#### 2.安装依赖
```bash
# cd到项目目录
cd my-mav
#安装依赖
go mod tidy
```

#### 3.运行
```bash
# 在项目目录下
go run main.go
```

### 打包成可执行文件
```bash
windows:
go build -o ma_nav.exe

linux:
go build
```

### 自行编译前后端成可执行文件
#### 1.克隆项目
```bash
git clone git@github.com:romer12/my-nav.git
```

#### 2.安装依赖
```bash
# cd到项目目录
cd my-mav
#安装依赖
go mod tidy
```

#### 3.打包前端
```bash
cd vue-navigation
```
```bash
# 使用npm安装依赖
npm install
#使用pnpm安装依赖
pnpm install
```
```bash
# 打包
#使用npm
npm build

#使用pnpm
pnpm build
```

#### 4.复制前端打包好的文件到相应目录
前端打包好之后，会在前端项目根目录(vue-navigation)下生成一个dist文件夹，将
favicon.ico
index.js
style.css
这三个文件，复制到my-nav目录下的assets/static目录，将index.html文件复制到assets/templates目录，用记事本或任一款编辑器打开index.html
修改
```html
<script type="module" crossorigin src="/index.js"></script>
<link rel="stylesheet" crossorigin href="/style.css">
```
这两行为下面的路径
```html
<script type="module" crossorigin src="/assets/static/index.js"></script>
<link rel="stylesheet" crossorigin href="/assets/static/style.css">
```

#### 5.打包为可执行文件
回到my-nav项目根目录，执行下面的命令进行交叉编译生成可执行文件
```bash
#编译为windows可执行文件
GOOS=windows GOARCH=amd64 go build -o windows-mynav.exe

#编译为linux可执行文件
GOOS=linux GOARCH=amd64 go build -o linux-mynav

#编译为mac os可执行文件
GOOS=darwin GOARCH=amd64 go build -o mac-mynav

```