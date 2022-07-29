package report

import (
	"estemplate/file"
	"estemplate/pdf"
	"estemplate/tpl"
	"estemplate/zip"
	"os"
)

const unzipDestDir = "rend"
const unzipDestDirWithSlash = "rend/"
const odtContentFile = "rend/content.xml"
const renderedFileOut = "rendered.odt"
const renderedPdfOut = "rendered.pdf"

// Only .odt files are supported
func Render(filename string, data map[string]any) (string, error) {
	err := zip.Decompress(filename, unzipDestDir)
	if err != nil {
		return "", err
	}

	content, f, err := file.Read(odtContentFile)

	f, err = file.Truncate(f)
	defer f.Close()

	err = render(f, content, data)
	if err != nil {
		return "", err
	}

	err = zip.Compress(unzipDestDir, renderedFileOut)
	if err != nil {
		return "", err
	}

	err = pdf.Create(renderedFileOut)
	if err != nil {
		return "", err
	}

	err = os.RemoveAll(unzipDestDirWithSlash)
	if err != nil {
		return "", err
	}

	return renderedPdfOut, nil
}

func CleanUp() {
	os.Remove(renderedFileOut)
	os.Remove(renderedPdfOut)
}

func render(f *os.File, content string, data map[string]any) error {
	err := tpl.Execute(content, data, f)
	if err != nil {
		return err
	}
	_, err = f.Write([]byte("\n"))
	return err
}
