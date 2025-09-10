/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package main

import (
	"context"
	"gitee.com/dn-jinmin/tlog"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"imooc.com/easy-chat/pkg/zlog"
	"testing"
)

func TestFile_log(t *testing.T) {
	logx.SetUp(logx.LogConf{
		ServiceName: "file.log",
		Mode:        "file",
		Encoding:    "json",
		Path:        "./zlog",
	})

	logx.Info("test file log")
	logx.Info("这是测试go-zero日志")
	logx.Info("将日志推送至elk")

	for {
	}
}

func TestRedis_log_ioWriter(t *testing.T) {
	io := zlog.NewRedisIoWriter("redis.io.writer", redis.RedisConf{
		Host:        "192.168.117.24:16379",
		Type:        "node",
		Pass:        "easy-chat",
		Tls:         false,
		NonBlock:    false,
		PingTimeout: 0,
	})

	logx.SetWriter(logx.NewWriter(io))

	logx.Infow("test file log", logx.LogField{
		Key:   "rid",
		Value: "1111",
	})
	logx.Info("这是测试go-zero日志")
	logx.Info("将日志推送至elk")

	//for {
	//}
}

func TestRedis_logx_ioWriter(t *testing.T) {
	io := zlog.NewRedisLogxWriter("redis.io.writer", redis.RedisConf{
		Host:        "192.168.117.24:16379",
		Type:        "node",
		Pass:        "easy-chat",
		Tls:         false,
		NonBlock:    false,
		PingTimeout: 0,
	})

	logx.SetWriter(io)

	logx.Info("test file log")
	logx.Info("这是测试go-zero日志")
	logx.Info("将日志推送至elk")

	//for {
	//}
}

func TestCtx_Log(t *testing.T) {
	io := zlog.NewRedisIoWriter("redis.io.writer", redis.RedisConf{
		Host:        "192.168.117.24:16379",
		Type:        "node",
		Pass:        "easy-chat",
		Tls:         false,
		NonBlock:    false,
		PingTimeout: 0,
	})

	logx.SetWriter(logx.NewWriter(io))

	ctx, _ := sdktrace.NewTracerProvider().Tracer(trace.TraceName).Start(context.Background(), "a")

	log := logx.WithContext(ctx)
	log.Info("test file log")
	log.Info("这是测试go-zero日志")
	log.Info("将日志推送至elk")

	logx.Info("test file log")
	logx.Info("这是测试go-zero日志")
	logx.Info("将日志推送至elk")
}

func TestCtx_TLog(t *testing.T) {
	ctx := tlog.TraceStart(context.Background())

	tlog.InfoCtx(ctx, "1", "测试")
	tlog.InfoCtx(ctx, "1", "测试2")
	tlog.InfoCtx(ctx, "1", "测试3")

	for {
	}
}

func TestCtx_TLog_Logx(t *testing.T) {
	tlog.Init(&tlog.Config{
		LoggerWriter: []tlog.LoggerWriter{zlog.NewTlog(redis.RedisConf{
			Host:        "192.168.117.24:16379",
			Type:        "node",
			Pass:        "easy-chat",
			Tls:         false,
			NonBlock:    false,
			PingTimeout: 0,
		})},
	})

	ctx, _ := sdktrace.NewTracerProvider().Tracer(trace.TraceName).Start(context.Background(), "a")
	traceId := trace.TraceIDFromContext(ctx)
	if traceId == "" {
		ctx = tlog.TraceStart(ctx)
	} else {
		ctx = tlog.TraceStartSetTraceId(ctx, traceId)
	}

	tlog.InfoCtx(ctx, "1", "测试")
	tlog.InfoCtx(ctx, "1", "测试2")
	tlog.InfoCtx(ctx, "1", "测试3")

	for {
	}
}
