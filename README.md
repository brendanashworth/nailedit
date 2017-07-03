# nailedit

> `nailedit` generates thumbnail images from the command line. It resizes
images and changes their format according to passed options and uses the
fastest available library to do so.

## Usage

`nailedit` takes input from STDIN and pipes the thumbnail to STDOUT.

```sh
$ cat image.jpeg | nailedit > thumbnail.jpeg

$ nailedit -help
```

## Options

- *`-dimensions`*: dimensions of the output thumbnail image, in format `WIDTHxHEIGHT`
- *`-format`*: format of the output thumbnail image (i.e., `jpeg`, `png`)
