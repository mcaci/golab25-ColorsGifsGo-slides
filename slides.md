---
# You can also start simply with 'default'
theme: seriph
# random image from a curated Unsplash collection by Anthony
# like them? see https://unsplash.com/collections/94734566/slidev
# background: https://cover.sli.dev
background: /images/first-cave.gif
# some information about your slides (markdown enabled)
title: Colors, images, gifs and Go!
info: |
  ## Colors, images and gifs: bring on the fun with Go!"
  Slides for GoLab 2025
# apply unocss classes to the current slide
class: text-center
# https://sli.dev/features/drawing
drawings:
  persist: false
# slide transition: https://sli.dev/guide/animations.html#slide-transitions
transition: slide-left
# enable MDC Syntax: https://sli.dev/features/mdc
mdc: true
# open graph
# seoMeta:
#  ogImage: https://cover.sli.dev
---

# Colors, images and gifs: bring on the fun with Go!

<div class="absolute bottom-10 text-left">
    <div>Michele Caci</div>
    <div>Senior Software Engineer at Amadeus</div>
    <div class="flex m-0 gap-1">
      <a href="https://github.com/mcaci" target="_blank" alt="Michele's GitHub" title="Michele's GitHub"
        class="text-xl slidev-icon-btn opacity-50 !border-none !hover:text-white">
        <carbon-logo-github />
      </a>
      <a href="https://x.com/goMicheleCaci" target="_blank" alt="Michele's X" title="Michele's X"
        class="text-xl slidev-icon-btn opacity-50 !border-none !hover:text-white">
        <carbon-logo-x />
      </a>
      <a href="https://www.linkedin.com/in/michele-caci-47770132/" target="_blank" alt="Michele's Linkedin" title="Michele's Linkedin"
        class="text-xl slidev-icon-btn opacity-50 !border-none !hover:text-white">
        <carbon-logo-linkedin />
      </a>
    </div>
</div>

---
transition: fade-out
backgroundSize: 80%
hide: true
---

# Abstract

<!-- layout: image-right -->

Let's walk through how we can put some color to our programming with Go.

We will explore how to work with colors to create palettes and images, we will see how to fill pixels in a rectangle and make a matrix of numbers become a forest or a map, we will see how to layer images and texts on top of each other to make some colorful banners and add a pinch of creativity to animate them into playful gifs.

And all of this with Go and (almost) always with the standard library!

Creating images and gifs is for sure about having fun and unleashing your creativity so join me to this journey of giving life to images and gifs with Go.

The outline of the talk should go this way

- Introduce myself and my passion for board games
- Explain the game Ticket to Ride and why I decided to implement it in Go
- Show the creation of a Ticket to Ride board in Go
- Explore some Graph algorithms that can give us some insights on
- If time permits also show some gameplay examples

LEVEL: Introductory and Overview
Monday October 6th, 2025 (45 min)

---

# 🧠 What do these have in common?

<img src="/images/cave-with-sand.png" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>
<img v-click="1" src="/images/forest.png" class="absolute top-50 left-25" style="width: 30%; height: auto;"/>
<img v-click="2" src="/images/bgGradient.png" class="absolute top-50 left-50" style="width: 30%; height: auto;"/>
<img v-click="3" src="/images/LunchBreak.gif" class="absolute bottom-50 left-50" style="width: 30%; height: auto;"/>
<img v-click="4" src="/images/game_of_life.gif" class="absolute bottom-25 right-50" style="width: 30%; height: auto;"/>

<!-- Show a sequence of 3 images or gifs created with Go -->

---
layout: lblue-fact
---

They are all made in Go

---

# ⚠️ Disclaimer

In this talk you can expect a lot of:
- Images 🖼️
- GIFs 🎞️
- And Go 🐹

The code will be shared so your creativity can take over!
- Light up those pixels! 🧮
- Let's Put Some Color to Our Programming with Go
  - Unleashing creativity with pixels, palettes, and Go!

I tried to create **all** the images in this presentation using Go.

---

# 🎨 Why would you even create images or gifs in Go?

- I love Go
- It’s fun
- It’s weirdly satisfying
- It shows that Go can hold its ground when working with images
  - Go is not the first language mentioned for image creation, so of course I had to do that (wink to the way Ron Evans says, Go is not for that, so of course I had to do it in Go)

- 🎨 Why Color in Code?
  - Programming isn't just logic—it's also art
  - Go is fast, simple, and surprisingly good at graphics
  - We'll use (almost) only the standard library!

