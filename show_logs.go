package show_logs

import (
	"fmt"
	"log"
)

var showLogs bool

func CheckParam(param string) {
	if param == "show_logs" {
		showLogs = true
	}
}

func Enable() {
	showLogs = true
}

func Disable() {
	showLogs = false
}

func IsEnabled() bool {
	return showLogs
}

func Print(a ...interface{}) {
	if showLogs {
		fmt.Print(a...)
	}

	log.Print(a...)
}

func Println(a ...interface{}) {
	if showLogs {
		fmt.Println(a...)
	}

	log.Println(a...)
}

func Printf(str string, a ...interface{}) {
	if showLogs {
		Print(fmt.Sprintf(str, a...))
	}
}

func Printfln(str string, a ...interface{}) {
	if showLogs {
		Println(fmt.Sprintf(str, a...))
	}
}

func Fatal(a ...interface{}) {
	if showLogs {
		fmt.Print(a...)
	}

	log.Fatal(a...)
}

func Fatalln(a ...interface{}) {
	if showLogs {
		fmt.Println(a...)
	}

	log.Fatalln(a...)
}

func Fatalf(str string, a ...interface{}) {
	if showLogs {
		Fatal(fmt.Sprintf(str, a...))
	}
}

func Fatalfln(str string, a ...interface{}) {
	if showLogs {
		Fatalln(fmt.Sprintf(str, a...))
	}
}

func Panic(a ...interface{}) {
	if showLogs {
		fmt.Print(a...)
	}

	log.Panic(a...)
}

func Panicln(a ...interface{}) {
	if showLogs {
		fmt.Println(a...)
	}

	log.Panicln(a...)
}

func Panicf(str string, a ...interface{}) {
	if showLogs {
		Panic(fmt.Sprintf(str, a...))
	}
}

func Panicfln(str string, a ...interface{}) {
	if showLogs {
		Panicln(fmt.Sprintf(str, a...))
	}
}

func Show(a ...interface{}) {
	if showLogs {
		fmt.Print(a...)
	}
}

func Showln(a ...interface{}) {
	if showLogs {
		fmt.Println(a...)
	}
}

func Showf(str string, a ...interface{}) {
	if showLogs {
		Show(fmt.Sprintf(str, a...))
	}
}

func Showfln(str string, a ...interface{}) {
	if showLogs {
		Showln(fmt.Sprintf(str, a...))
	}
}
