package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
	"errors"
)

const prefix = "/"

func HandleFileList (writer http.ResponseWriter, request *http.Request)error {
	if strings.Index(request.URL.Path,prefix) != 0{
		return errors.New("path must start " + "with" + prefix )
	}

	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}