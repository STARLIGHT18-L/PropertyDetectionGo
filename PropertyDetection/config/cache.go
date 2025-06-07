package config

import (
	"sync"
	"time"
)

// Cache 结构体用于存储键值对及其过期时间
type Cache struct {
	items map[string]valueWithExpiry
	mu    sync.RWMutex
}

// valueWithExpiry 结构体用于存储值和过期时间
type valueWithExpiry struct {
	value      interface{}
	expireTime int64
}

// InitCache 创建一个新的缓存实例
func InitCache() *Cache {
	return &Cache{
		items: make(map[string]valueWithExpiry),
	}
}

// SetValue 向缓存中存储一个键值对
func (c *Cache) SetValue(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = valueWithExpiry{
		value:      value,
		expireTime: -1,
	}
}

// GetValue 从缓存中获取指定键的值
func (c *Cache) GetValue(key string) interface{} {
	c.mu.RLock()
	item, exists := c.items[key]
	// 如果键不存在，直接释放读锁并返回 nil
	if !exists {
		c.mu.RUnlock()
		return nil
	}
	// 检查是否过期
	if c.isExpired(item) {
		// 过期了，升级为写锁进行删除操作
		c.mu.RUnlock()
		c.mu.Lock()
		defer c.mu.Unlock()
		// 双重检查，避免在升级锁的过程中其他协程已经删除了该键
		if item, ok := c.items[key]; ok && c.isExpired(item) {
			delete(c.items, key)
		}
		return nil
	}
	c.mu.RUnlock()
	return item.value
}

// GetValueOrDefault 从缓存中获取指定键的值，如果不存在或已过期则返回默认值
func (c *Cache) GetValueOrDefault(key string, defaultValue interface{}) interface{} {
	value := c.GetValue(key)
	if value == nil {
		return defaultValue
	}
	return value
}

// SetValueExpiry 向缓存中存储一个键值对，并设置其过期时间
func (c *Cache) SetValueExpiry(key string, value interface{}, timeout time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	expireTime := time.Now().Add(timeout).UnixNano()
	c.items[key] = valueWithExpiry{
		value:      value,
		expireTime: expireTime,
	}
}

// DeleteValue 从缓存中删除指定键及其对应的值
func (c *Cache) DeleteValue(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

// HasKey 检查缓存中是否存在指定的键
func (c *Cache) HasKey(key string) bool {
	c.mu.RLock()
	item, exists := c.items[key]
	// 如果键不存在，直接释放读锁并返回 false
	if !exists {
		c.mu.RUnlock()
		return false
	}
	// 检查是否过期
	if c.isExpired(item) {
		// 过期了，升级为写锁进行删除操作
		c.mu.RUnlock()
		c.mu.Lock()
		defer c.mu.Unlock()
		// 双重检查，避免在升级锁的过程中其他协程已经删除了该键
		if item, ok := c.items[key]; ok && c.isExpired(item) {
			delete(c.items, key)
		}
		return false
	}
	c.mu.RUnlock()
	return true
}

// GetExpire 获取指定键的剩余过期时间
func (c *Cache) GetExpire(key string) time.Duration {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, exists := c.items[key]
	if exists && item.expireTime != -1 {
		return time.Duration(item.expireTime-time.Now().UnixNano()) * time.Nanosecond
	}
	return -1
}

// Expire 为已经存在的键设置新的过期时间
func (c *Cache) Expire(key string, timeout time.Duration) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, exists := c.items[key]
	if exists && !c.isExpired(item) {
		expireTime := time.Now().Add(timeout).UnixNano()
		c.items[key] = valueWithExpiry{
			value:      item.value,
			expireTime: expireTime,
		}
		return true
	}
	return false
}

// isExpired 检查值是否已过期
func (c *Cache) isExpired(item valueWithExpiry) bool {
	return item.expireTime != -1 && time.Now().UnixNano() > item.expireTime
}
