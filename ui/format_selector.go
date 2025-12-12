package ui

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"re-asmr-spider/i18n"
	"re-asmr-spider/spider"
	"re-asmr-spider/utils"
)

// getExtensionsKey 获取格式组合的唯一标识（排序后的格式列表）
func getExtensionsKey(extensions []string) string {
	sorted := make([]string, len(extensions))
	copy(sorted, extensions)
	sort.Strings(sorted)
	return strings.Join(sorted, "|")
}

// ShowFormatConflictPrompt 逐个显示格式冲突并获取用户选择
// saveCallback: 每次选择后的保存回调函数
func ShowFormatConflictPrompt(analysis *spider.FormatAnalysis, saveCallback func(*spider.FilterStrategy)) *spider.FilterStrategy {
	utils.Warning(i18n.T("format_conflict_detected", len(analysis.ConflictGroups)))
	fmt.Println()

	strategy := &spider.FilterStrategy{
		Mode:             "manual",
		ManualSelections: make(map[string][]string),
	}

	// 按格式组合分组冲突
	typeGroups := make(map[string][]*spider.FormatGroup)
	for _, group := range analysis.ConflictGroups {
		key := getExtensionsKey(group.Extensions)
		typeGroups[key] = append(typeGroups[key], group)
	}

	// 记录每种格式组合的批量选择
	batchSelections := make(map[string][]string)
	downloadAllForType := make(map[string]bool)

	// 获取所有冲突组并排序
	conflictKeys := make([]string, 0, len(analysis.ConflictGroups))
	for key := range analysis.ConflictGroups {
		conflictKeys = append(conflictKeys, key)
	}
	sort.Strings(conflictKeys)

	// 逐个处理冲突
	for _, key := range conflictKeys {
		group := analysis.ConflictGroups[key]
		extKey := getExtensionsKey(group.Extensions)

		// 检查是否该类型已设置为下载所有格式
		if downloadAllForType[extKey] {
			strategy.ManualSelections[group.BaseName] = group.Extensions
			utils.Success(i18n.T("format_batch_applied", group.BaseName, strings.Join(group.Extensions, ", ")))
			// 立即保存
			if saveCallback != nil {
				saveCallback(strategy)
			}
			continue
		}

		// 检查是否已有该类型的批量选择
		if batchChoice, exists := batchSelections[extKey]; exists {
			// 直接应用批量选择
			strategy.ManualSelections[group.BaseName] = batchChoice
			utils.Success(i18n.T("format_batch_applied", group.BaseName, strings.Join(batchChoice, ", ")))
			// 立即保存
			if saveCallback != nil {
				saveCallback(strategy)
			}
		} else {
			// 询问用户（可能设置批量选择）
			remainingCount := countRemainingInGroup(typeGroups[extKey], strategy.ManualSelections)
			selectedExts, selectionType := promptForFileGroupWithBatch(group, remainingCount)

			// 根据选择类型处理
			switch selectionType {
			case "all-global":
				// c选项：全部文件都下载所有格式
				utils.Info(i18n.T("format_all_global_enabled"))
				// 对所有未处理的冲突都设置为下载所有格式
				for _, k := range conflictKeys {
					g := analysis.ConflictGroups[k]
					strategy.ManualSelections[g.BaseName] = g.Extensions
				}
				// 保存策略
				if saveCallback != nil {
					saveCallback(strategy)
				}
				return strategy

			case "all-remaining":
				// b选项：余下的同类型都下载所有格式
				strategy.ManualSelections[group.BaseName] = selectedExts
				downloadAllForType[extKey] = true
				utils.Info(i18n.T("format_all_remaining_enabled", remainingCount-1))
				// 立即保存当前策略
				if saveCallback != nil {
					saveCallback(strategy)
				}

			case "batch":
				// 批量选择逻辑
				strategy.ManualSelections[group.BaseName] = selectedExts
				batchSelections[extKey] = selectedExts
				utils.Info(i18n.T("format_batch_enabled", strings.Join(selectedExts, ", "), remainingCount-1))
				// 立即保存当前策略
				if saveCallback != nil {
					saveCallback(strategy)
				}

			default:
				// normal 或其他情况
				strategy.ManualSelections[group.BaseName] = selectedExts
				// 立即保存当前策略
				if saveCallback != nil {
					saveCallback(strategy)
				}
			}
		}
	}

	return strategy
}

// countRemainingInGroup 计算该格式组合中还有多少未处理的冲突
func countRemainingInGroup(groups []*spider.FormatGroup, processed map[string][]string) int {
	count := 0
	for _, g := range groups {
		if _, exists := processed[g.BaseName]; !exists {
			count++
		}
	}
	return count
}

