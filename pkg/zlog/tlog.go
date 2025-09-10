/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package zlog

import (
	"context"
	"gitee.com/dn-jinmin/tlog"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var (
	tlogTraceId   = "trace"
	tlogRelatedId = "tlog_relatedId"
	tlogPath      = "tlog_path"
	tlogMsg       = "tlog_msg"
	tlogLabel     = "tlog_label"
)

type Tlog struct{}

func NewTlog(redisCfg redis.RedisConf) *Tlog {
	io := NewRedisIoWriter("www.imooc.com", redisCfg)
	logx.SetWriter(logx.NewWriter(io))
	return &Tlog{}
}

func (z *Tlog) Write(ctx context.Context, content *tlog.Content, fs ...tlog.Field) {

	fields := z.buildFields(content)

	switch content.Level {
	case tlog.INFO:
		logx.Infow(content.Msg, fields...)
	case tlog.DEBUG:
		logx.Debugw(content.Msg, fields...)
	case tlog.ERROR, tlog.FATAL:
		logx.Errorw(content.Msg, fields...)
	case tlog.SEVERE:
		logx.Severe(content.Msg)
	case tlog.ALERT:
		logx.Alert(content.Msg)
	case tlog.STAT:
		logx.Stat(content.Msg)
	case tlog.SLOW:
		logx.Sloww(content.Msg, fields...)
	}
}

func (z *Tlog) buildFields(content *tlog.Content) []logx.LogField {
	fields := make([]logx.LogField, 0, 4)
	fields = append(fields, logx.LogField{
		Key:   tlogLabel,
		Value: content.Label,
	})
	fields = append(fields, logx.LogField{
		Key:   tlogTraceId,
		Value: content.TraceId,
	})
	fields = append(fields, logx.LogField{
		Key:   tlogRelatedId,
		Value: content.RelatedId,
	})

	fields = append(fields, logx.LogField{
		Key:   tlogPath,
		Value: content.Path,
	})
	return fields
}
