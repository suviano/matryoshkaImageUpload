package storageImage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func imageHandle() http.Handler {
	handlers := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			WriteImage(w, r)
		}
	}
	return applyMiddlewares(http.HandlerFunc(handlers))
}

// A model of body and query params to addProducts
//
// This model is used to hold product list to be added into database.
//
// swagger:parameters WriteImage
type writeImageParams struct {
	// List of binary images
	//
	// required: true
	// in: body
	Images []struct {
		Path    string `json:"path"`
		Content string `json:"content"`
	} `json:"images"`
}

// ImageRes is a response model
//
// response for successfull operation in /image endpoint
//
// swagger:response imageResult
type ImageRes struct {
}

// ImageErrRes is a response model
//
// response for error of /image endpoint
//
// swagger:response imageErrResult
type ImageErrRes struct {
	Message string `json:"msg"`
}

// WriteImage swagger:route POST /image Image
//
// Return path to save images
//
// 	Consumes:
//	- multipart/form-data
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Responses:
//		201: imageResult
//		400: imageErrResult
//		500: imageErrResult
func WriteImage(w http.ResponseWriter, r *http.Request) {
	multipartFile, header, err := r.FormFile("object")
	if err != nil {
		log.Warningf("{WriteImage}{Error reading Formfile: %v}", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ImageErrRes{Message: "error reading formfile"})
		return
	}
	defer multipartFile.Close()

	fileName, mimeTyp, err := fixImgExtension(header.Filename)
	if err != nil {
		log.Warningf("{WriteImage}{throws: %v}", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ImageErrRes{
			Message: fmt.Sprintf("%s does not have a valid format", header.Filename),
		})
		return
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, multipartFile)
	if err != nil {
		log.Warningf("{WriteImage}{error reading file data: %v}", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ImageErrRes{Message: ""})
		return
	}

	imgMapBuf, err := generateImgsByScale(buf, mimeTyp)
	if err != nil {
		log.Warningf("{WriteImage}{error generating images by size: %v}", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ImageErrRes{Message: "error saving image"})
		return
	}

	ctx := context.Background()
	err = storageClient.SaveImg(ctx, imgMapBuf["original"].Buf, bucket, fileName, true)
	if err != nil {
		log.Warningf("{WriteImage}{error saving image on storage: %v}", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ImageErrRes{Message: "error saving image"})
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, buf.String(), header.Size)
}
