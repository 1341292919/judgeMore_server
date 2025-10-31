package cache

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/redis/go-redis/v9"
	"judgeMore_server/app/user/domain/repository"
	"judgeMore_server/pkg/errno"
	"math/rand"
	"strings"
	"time"
)

type UserCache struct {
	client *redis.Client
}

func NewUserCache(client *redis.Client) repository.UserCache {
	return &UserCache{
		client: client,
	}
}

func (c *UserCache) IsKeyExist(ctx context.Context, key string) bool {
	return c.client.Exists(ctx, key).Val() == 1
}

func (c *UserCache) GetCodeCache(ctx context.Context, key string) (code string, err error) {
	value, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", errno.NewErrNo(errno.InternalRedisErrorCode, "write code to cache error:"+err.Error())
	}
	var storedCode string
	hlog.Info(value)
	hlog.Info(value)
	parts := strings.Split(value, "_")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid code format, expected 2 parts, got %d", len(parts))
	}
	storedCode = parts[0]
	return storedCode, nil
}
func (c *UserCache) PutCodeToCache(ctx context.Context, key string) (code string, err error) {
	code = generateRandomCode(6)
	timeNow := time.Now().Unix()
	value := fmt.Sprintf("%s_%d", code, timeNow)
	expiration := 10 * time.Minute
	err = c.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return "", errno.NewErrNo(errno.InternalRedisErrorCode, "write code to cache error:"+err.Error())
	}
	return code, nil
}

// 生成指定位数的随机验证码（字母+数字）
func generateRandomCode(length int) string {
	// 字符集：26个小写字母 + 26个大写字母 + 10个数字
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 初始化随机数生成器
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, length)
	for i := range code {
		code[i] = charSet[r.Intn(len(charSet))]
	}

	return string(code)
}
