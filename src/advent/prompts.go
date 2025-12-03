package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)


type PromptSolver struct {
	ops map[string]func()
	reader *bufio.Reader
}

func MakePromptSolver() *PromptSolver {
	ps := &PromptSolver{
		reader: bufio.NewReader(os.Stdin),
	}
	ps.ops = map[string]func() {
		"c": ps.createTemplatePrompt,
		"t": ps.testSolverPrompt,
		"?": ps.showMenuPrompt,
		"q": ps.exitPrompt,
	}
	ps.showMenuPrompt()
	return ps
}

func (ps *PromptSolver) Exec(op string) {
	if action, exists := ps.ops[op]; exists {
		action()
	} else {
		log.Printf("Operation %s not recognized\n", op)
	}
	fmt.Println()	
	fmt.Println("=+=+=+=+=+=+=+==+=+=+=+=+=+=+==+=+=+=+=+=+=+=+=+=+=+=")
	fmt.Print(">")
}

func (ps *PromptSolver) showMenuPrompt() {
	fmt.Println("=+=+=+=+=+=+=+==+=+=+=+=+=+=+==+=+=+=+=+=+=+=+=+=+=+=")
	fmt.Println("=+=+=+=+=+=+=+= Advent of Code solver =+=+=+=+=+=+=+=")
	fmt.Println("=+=+=+=+=+=+=+==+=+=+=+=+=+=+==+=+=+=+=+=+=+=+=+=+=+=")
	fmt.Println("Select an option:")
	fmt.Println("c: Create template files for year-day")
	fmt.Println("t: Test solver for year-day")
	fmt.Println("q: Quit")
	fmt.Println("?: Show this menu again")
	fmt.Println()
	fmt.Println("Enter option:")
	fmt.Print("> ")
}

func (ps *PromptSolver) createTemplatePrompt() {
	ans := getUserInput("Enter yyyy-dd: ")
	
	inputs := strings.Split(ans, "-")
	year, day := inputs[0], inputs[1]
	nday := day
	if nday[0] == '0' {
		nday = day[1:]
	}
	wd, err := os.Getwd()
	
	if err != nil {
		log.Fatal(err)
	}
	
	log.Printf("Creating template for %s\n", ans)
	
	destination_path := filepath.Join(wd, year, fmt.Sprintf("day%s", day))

	if err := os.MkdirAll(destination_path, os.ModeAppend); err != nil {
		log.Fatal("Error creating directories.\n",err)
	}
	log.Printf("Created directory: %s\n", destination_path)

	templatePath := filepath.Join(wd, "templates")
	templateDir, err := os.ReadDir(templatePath)
	if err != nil {
		log.Fatal("Error reading template directory.\n", err)
	}

	for _, entry := range templateDir {
		
		content, err := os.ReadFile(filepath.Join(templatePath, entry.Name()))
		if err != nil {
			log.Fatal("Error reading template file ", entry.Name(), ".\n", err)
		}
		template_content := string(content)
		template_content = strings.ReplaceAll(template_content, "$year$", year)
		template_content = strings.ReplaceAll(template_content, "$day$", nday)
		template_content = strings.ReplaceAll(template_content, "$go_version$", runtime.Version()[2:]) // remove "go" prefix
		template_content = strings.ReplaceAll(template_content, "$module_name$", fmt.Sprintf("github.com/ValdemarPCAntunes/src/advent/%s/day%s", year, day))
		
		file_name := entry.Name()
		// remove .txt extension and "template_" prefix
		file_name = file_name[len("template_"):len(file_name)-len(".txt")] 

		path := filepath.Join(destination_path, file_name)

		os.WriteFile(path, []byte(template_content), 0644)
		log.Println("File created at ", path)
	}

	addNewModuleToGoWork(year, day)

	log.Println("Template creation completed.")
	fmt.Println("Template creation completed.")
}

func addNewModuleToGoWork(year string, day string) {
	if _, filename, _, ok := runtime.Caller(0); !ok {
		log.Println("You must update go.work manually, the program was not able to retrive information about where the file is")
	} else {
		goWorkPath := filepath.Join(filepath.Dir(filename), "go.work")
		contentBytes, err := os.ReadFile(goWorkPath)
		if err != nil {
			log.Println("Error reading go.work file:", err)
		} else {
			content := string(contentBytes)
			newModulePath := fmt.Sprintf("\t./%s/day%s%s%s", year, day, GetLineEnding(), ")")
			content = strings.Replace(content, ")", newModulePath, 1)
			err = os.WriteFile(goWorkPath, []byte(content), 0644)
			if err != nil {
				log.Println("Error writing to go.work file:", err)
			} else {
				log.Println("go.work file updated successfully.")
			}
		}
	}
}


func (ps *PromptSolver) testSolverPrompt() {
	ans := getUserInput("Enter yyyy-dd: ")
	ans_parts := strings.Split(ans, "-")
	year, day := ans_parts[0], ans_parts[1]
	log.Printf("Testing solver for %s-%s\n", year, day)

	wd, _ := os.Getwd()
	if _, err := os.OpenFile(filepath.Join(wd,year, "day"+day, "solver.go"), os.O_RDWR|os.O_RDONLY|os.O_EXCL, 0666); err != nil {
		fmt.Printf("Directory for %s-day%s does not exist. Please create the template first.", year, day)
		return
	}
	
	cmd := exec.Command("go", "test", "./"+year+"/day"+day, "-v")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running tests: %v", err)
	}
}

func (ps *PromptSolver) exitPrompt() {
	os.Exit(0)
}

func getUserInput(question string) string {
	fmt.Print(question)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
