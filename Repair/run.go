package repair

import (
	"flag"
	"fmt"
)

// 主函数
func Run() {
	// 1.解析命令行
	Flag()
	// 2.处理命令行信息
	// 初始化数据
	// 判断是否显示工具涵盖的漏洞名字
	DebugLog("判断是否显示工具涵盖的漏洞名字")
	if show {
		showVul()
		return
	}
	DebugLog("处理输入的数据参数")
	err := parseFlag()
	if err != nil {
		flag.Usage()
		return
	}
	// 3.执行操作
	DebugLog("判断使用的搜索模式")
	ruledata, err := getRulesData(Rules, SearchData)
	if err != nil {
		ErrorLog(err.Error())
		return
	}
	// 4.返回结果
	for _, rd := range ruledata {
		InfoLog("============= 漏洞名字：%v =============", rd.Name)
		PrintLog("漏洞描述：%v", rd.Describe)
		PrintLog("漏洞等级：%v", VulLevel[rd.Level])
		tptmp := ""
		for _, tp := range rd.Type {
			tptmp += VulType[tp] + " "
		}
		PrintLog("漏洞类型：%v", tptmp)
		PrintLog("风险分析：%v", rd.RiskAnalysis)
		PrintLog("修复建议：")
		for i, r := range rd.RepairSug {
			PrintLog("  %v.%v", i+1, r)
		}
		fmt.Println()
	}
	if len(ruledata) == 0 {
		ErrorLog("未发现您搜索的漏洞建议说明，可以在https://github.com/a7cc/repair上提个issues，祝各位师傅天天挖到0day，目前支持的漏洞如下：")
		showVul()
	}
}
