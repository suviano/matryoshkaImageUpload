# IN DEVELOPMENT

# nameless-storage-image-saver

[![Build Status](https://travis-ci.org/suviano/matryoshkaImageUpload.svg?branch=master)](https://travis-ci.org/suviano/matryoshkaImageUpload)

## Overview

Save image file in different sizes into google cloud storage.

### execution test

the folder `cmd` contains a file to execute a test

``` sh
export GOOGLE_CLOUD_PROJECT="<project id>" && go run cmd/main.go -bucket="<bucket name>"
```

