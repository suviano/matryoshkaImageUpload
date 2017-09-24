# nameless-storage-image-saver

## Overview

Save image file in different sizes into google cloud storage.

### execute directly

``` sh
export GOOGLE_CLOUD_PROJECT="<project id>" && go run cmd/main.go -bucket="<bucket name>"
```

### request sample

``` sh
curl --request PUT \
  --url http://localhost:8080/image/tcc \
  --header 'content-type: multipart/form-data' \
  --form object=@/path/to/the/file/an-image.png
```
