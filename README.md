# IceLogrusEnhance

封装了一些用于增强/改善Logrus使用体验的小工具。

Encapsulates some widgets to enhance/improve the Logrus experience.

## Install 安装

```shell
go get -u github.com/Iceinu-Project/IceLogrusEnhance
```

## Usage 用法

### GORM logger Converter

将Logrus的logger实例转换为gorm可用的logger实例

Convert a logger instance from Logrus to a logger instance that can be used by gorm

```go
import "github.com/Iceinu-Project/IceLogrusEnhance"

func main() {
    gormLogger := icelogrus.NewGormLogrusLogger(logrus.New())
    gormLogger.LogMode(logger.Info)
    db, err := gorm.Open(sqlite.Open("iceinu.db"), &gorm.Config{
        Logger: gormLogger,
    })
}
```
