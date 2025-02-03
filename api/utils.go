package api

import (
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krau/shisoimg/dao"
	"github.com/krau/shisoimg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	md5Pattern = regexp.MustCompile(`^[a-f0-9]{32}$`)
)

func applyRules(path string) (match bool, newUrl string) {
	for _, rule := range dao.Rules() {
		if strings.HasPrefix(path, rule.Path) {
			parsedUrl, err := url.JoinPath(rule.Prefix, strings.TrimPrefix(path, rule.Path))
			if err != nil {
				utils.L.Errorf("Failed to join path %s with %s: %v", rule.Prefix, path, err)
				continue
			}
			return true, parsedUrl
		}
	}
	return false, ""
}

type RestfulCommonResponse[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func GinErrorResponse(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, &RestfulCommonResponse[any]{
		Status:  status,
		Message: message,
		Data:    nil,
	})
	ctx.Abort()
}

func GinBindError(ctx *gin.Context, err error) {
	GinErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
}

type GetRandomArtworksRequest struct {
	Limit int `form:"limit,default=1" binding:"gte=1,lte=200" json:"limit"`
}

type GetArtworkListRequest struct {
	Page     int `form:"page,default=1" binding:"omitempty,gte=1" json:"page"`
	PageSize int `form:"page_size,default=20" binding:"omitempty,gte=1,lte=200" json:"page_size"`
}

type Artwork struct {
	ID          string    `json:"id" bson:"_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	R18         bool      `json:"r18" bson:"r18"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	SourceType  string    `json:"source_type" bson:"source_type"`
	SourceURL   string    `json:"source_url" bson:"source_url"`
	LikeCount   uint      `json:"like_count" bson:"like_count"`

	Artist   *Artist            `json:"artist" bson:"artist"`
	Tags     []string           `json:"tags" bson:"tags"`
	Pictures []*PictureResponse `json:"pictures" bson:"pictures"`
}

type Artist struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Type     string `json:"type" bson:"type"`
	UID      string `json:"uid" bson:"uid"`
	Username string `json:"username" bson:"username"`
}

type PictureResponse struct {
	ID        string `json:"id"`
	Width     uint   `json:"width"`
	Height    uint   `json:"height"`
	Index     uint   `json:"index"`
	Hash      string `json:"hash"`
	FileName  string `json:"file_name"`
	MessageID int    `json:"message_id"`
	Thumbnail string `json:"thumbnail"`
	Regular   string `json:"regular"`
}

var fakeArtist = &Artist{
	ID:       primitive.NewObjectID().Hex(),
	Name:     "shisoimg",
	Type:     "shisoimg",
	UID:      "shisoimg",
	Username: "shisoimg",
}

func ResponseDataFromImage(image dao.Image) *Artwork {
	match, newUrl := applyRules(image.Path)
	if !match {
		newUrl = "/images/" + image.Md5
	}
	return &Artwork{
		ID:         image.Md5,
		Artist:     fakeArtist,
		Title:      image.Md5,
		Tags:       []string{},
		SourceType: "shisoimg",
		SourceURL:  "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
		CreatedAt:  image.CreatedAt,
		Pictures: []*PictureResponse{
			{
				ID:        image.Md5,
				Index:     0,
				Thumbnail: newUrl,
				Regular:   newUrl,
				Width:     uint(image.Width),
				Height:    uint(image.Height),
				Hash:      image.Md5,
				FileName:  path.Base(image.Path),
				MessageID: 0,
			},
		},
	}
}

func ResponseFromImages(images []dao.Image) RestfulCommonResponse[any] {
	artworks := make([]*Artwork, 0, len(images))
	for _, image := range images {
		artworks = append(artworks, ResponseDataFromImage(image))
	}
	return RestfulCommonResponse[any]{
		Status:  200,
		Message: "Success",
		Data:    artworks,
	}
}