// promptForFileGroupWithBatch 为单个文件组提示用户选择（支持批量应用）
// 返回: (选中的扩展名列表, 选择类型)
// 选择类型: "normal" - 普通选择, "batch" - 批量应用, "all-remaining" - 余下的都下载所有, "all-global" - 全部都下载所有
func promptForFileGroupWithBatch(group *spider.FormatGroup, remainingCount int) ([]string, string) {
	// 显示文件信息
	utils.Info(i18n.T("format_conflict_prompt", group.BaseName))

	// 排序扩展名以保持一致性
	sortedExts := make([]string, len(group.Extensions))
	copy(sortedExts, group.Extensions)
	sort.Strings(sortedExts)

	// 显示选项
	fmt.Println()
	fmt.Println("  " + i18n.T("format_option_all"))
	optionNum := 1
	for _, ext := range sortedExts {
		fmt.Println("  " + i18n.T("format_option_number", optionNum, ext))
		optionNum++
	}

	// 如果还有剩余的同类型冲突，显示批量选项
	batchStartNum := optionNum
	if remainingCount > 1 {
		fmt.Println()
		// b选项：余下的都下载所有格式
		fmt.Println("  " + i18n.T("format_option_all_remaining"))

		// 数字批量选项：余下的都使用某个格式
		for i, ext := range sortedExts {
			fmt.Println("  " + i18n.T("format_option_batch", batchStartNum+i, ext))
		}
	}

	// c选项：全部文件都下载所有格式（始终显示）
	fmt.Println()
	fmt.Println("  " + i18n.T("format_option_all_global"))
	fmt.Println()

	// 读取用户输入
	choice := ReadInput(i18n.T("format_prompt_choice"))
	choice = strings.TrimSpace(strings.ToLower(choice))

	// 解析选择
	if choice == "a" || choice == "" {
		// 下载所有格式（仅当前文件）
		utils.Success(i18n.T("format_selection_confirmed", strings.Join(sortedExts, ", ")))
		return sortedExts, "normal"
	}

	// b选项：余下的都下载所有格式
	if choice == "b" && remainingCount > 1 {
		utils.Success(i18n.T("format_selection_confirmed", strings.Join(sortedExts, ", ")))
		return sortedExts, "all-remaining"
	}

	// c选项：全部文件都下载所有格式
	if choice == "c" {
		utils.Success(i18n.T("format_selection_confirmed", strings.Join(sortedExts, ", ")))
		return sortedExts, "all-global"
	}

	// 检查是否是批量选择数字
	choiceNum, err := strconv.Atoi(choice)
	if err == nil && remainingCount > 1 && choiceNum >= batchStartNum && choiceNum < batchStartNum+len(sortedExts) {
		// 批量选择某个格式
		selectedExt := sortedExts[choiceNum-batchStartNum]
		utils.Success(i18n.T("format_selection_confirmed", selectedExt))
		return []string{selectedExt}, "batch"
	}

	// 解析数字选择 (支持多选，用逗号或空格分隔)
	choice = strings.ReplaceAll(choice, ",", " ")
	parts := strings.Fields(choice)

	// 检查是否以 0 结尾（表示批量应用到余下的同类型冲突）
	applyToBatch := false
	if len(parts) > 0 && parts[len(parts)-1] == "0" {
		applyToBatch = true
		parts = parts[:len(parts)-1] // 移除尾部的 0
	}

	// 使用map去重，避免O(n²)复杂度
	selectedMap := make(map[string]bool)
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil && num >= 1 && num < batchStartNum {
			ext := sortedExts[num-1]
			selectedMap[ext] = true
		}
	}

	// 转换为切片
	selectedExts := make([]string, 0, len(selectedMap))
	for ext := range selectedMap {
		selectedExts = append(selectedExts, ext)
	}

	if len(selectedExts) == 0 {
		// 无效选择，下载所有
		utils.Warning(i18n.T("format_selection_invalid"))
		return sortedExts, "normal"
	}

	// 根据是否批量应用决定返回类型
	selectionType := "normal"
	if applyToBatch && remainingCount > 1 {
		selectionType = "batch"
		utils.Info(i18n.T("format_batch_enabled", strings.Join(selectedExts, ", "), remainingCount-1))
	} else {
		utils.Success(i18n.T("format_selection_confirmed", strings.Join(selectedExts, ", ")))
	}

	return selectedExts, selectionType
}

// BuildFormatSelectorCallback 构建格式选择回调函数
func BuildFormatSelectorCallback(saveCallback func(*spider.FilterStrategy)) func(*spider.FormatAnalysis) *spider.FilterStrategy {
	return func(analysis *spider.FormatAnalysis) *spider.FilterStrategy {
		return ShowFormatConflictPrompt(analysis, saveCallback)
	}
}
