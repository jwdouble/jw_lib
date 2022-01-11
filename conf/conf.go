package conf

type Connector struct {
	appName    string
	driverName string
	dns        string
}

var (
	AppPgConn    = Connector{appName: "app_pg", driverName: "postgres"}
	AppRedisConn = Connector{appName: "app_redis", driverName: "redis"}
)

func (c *Connector) Value(s string) *Connector {
	c.dns = s
	return c
}

func (c *Connector) GetAppName() string {
	return c.appName
}

func (c *Connector) GetDriverName() string {
	return c.driverName
}

func (c *Connector) GetDSN() string {
	return c.dns
}
