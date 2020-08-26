//This package supports the creation of runner files for languages.
//A runner file is a file properly formatted in a given language that can be
//executed by the program package.
package runner

import (
	"os"
	"strings"
)

var supportedLanguages = map[string]func(string) (string, string){
	"python": python,
	"java":   java,
}

//IsSupportedLanguage checks if the given language is supported
func IsSupportedLanguage(lang string) bool {
	_, found := supportedLanguages[lang]
	return found
}

//GetFunctor returns the create file function for the given language
func GetFunctor(lang string) func(string) (string, string) {
	if IsSupportedLanguage(lang) {
		return supportedLanguages[lang]
	}
	return nil
}

func createFileAndAddCode(outFileName string, code string) error {
	runnerFile, err := os.Create(outFileName)
	if err == nil {
		runnerFile.WriteString(code)
		runnerFile.Close()
	}

	return err
}

func getRunnerFileLocation(suffix string) string {
	var location strings.Builder
	location.WriteString("../runner_files/")
	location.WriteString(suffix)
	return location.String()
}