const puppeteer = require("puppeteer");
const { Storage } = require("@google-cloud/storage");

async function main(urls) {
  console.log(`Passed in urls`);
  console.log("Upload complete!");
}

main(process.argv.slice(2)).catch((err) => {
  console.error(JSON.stringify({ severity: "ERROR", message: err.message }));
  process.exit(1);
});
