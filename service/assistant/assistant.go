package assistant

import (
	"encoding/json"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"os"
	"voice-assistant-manager/global"
)

type Service interface {
	List() (data []*global.Assistant, err error)
}

func NewService() Service {
	return &service{}
}

type service struct {
}

func (s *service) List() (data []*global.Assistant, err error) {
	// 文件路径
	filePath := "voice-assistant.json"

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 文件不存在，将 DefaultAssistant 写入文件
		err := ioutil.WriteFile(filePath, []byte(global.DefaultAssistant), 0644)
		if err != nil {
			logger.Error("写入文件失败: %v\n", err)
			return nil, err
		}
		logger.Info("文件不存在，已写入默认内容到 voice-assistant.json")
	}

	// 读取文件内容
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Error("读取文件失败: %v\n", err)
		return nil, err
	}

	// 解析 JSON 数据
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		logger.Error("解析 JSON 数据失败: %v\n", err)
		return nil, err
	}

	return data, nil
}
