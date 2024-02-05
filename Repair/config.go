package repair

// 修复建议类型
type RepairType struct {
	// 漏洞名字
	Name string
	// 漏洞描述
	Describe string
	// 漏洞风险等级
	Level int
	// 漏洞类型
	Type []int
	// 风险分析
	RiskAnalysis string
	// 修复建议
	RepairSug []string
	// 参考链接
	Link []string
}

// 用户输入指令的解析
var (
	// 版本
	Version string = "1.0.0"
	// 要搜索的漏洞信息
	Search string
	// 输出日志
	outfile string
	// 日志等级
	debugLevel int
	// 搜索的模式
	mode string
	// 搜索的数据
	SearchData string
	// 显示漏洞修复名字
	show bool
)
