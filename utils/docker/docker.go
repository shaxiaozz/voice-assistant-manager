package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/wonderivan/logger"
)

var dockerClient *client.Client

// InitDocker 初始化 Docker 客户端
func InitDocker() error {
	var err error
	// 连接到 Docker daemon
	dockerClient, err = client.NewClientWithOpts(
		client.FromEnv,
		client.WithHost("unix:///var/run/docker.sock"), // Docker socket 路径
	)
	if err != nil {
		return fmt.Errorf("初始化 Docker 客户端失败: %v", err)
	}

	// 测试连接
	ctx := context.Background()
	_, err = dockerClient.Ping(ctx)
	if err != nil {
		return fmt.Errorf("Docker 连接测试失败: %v", err)
	}

	logger.Info("Docker 客户端初始化成功")
	return nil
}

// ExecuteCommand 在指定容器中执行命令
func ExecuteCommand(containerName string, command ...string) error {
	if dockerClient == nil {
		return fmt.Errorf("Docker 客户端未初始化")
	}

	ctx := context.Background()

	// 创建执行配置
	execConfig := types.ExecConfig{
		Cmd:          command,
		AttachStdout: true,
		AttachStderr: true,
	}

	// 创建执行实例
	execResp, err := dockerClient.ContainerExecCreate(ctx, containerName, execConfig)
	if err != nil {
		return fmt.Errorf("创建执行实例失败: %v", err)
	}

	// 启动执行
	err = dockerClient.ContainerExecStart(ctx, execResp.ID, types.ExecStartCheck{})
	if err != nil {
		return fmt.Errorf("执行命令失败: %v", err)
	}

	// 获取执行结果
	execInspect, err := dockerClient.ContainerExecInspect(ctx, execResp.ID)
	if err != nil {
		return fmt.Errorf("获取执行结果失败: %v", err)
	}

	// 检查执行状态
	if execInspect.ExitCode != 0 {
		return fmt.Errorf("命令执行失败，退出码: %d", execInspect.ExitCode)
	}

	return nil
} 