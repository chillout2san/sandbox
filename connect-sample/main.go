// connect v2 で立てた最小構成の API サーバー。
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/connect/v2"
	"connectrpc.com/connect/v2/connecthttp"

	samplev1 "connect-sample/gen/sample/v1"
	"connect-sample/gen/sample/v1/samplev1connect"
)

// sampleServer は samplev1connect.SampleServiceHandler の実装。
// リクエストは読み捨てて固定値を返す。
type sampleServer struct{}

func (s *sampleServer) CreateHoge(_ context.Context, _ *samplev1.CreateHogeRequest) (*samplev1.CreateHogeResponse, error) {
	return &samplev1.CreateHogeResponse{
		Hoge: &samplev1.Hoge{
			Id:   "hoge-1",
			Name: "created hoge",
		},
	}, nil
}

func (s *sampleServer) GetHoge(_ context.Context, _ *samplev1.GetHogeRequest) (*samplev1.GetHogeResponse, error) {
	return &samplev1.GetHogeResponse{
		Hoge: &samplev1.Hoge{
			Id:   "hoge-1",
			Name: "fetched hoge",
		},
	}, nil
}

func main() {
	// connect.Server が procedure とハンドラの対応を持つ。トランスポート非依存。
	server := connect.NewServer(logInterceptor)
	samplev1connect.RegisterSampleServiceHandler(server, &sampleServer{})

	// connecthttp.Mount が net/http のルーターに各 procedure を登録する。
	// gRPC / gRPC-Web でも叩けるようにしておく。
	mux := http.NewServeMux()
	connecthttp.Mount(mux, server,
		connecthttp.WithGRPC(),
		connecthttp.WithGRPCWeb(),
	)

	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	// Ctrl-C で graceful shutdown する。
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("listening on %s", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen failed: %v", err)
			stop()
		}
	}()

	<-ctx.Done()
	log.Println("shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown failed: %v", err)
	}
}

// logInterceptor は呼ばれた procedure をログに出すだけのインターセプター。
func logInterceptor(next connect.ServerFunc) connect.ServerFunc {
	return func(ctx context.Context, spec connect.Spec, stream connect.ServerStream) error {
		log.Printf("rpc: %s", spec.Procedure)
		return next(ctx, spec, stream)
	}
}
