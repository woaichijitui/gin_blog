package models

type TagModel struct {
	MODEL
	Title    string         `json:"title" gorm:"size:16"` //文章标签的名称
	Articles []ArticleModel `json:"-" gorm:"many2many:article_tag"`
}
