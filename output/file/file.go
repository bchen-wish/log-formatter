package file

import (
	"encoding/json"
	"github.com/VertexC/log-formatter/util"
	"os"
)

type FileOutput struct {
	docCh  chan util.Doc
	logger *util.Logger
	f      *os.File
}

func NewFileOutput(filePath string, docCh chan util.Doc) *FileOutput {
	logger := util.NewLogger("file-output")

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0667)
	if err != nil {
		logger.Error.Fatalf("Failed to open file %s with error:%s \n", filePath, err)
	}

	output := &FileOutput{
		docCh: docCh,
		f:     f,
	}

	return output
}

func (output *FileOutput) Run() {
	defer output.f.Close()
	logger := output.logger
	for doc := range output.docCh {
		data, err := json.Marshal(doc)
		if err != nil {
			logger.Warning.Printf("Failed to marshal doc: %+v into json. %s", doc, err)
		}
		output.f.WriteString(string(data) + "\n")
	}
}
