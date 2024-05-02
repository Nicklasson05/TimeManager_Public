package main

import (
	"fmt"
	"os"
	"os/user"
)

func setup() {
	// Create directory on current user
	user, _ := user.Current()
	userDir := user.HomeDir
	userDir += "/TimeManagerSaves"
	err := os.Mkdir(userDir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
	}

	//log.txt
	file1, err := os.Create(userDir + "/log.txt")
	if err != nil {
		fmt.Println("Error creating .txt file:", err)
	}
	defer file1.Close()

	//markedDates.txt
	file2, err := os.Create(userDir + "/markedDates.txt")
	if err != nil {
		fmt.Println("Error creating .txt file:", err)
	}
	defer file2.Close()

	//response.txt
	file3, err := os.Create(userDir + "/response.txt")
	if err != nil {
		fmt.Println("Error creating .txt file:", err)
	}
	defer file3.Close()

	// Tasks.txt
	/*
		file6, err := os.Create(userDir + "/TaskID.txt")
		if err != nil {
			fmt.Println("Error creating .txt file:", err)
		}

		// useroptions.txt
		defer file6.Close()
	*/
	file7, err := os.Create(userDir + "/useroptions.txt")
	if err != nil {
		fmt.Println("Error creating .txt file:", err)
	}
	defer file7.Close()

	// Write to useroptions.txt
	writeToTasks("09:00", userDir+"/useroptions.txt")
	writeToTasks("Light", userDir+"/useroptions.txt")
	writeToTasks("https://gitlab.com/api/v4", userDir+"/useroptions.txt")
	writeToTasks(userDir, userDir+"/useroptions.txt")

	// Carry the Dictionary path to start
	restart()
}

func SetupDone(dir, folder string) bool {
	fileInfo, err := os.Stat(dir)
	if err != nil {
		fmt.Println("Directory not found")
		return false
	}

	if !fileInfo.IsDir() {
		fmt.Println("fileInfo is not a directory")
		return false
	}

	folderDir := dir + string(os.PathSeparator) + folder
	_, err = os.Stat(folderDir)
	if err != nil {
		fmt.Println("file does not exist in directory")
		return false
	}
	return true
}
