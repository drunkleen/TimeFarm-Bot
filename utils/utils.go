package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	Green  = color.New(color.FgGreen).SprintFunc()
	Cyan   = color.New(color.FgCyan).SprintFunc()
	Red    = color.New(color.FgRed).SprintFunc()
)

func PrintLogo() {
	fmt.Printf(yellow(
		"                    ┏┳┓•     ┏┓       \n" +
			"                     ┃ ┓┏┳┓┏┓┣ ┏┓┏┓┏┳┓\n" +
			"                     ┻ ┗┛┗┗┗ ┻ ┗┻┛ ┛┗┗\n"),
	)

}

func ClearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		return
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func createConfigDir() error {
	configDir := filepath.Join(".", "configs")

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.Mkdir(configDir, 0755)
		if err != nil {
			fmt.Println("Error creating configs folder:", err)
			return err
		}
		fmt.Println("Created configs folder:", configDir)
	} else if err != nil {
		fmt.Println("Error checking configs folder status:", err)
		return err
	}
	return nil
}

func loadListFile(fileName string) ([]string, error) {
	file, err := os.Open("./configs/" + fileName)
	if err != nil {
		return nil, createConfigDir()
	}
	defer file.Close()

	var parseList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			parseList = append(parseList, strings.TrimSpace(scanner.Text()))
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return parseList, nil
}

func ParseQueries() ([]string, error) {
	queryList, err := loadListFile("query.conf")
	if err != nil {
		return nil, err
	}
	if len(queryList) < 1 {
		return nil, errors.New(fmt.Sprintf("\"%v\" is empty!", "./configs/query.conf"))
	}
	return queryList, nil
}

func DeleteQuery(queryID string) error {
	filePath := "./configs/query.conf"

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read the file: %v", err)
	}

	// Convert content to string and replace the target string
	newContent := strings.ReplaceAll(string(content), queryID, "")

	// Write the modified content back to the file
	err = ioutil.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to the file: %v", err)
	}

	return nil
}

func ParseTokens() ([]string, error) {
	tokenList, err := loadListFile("tokens.conf")
	if err != nil {
		return nil, err
	}
	if len(tokenList) < 1 {
		return nil, errors.New(
			"no token or query found in ./configs/tokens.conf\n" +
				"Write your queries in./configs/query.conf like this:\nquery1\nquery2\n...\n\n",
		)
		os.Exit(1)
	}
	return tokenList, nil
}

func FormatUpTime(d time.Duration) string {
	totalSeconds := int(d.Seconds())

	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	return fmt.Sprintf("%dh%dm%ds", hours, minutes, seconds)
}

func FormatLeftDuration(d time.Duration) string {
	if d < 0 {
		return "0s"
	}
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	return fmt.Sprintf("%dh %dm", hours, minutes)
}
