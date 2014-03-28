package service

import (
	"flag"
	"fmt"
)

type Options struct {
	PidPath   string
	Port      int
	Interface string
	CertPath  string
	EtcdNodes listOptions
	EtcdPath  string
}

// Helper to parse options that can occur several times, e.g. cassandra nodes
type listOptions []string

func (o *listOptions) String() string {
	return fmt.Sprint(*o)
}

func (o *listOptions) Set(value string) error {
	*o = append(*o, value)
	return nil
}

func ParseCommandLine() (options Options, err error) {
	flag.Var(&options.EtcdNodes, "etcd", "Etcd discovery service API endpoints")
	flag.StringVar(&options.EtcdPath, "etcdPath", "vulcand", "Etcd path for storing configuration")
	flag.StringVar(&options.PidPath, "pidPath", "", "Path to write PID file to")
	flag.IntVar(&options.Port, "port", 8181, "Port to listen on")
	flag.StringVar(&options.Interface, "interface", "", "Interface to bind to")
	flag.StringVar(&options.CertPath, "certPath", "", "Certificate to use (enables TLS)")
	flag.Parse()
	return options, nil
}