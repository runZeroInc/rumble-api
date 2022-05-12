const yaml = require('js-yaml');
const fs   = require('fs');

const indexFile = 'rumble-api/index.html';
const specYAML = 'rumble-api.yml';

try {
  const doc = yaml.load(fs.readFileSync(specYAML, 'utf8'));
  const jsonData = JSON.stringify(doc, null, 2);
  const index = fs.readFileSync(indexFile, 'utf8');
  const newIndex = index.replace(/url:.*swagger.io.*\.json",/m, `  spec: ${jsonData},`);
  if (index === newIndex) {
    console.log("Index replacement failed");
    process.exit(1);
  }
  fs.writeFileSync(indexFile, newIndex, 'utf8');
  console.log(`Updated ${indexFile}`);
} catch (e) {
  console.log(e);
}
