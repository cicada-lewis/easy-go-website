package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func streamHandler(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	vid := params.ByName("vid-id")
	vl := VIDEO_DIR + vid

	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Error when try to open file: %v", err)
		sendErrorResponse(writer, http.StatusInternalServerError, "Video Not Found or Internal Error\n")
	}

	writer.Header().Set("Conntent Type", "video/mp4")
	http.ServeContent(writer, req, string(vid), time.Now(), video)
	defer video.Close()
}


func uploadHandler(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	req.Body = http.MaxBytesReader(writer, req.Body, MAX_UPLOAD_SIZE)
	if err := req.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {

		sendErrorResponse(writer, http.StatusBadRequest, "File is too Big!")
		return
	}

	file, _, err := req.FormFile("file")
	if err != nil {
		sendErrorResponse(writer, http.StatusInternalServerError, "Read File Failed!")
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file Failed: %v", err)
		sendErrorResponse(writer, http.StatusInternalServerError, "Read File Failed!")
	}
	fn := params.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR + fn, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(writer, http.StatusInternalServerError, "")
	}
	writer.WriteHeader(http.StatusCreated)
	io.WriteString(writer, "Upload Successfully!")

	return
}