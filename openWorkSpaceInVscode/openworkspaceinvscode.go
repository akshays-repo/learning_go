package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/manifoldco/promptui"
)

const ROOTWORKINPATH = "/home/brain/Desktop/Projects"

func main() {
	files, err := ioutil.ReadDir(ROOTWORKINPATH)
	if err != nil {
		log.Fatal(err)
	}

	folders := []string{}
	for _, f := range files {
		if f.IsDir() {
			folders = append(folders, f.Name())
		}
	}
	promtOption(folders)
}

func promtOption(folders []string) {
	prompt := promptui.Select{
		Label: "Select a folder to open on VSCODE",
		Items: folders,
		Size:  20,
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	openInVsCode(ROOTWORKINPATH + "/" + result)

}

func openInVsCode(path string) {
	cmd := exec.Command("/usr/bin/code", path)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	main()
}
