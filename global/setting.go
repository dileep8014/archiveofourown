package global

import (
	"github.com/shyptr/archiveofourown/pkg/errwrap"
	"github.com/shyptr/archiveofourown/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	JWTSetting      *setting.JwtSetting
	EmailSetting    *setting.EmailSetting
)

func SetupSetting() (err error) {
	defer errwrap.Add(&err, "init.SetupSetting")

	st, err := setting.NewSetting()
	if err != nil {
		return
	}
	err = st.ReadSection("Server", &ServerSetting)
	if err != nil {
		return
	}
	err = st.ReadSection("App", &AppSetting)
	if err != nil {
		return
	}
	err = st.ReadSection("Database", &DatabaseSetting)
	if err != nil {
		return
	}
	err = st.ReadSection("JWT", &JWTSetting)
	if err != nil {
		return
	}
	err = st.ReadSection("Email", &EmailSetting)
	if err != nil {
		return
	}
	return
}
