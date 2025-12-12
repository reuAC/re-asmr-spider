package spider

import (
	"path/filepath"
	"strings"
)

// FileInfo 文件信息
type FileInfo struct {
	Track     *track
	Path      string
	BaseName  string // 不含扩展名的文件名
	Extension string // 扩展名（含点号，如 .mp3）
}

// FormatGroup 同名不同格式的文件组
type FormatGroup struct {
	BaseName   string
	Extensions []string     // 该组中的所有扩展名
	Files      []*FileInfo  // 该组中的所有文件
}

// FormatAnalysis 格式分析结果
type FormatAnalysis struct {
	AllFiles       []*FileInfo              // 所有文件
	ConflictGroups map[string]*FormatGroup  // 有冲突的文件组 (key: baseName)
	AllExtensions  map[string]int           // 所有扩展名及其出现次数
}

// AnalyzeFormats 分析所有文件的格式
func AnalyzeFormats(tracks []track, basePath string) *FormatAnalysis {
	analysis := &FormatAnalysis{
		AllFiles:       make([]*FileInfo, 0),
		ConflictGroups: make(map[string]*FormatGroup),
		AllExtensions:  make(map[string]int),
	}

	collectFiles(tracks, basePath, analysis)
	findConflicts(analysis)

	return analysis
}

// collectFiles 递归收集所有文件信息
func collectFiles(tracks []track, currentPath string, analysis *FormatAnalysis) {
	for i := range tracks {
		t := &tracks[i]
		if t.Type != "folder" {
			// 提取文件名和扩展名
			ext := strings.ToLower(filepath.Ext(t.Title))
			baseName := strings.TrimSuffix(t.Title, ext)

			fileInfo := &FileInfo{
				Track:     t,
				Path:      currentPath,
				BaseName:  baseName,
				Extension: ext,
			}

			analysis.AllFiles = append(analysis.AllFiles, fileInfo)
			analysis.AllExtensions[ext]++
		} else {
			// 递归处理文件夹
			collectFiles(t.Children, filepath.Join(currentPath, t.Title), analysis)
		}
	}
}

// findConflicts 查找同名不同格式的文件 (全局检测，不区分目录)
func findConflicts(analysis *FormatAnalysis) {
	// 按 baseName 全局分组 (不包含路径)
	groups := make(map[string]*FormatGroup)

	for _, file := range analysis.AllFiles {
		key := file.BaseName // 只用文件名，不包含路径

		if group, exists := groups[key]; exists {
			// 检查是否已有此扩展名
			hasExt := false
			for _, ext := range group.Extensions {
				if ext == file.Extension {
					hasExt = true
					break
				}
			}
			if !hasExt {
				group.Extensions = append(group.Extensions, file.Extension)
			}
			group.Files = append(group.Files, file)
		} else {
			groups[key] = &FormatGroup{
				BaseName:   file.BaseName,
				Extensions: []string{file.Extension},
				Files:      []*FileInfo{file},
			}
		}
	}

	// 只保留有冲突的组（多于1个扩展名）
	for key, group := range groups {
		if len(group.Extensions) > 1 {
			analysis.ConflictGroups[key] = group
		}
	}
}

// FilterStrategy 过滤策略
type FilterStrategy struct {
	Mode              string              // "all", "priority", "manual"
	PriorityFormats   []string            // 优先格式列表（优先级从高到低）
	ManualSelections  map[string][]string // 手动选择: baseName -> 选中的扩展名列表
	IncludeFormats    []string            // 额外包含的扩展名（冲突解决后追加下载）
}

// ApplyFilter 应用过滤策略
func (analysis *FormatAnalysis) ApplyFilter(strategy *FilterStrategy) []*FileInfo {
	if strategy.Mode == "all" || len(analysis.ConflictGroups) == 0 {
		return analysis.AllFiles
	}

	filtered := make([]*FileInfo, 0)
	processedFiles := make(map[*FileInfo]bool) // 记录已处理的文件

	for _, file := range analysis.AllFiles {
		// 检查是否已被处理
		if processedFiles[file] {
			continue
		}

		key := file.BaseName // 全局检测，只用文件名

		// 如果这个文件属于冲突组
		if group, hasConflict := analysis.ConflictGroups[key]; hasConflict {
			// 对于该文件名的所有实例，应用相同的策略
			selectedFiles := selectFiles(group, strategy)

			// 标记所有相关文件为已处理
			for _, f := range group.Files {
				processedFiles[f] = true
			}

			// 只添加被选中的文件
			filtered = append(filtered, selectedFiles...)
		} else {
			// 非冲突文件直接添加
			filtered = append(filtered, file)
			processedFiles[file] = true
		}
	}

	// 额外下载指定扩展名的文件（冲突解决后追加）
	if len(strategy.IncludeFormats) > 0 {
		// 创建扩展名快速查找map
		includeMap := make(map[string]bool)
		for _, ext := range strategy.IncludeFormats {
			includeMap[ext] = true
		}

		// 遍历所有文件，添加指定扩展名的文件
		for _, file := range analysis.AllFiles {
			// 检查扩展名是否匹配且未被添加过
			if includeMap[file.Extension] && !processedFiles[file] {
				filtered = append(filtered, file)
				processedFiles[file] = true
			}
		}
	}

	return filtered
}

// selectFiles 根据策略从文件组中选择文件
func selectFiles(group *FormatGroup, strategy *FilterStrategy) []*FileInfo {
	if strategy.Mode == "manual" {
		// 手动选择模式 - 根据用户的选择
		selectedExts, ok := strategy.ManualSelections[group.BaseName]
		if !ok || len(selectedExts) == 0 {
			// 如果没有选择记录，返回所有文件
			return group.Files
		}

		result := make([]*FileInfo, 0)
		for _, file := range group.Files {
			for _, ext := range selectedExts {
				if file.Extension == ext {
					result = append(result, file)
					break
				}
			}
		}
		return result

	} else if strategy.Mode == "priority" {
		// 按优先级选择 - 选择所有匹配该优先级的文件（跨目录）
		for _, priorityExt := range strategy.PriorityFormats {
			result := make([]*FileInfo, 0)
			for _, file := range group.Files {
				if file.Extension == priorityExt {
					result = append(result, file)
				}
			}
			if len(result) > 0 {
				return result // 返回所有该格式的文件
			}
		}
		// 如果没有匹配的优先格式，返回所有格式的文件
		return group.Files
	}

	return group.Files
}
