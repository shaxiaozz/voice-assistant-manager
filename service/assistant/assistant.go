package assistant

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
	"voice-assistant-manager/global"
	"voice-assistant-manager/utils/docker"

	"github.com/wonderivan/logger"
)

type containerOperation struct {
	containerName string
	operation     func() error
}

type Service interface {
	List() (data []*global.Assistant, err error)
	UpdateVoiceID(voiceID string) error
	UpdateVoiceSpeedEmotio(voiceSpeed, voiceEmotio string) error
}

func NewService() Service {
	s := &service{
		operationChan:  make(chan containerOperation, 100),
		containerLocks: make(map[string]*sync.Mutex),
	}

	// 启动 worker
	go s.processDockerOperations()

	return s
}

type service struct {
	operationChan  chan containerOperation
	containerLocks map[string]*sync.Mutex
	lockMutex      sync.RWMutex
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

	var wg sync.WaitGroup
	errChan := make(chan error, len(data))

	for key, value := range data {
		// 修改语音ID
		data[key].Voice = voiceID

		// 容器名称
		containerName := "voice-pipeline-" + value.ID + "-agent-python"

		wg.Add(1)
		go func(containerName string) {
			defer wg.Done()

			// 执行命令
			s.operationChan <- containerOperation{
				containerName: containerName,
				operation: func() error {
					command := []string{"python", "update_env_file.py", "CARTESIA_VOICE_ID", voiceID}
					if err := docker.ExecuteCommand(containerName, command...); err != nil {
						errChan <- err
						return err
					}

					// 重启容器
					if err := docker.RestartContainer(containerName); err != nil {
						errChan <- err
						return err
					}

					logger.Info("容器 %s 更新并重启成功", containerName)
					return nil
				},
			}
		}(containerName)
	}

	// 等待所有操作完成
	wg.Wait()
	close(errChan)

	// 检查是否有错误
	for err := range errChan {
		if err != nil {
			return err
		}
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

// 修改助手语音速度与情感
func (s *service) UpdateVoiceSpeedEmotio(voiceSpeed, voiceEmotio string) error {
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

	var wg sync.WaitGroup
	errChan := make(chan error, len(data))

	for key, value := range data {
		// 修改语音速度与情感
		data[key].VoiceSpeed = voiceSpeed
		data[key].VoiceEmotion = voiceEmotio

		// 容器名称
		containerName := "voice-pipeline-" + value.ID + "-agent-python"

		wg.Add(1)
		go func(containerName string) {
			defer wg.Done()

			// 执行命令
			s.operationChan <- containerOperation{
				containerName: containerName,
				operation: func() error {
					command1 := []string{"python", "update_env_file.py", "CARTESIA_SPEED", voiceSpeed}
					if err := docker.ExecuteCommand(containerName, command1...); err != nil {
						errChan <- err
						return err
					}

					command2 := []string{"python", "update_env_file.py", "CARTESIA_EMOTION", voiceEmotio}
					if err := docker.ExecuteCommand(containerName, command2...); err != nil {
						errChan <- err
						return err
					}

					// 重启容器
					if err := docker.RestartContainer(containerName); err != nil {
						errChan <- err
						return err
					}

					logger.Info("容器 %s 更新并重启成功", containerName)
					return nil
				},
			}
		}(containerName)
	}

	// 等待所有操作完成
	wg.Wait()
	close(errChan)

	// 检查是否有错误
	for err := range errChan {
		if err != nil {
			return err
		}
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

	logger.Info("成功更新语音速度和情感并重启所有相关容器")
	return nil
}

// 添加处理 Docker 操作的方法
func (s *service) processDockerOperations() {
	for op := range s.operationChan {
		go func(op containerOperation) {
			s.lockMutex.RLock()
			lock, exists := s.containerLocks[op.containerName]
			if !exists {
				s.lockMutex.RUnlock()
				s.lockMutex.Lock()
				// 双重检查
				if lock, exists = s.containerLocks[op.containerName]; !exists {
					lock = &sync.Mutex{}
					s.containerLocks[op.containerName] = lock
				}
				s.lockMutex.Unlock()
			} else {
				s.lockMutex.RUnlock()
			}

			lock.Lock()
			defer lock.Unlock()

			if err := op.operation(); err != nil {
				logger.Error("执行Docker操作失败 %s: %v", op.containerName, err)
			}
		}(op)
	}
}
