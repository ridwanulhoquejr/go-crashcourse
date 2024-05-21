package talking

import "fmt"

func Talk(name string) string {
	talk := fmt.Sprintf("Hi, %v. Welcome to the talking package!", name)

	return talk
}
