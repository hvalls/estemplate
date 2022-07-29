package pdf

import "os/exec"

// Creates filename.pdf (without old extension)
func Create(filename string) error {
	cmd := exec.Command("soffice", "--headless", "--convert-to", "pdf", filename)
	return cmd.Run()
}