---

# 👋 Welcome!
- Who I am
- Why I love Go
- What this talk is about: creativity, color, and fun!

---

# 🌈 Creating Images

The basic stuff

A basic image starts with a rectangle

```go
r := image.Rect(0, 0, 1024, 768)
```

The images is created with the boundaries given by the rectangle

```go
img := image.NewRGBA(r)
```

We set the pixels of the image to the color we want

```go
for x := range r.Max.X {
  for y := range r.Max.Y {
    img.Set(x, y, color.RGBA{G: 150, A: 255})
  }
}
```

And then we save it into a file

```go
f, _ := os.Create("green.png"); png.Encode(f, img)
```

---
layout: two-cols
---

# 🌈 Creating Images

The code put altogether gives us this basic green image

```go
package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	r := image.Rect(0, 0, 1024, 768)
	img := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			img.Set(x, y, color.RGBA{G: 150, A: 255})
		}
	}
	f, _ := os.Create("green.png")
	png.Encode(f, img)
}

```

::right::

<img src="/images/green.png" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>
<!-- 
<img v-click v-click.hide src="/images/very-dull-page.png" class="absolute bottom-10 right-100" style="width: 30%; height: auto;"/>
<img v-click src="/images/very-dull-page.png" class="absolute bottom-10 right-10" style="width: 30%; height: auto;"/> -->

---

# 🌈 Creating Images

Let's focus on

_We set the pixels of the image to the color we want_

```go
img.Set(x, y, color.RGBA{G: 150, A: 255})
```

This sets a pixel at coordinates `(x, y)` to a fully opaque medium green (not too bright, not too dark)

```go
img.Set(x, y, color.RGBA{B: 150, A: 255})
```

This sets a pixel at coordinates `(x, y)` to a fully opaque medium blue (not too bright, not too dark)

```go
b := float64(x) / float64(r.Max.X) * 255
g := float64(y) / float64(r.Max.Y) * 255
img.Set(x, y, color.RGBA{B: uint8(b), G: uint8(g), A: 255})
```

This sets a pixel at coordinates `(x, y)` to a combination of blue and green that depend on the coordinates

---
layout: two-cols
---

# 🌈 Creating Images

Which on the same program above gives us this image

```go
package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	r := image.Rect(0, 0, 1024, 768)
	img := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
    b := float64(x) / float64(r.Max.X) * 255
    g := float64(y) / float64(r.Max.Y) * 255
    img.Set(x, y, color.RGBA{B: uint8(b), G: uint8(g), A: 255})
		}
	}
	f, _ := os.Create("bgGradient.png")
	png.Encode(f, img)
}

```

::right::

<img src="/images/bgGradient.png" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>

---
layout: two-cols
---

# 🌈 Creating Images

A function of anything

````md magic-move
```go
func cFuncBGGradient(x, y int, r image.Rectangle) color.Color {
  b := float64(x) / float64(r.Max.X) * 255
  g := float64(y) / float64(r.Max.Y) * 255
  return color.RGBA{B: uint8(b), G: uint8(g), A: 255}
}
```

```go
func cfunc(x, y int) color.Color {
  r := math.Abs(math.Cos(float64(x))) * 50
  g := math.Abs(math.Sin(float64(y))) * 200
  b := math.Abs(math.Cos(float64(x*10-y*10))) * 50
  return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}
```

```go
func cfuncT(t time.Time) color.Color {
  b := t.Nanosecond() % 200
  g := t.Nanosecond() % 256
  return color.RGBA{B: uint8(b), G: uint8(g), A: 255}
}
```
````

::right::

<img v-click v-click.hide src="/images/bgGradient.png" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>
<img v-click v-click.hide src="/images/stripes.png" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>
<img v-click v-click.hide src="/images/timedots.png" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>

---
layout: two-cols
---

# 🌈 Creating Palettes and Images

- How to define and use colors in Go
- Building palettes
- Drawing basic shapes and rectangles

````md magic-move
```go
r := image.Rect(0, 0, 1024, 768)
img := image.NewRGBA(r)
for x := range r.Max.X {
  for y := range r.Max.Y {
    img.Set(x, y, color.RGBA{G: 150, A: 255})
  }
}
f, _ := os.Create("green.png")
png.Encode(f, img)
```

```go
r := image.Rect(0, 0, 1024, 768)
img := image.NewRGBA(r)
for x := range r.Max.X {
  for y := range r.Max.Y {
    img.Set(x, y, color.RGBA{B: 150, A: 255})
  }
}
f, _ := os.Create("blue.png")
png.Encode(f, img)
```

