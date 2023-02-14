package models

type User struct {
	FullName string `form:"full_name" json:"full_name" bson:"full_name" validate:"required"`
	Email string `form:"email" json:"email" bson:"email" validate:"email"`
	PhoneNumber int `form:"phone_number" json:"phone_number" bson:"phone_number" validate:"min=10,max=10,number"`
	PictureUpload *multipart.FileHeader  `form:"picture_upload" json:"picture_upload" bson:"picture_upload"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdateAt string `json:"update_at" bson:"update_at"`
}

type Pagination struct {
	Limit int `json:"limit"`
	Page int `json:"page"`
}
