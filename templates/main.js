#!/usr/bin/env node

const fs = require('fs');
const readline = require('readline');

const args = process.argv.slice(2);
const inputIndex = args.indexOf('-i');

if (inputIndex === -1 || !args[inputIndex + 1]) {
  console.error('Usage: node main.js -i <input-file>');
  process.exit(1);
}

const inputFileName = args[inputIndex + 1];

async function main() {
  const fileStream = fs.createReadStream(inputFileName);
  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });

  for await (const line of rl) {
    console.log(line);
  }
}

main().catch(err => {
  console.error('Error:', err);
  process.exit(1);
});
