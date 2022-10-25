package models

type RegisterUserInput struct {
	Age      uint   `json:"age" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
}

type PhotoForm struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" binding:"required"`
}

type UpdatePhotoForm struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}

type CommentInput struct {
	Message string `json:"message" binding:"required"`
	PhotoID uint   `json:"photo_id"`
}

type UpdateCommentInput struct {
	Message string `json:"message" binding:"required"`
}

type SosmedInput struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" binding:"required"`
}
