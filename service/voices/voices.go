package voices

import (
	"time"
	"voice-assistant-manager/global"
	"voice-assistant-manager/utils/cache"
	"voice-assistant-manager/utils/cartesia"
)

const (
	CartesiaVoicesCacheKey = "cartesia:voices"
	CacheExpiration        = time.Hour * 2 // 缓存时间设置为2小时
)

type Service interface {
	ListCartesiaVoices() ([]*global.CartesiaVoices, error)
	StartVoiceUpdateTask()
}

type service struct{}

func NewService() Service {
	s := &service{}
	// 启动服务时开始定时任务
	go s.StartVoiceUpdateTask()
	return s
}

func (s *service) ListCartesiaVoices() ([]*global.CartesiaVoices, error) {
	// 先从缓存获取
	if data, exists := cache.Get(CartesiaVoicesCacheKey); exists {
		return data.([]*global.CartesiaVoices), nil
	}

	// 缓存不存在则从API获取
	return s.updateVoicesCache()
}

func (s *service) updateVoicesCache() ([]*global.CartesiaVoices, error) {
	voices, err := cartesia.ListVoices(global.CartesiaApiKey)
	if err != nil {
		return nil, err
	}

	// 更新缓存
	cache.Set(CartesiaVoicesCacheKey, voices, CacheExpiration)
	return voices, nil
}

func (s *service) StartVoiceUpdateTask() {
	ticker := time.NewTicker(time.Hour) // 每小时更新一次
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			_, err := s.updateVoicesCache()
			if err != nil {
				// 这里可以添加日志记录
				continue
			}
		}
	}
}
