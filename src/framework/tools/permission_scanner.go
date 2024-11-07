package tools

import (
	"go/ast"
	"go/parser"
	"go/token"
	"matuto.com/GoPure/src/app/model"
	"os"
	"path/filepath"
	"strings"
)

// ScanPermissions 扫描项目中的权限注解
func ScanPermissions(rootPath string) []model.Menu {
	var menus []model.Menu

	// 遍历项目文件
	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		// 解析Go文件
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return nil
		}
		// 查找权限注解
		ast.Inspect(node, func(n ast.Node) bool {
			if fn, ok := n.(*ast.FuncDecl); ok {
				if fn.Doc != nil {
					for _, comment := range fn.Doc.List {
						if strings.Contains(comment.Text, "@RequirePermission") {
							menu := parsePermissionAnnotation(comment.Text)
							menus = append(menus, menu)
						}
					}
				}
			}
			return true
		})
		return nil
	})

	return menus
}

func parsePermissionAnnotation(text string) model.Menu {
	// 解析注解内容
	annotation := strings.TrimPrefix(text, "//")
	annotation = strings.TrimPrefix(annotation, "@")

	// 解析注解内容
	annotation = strings.TrimSpace(annotation)
	parts := strings.Split(annotation, " ")
	if len(parts) < 2 {
		return model.Menu{}
	}
	return model.Menu{
		Url:  parts[0],
		Name: strings.Join(parts[1:], " "),
	}
}
