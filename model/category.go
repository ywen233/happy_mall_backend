package model

type Category struct {
	CategoryID string `json:"categoryID" gorm:"column:category_id"`
	Name       string `json:"name" gorm:"column:name"`
	Desc       string `json:"desc" gorm:"column:desc"`
	Order      int    `json:"order" gorm:"column:order"`
	ParentId   string `gorm:"column:parent_id"`
	IsDeleted  bool   `json:"isDeleted" gorm:"column:is_deleted"`
}