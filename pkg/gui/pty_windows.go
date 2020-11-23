// +build windows

package gui

import "os/exec"

func (gui *Gui) onResize() error {
	return nil
}

func (gui *Gui) newPtyTask(viewName string, cmd *exec.Cmd, prefix string) error {
	return gui.newCmdTask(viewName, cmd, prefix)
}
