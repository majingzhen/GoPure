package router

import (
	"encoding/gob"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/routers"
	"matuto.com/GoPure/src/framework/middleware"
	"matuto.com/GoPure/src/global"
	"os"
	"path/filepath"
	"strings"
)

type Router struct{}

func (r *Router) InitRouter() *gin.Engine {
	e := gin.New()
	// 前端文件
	e.HTMLRender = LoadTemplateFiles("public/templates", ".html")
	e.Static("/static", "./public/static")
	// 使用中间件
	e.Use(middleware.GinLogger())
	e.Use(gin.Recovery())
	// 将LoginUserVO 类型 注册到gob中，允许在session中存储该类型
	gob.Register(view.LoginUserVO{})
	store := cookie.NewStore([]byte(global.Config.Session.Secret))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: global.Viper.GetInt("session.expire"),
	})
	e.Use(sessions.Sessions("GoPure", store))

	// 跨域处理
	// 使用Cors中间件处理跨域请求
	e.Use(middleware.Cors())
	{
		routers.InitRouter(e)
		//router.baseRouter.InitBaseRouter(api)
		//router.genRouter.InitGenRouter(api)
	}

	return e
}

func getFilelist(path string, stuffix string) (files []string) {
	// 遍历目录
	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		// 将模板后缀的文件放到列表
		if strings.HasSuffix(path, stuffix) {
			files = append(files, path)
		}
		return nil
	})
	return
}

// LoadTemplateFiles 加载模板
func LoadTemplateFiles(templateDir, stuffix string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	rd, _ := ioutil.ReadDir(templateDir)
	for _, fi := range rd {
		if fi.IsDir() {
			// 如果是目录
			for _, f := range getFilelist(filepath.Join(templateDir, fi.Name()), stuffix) {
				// 添加到模板的时候，去掉根路径，并确保使用正斜杠
				templatePath := filepath.ToSlash(f[len(templateDir)+1:])
				r.AddFromFiles(templatePath, f)
			}
		} else {
			if strings.HasSuffix(fi.Name(), stuffix) {
				// 如果在根目录底下的文件直接添加
				filePath := filepath.Join(templateDir, fi.Name())
				r.AddFromFiles(fi.Name(), filePath)
			}
		}
	}
	return r
}
