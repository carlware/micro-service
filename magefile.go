// +build mage

package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	Default    = Build
	goFiles    = getGoFiles()
	goSrcFiles = getGoSrcFiles()
	// settings
	appName   = "carlware/accounts"
	binName   = "app"
	grpcPath  = "proto"
	mockPaths = []string{
		"/internal/interfaces",
	}
)

var curDir = func() string {
	name, _ := os.Getwd()
	return name
}()

// Calculate file paths
var toolsBinDir = normalizePath(path.Join(curDir, "tools", "bin"))

func init() {
	// TODO: eerrror
	time.Local = time.UTC

	// Add local bin in PATH
	err := os.Setenv("PATH", fmt.Sprintf("%s:%s", toolsBinDir, os.Getenv("PATH")))
	if err != nil {
		panic(err)
	}
}

func Build() {
	banner := figure.NewFigure("ACCOUNTS", "", true)
	banner.Print()

	fmt.Println("")
	color.Red("# Build Info ---------------------------------------------------------------")
	fmt.Printf("Go version : %s\n", runtime.Version())
	fmt.Printf("Git revision : %s\n", hash())
	fmt.Printf("Git branch : %s\n", branch())
	fmt.Printf("Tag : %s\n", tag())

	fmt.Println("")

	color.Red("# Core packages ------------------------------------------------------------")
	// mg.SerialDeps(Go.Deps, Go.Generate, Go.Lint, Test.All)

	fmt.Println("")
	color.Red("# Artifacts ----------------------------------------------------------------")
	mg.Deps(Bin.App)
}

func Generate() {
	color.Cyan("## Generate ")
	mg.SerialDeps(Go.Generate)
}

func Lint() {
	color.Cyan("## Execute linter")
	mg.SerialDeps(Go.Lint)
}

func Tidy() {
	color.Cyan("## Execute tidy")
	mg.SerialDeps(Go.Tidy)
}

func Format() {
	color.Cyan("## Execute format")
	mg.SerialDeps(Go.Format)
}

func Ver() {
	fmt.Printf("%s\n", tag())
}

// -----------------------------------------------------------------------------

type Gen mg.Namespace

// Generate protobuf
func (Gen) Protobuf() error {
	color.Blue("### Protobuf")

	return sh.RunV("prototool", "all", "--fix", grpcPath)
}

// Generate mocks for tests
func (Gen) Mocks() {
	color.Blue("### Mocks")

	for _, path := range mockPaths {
		mustGoGenerate("Interfaces", appName+path)
	}
}

// -----------------------------------------------------------------------------

type Go mg.Namespace

// Generate go code
func (Go) Generate() error {
	color.Cyan("## Generate code")
	mg.SerialDeps(Gen.Mocks)
	return nil
}

// Tidy add/remove depenedencies.
func (Go) Tidy() error {
	fmt.Println("## Cleaning go modules")
	return sh.RunV("go", "mod", "tidy", "-v")
}

// Deps install dependency tools.
func (Go) Deps() error {
	color.Cyan("## Vendoring dependencies")
	return sh.RunV("go", "mod", "vendor")
}

// Lint run linter.
func (Go) Lint() error {
	mg.Deps(Go.Format)
	color.Cyan("## Lint go code")
	return sh.RunV("golangci-lint", "run", "--deadline=10m")
}

// Format runs gofmt on everything
func (Go) Format() error {
	color.Cyan("## Format everything")
	args := []string{"-s", "-w"}
	args = append(args, goFiles...)
	return sh.RunV("gofumpt", args...)
}

// -----------------------------------------------------------------------------

type Test mg.Namespace

// Test run go test
func (Test) Qa() error {
	return TestTag("qa")
}

func (Test) Unit() error {
	return TestTag("unit")
}

func (Test) Integration() error {
	return TestTag("integration")
}

func (Test) All() error {
	return TestTag("unit,integration,qa")
}

// -----------------------------------------------------------------------------

type Bin mg.Namespace

// Build microservice
func (Bin) App() error {
	return goBuild(appName+"/cli", binName)
}

// -----------------------------------------------------------------------------

func goBuild(packageName, out string) error {
	fmt.Printf(" > Building %s [%s]\n", out, packageName)

	varsSetByLinker := map[string]string{
		appName + "/internal/version.Version":   tag(),
		appName + "/internal/version.Revision":  hash(),
		appName + "/internal/version.BuildDate": time.Now().Format(time.RFC3339),
		appName + "/internal/version.GoVersion": runtime.Version(),
	}
	var linkerArgs string
	for name, value := range varsSetByLinker {
		linkerArgs += fmt.Sprintf(" -X %s=%s", name, value)
	}
	linkerArgs = fmt.Sprintf("-ldflags=\"-s -w %s\"", linkerArgs)

	return sh.Run("go", "build", "-tags", linkerArgs, "-o", fmt.Sprintf("bin/%s", out), packageName)
}

func getGoFiles() []string {
	var goFiles []string

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "vendor/") {
			return filepath.SkipDir
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		goFiles = append(goFiles, path)
		return nil
	})

	return goFiles
}

func getGoSrcFiles() []string {
	var goSrcFiles []string

	for _, path := range goFiles {
		if !strings.HasSuffix(path, "_test.go") {
			continue
		}

		goSrcFiles = append(goSrcFiles, path)
	}

	return goSrcFiles
}

// tag returns the git tag for the current branch or "" if none.
func tag() string {
	s, _ := sh.Output("git", "describe", "--tags", "--abbrev=0")
	return s
}

// hash returns the git hash for the current repo or "" if none.
func hash() string {
	hash, _ := sh.Output("git", "rev-parse", "--short", "HEAD")
	return hash
}

// branch returns the git branch for current repo
func branch() string {
	hash, _ := sh.Output("git", "rev-parse", "--abbrev-ref", "HEAD")
	return hash
}

func mustStr(r string, err error) string {
	if err != nil {
		panic(err)
	}
	return r
}

func mustGoGenerate(txt, name string) {
	fmt.Printf(" > %s [%s]\n", txt, name)
	err := sh.RunV("go", "generate", name)
	if err != nil {
		panic(err)
	}
}

// normalizePath turns a path into an absolute path and removes symlinks
func normalizePath(name string) string {
	absPath := mustStr(filepath.Abs(name))
	return absPath
}

func TestTag(tag string) error {
	color.Cyan("## Running tests")
	sh.Run("mkdir", "-p", "test-results/junit")
	return sh.RunV("gotestsum", "--no-summary=skipped", "--", "-short", "-race", "-cover", "-tags="+tag, "./...")
}
