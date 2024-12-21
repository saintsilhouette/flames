# The Fractal Flame Algorithm

## Table of Contents
- [Table of Contents](#table-of-contents)
- [Implemented Affine Transformation](#implemented-affine-transformation)
- [Usage](#usage)
- [Installation](#installation)
- [References](#references)

The *Fractal Flame* Algorithm is a member of the Iterated Function System (IFS) class of fractal algorithms. A 2D IFS creates images by plotting the output of a chaotic attractor directly on the image plane.

## Implemented Affine Transformation 
- Linear
- Sinusoidal
- Spherical
- Polar 
- Heart
- Disk

## Usage
There are a few options available to run this application:

```sh
  -g string
        maximum number of goroutines (default "32")
  -h string
        image height (default "1080")
  -i string
        number of iterations per point (default "512")
  -s string
        number of event loop iterations (default "32768")
  -w string
        image width (default "1920")
```

For instance:

```sh
./fractal -w 2560 -h 1440
```

The options of generating multiple images are available: by default, the results will be stored in folder `images` in the root of the project.

## Installation
You should have Go 1.22 or newer installed.
```sh
git clone git@github.com:voidsilhouette/flames.git
```

## References
- https://en.wikipedia.org/wiki/Fractal_flame
- https://flam3.com/flame_draves.pdf
- https://habr.com/ru/articles/251537/