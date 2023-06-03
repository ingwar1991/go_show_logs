package show_logs

import "testing"

func TestCheckParam(t *testing.T) {
    if showLogs {
        t.Error("showLogs should be false")
    }

    CheckParam("show_logs")

    if !showLogs {
        t.Error("showLogs should be true")
    }
}
