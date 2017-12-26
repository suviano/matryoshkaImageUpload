package matryoshka

import (
	"bytes"
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
)

var (
	storageClient IStorage
)

func init() {
	storageClient = &StorageClient{}
}

// WriteImage save images
func WriteImage(prefix, filePath, bucket string, buf *bytes.Buffer) (map[string]*BufMedia, error) {
	if prefix == "" {
		return nil, errors.New("empty prefix")
	}
	if bucket == "" {
		return nil, errors.New("empty bucket")
	}
	if buf == nil || buf.Len() == 0 {
		return nil, errors.New("empty buffer")
	}

	fileName, ext, mimeTyp, err := solveImgInfo(filePath)
	if err != nil {
		log.Warningf("{WriteImage}{error solving image attr: %v}", err)
		return nil, err
	}

	imgMapBuf, err := generateImgsByScale(buf, prefix, fileName, ext, mimeTyp)
	if err != nil {
		log.Warningf("{WriteImage}{error generating imgs: %v}", err)
		return imgMapBuf, err
	}

	ctx := context.Background()
	for i, imgBuffer := range imgMapBuf {
		err = storageClient.SaveImg(ctx, prefix, bucket, imgBuffer)
		if err != nil {
			log.Warningf("{WriteImage}{error saving image on storage: %v}", err)
			return imgMapBuf, err
		}
		imgMapBuf[i].Buf = nil
	}
	return imgMapBuf, err
}
