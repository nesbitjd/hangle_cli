#!/bin/bash
#https://docs.github.com/en/rest/releases/releases?apiVersion=2022-11-28#create-a-release

if [ -z "$GITHUB_TOKEN" ]; then
    echo "GITHUB_TOKEN is unset"
    exit 1
fi

if [ -z "$OWNER" ]; then
    echo "OWNER is unset"
    exit 1
fi

if [ -z "$REPO" ]; then
    echo "REPO is unset"
    exit 1
fi

RELEASE_ID=$(curl \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer $GITHUB_TOKEN"\
  -H "X-GitHub-Api-Version: 2022-11-28" \
  https://api.github.com/repos/$OWNER/$REPO/releases/latest \
  | jq '.id') 

if [ -z "$RELEASE_ID" ]; then
    echo "RELEASE_ID is unset"
    exit 1
fi

for filename in release/*
do
curl \
  -X POST \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer $GITHUB_TOKEN"\
  -H "X-GitHub-Api-Version: 2022-11-28" \
  -H "Content-Type: application/octet-stream" \
  "https://uploads.github.com/repos/$OWNER/$REPO/releases/$RELEASE_ID/assets?name=${filename#"release/"}" \
  --data-binary @$filename
done