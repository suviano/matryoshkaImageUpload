package storageImage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
)

func imageHandle() http.Handler {
	handlers := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			WriteImage(w, r)
		}
	}
	return http.HandlerFunc(handlers)
}

type imageRes struct {
	Message string `json:"msg"`
}

// WriteImage save images.
func WriteImage(w http.ResponseWriter, r *http.Request) {
	prefix := strings.Trim(path.Base(r.URL.Path), " ")
	multipartFile, header, err := r.FormFile("object")
	if err != nil {
		log.Warningf("{WriteImage}{Error reading Formfile: %v}", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(imageRes{Message: "error reading formfile"})
		return
	}
	defer multipartFile.Close()

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, multipartFile)
	if err != nil {
		log.Warningf("{WriteImage}{error reading file data: %v}", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(imageRes{Message: ""})
		return
	}

	fileName, ext, mimeTyp, err := fixImgExtension(header.Filename)
	if err != nil {
		log.Warningf("{WriteImage}{throws: %v}", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(imageRes{
			Message: fmt.Sprintf("%s does not have a valid format", header.Filename),
		})
		return
	}

	imgMapBuf, err := generateImgsByScale(buf, prefix, fileName, ext, mimeTyp)
	if err != nil {
		log.Warningf("{WriteImage}{error generating images by size: %v}", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(imageRes{Message: "error saving image"})
		return
	}

	ctx := context.Background()
	for _, imgBuffer := range imgMapBuf {
		err = storageClient.SaveImg(ctx, prefix, bucket, imgBuffer)
		if err != nil {
			log.Warningf("{WriteImage}{error saving image on storage: %v}", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(imageRes{Message: "error saving image"})
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(imageRes{Message: "Success"})
}
