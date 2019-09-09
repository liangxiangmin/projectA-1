package controller

import (
	"dao"
	"github.com/gin-gonic/gin"
	"log"
)

// 负责 分类相关操作的函数集合文件

// 分类树
func CategoryTree(c *gin.Context) {
	//连接数据库，获取全部的分类内容
	config := map[string]string{
		"username": "projectAUser",
		"password": "hellokang",
		"host": "127.0.0.1",
		"port": "3306",
		"dbname": "projectA",
		"collation": "utf8mb4_general_ci",
	}
	db, err := dao.NewDao(config)
	if err != nil {
		log.Println(err)
		return
	}

	// 查询分类的全部数据
	rows, err := db.Table("a_categories").Rows()
	log.Println(rows, err)
	if err == nil {
		// 响应结果
		c.JSON(200, gin.H{
			"error": "",
			"data": rows,
		})
	} else {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
	}

}

//动作 2
func CategoryDelete() {

}