#!/bin/bash

DIST_DIR="rumble-api"
ZIP_NAME="swagger_latest.zip"
DIST_ZIP="rumble-api.zip"
OAS3_FILE="../rumble-api.yml"

## Clean up any directories that may have lingered
if [[ -d "$DIST_DIR" ]]; then
  rm -r "${DIST_DIR:?}"
fi

if [[ -f "$ZIP_NAME" ]]; then
  rm "${ZIP_NAME:?}"
fi

if [[ -f "$DIST_ZIP" ]]; then
  rm "${DIST_ZIP:?}"
fi

find . -name "swagger-api*" -print0 | xargs -r0 rm -r

## Get latest swagger zip from Github
LINK="$(curl -s https://api.github.com/repos/swagger-api/swagger-ui/releases/latest | yq eval '.zipball_url' -)"

## Download an unzip
echo "Downloading latest Swagger UI from: ${LINK}"
curl -s -o swagger_latest.zip -L "${LINK}"
unzip -q swagger_latest.zip
mkdir -p "${DIST_DIR}"
cp -r swagger-api*/dist/* "${DIST_DIR}"

## Copy rumble API spec
cp "$OAS3_FILE" "${DIST_DIR}/rumble-api.yml"
cp README.md "${DIST_DIR}/README.md"

## Convert yaml to json
SPEC_JSON="$(yq eval -o=j -I=-0 "${DIST_DIR}"/rumble-api.yml | sed -e 's/[\/&]/\\&/g' -e 's/console.rumble.run/localhost/g')"

## Edit index to use the rumble API spec
URL_VAL="$(grep url rumble-api/index.html | sed -e 's/[\/&]/\\&/g')"
sed -i "s/${URL_VAL}/spec:${SPEC_JSON},/" rumble-api/index.html

## Zip up the files for distribution
zip -rq rumble-api.zip rumble-api

## Clean up
find . -name "swagger-api*" -print0 | xargs -r0 rm -r
rm swagger_latest.zip
rm -r rumble-api

echo "Finished... ${DIST_ZIP} is ready for distribution."
