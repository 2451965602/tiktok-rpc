package constants

var (
	MySQLUserName string
	MySQLPassWord string
	MySQLHost     string
	MySQLPort     string
	MySQLName     string
	MySQLDSN      string

	ServiceHost string
	ServicePort string
	ServiceAddr string

	EtcdHost string
	EtcdPort string
	EtcdAddr string
)

const (
	SocialTable = "social"

	DefaultTime = "1970-01-01 08:00:00"
	TimeFormat  = "2006-01-02 15:04:05"
)
