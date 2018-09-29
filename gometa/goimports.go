package gometa

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/gwaylib/goget/gometa/config"
)

type IOReadCloser struct {
	*strings.Reader
}

func (io *IOReadCloser) Close() error {
	// ignore
	return nil
}

type goImport struct {
	Prefix   string `json:"Prefix"`
	VCS      string `json:"VCS"`
	RepoRoot string `json:"RepoRoot"`
}

var (
	gogetLock    = sync.RWMutex{}
	gogetOptions = map[string]*regexp.Regexp{}
	gogetImports = map[string]goImport{}
)

func init() {
	gogetLock.Lock()
	defer gogetLock.Unlock()
	cfg, err := config.ReadDefault("./.goget")
	if err != nil {
		errStr := err.Error()
		// for linux and windows
		if !strings.Contains(errStr, "no such file or directory") && !strings.Contains(errStr, "The system cannot find the file specified") {
			log.Printf("%s,%s\n", "./.goget", err.Error())
			return
		}
		root := os.Getenv("PJ_ROOT")
		if len(root) == 0 {
			return
		}
		cfg, err = config.ReadDefault(root + "/.goget")
		if err != nil {
			errStr := err.Error()
			if !strings.Contains(errStr, "no such file or directory") && !strings.Contains(errStr, "The system cannot find the file specified") {
				log.Printf("%s, %s\n", root+"/.goget", err.Error())
			}
			return
		}

	}

	// 解析配置文件
	sessionKey := "DEFAULT"
	options, err := cfg.Options(sessionKey)
	if err != nil {
		panic(err)
	}
	for _, opt := range options {
		gogetOptions[opt] = regexp.MustCompile(opt)
		data, _ := cfg.String(sessionKey, opt)
		gImport := goImport{}
		if err := json.Unmarshal([]byte(data), &gImport); err != nil {
			panic(err)
		}
		gogetImports[opt] = gImport
	}
	return
}

// export goget function
func Local(importPath string) (body io.ReadCloser) {
	return goget(importPath)
}

func goget(importPath string) (body io.ReadCloser) {
	gogetLock.RLock()
	defer gogetLock.RUnlock()

	for key, opt := range gogetOptions {
		// 正则查找
		if !opt.MatchString(importPath) {
			continue
		}
		// 包信息数据
		gImport, _ := gogetImports[key]
		return &IOReadCloser{Reader: strings.NewReader(fmt.Sprintf(`<html><head><meta content='%s %s %s' name='go-import'></head></html>`, gImport.Prefix, gImport.VCS, gImport.RepoRoot))}
	}

	// 若不存在本地配置数据，默认将golang.org/x包强转至github/golang
	if strings.Contains(importPath, "golang.org/x/") {
		// 解析包
		paths := strings.Split(importPath, "/")
		if len(paths) > 2 {
			pkgName := strings.Join(paths[:3], "/")
			gitUrl := "https://github.com/golang/" + paths[2] + ".git"
			return &IOReadCloser{Reader: strings.NewReader(fmt.Sprintf(`<html><head><meta content='%s git %s' name='go-import'></head></html>`, pkgName, gitUrl))}
		}
	}
	return
}
