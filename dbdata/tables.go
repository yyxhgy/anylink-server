package dbdata

import (
	"encoding/json"
	"time"
)

type Group struct {
	Id           int            `json:"id" xorm:"pk autoincr not null"`
	Name         string         `json:"name" xorm:"varchar(60) not null unique"`
	Note         string         `json:"note" xorm:"varchar(255)"`
	AllowLan     bool           `json:"allow_lan" xorm:"Bool"`
	ClientDns    []ValData      `json:"client_dns" xorm:"Text"`
	RouteInclude []ValData      `json:"route_include" xorm:"Text"`
	RouteExclude []ValData      `json:"route_exclude" xorm:"Text"`
	LinkAcl      []GroupLinkAcl `json:"link_acl" xorm:"Text"`
	Bandwidth    int            `json:"bandwidth" xorm:"Int"` // 带宽限制
	Status       int8           `json:"status" xorm:"Int"`    // 1正常
	CreatedAt    time.Time      `json:"created_at" xorm:"DateTime created"`
	UpdatedAt    time.Time      `json:"updated_at" xorm:"DateTime updated"`
}

type User struct {
	Id       int    `json:"id" xorm:"pk autoincr not null"`
	Username string `json:"username" xorm:"varchar(60) not null unique"`
	Nickname string `json:"nickname" xorm:"varchar(255)"`
	Email    string `json:"email" xorm:"varchar(255)"`
	// Password  string    `json:"password"`
	PinCode    string    `json:"pin_code" xorm:"varchar(32)"`
	OtpSecret  string    `json:"otp_secret" xorm:"varchar(255)"`
	DisableOtp bool      `json:"disable_otp" xorm:"Bool"` // 禁用otp
	Groups     []string  `json:"groups" xorm:"Text"`
	Status     int8      `json:"status" xorm:"Int"` // 1正常
	SendEmail  bool      `json:"send_email" xorm:"Bool"`
	CreatedAt  time.Time `json:"created_at" xorm:"DateTime created"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"DateTime updated"`
}

type IpMap struct {
	Id        int       `json:"id" xorm:"pk autoincr not null"`
	IpAddr    string    `json:"ip_addr" xorm:"varchar(32) not null unique"`
	MacAddr   string    `json:"mac_addr" xorm:"varchar(32) not null unique"`
	Username  string    `json:"username" xorm:"varchar(60)"`
	Keep      bool      `json:"keep" xorm:"Bool"` // 保留 ip-mac 绑定
	KeepTime  time.Time `json:"keep_time" xorm:"DateTime"`
	Note      string    `json:"note" xorm:"varchar(255)"` // 备注
	LastLogin time.Time `json:"last_login" xorm:"DateTime updated"`
	UpdatedAt time.Time `json:"updated_at" xorm:"DateTime updated"`
}

type Setting struct {
	Id        int             `json:"id" xorm:"pk autoincr not null"`
	Name      string          `json:"name" xorm:"varchar(60) not null unique"`
	Data      json.RawMessage `json:"data" xorm:"Text"`
	UpdatedAt time.Time       `json:"updated_at" xorm:"DateTime updated"`
}

type AccessAudit struct {
	Id        int       `json:"id" xorm:"pk autoincr not null"`
	Username  string    `json:"username" xorm:"varchar(60) not null"`
	Protocol  uint8     `json:"protocol" xorm:"not null"`
	Src       string    `json:"src" xorm:"varchar(60) not null"`
	SrcPort   uint16    `json:"src_port" xorm:"not null"`
	Dst       string    `json:"dst" xorm:"varchar(60) not null"`
	DstPort   uint16    `json:"dst_port" xorm:"not null"`
	CreatedAt time.Time `json:"created_at" xorm:"DateTime"`
}
