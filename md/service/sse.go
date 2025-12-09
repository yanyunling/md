package service

import (
	"encoding/json"
	"fmt"
	"md/model/common"
	"md/util"
	"sync"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/muesli/cache2go"
)

// 用户id对应sseIds缓存
var userIdToSSEIds = make(map[string][]string)

// SSE通道缓存
var sseChanMap = make(map[string]chan common.SSEMessage)

// SSE缓存锁
var mu sync.RWMutex

// SSE处理器
func SSEHandler(ctx iris.Context) {
	// 验证SSE临时token，此token作为sseId
	token := ctx.Params().Get("token")
	res, err := cache2go.Cache(common.SSETokenCache).Value(token)
	if err != nil {
		return
	}
	cache2go.Cache(common.SSETokenCache).Delete(token)
	userId := res.Data().(string)

	// 连接成功后发送一次心跳
	sendHeart(ctx)
	heartTicker := time.NewTicker(30 * time.Second)
	defer heartTicker.Stop()

	// 缓存连接
	c := cacheConnection(userId, token)

	// 清空连接缓存
	defer clearCacheConnection(userId, token, c)

	// 监听通道消息并推送
	for {
		select {
		case <-ctx.Request().Context().Done():
			return
		case <-heartTicker.C:
			// 心跳
			sendHeart(ctx)
		case message := <-c:
			// 消息
			messageBytes, err := json.Marshal(message)
			if err == nil {
				msg := fmt.Sprintf("data: %v\n\n", string(messageBytes))
				ctx.WriteString(msg)
				ctx.ResponseWriter().Flush()
			}
		}
	}
}

// 获取SSE临时token
func GetSSEToken(userId string) string {
	// 生成token
	token := util.RandomString(32)

	// 缓存token
	cache2go.Cache(common.SSETokenCache).Add(token, time.Minute, userId)

	return token
}

// 向用户发送SSE消息
func SendMessageToUser(userId string, message common.SSEMessage) {
	mu.RLock()
	defer mu.RUnlock()

	sseIds, ok := userIdToSSEIds[userId]
	if ok {
		for _, sseId := range sseIds {
			c, ok2 := sseChanMap[sseId]
			if ok2 {
				c <- message
			}
		}
	}
}

// 广播发送SSE消息
func SendMessageToAll(message common.SSEMessage) {
	mu.RLock()
	defer mu.RUnlock()

	for _, c := range sseChanMap {
		c <- message
	}
}

// 缓存连接
func cacheConnection(userId, sseId string) chan common.SSEMessage {
	mu.Lock()
	defer mu.Unlock()

	sseIds, ok := userIdToSSEIds[userId]
	if ok {
		userIdToSSEIds[userId] = append(sseIds, sseId)
	} else {
		userIdToSSEIds[userId] = []string{sseId}
	}

	c := make(chan common.SSEMessage)
	sseChanMap[sseId] = c

	return c
}

// 清空连接缓存
func clearCacheConnection(userId, sseId string, c chan common.SSEMessage) {
	mu.Lock()
	defer mu.Unlock()

	delete(sseChanMap, sseId)

	sseIds, ok := userIdToSSEIds[userId]
	if ok {
		userIdToSSEIds[userId] = util.RemoveAllMatches(sseIds, sseId)
	}

	close(c)
}

// 发送心跳
func sendHeart(ctx iris.Context) {
	messageBytes, err := json.Marshal(common.SSEMessage{Type: "heart"})
	if err == nil {
		msg := fmt.Sprintf("data: %v\n\n", string(messageBytes))
		ctx.WriteString(msg)
		ctx.ResponseWriter().Flush()
	}
}
