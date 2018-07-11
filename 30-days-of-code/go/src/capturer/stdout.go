package capturer

import (
    "io/ioutil"
    "os"
    "path/filepath"
)


type stdoutCapturer struct {
    temporaryStdoutPath string
    originalStdout *os.File
    temporaryStdout *os.File
}

func NewStdoutCapturer () *stdoutCapturer {
    s := stdoutCapturer {
        filepath.Join(os.TempDir(), "stdout"),
        os.Stdout,
        nil,
    }
    s.temporaryStdout, _ = os.Create(s.temporaryStdoutPath)

    return &s
}

func (s *stdoutCapturer) StartCapture () {
    os.Stdout = s.temporaryStdout
}

func (s *stdoutCapturer) StopCapture () {
    s.temporaryStdout.Close()
    os.Stdout = s.originalStdout
}

func (s *stdoutCapturer) GetCapturedString () string {
    stdoutCaptured, _ := ioutil.ReadFile(s.temporaryStdoutPath)
    return string(stdoutCaptured)
}
