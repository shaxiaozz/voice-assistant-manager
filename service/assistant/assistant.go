package assistant

import (
	"encoding/json"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"os"
	"voice-assistant-manager/global"
	"voice-assistant-manager/utils/docker"
)

type Service interface {
	List() (data []*global.Assistant, err error)
	UpdateVoiceID(voiceID string) error
}

func NewService() Service {
	return &service{}
}

type service struct {
}

// 获取助手列表
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

// 修改助手语音ID
func (s *service) UpdateVoiceID(voiceID string) error {
	// 文件路径
	filePath := "voice-assistant.json"
	// 读取文件内容
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Error("读取文件失败: %v\n", err)
		return err
	}

	// 解析 JSON 数据
	var data []global.Assistant
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		logger.Error("解析 JSON 数据失败: %v\n", err)
		return err
	}

	for key, value := range data {
		// 修改语音ID
		data[key].Voice = voiceID

		// 容器名称
		containerName := "voice-pipeline-" + value.ID + "-agent-python"
		
		// 执行 Docker 命令更新环境变量
		command := []string{"python", "update_env_file.py", "CARTESIA_VOICE_ID", voiceID}
		if err := docker.ExecuteCommand(containerName, command...); err != nil {
			logger.Error("执行Docker命令失败: %v\n", err)
			return err
		}

		// 重启容器
		if err := docker.RestartContainer(containerName); err != nil {
			logger.Error("重启容器失败: %v\n", err)
			return err
		}
		
		logger.Info("容器 %s 更新并重启成功", containerName)
	}

	// 转换为 JSON 字符串
	dataJSON, err := json.Marshal(data)
	if err != nil {
		logger.Error("转换 JSON 字符串失败: %v\n", err)
		return err
	}

	// 重新写入文件
	if err := ioutil.WriteFile(filePath, dataJSON, 0644); err != nil {
		logger.Error("写入文件失败: %v\n", err)
		return err
	}

	logger.Info("成功更新语音ID并重启所有相关容器")
	return nil
}
