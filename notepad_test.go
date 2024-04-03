package notepad

import (
    "testing"
)

func TestOpen(t *testing.T) {
    err := Open()
    if err != nil {
        t.Errorf("Error opening Notepad: %v", err)
    }
}
