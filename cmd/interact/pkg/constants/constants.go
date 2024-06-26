package constants

var (
	MySQLUserName string
	MySQLPassWord string
	MySQLHost     string
	MySQLPort     string
	MySQLName     string
	MySQLDSN      string

	RedisUserName string
	RedisPassWord string
	RedisHost     string
	RedisPort     string

	ServiceHost string
	ServicePort string
	ServiceAddr string

	EtcdHost string
	EtcdPort string
	EtcdAddr string

	JaegerHost string
	JaegerPort string
	JaegerAddr string
)

const (
	CommentTable = "comment"
	LikeTable    = "like"
)
