package utils

import (
	"os"
)

var (
	logger = GetLogger("utils")
	ProjectPath = getProjectPath()
)

func getProjectPath() string{
    // 获取当前工作目录路径
    projectPath, err := os.Getwd()
    if err != nil {
        logger.Warn("Error:", err)
        return "./"
    }
    logger.Info("Project path:", projectPath)
	return projectPath
}