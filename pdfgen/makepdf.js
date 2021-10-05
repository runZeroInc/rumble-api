const puppeteer = require('puppeteer');

(async () => {
    const browser = await puppeteer.launch();
    const page = await browser.newPage();
    await page.goto('file://' + __dirname + '/html/index.html', {
        waitUntil: 'networkidle0',
    });
    await page.pdf({
        path: 'rumble-api.pdf',
        format: 'letter',
        margin: {
            top: "1in",
            bottom: "1in",
            left: "0.75in",
            right: "0.74in"
        },
        headerTemplate: '<div style="width: 100%"><p style="text-align: center; font-size: 9px;">Rumble API documentation</p></div>',
        footerTemplate: '<div style="width: 100%"><p style="text-align: center; font-size: 9px;">Page&nbsp;<span class="pageNumber"></span>&nbsp;of&nbsp;<span class="totalPages"></span></p></div>',
        displayHeaderFooter: true,
        printBackground: true
    });
    await browser.close();
})();
