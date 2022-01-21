cd $(git rev-parse --show-toplevel)
npx openapi-generator-cli generate -i rumble-api.yml -g html -o pdfgen/html -c pdfgen/config.yaml
# Work around https://github.com/OpenAPITools/openapi-generator/issues/11292
sed -i '{ /<div class="param">.*Minus/s/Minus/-/g }' pdfgen/html/index.html
node pdfgen/makepdf.js
