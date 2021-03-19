package main

import (
	"github.com/placons/go-rest-mock/config"
	"github.com/placons/go-rest-mock/handler"
	"github.com/placons/go-rest-mock/model"
	"github.com/placons/go-rest-mock/reader"
	"github.com/placons/go-rest-mock/tlsserver"
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	cfg := config.Parse()

	files := reader.ReadFiles(cfg.Path)
	size := len(files)
	if size == 0 {
		fullPath, err := filepath.Abs(cfg.Path)
		if err == nil {
			log.Println("No mock definitions found in path " + fullPath)
		} else {
			log.Println("No mock definitions found in path " + cfg.Path)
		}
		return
	}
	for i, file := range files {

		if file.IsDir() {
			continue
		}

		md := reader.ReadDefinition(cfg.Path + "/" + file.Name())

		if i == size-1 {
			// last one blocks and prevents from exiting
			runServer(md, cfg)
		} else {
			// using non blocking listen & serve
			go func() {
				runServer(md, cfg)
			}()
		}
	}
}

func runServer(md *model.MockDefinition, cfg *config.Config) {
	log.Println("Starting mock on port:", md.Port, "for backend:", md.ID)
	log.SetFlags(log.Llongfile)
	if md.TLS {
		s := &tlsserver.EmbeddedServer{
			WebserverCertificate: "-----BEGIN CERTIFICATE-----\nMIICuTCCAkCgAwIBAgIUG652nDShCBKdp7BdQq2IPsoFecIwCgYIKoZIzj0EAwIw\ngZMxCzAJBgNVBAYTAkRFMQwwCgYDVQQIDANOUlcxFDASBgNVBAcMC0R1ZXNzZWxk\nb3JmMRMwEQYDVQQKDApQUk9EWU5BIFNFMQswCQYDVQQLDAJQUzEVMBMGA1UEAwwM\nTWFydGluIEtydXNlMScwJQYJKoZIhvcNAQkBFhhtYXJ0aW4ua3J1c2VAcHJvZHlu\nYS5jb20wHhcNMjEwMzE5MTgxNTUzWhcNMzEwMzE3MTgxNTUzWjCBkzELMAkGA1UE\nBhMCREUxDDAKBgNVBAgMA05SVzEUMBIGA1UEBwwLRHVlc3NlbGRvcmYxEzARBgNV\nBAoMClBST0RZTkEgU0UxCzAJBgNVBAsMAlBTMRUwEwYDVQQDDAxNYXJ0aW4gS3J1\nc2UxJzAlBgkqhkiG9w0BCQEWGG1hcnRpbi5rcnVzZUBwcm9keW5hLmNvbTB2MBAG\nByqGSM49AgEGBSuBBAAiA2IABKDSP0+RCjZOgcRD6hGB9BxFvOUaf7Z2c8Wt6Dqv\nFkWMpUcEPpDKzDS7lmHOC23WFnsaTUDOgBxxi5YMKhxyqTejqSZufPjdtRBzAQef\nkYeB1doiKzBgE3NBgx8+XDIlBKNTMFEwHQYDVR0OBBYEFJeJ8Etlxxvp3u5+rkGz\nHnAXHIBFMB8GA1UdIwQYMBaAFJeJ8Etlxxvp3u5+rkGzHnAXHIBFMA8GA1UdEwEB\n/wQFMAMBAf8wCgYIKoZIzj0EAwIDZwAwZAIwfCfwHOn40IIiJ8tA+dcc1dS4fspf\nqRGbj8+8ZE2o0D8+wdyRAWJ1oZtaK25wW/gSAjAnEv2XiXfplr4OC/z8XOTyViET\nZ444cs7LyYFEv1GqW/7YlXaGZClXmm5VHJsltxQ=\n-----END CERTIFICATE-----\n",
			WebserverKey:         "-----BEGIN EC PARAMETERS-----\nBgUrgQQAIg==\n-----END EC PARAMETERS-----\n-----BEGIN EC PRIVATE KEY-----\nMIGkAgEBBDBF/fTvwPk+31AxwgaBl4zGjL3Zs5p905xU/7PpZ1nEjfKYurcjVDNs\nzeLjc+RrkWWgBwYFK4EEACKhZANiAASg0j9PkQo2ToHEQ+oRgfQcRbzlGn+2dnPF\nreg6rxZFjKVHBD6Qysw0u5Zhzgtt1hZ7Gk1AzoAccYuWDCoccqk3o6kmbnz43bUQ\ncwEHn5GHgdXaIiswYBNzQYMfPlwyJQQ=\n-----END EC PRIVATE KEY-----",
		}

		_ = s.ListenAndServeTLS(":"+md.Port, handler.NewHandler(md, cfg))
	} else {
		log.Fatal(http.ListenAndServe(":"+md.Port, handler.NewHandler(md, cfg)))
	}
}
