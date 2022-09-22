/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-04-25 16:54:18
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-07-16 19:14:30
 */
package service

import (
	"sync"
	"time"

	cache "github.com/patrickmn/go-cache"
)

// GetCacheFunc 获取缓存委托
type GetCacheFunc func(cacheKey string) []byte

// SetCacheFunc 设置缓存委托
type SetCacheFunc func(cacheKey string, value []byte, expiration time.Duration) bool

var CacheObj *cache.Cache

//初始化缓存配置 只能初始化一次
func InitCache() {
	var initOnce sync.Once
	initOnce.Do(func() {
		CacheObj = cache.New(30*time.Second, 10*time.Second)
	})
	LOG.Info("缓存服务初始化成功")
}

//设置缓存
func SetCacheDefault(cacheKey string, val interface{}, d time.Duration) {
	if d > 0 {
		CacheObj.Set(cacheKey, val, d)
	}
}

//获取缓存
func GetCache(key string) (value interface{}, found bool) {
	value, found = CacheObj.Get(key)
	return
}

//删除缓存
func DelCache(key string) {
	CacheObj.Delete(key)
}
