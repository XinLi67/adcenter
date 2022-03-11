package policies

import (
	"adcenter/app/models/post"
	"adcenter/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyPost(c *gin.Context, postModel post.Post) bool {
	// return auth.CurrentUID(c) == postModel.UserID
	return auth.CurrentUID(c) == postModel.Title
}

// func CanViewPost(c *gin.Context, postModel post.Post) bool {}
// func CanCreatePost(c *gin.Context, postModel post.Post) bool {}
// func CanUpdatePost(c *gin.Context, postModel post.Post) bool {}
// func CanDeletePost(c *gin.Context, postModel post.Post) bool {}
