package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/krau/shisoimg/dao"
	"github.com/krau/shisoimg/utils"
	"gorm.io/gorm"
)

func randomImage(c *gin.Context) {
	image, err := dao.GetRandomImage()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.String(404, "no image found")
		return
	}
	if err != nil {
		utils.L.Errorf("Failed to get random image: %v", err)
		c.String(500, "internal server error")
		return
	}
	match, newUrl := applyRules(image.Path)
	if match {
		c.Redirect(302, newUrl)
	}
	c.Redirect(302, "/images/"+image.Md5)
}

func getImage(c *gin.Context) {
	md5 := c.Param("md5")
	image, err := dao.GetImageByMd5(md5)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.String(404, "no image found")
		return
	}
	if err != nil {
		utils.L.Errorf("Failed to get image %s: %v", md5, err)
		c.String(500, "internal server error")
		return
	}
	match, newUrl := applyRules(image.Path)
	if match {
		c.Redirect(302, newUrl)
	}
	c.File(image.Path)
}
