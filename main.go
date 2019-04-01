package trash_cli

import (
	"fmt"
	"os"
)

func MoveFile(source string, destination string) {
	err := os.Rename(source, destination)
	if err != nil {
		fmt.Println(err)
	}
}
