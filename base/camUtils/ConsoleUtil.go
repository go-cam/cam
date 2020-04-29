package camUtils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var Console = new(ConsoleUtil)

type ConsoleUtil struct {
}

// runs the command and returns its combined standard
// output and standard error.
func (util *ConsoleUtil) Run(cmd string) ([]byte, error) {
	name, args := util.parseCommand(cmd)

	c := exec.Command(name, args...)
	c.Stdout = os.Stdout
	err := c.Run()
	if err != nil {
		return nil, err
	}

	var bytes []byte
	_, err = os.Stdout.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// runs the command and returns its combined standard, and print output content realtime
// output and standard error.
func (util *ConsoleUtil) Start(cmd string) error {
	name, args := util.parseCommand(cmd)

	var err error
	c := exec.Command(name, args...)
	c.Stdout = os.Stdout
	reader := bufio.NewReader(os.Stdout)
	err = c.Start()
	if err != nil {
		return err
	}

	var index int
	var contentArray = make([]string, 0, 5)
	contentArray = contentArray[0:0]
	// print while cmd run
	for {

		bytes, err := reader.ReadBytes('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(string(bytes))
		index++
		contentArray = append(contentArray, string(bytes))
	}

	err = c.Wait()
	return err
}

// Run and no output to console
func (util *ConsoleUtil) RunNoOutput(cmd string) error {
	name, args := util.parseCommand(cmd)

	c := exec.Command(name, args...)
	c.Stdout = nil
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

// parse command
func (util *ConsoleUtil) parseCommand(command string) (name string, args []string) {
	splice := strings.Split(command, " ")
	if len(splice) <= 1 {
		var args []string
		return command, args
	}

	return splice[0], splice[1:]
}

func (util *ConsoleUtil) IsLinux() bool {
	return runtime.GOOS == "linux"
}

func (util *ConsoleUtil) IsWindows() bool {
	return runtime.GOOS == "windows"
}

// check user whether press y
func (util *ConsoleUtil) IsPressY() bool {
	input := bufio.NewScanner(os.Stdin)
	if !input.Scan() {
		return false
	}
	str := strings.ToLower(input.Text())
	str = strings.TrimSpace(str)
	if str != "y" {
		return false
	}

	return true
}

// check is run by command mode.
func (util *ConsoleUtil) IsRunByCommand() bool {
	return len(os.Args) >= 2
}

// Check whether the system has the command
func (util *ConsoleUtil) HasCommand(cmd string) bool {
	if util.IsWindows() {
		return util.hasCommandWindow(cmd)
	}
	// TODO linux support
	return false
}

func (util *ConsoleUtil) hasCommandWindow(cmd string) (has bool) {
	defer func() {
		rec := recover()
		has = rec == nil
	}()

	err := util.RunNoOutput("where " + cmd)
	if err != nil {
		panic(err)
	}

	return true
}
