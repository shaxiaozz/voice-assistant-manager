package cartesia

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"voice-assistant-manager/global"
)

// ListVoices 获取可用的语音列表
func ListVoices(apiKey string) ([]*global.CartesiaVoices, error) {
	url := "https://api.cartesia.ai/voices/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	req.Header.Add("Cartesia-Version", "2024-06-10")
	req.Header.Add("X-API-Key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	var voices []*global.CartesiaVoices
	if err := json.Unmarshal(body, &voices); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %v", err)
	}

	return voices, nil
}
