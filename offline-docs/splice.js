const yaml = require('js-yaml');
const fs   = require('fs');

const indexFile = 'rumble-api/index.html';
const specYAML = 'rumble-api.yml';

try {
  const doc = yaml.load(fs.readFileSync(specYAML, 'utf8'));
  const jsondata = JSON.stringify(doc, null, 2);
  const index = fs.readFileSync(indexFile, 'utf8');
  const newindex = index.replace(/url:.*swagger.io.*\.json",/m, `  spec: ${jsondata},`);
  console.log(newindex);
  fs.writeFileSync(indexFile, newindex, 'utf8');
} catch (e) {
  console.log(e);
}
