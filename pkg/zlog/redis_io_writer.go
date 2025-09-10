/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package zlog

import (
	"fmt"
	"io"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type redisIoWriter struct {
	redisKey string
	redis    *redis.Redis
}

func NewRedisIoWriter(redisKey string, redisCfg redis.RedisConf) io.Writer {
	return &redisIoWriter{
		redisKey: redisKey,
		redis:    redis.MustNewRedis(redisCfg),
	}
}

func (r *redisIoWriter) Write(p []byte) (n int, err error) {
	// 写日志
	go r.redis.Rpush(r.redisKey, string(p))

	fmt.Println(string(p))

	return 0, nil
}
