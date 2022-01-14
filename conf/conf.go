package conf

type Connector struct {
	appName    string
	driverName string
	addr       string
}

var (
	AppPgConn    = Connector{appName: "app_pg", driverName: "postgres"}
	AppRedisConn = Connector{appName: "app_redis", driverName: "redis"}
)

func (c *Connector) Value(s string) *Connector {
	c.addr = s
	return c
}

func (c *Connector) GetAppName() string {
	return c.appName
}

func (c *Connector) GetDriverName() string {
	return c.driverName
}

func (c *Connector) GetAddr() string {
	return c.addr
}
