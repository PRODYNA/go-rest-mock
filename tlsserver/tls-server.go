package tlsserver

import (
	"crypto/tls"
	"net"
	"net/http"
)

type EmbeddedServer struct {
	http.Server
	WebserverCertificate string
	WebserverKey         string
}

func (srv *EmbeddedServer) ListenAndServeTLS(addr string, handler http.Handler) error {

	/*
		This is where we "hide" or "override" the default "ListenAndServeTLS" method so we modify it to accept
		hardcoded certificates and keys rather than the default filenames
		The default implementation of ListenAndServeTLS was obtained from:
		https://github.com/zenazn/goji/blob/master/graceful/server.go#L33
		and tls.X509KeyPair (http://golang.org/pkg/crypto/tls/#X509KeyPair) is used,
		rather than the default tls.LoadX509KeyPair
	*/
	srv.Handler = handler

	config := &tls.Config{
		MinVersion: tls.VersionTLS10,
	}
	if srv.TLSConfig != nil {
		*config = *srv.TLSConfig
	}
	if config.NextProtos == nil {
		config.NextProtos = []string{"http/1.1"}
	}

	var err error
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], err = tls.X509KeyPair([]byte(srv.WebserverCertificate), []byte(srv.WebserverKey))
	if err != nil {
		return err
	}

	conn, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	tlsListener := tls.NewListener(conn, config)
	return srv.Serve(tlsListener)
}
