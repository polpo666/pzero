package mod

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/execx"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"golang.org/x/mod/modfile"

	"github.com/polpo666/pzero/cmd/pzero/internal/config"
	"github.com/polpo666/pzero/cmd/pzero/internal/pkg/console"
)

// ModuleStruct contains the relative data of go module,
// which is the result of the command go list
type ModuleStruct struct {
	Path      string
	Dir       string
	GoVersion string
}

// GetParentPackage if is submodule project, root package is based on go.mod and add its dir
func GetParentPackage(workDir string) (string, error) {
	mod, err := GetGoMod(workDir)
	if err != nil {
		return "", err
	}
	trim := strings.TrimPrefix(workDir, mod.Dir)
	return filepath.ToSlash(filepath.Join(mod.Path, trim)), nil
}

func GetGoVersion() (string, error) {
	resp, err := execx.Run("go env GOVERSION", "")
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(resp, "go"), nil
}

func getNearestGoModDir(workDir string) (string, bool, error) {
	abs, err := filepath.Abs(workDir)
	if err != nil {
		return "", false, err
	}

	dir := abs
	for {
		if pathx.FileExists(filepath.Join(dir, "go.mod")) {
			return dir, true, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", false, nil
}

func parseGoMod(dir string) (*ModuleStruct, error) {
	goModBytes, err := os.ReadFile(filepath.Join(dir, "go.mod"))
	if err != nil {
		return nil, err
	}

	mod, err := modfile.Parse("", goModBytes, nil)
	if err != nil {
		return nil, err
	}

	abs, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	return &ModuleStruct{
		Path:      mod.Module.Mod.Path,
		Dir:       abs,
		GoVersion: mod.Go.Version,
	}, nil
}

// GetGoMod is used to determine whether workDir is a go module project through command `go list -json -m`
func GetGoMod(workDir string) (*ModuleStruct, error) {
	// 优先使用最近的上层 go.mod，避免外层 go.work 把生成项目解析成宿主仓库模块。
	if modDir, ok, err := getNearestGoModDir(workDir); err != nil {
		return nil, err
	} else if ok {
		return parseGoMod(modDir)
	}

	// 通过 go list -json -m 获取
	ms, err := GetGoMods(workDir)
	if err != nil {
		return nil, err
	}

	if len(ms) == 0 {
		return nil, errors.New("not go module project")
	}

	// mono project
	for _, m := range ms {
		if filepath.Clean(workDir) == filepath.Clean(m.Dir) {
			return &m, nil
		}
	}

	// unknown
	return &ms[0], nil
}

func GetGoMods(workDir string) ([]ModuleStruct, error) {
	command := exec.Command("go", "list", "-json", "-m")
	command.Dir = workDir
	data, err := command.CombinedOutput()
	if err != nil {
		if strings.Contains(string(data), "go mod tidy") {
			if !config.C.Quiet {
				fmt.Printf("%s go mod tidy. Please wait...\n", console.Green("Running"))
			}
			if _, err = execx.Run("go mod tidy", workDir); err != nil {
				return nil, err
			}
			command = exec.Command("go", "list", "-json", "-m")
			command.Dir = workDir
			if data, err = command.CombinedOutput(); err != nil {
				return nil, errors.New(string(data))
			}
		} else {
			return nil, errors.New(string(data))
		}
	}

	var ms []ModuleStruct
	decoder := json.NewDecoder(bytes.NewReader(data))
	for {
		var m ModuleStruct
		err = decoder.Decode(&m)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		}
		ms = append(ms, m)
	}
	return ms, nil
}
