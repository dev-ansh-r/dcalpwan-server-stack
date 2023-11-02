// NOTE: Set CAROOT=. when running go generate, i.e.:
// $ CAROOT=. go generate .

//go:generate mkcert -cert-file servercert.pem -key-file serverkey.pem localhost 127.0.0.1 ::1
//go:generate mkcert -cert-file clientcert.pem -key-file clientkey.pem -client localhost

package testdata
