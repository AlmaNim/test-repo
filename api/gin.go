package api

import (
	"go_banners/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	router.GET("/user_banner", getUserBanner)
	router.GET("/banner", getBanner)
	router.POST("/banner", postBanner)
	router.PATCH("/banner/:id", patchBanner)
	router.DELETE("/banner/:id", deleteBanner)
	router.Run("localhost:8080")
}

func getUserBanner(c *gin.Context) {
	tagId, tagOk := c.GetQuery("tag_id")
	featureId, featureOk := c.GetQuery("feature_id")
	lastRevision, lastRevisionOk := c.GetQuery("use_last_revision")
	if !tagOk || !featureOk {
		c.JSON(http.StatusBadRequest, struct{ error string }{error: "parameters not found"})
		return
	}
	tagIdInt, tagError := strconv.Atoi(tagId)
	featureIdInt, featureError := strconv.Atoi(featureId)
	if tagError != nil || featureError != nil {
		c.JSON(http.StatusBadRequest, struct{ error string }{error: "parameters are not integers"})
		return
	}
	lastRevisionBool := false
	if lastRevisionOk && lastRevision == "true" {
		lastRevisionBool = true
	}
	result, error := db.GetBanner(tagIdInt, featureIdInt, lastRevisionBool)
	if error != nil && error.Error() == "not found" {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, result)
}

func getBanner(c *gin.Context) {
	tagId, tagOk := c.GetQuery("tag_id")
	featureId, featureOk := c.GetQuery("feature_id")
	limit, limitOk := c.GetQuery("limit")
	offset, offsetOk := c.GetQuery("offset")
	var tagIdInt *int
	var featureIdInt *int
	var limitInt *int
	var offsetInt *int
	if tagOk {
		tagId, tagError := strconv.Atoi(tagId)
		if tagError == nil {
			tagIdInt = &tagId
		}
	}
	if featureOk {
		featureId, featureError := strconv.Atoi(featureId)
		if featureError == nil {
			featureIdInt = &featureId
		}
	}
	if limitOk {
		limit, limitError := strconv.Atoi(limit)
		if limitError == nil {
			limitInt = &limit
		}
	}
	if offsetOk {
		offset, offsetError := strconv.Atoi(offset)
		if offsetError == nil {
			offsetInt = &offset
		}
	}
	result, err := db.GetBanners(featureIdInt, tagIdInt, limitInt, offsetInt)
	if err != nil {
		switch err.Error() {
		case "internal error":
			c.Writer.WriteHeader(http.StatusInternalServerError)
		case "not found":
			c.Writer.WriteHeader(http.StatusNotFound)
		default:
			c.Writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	c.JSON(http.StatusOK, result)
}

func postBanner(c *gin.Context) {
	body := PostBannerRequest{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, struct{ error string }{error: "parameters are not integers"})
		return
	}
	result, err := db.CreateBanner(body)

}

func patchBanner(c *gin.Context) {
	id := c.Param("id")
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(id))
}

func deleteBanner(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, struct{ error string }{error: "parameters are not integers"})
		return
	}
	err = db.DeleteBanner(idInt)
	if err != nil {
		switch err.Error() {
		case "internal error":
			c.Writer.WriteHeader(http.StatusInternalServerError)
		case "not found":
			c.Writer.WriteHeader(http.StatusNotFound)
		default:
			c.Writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent)
}
