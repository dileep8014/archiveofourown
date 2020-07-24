package setting

import "time"

type ServerSetting struct {
	RunMode      string
	HttpPort     string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

type AppSetting struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
	ContextTimeout  int
}

type DatabaseSetting struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdConns   int
	MaxOpenConns int
}

type JwtSetting struct {
	Secret string
	Issuer string
	Expire int
}

type EmailSetting struct {
	Host     string
	Port     int
	UserName string
	Password string
	From     string
	IsSSL    bool
	To       []string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	return s.vp.UnmarshalKey(k, v)
}
