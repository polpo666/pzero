package addapi

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/api/format"

	"github.com/polpo666/pzero/cmd/pzero/internal/config"
	"github.com/polpo666/pzero/cmd/pzero/internal/embeded"
	"github.com/polpo666/pzero/cmd/pzero/internal/pkg/filex"
	"github.com/polpo666/pzero/cmd/pzero/internal/pkg/stringx"
	"github.com/polpo666/pzero/cmd/pzero/internal/pkg/templatex"
)

func Run(args []string) error {
	baseDir := filepath.Join("desc", "api")

	apiName := args[0]

	if strings.HasSuffix(apiName, ".api") {
		apiName = strings.TrimSuffix(apiName, ".api")
	}

	// fix https://github.com/polpo666/pzero/issues/405.
	// For jzero, each api file, the server name can be different.
	template, err := templatex.ParseTemplate(filepath.Join("api", "template.api.tpl"), map[string]any{
		"Service": stringx.ToCamel(apiName),
		"Group":   apiName,
	}, embeded.ReadTemplateFile(filepath.Join("api", "template.api.tpl")))
	if err != nil {
		return err
	}

	if config.C.Add.Output == "file" {
		if filex.FileExists(filepath.Join(baseDir, apiName+".api")) {
			return fmt.Errorf("%s already exists", apiName)
		}

		_ = os.MkdirAll(filepath.Dir(filepath.Join(baseDir, apiName)), 0o755)

		err = os.WriteFile(filepath.Join(baseDir, apiName+".api"), template, 0o644)
		if err != nil {
			return err
		}

		// format
		return format.ApiFormatByPath(filepath.Join(baseDir, apiName+".api"), false)
	}
	fmt.Println(string(template))
	return nil
}
