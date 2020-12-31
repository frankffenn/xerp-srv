package user

import "time"

type LoginType string

const (
	GuestLogin  LoginType = "guest"
	PhoneLogin  LoginType = "phone"
	WechatLogin LoginType = "wechat"
)

type LoginVar struct {
	LoginType LoginType `form:"login_type" json:"login_type" binding:"required"`
	Username  string    `form:"username" json:"username" binding:"required"`
	Password  string    `form:"password" json:"password" binding:"required"`
}

type LoginResp struct {
	GUID   string `json:"guid"`
	UserID uint64 `json:"user_id"`
	Level  uint64 `json:"level"`
}

// User ...
type User struct {
	ID        uint64     `xorm:"pk autoincr"`
	Username  string     `xorm:"'username' varchar(25) unique "`
	Password  string     `xorm:"'password' not null"`
	LoginType LoginType  `xorm:"'login_type'"`
	GUID      string     `json:"'guid'"`
	Level     uint64     `xorm:"'level'"`
	IsBanned  bool       `xorm:"'is_banned'`
	CreatedAt time.Time  `xorm:"created"`
	UpdatedAt time.Time  `xorm:"updated" `
	DeletedAt *time.Time `xorm:"-" deleted json:"omitempty"`
}

func (u *User) TableName() string {
	return "user"
}
