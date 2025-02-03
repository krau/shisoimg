package api

import (
	"errors"
	"net/http"

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

func v1RandomArtworks(c *gin.Context) {
	var request GetRandomArtworksRequest
	if err := c.ShouldBind(&request); err != nil {
		GinBindError(c, err)
		return
	}
	images, err := dao.GetImageListRandom(request.Limit)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			GinErrorResponse(c, http.StatusNotFound, "Artworks not found")
			return
		}
		utils.L.Errorf("Failed to get random artworks: %v", err)
		GinErrorResponse(c, http.StatusInternalServerError, "Failed to get random artworks")
		return
	}
	if len(images) == 0 {
		GinErrorResponse(c, http.StatusNotFound, "Artworks not found")
		return
	}
	c.JSON(200, ResponseFromImages(images))
}

func v1ListArtworks(c *gin.Context) {
	var request GetArtworkListRequest
	if err := c.ShouldBind(&request); err != nil {
		GinBindError(c, err)
		return
	}
	images, err := dao.GetImageList(request.Page, request.PageSize)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			GinErrorResponse(c, http.StatusNotFound, "Artworks not found")
			return
		}
		utils.L.Errorf("Failed to get random artworks: %v", err)
		GinErrorResponse(c, http.StatusInternalServerError, "Failed to get random artworks")
		return
	}
	if len(images) == 0 {
		GinErrorResponse(c, http.StatusNotFound, "Artworks not found")
		return
	}
	c.JSON(200, ResponseFromImages(images))
}

func v1GetArtwork(c *gin.Context) {
	md5 := c.Param("id")
	if len(md5) != 32 || !md5Pattern.MatchString(md5) {
		GinErrorResponse(c, http.StatusBadRequest, "Invalid artwork ID")
		return
	}
	image, err := dao.GetImageByMd5(md5)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			GinErrorResponse(c, http.StatusNotFound, "Artwork not found")
			return
		}
		GinErrorResponse(c, http.StatusInternalServerError, "Failed to get artwork")
		return
	}
	c.JSON(200, &RestfulCommonResponse[*Artwork]{
		Status:  200,
		Message: "Success",
		Data:    ResponseDataFromImage(*image),
	})
}
