// main.go
package main

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"html/template"
	"io"
	"nav/DTO"
	"nav/assets"
	"nav/middleware"
	"nav/utils"
	"net/http"
	"strconv"
)

var templates embed.FS

var db *gorm.DB

func main() {
	// 初始化数据库
	var err error
	db, err = gorm.Open(sqlite.Open("nav.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	err = db.AutoMigrate(&DTO.Group{}, &DTO.Link{})
	if err != nil {
		fmt.Println("数据库迁移失败：" + err.Error())
		return
	}

	// 初始化Gin
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	//初始化静态资源
	r.StaticFS("assets", http.FS(assets.Static))
	// 设置模板资源
	r.SetHTMLTemplate(template.Must(template.New("").ParseFS(assets.Template, "templates/*")))
	// r.LoadHTMLGlob("templates/*") //dev模式下使用这种方式

	// 全局中间件配置
	r.Use(
		middleware.NewCORS(),
	)

	// 路由配置
	r.GET("/", indexPage)
	r.GET("/groups", getGroups)
	r.POST("/group/:id", updateGroupById)
	r.GET("/groups/:id/links", getLinksByGroup)
	r.POST("/groups", createGroup)
	r.POST("/links", createLink)
	r.POST("/link", updateLink)
	r.DELETE("/groups/:id", deleteGroup)
	r.DELETE("/links/:id", deleteLink)
	r.POST("/import", importJson)

	go func() {
		err = r.Run(":9120")
		if err != nil {
			fmt.Println("服务器启动失败：" + err.Error())
			return
		}
	}()

	fmt.Println("【启动成功，请打开浏览器，在地址栏输入 http://127.0.0.1:9120/ 并按下回车键进行访问】")
	fmt.Println("【按下ctrl + c或者直接关闭命令行窗口结束运行】")

	select {}
}

func indexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// API Handlers
func getGroups(c *gin.Context) {
	var groups []DTO.Group
	db.Model(&groups).Preload("Links").Find(&groups)

	utils.ReturnSuccess(c, "", groups)
	return
}

func updateGroupById(c *gin.Context) {
	groupID := c.Param("id")
	var group DTO.Group
	err := c.ShouldBindJSON(&group)
	if groupID == "" {
		utils.ReturnError(c, "ID参数错误")
		return
	}
	if group.Name == "" {
		utils.ReturnError(c, "请输入分组名")
		return
	}

	if err != nil {
		utils.ReturnError(c, "参数错误")
		return
	}
	db.Model(&group).Where("id = ?", groupID).Update("name", group.Name)
	utils.ReturnSuccess(c, "更新成功", nil)
}

func getLinksByGroup(c *gin.Context) {
	groupID := c.Param("id")
	var links []DTO.Link
	db.Where("group_id = ?", groupID).Find(&links)

	utils.ReturnSuccess(c, "", links)
}

func createGroup(c *gin.Context) {
	var group DTO.Group
	var info DTO.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		utils.ReturnError(c, err.Error())
		return
	}
	if group.Name == "" {
		utils.ReturnError(c, "请输入分组名")
		return
	}

	// 查询分组名是否已存在
	result := db.Where("name = ?", group.Name).First(&info)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			db.Create(&group)
			utils.ReturnSuccess(c, "新增分组成功", nil)
		} else {
			utils.ReturnError(c, "查询出错："+result.Error.Error())
			return
		}
	} else {
		utils.ReturnError(c, "该分组已经存在")
		return
	}

}

func createLink(c *gin.Context) {
	var link DTO.Link
	if err := c.ShouldBindJSON(&link); err != nil {
		utils.ReturnError(c, err.Error())
		return
	}

	db.Create(&link)

	utils.ReturnSuccess(c, "新增链接成功", link)
}

func updateLink(c *gin.Context) {
	var link DTO.Link
	if err := c.ShouldBindJSON(&link); err != nil {
		utils.ReturnError(c, err.Error())
		return
	}
	if link.ID <= 0 {
		utils.ReturnError(c, "链接id必须传递")
		return
	}

	db.Model(&link).Save(&link)

	utils.ReturnSuccess(c, "更新链接成功", link)
}

func deleteGroup(c *gin.Context) {
	id := c.Param("id")
	var links []DTO.Link
	db.Where("group_id = ?", id).Find(&links)

	var linkIds []int
	if len(links) > 0 {
		for _, link := range links {
			linkIds = append(linkIds, link.ID)
		}
		doDeleteLink(linkIds)
	}
	db.Delete(&DTO.Group{}, id)

	utils.ReturnSuccess(c, "已成功删除分组", nil)
}

func deleteLink(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	doDeleteLink([]int{id})

	utils.ReturnSuccess(c, "已删除该链接", nil)
}

func doDeleteLink(ids []int) {
	db.Delete(&DTO.Link{}, ids)
}

func importJson(c *gin.Context) {
	opType := c.Query("type")
	if opType == "" {
		utils.ReturnError(c, "操作类型参数错误")
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		utils.ReturnError(c, err.Error())
		return
	}
	if file == nil {
		utils.ReturnError(c, "请上传文件")
		return
	}

	fileHandler, err := file.Open()
	if err != nil {
		utils.ReturnError(c, "打开文件失败："+err.Error())
		return
	}

	content, err := io.ReadAll(fileHandler)
	if err != nil {
		utils.ReturnError(c, "读取文件失败："+err.Error())
		return
	}

	// 分类型处理
	switch opType {
	case "full":
		// 从chrome导出的json书签文件导入
		err = doImportFull(content)
	case "group":
		// 导入到某个组
		// 获取group_id
		var groupId int
		groupId, err = strconv.Atoi(c.Query("group_id"))
		if groupId <= 0 {
			utils.ReturnError(c, "分组id参数错误")
			return
		}

		if err != nil {
			utils.ReturnError(c, "类型转换出错："+err.Error())
			return
		}
		err = doImportGroup(content, groupId)
	}
	if err != nil {
		utils.ReturnError(c, "操作失败："+err.Error())
		return
	}

	utils.ReturnSuccess(c, "导入成功", nil)
}

// 从chrome导出的json书签文件导入
func doImportFull(content []byte) error {
	var result DTO.GoogleJsonData
	err := json.Unmarshal(content, &result)
	if err != nil {
		return err
	}
	// 遍历到group结构体
	var groups []DTO.Group
	bookMarkData := result.Roots.BookmarkBar.Children
	if len(bookMarkData) > 0 {
		for _, item := range bookMarkData {
			if item.Type == "folder" {
				var childs []DTO.Link
				if len(item.Children) > 0 {
					for _, item2 := range item.Children {
						childs = append(childs, DTO.Link{
							URL:   item2.Url,
							Title: item2.Name,
						})
					}
				}
				// 将数据添加到groups中
				groups = append(groups, DTO.Group{
					Name:  item.Name,
					Links: childs,
				})
			}
		}

		// 批量插入
		db.Model(&DTO.Group{}).Create(&groups)
	}

	return nil
}

// json文件导入到某个组
func doImportGroup(content []byte, groupId int) error {
	var links []DTO.Link
	err := json.Unmarshal(content, &links)
	if err != nil {
		return err
	}

	if len(links) > 0 {
		for _, link := range links {
			if link.Title == "" || link.URL == "" {
				return errors.New("json数据格式错误，请仔细阅读示例")
			}
		}

		for i, item := range links {
			if item.Title != "" && item.URL != "" {
				links[i].GroupID = groupId
			}
		}
		db.Model(&DTO.Link{}).Create(&links)
	} else {
		return errors.New("解析文件数据为空，有可能是格式不正确")
	}

	return nil
}