```go
r := image.Rect(0, 0, 1024, 768)
img := image.NewRGBA(r)
for x := range r.Max.X {
  for y := range r.Max.Y {
    b := float64(x) / float64(r.Max.X) * 255
    g := float64(y) / float64(r.Max.Y) * 255
    img.Set(x, y, color.RGBA{B: uint8(b), G: uint8(g), A: 255})
  }
}
f, _ := os.Create("bgGradient.png")
png.Encode(f, img)
```
````

::right::

<img src="/images/green.png" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>
<img v-click="1" src="/images/blue.png" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>
<img v-click="2" src="/images/bgGradient.png" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>

---
layout: two-cols
---

## 🧮 From Numbers to Maps
- Visualizing matrices as forests, terrain, or heatmaps
- Mapping values to colors
- Example: turning a grid of numbers into a landscape

````md magic-move
```go
r := image.Rect(0, 0, 1024, 768)
img := image.NewRGBA(r)
for x := range r.Max.X {
  for y := range r.Max.Y {
    img.Set(x, y, color.RGBA{G: 150, A: 255})
  }
}
f, _ := os.Create("green.png")
png.Encode(f, img)
```

```go
r := image.Rect(0, 0, 1024, 768)
img := image.NewRGBA(r)
for x := range r.Max.X {
  for y := range r.Max.Y {
    img.Set(x, y, color.RGBA{B: 150, A: 255})
  }
}
f, _ := os.Create("blue.png")
png.Encode(f, img)
```

```go
r := image.Rect(0, 0, 1024, 768)
img := image.NewRGBA(r)
for x := range r.Max.X {
  for y := range r.Max.Y {
    b := float64(x) / float64(r.Max.X) * 255
    g := float64(y) / float64(r.Max.Y) * 255
    img.Set(x, y, color.RGBA{B: uint8(b), G: uint8(g), A: 255})
  }
}
f, _ := os.Create("bgGradient.png")
png.Encode(f, img)
```
````

::right::

<img v-click="[1, '+1']" src="/images/forest.png" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>
<img v-click="[2, '+1']" src="/images/hill.png" class="absolute top-50 right-25" style="width: 35%; height: auto;"/>
<img v-click="3" src="/images/cave-with-sand.png" class="absolute top-50 right-25" style="width: 40%; height: auto;"/>


<!-- Add joke about the matrix movie and or bitmaps -->

---

## 🖼️ Layering Images and Text
- Composing visuals with layers
- Adding text overlays
- Creating banners and posters

---

## 🎞️ Animating with Go
- Making playful GIFs
- Frame-by-frame animation
- Example: animated banner or pixel art

---

## 🧰 Tools & Techniques
- Go packages used (mostly standard library)
- Tips for working with images and pixels
- How to stay creative while coding
---

## 🧠 What You Can Build
- Fun projects: maps, games, visualizations
- Ideas to explore on your own
- Encouragement to experiment!

---

## 🙌 Wrapping Up
- Recap of what we covered
- Resources and links
- Final thoughts: code is a canvas!

---

## 💬 Q&A
- Ask me anything!
- Connect with me online

---
layout: lblue-end
---

<div class="text-white font-size-10">
Thank you very much!
</div>

<div class="absolute bottom-10">
  <div  class="text-white">Michele Caci</div>
  <div class="flex m-0 gap-1">
    <a href="https://github.com/mcaci" target="_blank" alt="Michele's GitHub" title="Michele's GitHub"
      class="text-xl slidev-icon-btn opacity-50 !border-none !hover:text-white">
      <carbon-logo-github />
    </a>
    <a href="https://x.com/goMicheleCaci" target="_blank" alt="Michele's X" title="Michele's X"
      class="text-xl slidev-icon-btn opacity-50 !border-none !hover:text-white">
      <carbon-logo-x />
    </a>
    <a href="https://www.linkedin.com/in/michele-caci-47770132/" target="_blank" alt="Michele's Linkedin" title="Michele's Linkedin"
      class="text-xl slidev-icon-btn opacity-50 !border-none !hover:text-white">
      <carbon-logo-linkedin />
    </a>
  </div>
</div>

<!-- HOW ABOUT A FUNNY IDEA FOR PEACE OUT? LIKE A ENDING CREDITS? OR SOMETHING ELSE? TO BE SEEN -->