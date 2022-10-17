#!/bin/bash
#
gsutil -m rsync -R ./docs/src/.vuepress/dist gs://docs.imversed.com
#https://cloud.google.com/storage/docs/gsutil/commands/web
gsutil web set -m index.html -e index.html gs://docs.imversed.com
