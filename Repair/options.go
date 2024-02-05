package repair

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

// Tag
func tag() {
	fmt.Println(`                         `)
	fmt.Println(` _____             _     `)
	fmt.Println(`| __  |___ ___ ___|_|___ `)
	fmt.Println(`|    -| -_| . | .'| |  _|`)
	fmt.Println(`|__|__|___|  _|__,|_|_|  `)
	fmt.Println(`          |_|            `)
	fmt.Printf("                ver: %s%-8s%s\n", YELLOW, Version, END)
	fmt.Println()
}

// 处理flag
func Flag() {
	tag()
	flag.StringVar(&Search, "s", "", `搜索的漏洞信息，使用以下模式name(漏洞名字)\describe(漏洞描述)\level(漏洞等级)\type(漏洞类型)\riskanalysis(风险分析)\repairsug(修复建议)，例如：搜索漏洞名字为sql，规则为：name:sql`)
	flag.StringVar(&outfile, "outfile", "log.txt", "保存日志文件")
	flag.IntVar(&debugLevel, "debug", 0, "debug等级日志,0(Basic)/1(Error)/3(Warn)/4(Debug)")
	flag.BoolVar(&show, "show", false, "显示工具涵盖的漏洞名字，默认不显示(false)")
	flag.Parse()
}

// 处理数据
func parseFlag() error {
	// 判断是否输入搜索的内容
	DebugLog("处理输入的搜索内容")
	if Search == "" || !strings.Contains(Search, ":") {
		return errors.New("没有输入正确的搜索内容")
	}
	// 分割:左右的内容
	tmplist := strings.SplitN(Search, ":", 2)
	// 赋值
	mode, SearchData = tmplist[0], tmplist[1]
	return nil
}
