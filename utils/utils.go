package utils

import (
	"os"
)

var (
	logger = GetLogger("Package-utils")
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

// 检查切片slice是否包含目标字符串target。如果包含则返回 true，否则返回 false
func ContainsString(slice []string, target string) bool {
    for _, v := range slice {
        if v == target {
            return true
        }
    }
    return false
}