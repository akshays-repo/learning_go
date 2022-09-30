package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

const ROOTWORKINPATH = "/home/brain/Desktop/Projects"

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetDirectorys() ([]string, error) {
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
	return folders, err
}
func (a *App) OpenInVsCode(path string) {

	finalPath := ROOTWORKINPATH + "/" + path
	cmd := exec.Command("/usr/bin/code", finalPath)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
