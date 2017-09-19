# nameless-storage-image-saver

## Overview

Save image file in different sizes into google cloud storage.


## Technical view


---

### Generate keys

openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

---

### execute directly

``` sh
export GOOGLE_CLOUD_PROJECT="<project id>" && go run cmd/main.go -bucket="<bucket name>"
```
