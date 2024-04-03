// Package notepad provides a function to open Notepad.
package notepad

import (
    "os/exec"
    "runtime"
)

// Open opens Notepad on the user's system.
func Open() error {
    var cmd *exec.Cmd
    if runtime.GOOS == "windows" {
        cmd = exec.Command("notepad")
    } else {
        return errors.New("unsupported operating system")
    }
    return cmd.Run()
}
