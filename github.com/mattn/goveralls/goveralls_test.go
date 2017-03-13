package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pborman/uuid"
)

func TestUsage(t *testing.T) {
	tmp := prepareTest(t)
	defer os.RemoveAll(tmp)
	cmd := exec.Command("goveralls", "-h")
	b, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatal("Expected exit code 1 bot 0")
	}
	s := strings.Split(string(b), "\n")[0]
	if !strings.HasPrefix(s, "Usage: goveralls ") {
		t.Fatalf("Expected %v, but %v", "Usage: ", s)
	}
}

func TestInvalidArg(t *testing.T) {
	tmp := prepareTest(t)
	defer os.RemoveAll(tmp)
	cmd := exec.Command("goveralls", "pkg")
	b, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatal("Expected exit code 1 bot 0")
	}
	s := strings.Split(string(b), "\n")[0]
	if !strings.HasPrefix(s, "Usage: goveralls ") {
		t.Fatalf("Expected %v, but %v", "Usage: ", s)
	}
}

/* FIXME: currently this dones't work because the command goveralls will run
 * another session for this session.
func TestGoveralls(t *testing.T) {
	wd, _ := os.Getwd()
	tmp := prepareTest(t)
	os.Chdir(tmp)
	defer func() {
		os.Chdir(wd)
		os.RemoveAll(tmp)
	}()
	runCmd(t, "go", "get", "github.com/mattn/goveralls/testergo-runewidth")
	b := runCmd(t, "goveralls", "-package=github.com/mattn/goveralls/tester")
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	s := lines[len(lines)-1]
	if s != "Succeeded" {
		t.Fatalf("Expected test of tester are succeeded, but failured")
	}
}
*/

func prepareTest(t *testing.T) (tmpPath string) {
	tmp := os.TempDir()
	tmp = filepath.Join(tmp, uuid.New())
	runCmd(t, "go", "build", "-o", filepath.Join(tmp, "bin", "goveralls"), "github.com/mattn/goveralls")
	os.Setenv("PATH", filepath.Join(tmp, "bin")+string(filepath.ListSeparator)+os.Getenv("PATH"))
	os.MkdirAll(filepath.Join(tmp, "src"), 0755)
	return tmp
}

func runCmd(t *testing.T, cmd string, args ...string) []byte {
	b, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		t.Fatalf("Expected %v, but %v: %v", nil, err, string(b))
	}
	return b
}
