# Rumble API offline documentation
This directory contains an offline version of the Rumble Swagger API docs. It can be opened locally in a browser,
or be served from a webserver such as Apache or NGINX.

## Using the offline documentation
You have two options for using this documentation.
### Open locally in your browser
You can open the provided `index.html` in your browser. This will expose all the API endpoints you can use. You will
need to change the server URL to be your self-hosted console or `console.rumble.run`.

### Hosting behind a webserver
Hosting the documentation with a webserver is very easy. You will need to configure a site to host static files and
add the directory contents to that site's root directory. We've included links with guides for a few webservers below.

- [NGINX](http://nginx.org/en/docs/beginners_guide.html#static)
- [Apache](https://httpd.apache.org/docs/trunk/getting-started.html)
- [IIS](https://docs.microsoft.com/en-us/iis/configuration/system.webserver/staticcontent/)

## Configuration
You can edit the URL key that's located behind `spec` in `index.html`. This will be the URL to the 
API server.
