package testmod

import (
	"fmt"
	"github.com/orphan/test"
)

func Hi(name string) string {
	return fmt.Sprintf("Hi, %s %s %s %s", name, test.Hello("roberto"), "fuck", "you bitch!!!")
}
