package mod

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetParentPackageUsesNearestGoMod(t *testing.T) {
	t.Parallel()

	rootDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(rootDir, "go.mod"), []byte("module github.com/polpo666/pzero\n\ngo 1.24.3\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(rootDir, "go.work"), []byte("go 1.24.3\nuse .\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	exampleDir := filepath.Join(rootDir, "examples", "simpleapi")
	if err := os.MkdirAll(filepath.Join(exampleDir, "plugins", "helloworld"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(exampleDir, "go.mod"), []byte("module github.com/jzero-io/examples/simpleapi\n\ngo 1.24.3\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	pkg, err := GetParentPackage(filepath.Join(exampleDir, "plugins", "helloworld"))
	if err != nil {
		t.Fatal(err)
	}

	want := "github.com/jzero-io/examples/simpleapi/plugins/helloworld"
	if pkg != want {
		t.Fatalf("GetParentPackage() = %q, want %q", pkg, want)
	}
}
