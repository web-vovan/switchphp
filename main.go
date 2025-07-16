package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"slices"
	"strings"
)

func main() {
	versions := []string{"7.4", "8.1", "8.2"}

	if len(os.Args) < 2 {
		fmt.Println("Укажите номер версии php")
		return
	}

	newVersion := os.Args[1]
	currentVersion, err := getPhpVersion()

	if err != nil {
		fmt.Println(err)
		return
	}

	if !slices.Contains(versions, newVersion) {
		fmt.Println("Нельзя переключиться на php " + newVersion)
		fmt.Println("Доступные версии: " + strings.Join(versions, ", "))
		return
	}

	cmd1 := exec.Command("brew", "unlink", "php@"+currentVersion)
	cmd2 := exec.Command("brew", "link", "php@"+newVersion)
	cmd3 := exec.Command("php", "-v")

	_, err1 := cmd1.Output()
	_, err2 := cmd2.Output()

	if err1 != nil || err2 != nil {
		fmt.Println("Возникла ошибка")
		return
	}

	output, err := cmd3.Output()

	if err != nil {
		fmt.Println("Возникла ошибка")
		return
	}

	fmt.Println(string(output))
}

func getPhpVersion() (string, error) {
	cmd := exec.Command("php", "-v")
	data, _ := cmd.Output()
	lines := strings.Split(string(data), "\n")

	pattern := `\d.\d`
	re := regexp.MustCompile(pattern)

	for _, line := range lines {
		match := re.FindString(line)

		if len(match) > 0 {
			return match, nil
		}
	}

	return "", fmt.Errorf("не найдена текущая версия php")
}
