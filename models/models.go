package models

import (
	"time"
)

// MailboxConf 存储邮箱配置信息
type MailboxConf struct {
	Title         string
	Body          string
	RecipientList []string
	Sender        string
	SPassword     string
	SMTPAddr      string
	SMTPPort      int
}

type Meet struct {
	Id          int       `gorm:"column:id;type:int(11);AUTO_INCREMENT;primary_key" json:"id"`
	Name        string    `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	Time        time.Time `gorm:"column:time;type:datetime;NOT NULL" json:"time"`
	Location    string    `gorm:"column:location;type:varchar(255);NOT NULL" json:"location"`
	Description string    `gorm:"column:description;type:text" json:"description"`
	Organizer   string    `gorm:"column:organizer;type:varchar(255);NOT NULL" json:"organizer"`
}

func (m *Meet) TableName() string {
	return "meet"
}

type Invite struct {
	Id        int       `gorm:"column:id;type:int(11);AUTO_INCREMENT;primary_key" json:"id"`
	MeetingId int       `gorm:"column:meeting_id;type:int(11);NOT NULL" json:"meeting_id"`
	Email     string    `gorm:"column:email;type:varchar(255);NOT NULL" json:"email"`
	SentTime  time.Time `gorm:"column:sent_time;type:datetime;NOT NULL" json:"sent_time"`
	Status    string    `gorm:"column:status;type:enum('sent','accepted','declined');NOT NULL" json:"status"`
}

func (m *Invite) TableName() string {
	return "invite"
}
