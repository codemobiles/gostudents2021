package model

type User struct {
	username string `form:"username" binding:"required"`
	password string `form:"password" binding:"required"`
}
