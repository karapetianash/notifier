package notify

import "os/exec"

var command = exec.Command

// Send method sends notification for linux
func (n *Notify) Send() error {
	notifyCmdName := "notify-send"

	notifyCmd, err := exec.LookPath(notifyCmdName)
	if err != nil {
		return err
	}

	notifyCommand := command(notifyCmd, "-u", n.severity.String(), n.title, n.message)
	return notifyCommand.Run()
}
