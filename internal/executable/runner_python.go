package executable

import (
	"log"
	"strings"
)

//Python creates a runnerFile for python languages.
func createRunnerFilePython(code string, settings *FileSettings) (string, string, error) {
	settings = fillRestOfFileSettings("python", settings)
	langCommand := "python3"
	var fileName strings.Builder
	fileName.WriteString(settings.FileNamePrefix)
	fileName.WriteString(settings.ClassName)
	fileName.WriteString(".py")
	outFileName := getRunnerFileLocation(fileName.String())

	var formattedCode strings.Builder
	insertImportsPython(&formattedCode, settings)
	formattedCode.WriteString(code)
	insertTrailingCodePython(&formattedCode, settings)

	err := createFileAndAddCode(outFileName, formattedCode.String())

	if err != nil {
		log.Fatal("Could not create runner file!")
	}
	return langCommand, outFileName, nil
}

func insertImportsPython(formattedCode *strings.Builder, settings *FileSettings) {
	formattedCode.WriteString(settings.Imports)
	formattedCode.WriteString("\n")
}

func insertTrailingCodePython(formattedCode *strings.Builder, settings *FileSettings) {
	formattedCode.WriteString("\n")
	formattedCode.WriteString(settings.TrailingCode)
}
