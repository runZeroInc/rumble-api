npx openapi-generator-cli generate -i rumble-api.yml -g html -o pdfgen/html -c pdfgen/config.yaml
node pdfgen/makepdf.js
