package entity

type User struct {
	id        int64  `boil:"id" json:"id"`
	name      string `boil:"name" json:"username"`
	email     string `boil:"email" json:"email"`
	createdAt string `boil:"created_at" json:"created_at"`
	updatedAt string `boil:"updated_at" json:"updated_at"`
	deletedAt string `boil:"deleted_at" json:"deleted_at"`
}
