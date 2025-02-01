package dao

import (
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/krau/shisoimg/utils"
	"gorm.io/gorm"
)

func CreateImagesFromDir(rootPath string) (int, error) {
	imageExts := []string{".jpg", ".jpeg", ".png", ".gif", ".webp", ".avif", ".heic", ".heif", ".bmp", ".tiff"}
	count := 0
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		utils.L.Debugf("Processing %s", path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		isImage := false
		for _, imgExt := range imageExts {
			if ext == imgExt {
				isImage = true
				break
			}
		}
		if !isImage {
			return nil
		}

		md5, err := utils.CalcFileMD5(path)
		if err != nil {
			return err
		}

		img := Image{
			Path: path,
			Md5:  md5,
		}
		count++
		if err := db.Where("md5 = ?", md5).First(&Image{}).Error; err == nil {
			return db.Model(&Image{}).Where("md5 = ?", md5).Update("path", path).Error
		}
		return db.Create(&img).Error
	})
	return count, err
}

func GetImageList(page, pageSize int) ([]Image, error) {
	var images []Image
	err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&images).Error
	return images, err
}

func GetImageListRandom(limit int) ([]Image, error) {
	var images []Image
	err := db.Limit(limit).Order("RANDOM()").Find(&images).Error
	return images, err
}

func GetRandomImage() (*Image, error) {
	var img Image
	var count int64
	err := db.Model(&Image{}).Count(&count).Error
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	randIdx := rand.Intn(int(count))
	err = db.Offset(randIdx).Limit(1).Find(&img).Error
	return &img, err
}

func GetImageByMd5(md5 string) (*Image, error) {
	var img Image
	err := db.Where("md5 = ?", md5).First(&img).Error
	return &img, err
}
