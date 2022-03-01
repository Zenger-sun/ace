package ace

type Context struct {
	Logger
	environ *Environ
	server  Listener
}

func (c *Context) Setup(environ *Environ) {

	c.Logger = GetLogItf("ace", environ.logPath)
	c.environ = environ

	certPath, privPath := environ.GetCertPath()
	switch environ.GetNetwork() {
	case NET_HTTP:
		c.server = NewHttpServer(environ.GetAddr(), certPath, privPath, nil)
	}

	go c.server.Serve()

	c.Debug("ace is rising...")
}

func (c *Context) AddPath(r Router, h HandlerFunc) {
	c.server.AddRouter(r, h)
}

func (c *Context) Shutdown() {
	c.Debug("ace shutdown.")
	c.server.Close()
	c.Logger.Close()
}

func NewContext() *Context {
	return &Context{}
}
