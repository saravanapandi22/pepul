package models

type User struct {
	FullName string `json:"full_name" bson:"full_name"`
	Email string `json:"email" bson:"email"`
	PhoneNumber int `json:"phone_number" bson:"phone_number"`
	PictureUpload string `json:"picture_upload" bson:"picture_upload"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdateAt string `json:"update_at" bson:"update_at"`
}
