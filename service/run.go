package service

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	opentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"google.golang.org/grpc"
	kitCfg "kit/config"
	"kit/log"
	pb "kit/model/example"
	"kit/net/etcd"
	"kit/net/httpsvr"
	"kit/net/pprof"
	"kit/net/router"
	"net"
	"net/http"
	"reply/config"
	"time"
)

var (
	svr *service
	// trace
	collector zipkin.Collector
	tracer    opentracing.Tracer
)

func Run(c *config.Config) (err error) {
	if err = initTrace(c.Trace, c.Common); err != nil {
		return
	}

	// run pprof
	if err = pprof.Init(c.Mhttp.Pprof); err != nil {
		return
	}

	// run http
	if err = runHttp(c.Mhttp, c.Router); err != nil {
		return
	}

	// run grpc
	if err = runRpc(c.Grpc); err != nil {
		return
	}
	return
}

func initTrace(c *kitCfg.Trace, cf *kitCfg.Common) (err error) {
	collector, err = zipkin.NewHTTPCollector(c.Addr)
	if err != nil {
		log.Error("zipkin.NewHTTPCollector err(%v)", err)
		return
	}
	recorder := zipkin.NewRecorder(collector, c.Debug, cf.HostPort, cf.Family)
	tracer, err = zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(c.SameSpan),
		zipkin.TraceID128Bit(c.TraceID128Bit),
	)
	if err != nil {
		log.Error("zipkin.NewTracer error(%v)", err)
	}
	opentracing.InitGlobalTracer(tracer)
	return
}

// rpc
func runRpc(c *kitCfg.Grpc) (err error) {
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", c.Port))
	if err != nil {
		log.Error("run rpc error(%v)", err)
		return
	} else {
		log.Info("run rpc success port:%d", c.Port)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads()),
		),
	)
	servi := &helloServer{svr: svr}
	pb.RegisterHelloServiceServer(s, servi)
	// go rpc
	go func() {
		if err = s.Serve(listen); err != nil {
			log.Error("run rpc fail error(%v)", err)
			return
		}
	}()
	// register to etcd
	err = etcd.Register(c.Name, c.Addr, c.Port, c.EtcdAddr, time.Second*10, 15)
	if err != nil {
		log.Error("etcd register error(%v)", err)
	}
	return
}

// http
func runHttp(c *kitCfg.Mhttp, cr *kitCfg.Router) (err error) {
	// internal
	inMux := http.NewServeMux()
	inRou := router.NewRouter(cr, inMux)
	initInner(inRou)
	if err = httpsvr.RunHttp(c.Inner, inMux); err != nil {
		log.Error("httpsvr.RunHttp error(%v)", err)
		return
	} else {
		log.Info("RunInnerHttp success port:%d", c.Inner.Port)
	}

	// outter
	outMux := http.NewServeMux()
	outRou := router.NewRouter(cr, outMux)
	initOutter(outRou)
	if err = httpsvr.RunHttp(c.Outter, outMux); err != nil {
		log.Error("RunOutterHttp error(%v)", err)
		return
	} else {
		log.Info("RunOutterHttp success port:%d", c.Outter.Port)
	}
	return
}

// service exit need do this thing
func EndTracing() {
	if collector != nil {
		collector.Close()
	}
}

func UnRegisterEtcd() {
	etcd.UnRegister()
}
