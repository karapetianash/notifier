package notify

import (
	"fmt"
	"os/exec"
)

var command = exec.Command

// Send method sends notification for macOS
func (n *Notify) Send() error {
	notifyCmdName := "terminal-notify"

	notifyCmd, err := exec.LookPath(notifyCmdName)
	if err != nil {
		return err
	}

	title := fmt.Sprintf("(%s) %s", n.severity, n.message)

	notifyCommand := command(notifyCmd, "-title", title, "-message", n.message)
	return notifyCommand.Run()
}
