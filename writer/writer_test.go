package writer

import "testing"

func TestWriteToPath(t *testing.T) {
	testwriter := Writer{}
	testwriter.SetOutput("test.txt")
	content := []string{"compton", "puppy", "mennya", "doota"}
	testwriter.WriteToPath(content)
}
