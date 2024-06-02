package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"io"
	"os"
	"path"
)

// 定义颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct {
	//	自定义前缀
	prefix string
}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var leverColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		leverColor = gray
	case logrus.WarnLevel:
		leverColor = yellow
	case logrus.ErrorLevel:
		leverColor = red
	default:
		leverColor = blue
	}

	//	设置字节缓冲区
	//	用来返回字节切片
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	//设置日期格式
	formatTime := entry.Time.Format("2006-01-02 15:04:05")

	//自定义输出
	//	Caller 函数是?
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		//	自定义日志所在的函数 和行数
		//		path.Base 去除重复的根路径
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//	自定义输出格式
		//	fmt.Fprintf 将格式化后的字符串写入到指定的 io.Writer 接口中。
		fmt.Fprintf(b, "%s [%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", f.prefix, formatTime, leverColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s [%s] \x1b[%dm[%s]\x1b[0m %s\n", f.prefix, formatTime, leverColor, entry.Level, entry.Message)
	}

	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	return NewLog()
}

func NewLog() *logrus.Logger {

	// 设置log文件输出路径
	file, err := os.OpenFile(global.Config.Logger.Director, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}

	//同时输出到控制台和文件
	writers := []io.Writer{
		file,
		os.Stdout,
	}

	fileAndStdout := io.MultiWriter(writers...)

	// 设置各类东西
	mLog := logrus.New()                                                  //	新建一个实例
	logrus.SetFormatter(&logrus.TextFormatter{})                          //	设置输出类型
	mLog.SetOutput(fileAndStdout)                                         //	日志输出
	mLog.SetReportCaller(global.Config.Logger.ShowLine)                   //	是否开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{prefix: global.Config.Logger.Prefix}) //	设置自己定义的Formatter

	level, err := logrus.ParseLevel(global.Config.Logger.Lever)
	//	若出错 选择默认等级
	if err != nil {
		level = logrus.InfoLevel

	}
	mLog.SetLevel(level) //设置最低的Level
	return mLog
}
