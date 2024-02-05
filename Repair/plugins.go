package repair

import (
	"errors"
	"fmt"
	"strings"
)

// 通过模式获取
func getRulesData(rules []RepairType, searchData string) (data []RepairType, err error) {
one:
	for _, rule := range rules {
		switch {
		case strings.HasPrefix("name", strings.ToLower(mode)):
			DebugLog("进入 name 规则模式")
			if strings.Contains(strings.ToLower(rule.Name), strings.ToLower(searchData)) {
				DebugLog("通过 name 规则模式，添加规则（%v）", rule.Name)
				data = append(data, rule)
			}
		case strings.HasPrefix("describe", strings.ToLower(mode)):
			DebugLog("进入 describe 规则模式")
			if strings.Contains(strings.ToLower(rule.Describe), strings.ToLower(searchData)) {
				DebugLog("通过 describe 规则模式，添加规则（%v）", rule.Name)
				data = append(data, rule)
			}
		case strings.HasPrefix("level", strings.ToLower(mode)):
			DebugLog("进入 level 规则模式")
			if strings.Contains(strings.ToLower(VulLevel[rule.Level]), strings.ToLower(searchData)) {
				DebugLog("通过 level 规则模式，添加规则（%v）", rule.Name)
				data = append(data, rule)
			}
		case strings.HasPrefix("type", strings.ToLower(mode)):
			DebugLog("进入 type 规则模式")
			for _, tp := range rule.Type {
				if strings.Contains(strings.ToLower(VulType[tp]), strings.ToLower(searchData)) {
					DebugLog("通过 type 规则模式，添加规则（%v）", rule.Name)
					data = append(data, rule)
					break
				}
			}
		case strings.HasPrefix("riskanalysis", strings.ToLower(mode)):
			DebugLog("进入 riskanalysis 规则模式")
			if strings.Contains(strings.ToLower(rule.RiskAnalysis), strings.ToLower(searchData)) {
				DebugLog("通过 riskanalysis 规则模式，添加规则（%v）", rule.Name)
				data = append(data, rule)
			}
		case strings.HasPrefix("repairsug", strings.ToLower(mode)):
			DebugLog("进入 repairsug 规则模式")
			for _, rs := range rule.RepairSug {
				if strings.Contains(strings.ToLower(rs), strings.ToLower(searchData)) {
					DebugLog("通过 repairsug 规则模式，添加规则（%v）", rule.Name)
					data = append(data, rule)
					break
				}
			}
		case strings.EqualFold(mode, "all"):
			DebugLog("进入 all 规则模式")
			tptmp := ""
			for _, tp := range rule.Type {
				tptmp += VulType[tp] + "=="
			}
			ruleData := fmt.Sprintf("%v==%v==%v==%v==%v==%v", rule.Name, rule.Describe, VulLevel[rule.Level], tptmp, rule.RiskAnalysis, rule.RepairSug)
			if strings.Contains(strings.ToLower(ruleData), strings.ToLower(searchData)) {
				DebugLog("通过 all 规则模式，添加规则（%v）", rule.Name)
				data = append(data, rule)
			}
		default:
			err = errors.New(`未发现存在该模式，可以使用以下模式：name(漏洞名字)\describe(漏洞描述)\level(漏洞等级)\type(漏洞类型)\riskanalysis(风险分析)\repairsug(修复建议)`)
			break one
		}
	}
	return
}
func showVul() {
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++工具涵盖的漏洞名字++++++++++++++++++++++++++++++++++++++++++++++")
	for i, sr := range Rules {
		fmt.Printf("|   %-.15s   ", sr.Name)
		if i%5 == 4 {
			fmt.Println(" |")
		}
	}
	fmt.Println("|")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}
