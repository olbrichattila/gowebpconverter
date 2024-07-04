## Image converter to webp

## This is a work in porgress....

Converts image from well known image formats to webp

usecases:
1. As command line converter
2. Web server, converted CDN

It caches the image

## Usage:
### Command:

Command line, outputs the converted image to standard output. You can pipe to a file, or use as a CGI script
```
go run ./cmd/cmd test.jpg > test.webp
```

### Web server:
```
go run ./cmd/web
```

it will listen on port 8000

Example URL:

```http://localhost:8000/test2.jpg```
The test2 image will be converted to webp and displayed. Create a cache folder before.

TODO:
- Cahe to database
- Cache to redis
- Cache to S3
- Read cachable from all above
- Image resize, Crop
- Cache TTL, 
- .env to setup cache on, off, ttl, cache type.

