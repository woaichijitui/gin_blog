package models

type CommentModel struct {
	MODEL
	SubComments        []*CommentModel `json:"sub_comments"  gorm:"foreignKey:ParentCommentID"`        //子评论列表
	ParentCommentModel *CommentModel   `json:"parent_comment_model" gorm:"foreignKey:ParentCommentID"` //父级评论
	ParentCommentID    *uint           `json:"parent_comment_id" gorm:"size:10"`

	Content      string       `json:"content" gorm:"size:256"`               //评论内容
	DiggCount    int          `json:"digg_count" gorm:"size:8;default:0"`    //点赞数
	CommentCount int          `json:"comment_count" gorm:"size:8;default:0"` //评论数
	Article      ArticleModel `json:"article" `                              //关联的文章
	ArticleID    uint         `json:"article_id" gorm:"size:10"`             //文章id
	User         UserModel    `json:"user" `                                 //用户的昵称
	UserID       uint         `json:"user_id" gorm:"size:10"`                //评论的用户
}
