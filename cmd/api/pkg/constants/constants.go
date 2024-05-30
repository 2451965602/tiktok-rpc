package constants

var (
	MySQLUserName string
	MySQLPassWord string
	MySQLHost     string
	MySQLPort     string
	MySQLName     string
	MySQLDSN      string

	QiNiuBucket    string
	QiNiuAccessKey string
	QiNiuSecretKey string
	QiNiuDomain    string

	ServiceHost   string
	ServicePort   string
	WebsocketPort string
	ServiceAddr   string
	WebsocketAddr string

	EtcdHost string
	EtcdPort string
	EtcdAddr string

	SentinelThreshold        float64
	SentinelStatIntervalInMs uint32
)

const (
	MsgTable = "messages"

	ContextUid = "userid"

	TimeFormat = "2006-01-02 15:04:05"
)
