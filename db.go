package main

type USER struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	UserName   string `gorm:"type:varchar(50);unique;not null" json:"user_name"`
	Password   string `gorm:"type:varchar(191);not null" json:"password"`
	Freeze     bool   `gorm:"tinyint(1);not null;comment:是否冻结" json:"freeze"`
	CreateTime string `gorm:"type:varchar(50)" json:"create_time"`
	Avatar     string `gorm:"type:varchar(90)" json:"avatar"`
	NickName   string `gorm:"type:varchar(90)" json:"nick_name"`
	Version    string `gorm:"type:varchar(191)" json:"version"`
}

func (t *USER) TableName() string {
	return "t_user"
}
