package common

import (
	"fmt"
	"os/exec"
)

func Notify(key string, jsonString string) {
	arg := fmt.Sprintf("%s=%s", key, jsonString)
	cmd := exec.Command("eww", "-c", "/home/kingsley/.config/eww", "update", arg)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
