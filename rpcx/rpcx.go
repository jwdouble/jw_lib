package rpcx

import (
	"crypto/tls"
	"net/http"
	"os"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"jw.lib/conf"
	"jw.lib/logx"
)

type GRPCServer struct {
	Server *http.Server
}

// New return server with tls
//if tls err or tls not exist just err
func New(addr string, f func(grpcServer *grpc.Server, gwmux *runtime.ServeMux, dopts []grpc.DialOption)) *GRPCServer {
	if conf.Get("app.network.tls_enable") == "false" {
		grpcServer := grpc.NewServer()
		mux := http.NewServeMux()
		gwmux := runtime.NewServeMux()
		dopts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

		if f != nil {
			f(grpcServer, gwmux, dopts)
		}

		mux.Handle("/", gwmux)
		srv := &http.Server{
			Addr:    addr,
			Handler: grpcHandleFunc(grpcServer, mux),
		}

		return &GRPCServer{srv}
	}

	tf, err := credentials.NewServerTLSFromFile("/static/tf.pem", "/static/tf.key")
	if err != nil {
		logx.Errorf(err, "NewServerTLSFromFile")
	}

	grpcServer := grpc.NewServer(grpc.Creds(tf))
	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(tf)}

	if f != nil {
		f(grpcServer, gwmux, dopts)
	}

	mux.Handle("/", gwmux)
	srv := &http.Server{
		Addr:      addr,
		Handler:   grpcHandleFunc(grpcServer, mux),
		TLSConfig: getTLSConfig(),
	}

	return &GRPCServer{srv}
}

func grpcHandleFunc(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有跨域请求
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func getTLSConfig() *tls.Config {
	pem, _ := os.ReadFile("/static/tls.pem")
	key, _ := os.ReadFile("/static/tls.key")

	cert, err := tls.X509KeyPair(pem, key)
	if err != nil {
		logx.Errorf(err, "tls.X509KeyPair")
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{http2.NextProtoTLS}, // HTTP2 TLS支持
	}
}

func (s *GRPCServer) Run() error {
	logx.Infof("server start finish")
	return s.Server.ListenAndServe()
}
