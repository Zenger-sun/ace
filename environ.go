package ace

import (
	"ace/utils/cert"
	"flag"
)

type Environ struct {
	addr    string
	network uint8
	logPath string
}

func (e *Environ) SetOpt() error {
	network := flag.Int("T", 0, "set net type.")
	addr := flag.String("I", "127.0.0.1:8000", "set listen addr.")
	flag.Parse()

	e.addr = *addr
	e.network = uint8(*network)

	return nil
}

func (e *Environ) GetAddr() string {
	return e.addr
}

func (e *Environ) GetNetwork() uint8 {
	return e.network
}

func (e *Environ) GetCertPath() (string, string) {
	return cert.GetCertificatePaths()
}
