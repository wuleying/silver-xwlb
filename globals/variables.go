package globals

import (
	"github.com/wuleying/silver-xwlb/utils"
	"time"
)

// 全局变量
var (
	// 根目录
	RootDir = utils.FileGetCurrentDirectory()
	// 模板目录
	TemplateDir = RootDir + "/admin/web/template"
	// 当前时间
	CurrentTime = time.Now().Local()
)