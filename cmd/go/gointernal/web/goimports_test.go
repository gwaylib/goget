package web

import (
	"io/ioutil"
	"strings"
	"testing"
)

// depend on .goget
var gogetCases = []struct {
	pkgName  string
	goImport goImport
}{
	{"golang.org/x/net", goImport{Prefix: "golang.org/x/net", VCS: "git", RepoRoot: "https://github.com/golang/net.git"}},
	{"golang.org/x/net/context", goImport{Prefix: "golang.org/x/net", VCS: "git", RepoRoot: "https://github.com/golang/net.git"}},
	{"git.ot24.net/go/engine/errors", goImport{Prefix: "git.ot24.net/go/engine", VCS: "git", RepoRoot: "https://git.ot24.net/go/engine.git"}},
}

func TestGoGet(t *testing.T) {
	// os.SetEnv("PJ_ROOT", "./")
	for _, c := range gogetCases {
		_, body := goget(c.pkgName)
		if body == nil {
			t.Fatal(c)
		}
		bodyData, err := ioutil.ReadAll(body)
		if err != nil {
			t.Fatal(err)
		}
		// TODO: parse html
		bodyStr := string(bodyData)
		if !strings.Contains(bodyStr, c.goImport.Prefix) {
			t.Fatal(bodyStr, c)
		}
		if !strings.Contains(bodyStr, c.goImport.VCS) {
			t.Fatal(bodyStr, c)
		}
		if !strings.Contains(bodyStr, c.goImport.RepoRoot) {
			t.Fatal(bodyStr, c)
		}
	}
}
