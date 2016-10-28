package presentj

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const working = false

func must(s string, err error) string {
	if err != nil {
		panic(err)
	}
	return s
}

func JavaExec(src string) {
	class := className(src)
	if len(class) == 0 {
		log.Fatalf("invalid Java source")
	}
	workspace, err := ioutil.TempDir("", "presentjava")
	if err != nil {
		log.Fatalf("failed to create workspace: %v", err) //create temporary workspace
	}
	defer os.RemoveAll(workspace)                          // delete workspace on exit
	err = os.Mkdir(filepath.Join(workspace, "main"), 0777) // create workspace dir structure
	if err != nil && !os.IsExist(err) {
		log.Fatalf("failed to create workspace dir structure: %v", err)
		return
	}
	err = ioutil.WriteFile(filepath.Join(workspace, "main", class+".java"), []byte(src), 0666)
	if err != nil {
		log.Fatalf("failed to write java source: %v", err)
		return
	}

	// compile java source
	err = (&exec.Cmd{
		Path:   must(exec.LookPath("javac")), // lookup javac binary
		Args:   []string{"javac", filepath.Join("main", class+".java")},
		Dir:    workspace,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}).Run()
	if err != nil {
		log.Fatalf("failed to compile java source: %v", err)
		return
	}

	// run java snippet
	err = (&exec.Cmd{
		Path:   must(exec.LookPath("java")), // lookup java binary
		Args:   []string{"java", "main." + class},
		Dir:    workspace,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}).Run()
	if err != nil {
		if working {
			log.Fatalf("failed to run java snippet: %v", err)
		} else {
			log.Printf("failed to run java snippet: %v", err)
			return
		}
	}
}

func MavenExec(src, pomFile string) {
	maven(src, pomFile, false)
}

func MavenTest(src, pomFile string) {
	maven(src, pomFile, true)
}

func className(src string) string {
	fields := strings.Fields(src)
	var class string
	for i, s := range fields {
		if s == "class" {
			class = fields[i+1]
			break
		}
	}
	return class
}

func maven(src, pomFile string, test bool) {
	class := className(src)
	if len(class) == 0 {
		log.Fatalf("invalid Java source")
	}
	b, err := ioutil.ReadFile(pomFile) // read maven pom file
	if err != nil {
		log.Fatalf("failed to read pom file: %v", err)
	}

	pom := string(b)                                    // convert content to a 'string'
	workspace, err := ioutil.TempDir("", "presentjava") // create temporary workspace for maven project
	if err != nil {
		log.Fatalf("failed to create workspace: %v", err)
	}
	defer os.RemoveAll(workspace) // delete workspace on exit
	folder := "main"
	if test {
		folder = "test"
	}
	err = os.MkdirAll(filepath.Join(workspace, "src", folder, "java", folder), 0777) // create workspace dir structure
	if err != nil && !os.IsExist(err) {
		log.Fatalf("failed to create workspace dir structure: %v", err)
		return
	}
	err = ioutil.WriteFile(filepath.Join(workspace, "pom.xml"), []byte(pom), 0666) // write pom file
	if err != nil {
		log.Fatalf("failed to write pom file: %v", err)
		return
	}
	err = ioutil.WriteFile(filepath.Join(workspace, "src", folder, "java", folder, class+".java"), []byte(src), 0666) // write java source
	if err != nil {
		log.Fatalf("failed to write java source: %v", err)
		return
	}

	if test {
		// run unit test
		err = (&exec.Cmd{
			Path:   must(exec.LookPath("mvn")), // lookup maven binary
			Args:   []string{"mvn", "-q", "-Dtest=" + class, "test"},
			Dir:    workspace,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}).Run()
		if err != nil {
			log.Printf("failed to run unit test: %v", err)
			return
		}
		return
	}

	// compile maven project
	err = (&exec.Cmd{
		Path:   must(exec.LookPath("mvn")), // lookup maven binary
		Args:   []string{"mvn", "-q", "compile"},
		Dir:    workspace,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}).Run()
	if err != nil {
		log.Fatalf("failed to compile maven project: %v", err)
		return
	}

	// run maven project
	err = (&exec.Cmd{
		Path:   must(exec.LookPath("mvn")), // lookup maven binary
		Args:   []string{"mvn", "-q", "exec:java", "-Dexec.mainClass=main." + class},
		Dir:    workspace,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}).Run()
	if err != nil {
		if working {
			log.Fatalf("failed to run java snippet: %v", err)
		} else {
			log.Printf("failed to run java snippet: %v", err)
			return
		}
	}

}
