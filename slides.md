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

# Colors, images and gifs: bring on the fun with Go

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

# Can you guess what do these images have in common

<img src="/images/actual-cave-with-sand.png" class="absolute top-25 left-20" style="width: 45%; height: auto;"/>
<img src="/images/forest.png" class="absolute bottom-5 left-10" style="width: 23%; height: auto;"/>
<img src="/images/bgGradient.png" class="absolute bottom-25 left-85" style="width: 20%; height: auto;"/>
<img src="/images/LunchBreak.gif" class="absolute bottom-8 left-67" style="width: 30%; height: auto;"/>
<img src="/images/game_of_life.gif" class="absolute bottom-10 right-5" style="width: 40%; height: auto;"/>

<!-- Show a sequence of 3 images or gifs created with Go -->

---
layout: lblue-fact
---

They are all made in Go

---
layout: fact
---

# ⚠️ Disclaimer

<v-click>
Most of the images in this presentation are made in Go
</v-click>

<!-- If an image is not made in Go you'll notice -->

---

# But why would I even start make images or gifs in Go?

<v-clicks>

- I love Go
- It’s fun
- The results can be weirdly satisfying

</v-clicks>

<img v-click src="/images/whatIsBestForImageProcessing.png" class="absolute top-25 left-30" style="width: 55%; height: auto;"/>
<img v-click src="/images/very-dull-page.png" class="absolute bottom-10 left-10" style="width: 40%; height: auto;"/>

<!-- 
- I love Go
  - It's a straightforward language that can generate performant programs
- It’s fun 
- The results can be weirdly satisfying
  - There are some images I couldn't believe myself to be able to generate
- When I googled "which programming language is best for pictures" I got this
  - So of course I had to do it in Go
-->

---
layout: center
---

# 👋 Hello

- I'm Michele
- I work in Amadeus
- I deploy and operate Splunk
- My hobbies include languages, board games and TBA (something fun)
- And I'm going to talk you about creativity, art and fun with Go!

<!-- Amadeus: we make travel working -->

---
layout: lblue-fact
---

Let's start creating our first Image

---

# 🖼️ Our First Image

Step by step

<v-click>
1. We start by defining the bounds of the image as a rectangle.

```go
r := image.Rect(0, 0, 1024, 768) // A 1024x768 image
```

</v-click>

<v-click>
2. We create the image using the rectangle and a specific color space (RGBA).

```go
img := image.NewRGBA(r)
```

</v-click>

<v-click>
3. We set the pixels of the image to the color we want

```go
for x := range r.Max.X {
  for y := range r.Max.Y {
    img.Set(x, y, color.RGBA{G: 150, A: 255})
  }
}
```

</v-click>

<v-click>
4. We encode the image into a file with a specific format

```go
f, _ := os.Create("green.png")
png.Encode(f, img)
```

</v-click>

<img v-click="1" src="/images/bounds.png" class="absolute top-18 right-10" style="width: 28%; height: auto;"/>
<img v-click="5" src="/images/green.png" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>

---

# 🖼️ Our First Image

Full code

```go{all|13-17|15}
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

<img src="/images/green.png" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>

---

# 🎨 Beyond our First Image

````md magic-move
```go
img.Set(x, y, color.RGBA{G: 150, A: 255})
```

```go
img.Set(x, y, color.RGBA{B: 150, A: 255})
```

```go
switch (x/32 + y/32) % 2 {
case 0:
  img.Set(x, y, color.RGBA{B: 150, A: 255})
default:
  img.Set(x, y, color.RGBA{G: 150, A: 255})
}
```

```go
img.Set(x, y, color.RGBA{
  B: uint8(float64(x) / float64(r.Max.X) * 255),
  G: uint8(float64(y) / float64(r.Max.Y) * 255),
  A: 255,
})
```

```go
img.Set(x, y, color.RGBA{
  R: uint8((1 + math.Cos(float64(x)/10)) * 255),
  G: uint8((1 + math.Sin(float64(y)/10)) * 255),
  A: 255,
})
```

```go
nano := uint32(time.Now().UnixNano())
img.Set(x, y, color.RGBA{
  R: uint8((nano >> 16) & 0xFF),
  G: uint8((nano >> 8) & 0xFF),
  B: uint8(nano & 0xFF),
  A: 255,
})
```
````

<img src="/images/green.png" class="absolute top-18 right-10" style="width: 28%; height: auto;"/>
<img v-click="+1" src="/images/blue.png" class="absolute top-35 right-20" style="width: 30%; height: auto;"/>
<img v-click="+2" src="/images/bgCheckerboard.png" class="absolute top-50 right-35" style="width: 30%; height: auto;"/>
<img v-click="+3" src="/images/bgGradient.png" class="absolute top-65 right-50" style="width: 28%; height: auto;"/>
<img v-click="+4" src="/images/plaid.png" class="absolute top-50 right-25" style="width: 40%; height: auto;"/>
<img v-click="+5" src="/images/timeflow.png" class="absolute top-50 right-25" style="width: 45%; height: auto;"/>

<!-- 
In the first image we started by setting all pixels to a fully opaque medium green (not too bright, not too dark)

We can easily change that to a fully opaque medium blue (not too bright, not too dark)

Or go even a bit further and create a checkerboard pattern or a gradient of blue and green where each pixel is a combination of blue and green that depend on the coordinates

We decide the color we want to paint the image the way we want and with the tools we want: our choice is just a function that decides what color goes where.

We can use math (yes, math) to create a nice plaid.

Or we can even visualize the flow of time. By the way if your really want to see how the fime flows, here it is... a malfunctioning TV screen: mystery solved, you're welcome.
-->

---
layout: lblue-fact
---

Fun fact

---
layout: fact
---

Piet Mondrian famously used Go to paint his "Composition with Red, Blue and Yellow" painting in 1930

<img src="/images/Piet_Mondriaan.jpg" class="absolute top-5 left-15" style="width: 15%; height: auto;"/>
<img src="/images/Piet_Mondriaan,_1930_-_Mondrian_Composition_II_in_Red,_Blue,_and_Yellow.jpg" class="absolute bottom-20 right-15" style="width: 10%; height: auto;"/>
<img v-click src="/images/pietGondrian.png" class="absolute top-10 left-60" style="width: 50%; height: auto;"/>

---

# 🧮 Using Inputs and Matrices

Guided creation of images

When we set a pixel to a specific color in this way

```go
img.Set(x, y, color.RGBA{G: 150, A: 255})
```

We are the ones deciding the color, either by hardcoding it or by computing it via a function

Another way to create images is to be guided by some input

---

# 🧮 Using Inputs and Matrices

From inputs

```
100021112110202312022010330204312040000111012143445142221414240220240442332040010320133120230011020
111110201332323210211143123214321343332124211413115514155115511033421001222101204330001300333011010
222100123300231230122203432310224441551434231352112532354252244334410042212441233243220102033110020
212210311310102334321243104133012535144151532555155341325352512544453402400411340202133231000102020
112002102322001003304233302040124544411225523533534235113353522233535550432232202401221022110311110
112102131331201432320312233124434232544144233241123334112232531521542551332434224211234133132330300
...
```

```
...
SbcccccccaaaaacaaaaaaaaccccccaaaaaccccccccciiinnntttxxxEzzzzyyyyvvqqqjjjdddccccc
abcccccccccccccaaaaaaaaaccccaaaaaaccccccccciiinnnntttxxxxyyyyyvvvvqqjjjdddcccccc
abcccccccccccccaaaaaaaaaacccaaaaaacccccccccciiinnnttttxxxyyyyyvvvqqqjjjdddcccccc
abccccccccccccccccaaaaaaacccaaaaaaccccccccccciiinnnntttwyyywyyyvvrrrkkjdddcccccc
abcccccccccccccccaaaaaaaaccccaaaccccccccccccciiihnnnttwwwywwyyywvrrrkkkeeccccccc
abcccccccccccccccaaaaaaaaccccccccccccccccccccchhhmmmsswwwwwwwwwwwvrrkkkeeccccccc
abcccccccaacccccccacaaacccccccccccccccccccaacchhhhmmsswwwwwswwwwwrrrkkkeeccccccc
abcccccccaaaccacccccaaacccccccccccccccaaccaaccchhhmmssswwwssrrwwwrrrkkkeeccccccc
abcccccccaaaaaaacccccccccccaaaccccccccaaaaaaccchhhmmssssssssrrrrrrrrkkkeeaaacccc
abcccccaaaaaaaaccccccccccccaaaaccccccccaaaaaaachhhmmmssssssllrrrrrrkkkeeeaaacccc
abccccaaaaaaaaaccccccccccccaaaacccccccccaaaaacchhhmmmmsssllllllllkkkkkeeeaaacccc
...
```

<arrow v-click x1="170" y1="200"  x2="70" y2="150" color="#F00" width="2" arrowSize="1" />
<arrow v-click x1="180" y1="200"  x2="80" y2="150" color="#F00" width="2" arrowSize="1" />
<arrow v-click x1="190" y1="200"  x2="90" y2="150" color="#F00" width="2" arrowSize="1" />

---

# 🧮 Using Inputs and Matrices

To matrices

```go
for x := range r.Max.X {
  for y := range r.Max.Y {
    img.Set(x, y, color.RGBA{
      G: 35 + uint8(175.0/float64(m[x][y]+1)),
      A: 255,
      })
  }
}
```

or

```go
for x := range r.Max.X {
  for y := range r.Max.Y {
    img.Set(x, y, color.RGBA{
      R: uint8(55 + 200*(float64(m[x][y])-('a'-1))/float64('z'+1-('a'-1))),
      G: uint8(50 + 150*(float64(m[x][y])-('a'-1))/float64('z'+1-('a'-1))),
      B: uint8(25 + 100*(float64(m[x][y])-('a'-1))/float64('z'+1-('a'-1))),
      A: 255,
    })
  }
}
```

<arrow v-click x1="380" y1="250" x2="320" y2="210" color="#F00" width="2" arrowSize="1" />
<arrow v-click="1" x1="380" y1="525" x2="320" y2="465" color="#F00" width="2" arrowSize="1" />

---

# 🧮 Using Inputs and Matrices

To images

```
100021112110202312022010330204312040000111012143445142221414240220240442332040010320133120230011020
111110201332323210211143123214321343332124211413115514155115511033421001222101204330001300333011010
222100123300231230122203432310224441551434231352112532354252244334410042212441233243220102033110020
212210311310102334321243104133012535144151532555155341325352512544453402400411340202133231000102020
112002102322001003304233302040124544411225523533534235113353522233535550432232202401221022110311110
112102131331201432320312233124434232544144233241123334112232531521542551332434224211234133132330300
...
```

```
...
SbcccccccaaaaacaaaaaaaaccccccaaaaaccccccccciiinnntttxxxEzzzzyyyyvvqqqjjjdddccccc
abcccccccccccccaaaaaaaaaccccaaaaaaccccccccciiinnnntttxxxxyyyyyvvvvqqjjjdddcccccc
abcccccccccccccaaaaaaaaaacccaaaaaacccccccccciiinnnttttxxxyyyyyvvvqqqjjjdddcccccc
abccccccccccccccccaaaaaaacccaaaaaaccccccccccciiinnnntttwyyywyyyvvrrrkkjdddcccccc
abcccccccccccccccaaaaaaaaccccaaaccccccccccccciiihnnnttwwwywwyyywvrrrkkkeeccccccc
abcccccccccccccccaaaaaaaaccccccccccccccccccccchhhmmmsswwwwwwwwwwwvrrkkkeeccccccc
abcccccccaacccccccacaaacccccccccccccccccccaacchhhhmmsswwwwwswwwwwrrrkkkeeccccccc
abcccccccaaaccacccccaaacccccccccccccccaaccaaccchhhmmssswwwssrrwwwrrrkkkeeccccccc
abcccccccaaaaaaacccccccccccaaaccccccccaaaaaaccchhhmmssssssssrrrrrrrrkkkeeaaacccc
abcccccaaaaaaaaccccccccccccaaaaccccccccaaaaaaachhhmmmssssssllrrrrrrkkkeeeaaacccc
abccccaaaaaaaaaccccccccccccaaaacccccccccaaaaacchhhmmmmsssllllllllkkkkkeeeaaacccc
...
```

<img v-click="1" src="/images/forest.png" class="absolute top-15 right-20" style="width: 25%; height: auto;"/>
<img v-click="1" src="/images/hill.png" class="absolute bottom-5 right-25" style="width: 40%; height: auto;"/>

<!-- 
Add joke about the matrix movie and or bitmaps 
-->

---

# 🧮 Using Inputs and Matrices

More complex input

```
525,119 -> 525,122 -> 523,122 -> 523,125 -> 529,125 -> 529,122 -> 528,122 -> 528,119
497,69 -> 497,73 -> 489,73 -> 489,78 -> 504,78 -> 504,73 -> 501,73 -> 501,69
480,38 -> 480,31 -> 480,38 -> 482,38 -> 482,35 -> 482,38 -> 484,38 -> 484,35 -> 484,38 -> 486,38 -> 486,28 -> ...
480,38 -> 480,31 -> 480,38 -> 482,38 -> 482,35 -> 482,38 -> 484,38 -> 484,35 -> 484,38 -> 486,38 -> 486,28 -> ...
...
```

<v-click>

`(x, y)` coordinates in a 2D space representing walls inside a cave
</v-click>

<v-clicks>

- `525,119 -> 525,122`
  - a vertical wall (only Y changes)
- `525,122 -> 523,122`
  - a horizontal wall (only X changes)
</v-clicks>

We transform these sequences into a byte matrix and encode each element of the cave as a specific char
- '0' -> empty space, '*' -> sand, 'x' -> wall

<!-- So we uses these rules to build a matrix with the coordinates representing this cave 
and use it as input to color our image-->

---

# 🧮 Using Inputs and Matrices

The cave

```go
// var cave [][]byte
for i := range cave {
  for j := range cave[i] {
    var r, g, b uint8
    switch cave[i][j] {
    case 0:
      r, g, b = 25, 10, 0 // ~ black
    case '*':
      r, g, b = 150, 150, 0 // yellow
    default:
      r, g, b = 200, 100, 0 // orange
    }
    img.Set(i, j, color.RGBA{R: r, G: g, B: b, A: 255})
  }
}
```

<img v-click src="/images/cave.png" class="absolute top-45 left-35" style="width: 70%; height: auto;"/>
<img v-click src="/images/first-cave.gif" class="absolute top-45 left-35" style="width: 70%; height: auto;"/>
<img v-click src="/images/cave-with-sand.png" class="absolute top-45 left-35" style="width: 70%; height: auto;"/>
<img v-click src="/images/actual-cave-with-sand.png" class="absolute top-45 left-35" style="width: 70%; height: auto;"/>

---
layout: center
---

# 🌍 A real world example

Github contribution table

<img v-click src="/images/actual-gh-contributions.png" class="absolute bottom-20 left-50" style="width: 60%; height: auto;"/>

---

# 🌍 A real world example

Github contribution table

<v-clicks>

We take the input from the HTML of a github user's homepage

```html
<td ... id="contribution-day-component-3-6" data-level="3" ...></td>
```

- `id` attribute has the (x, y) coordinates
- `data-level` has the index of the color from a __palette__

</v-clicks>

<br/>

<v-click>

And a palette is easily definable in Go

```go
// type color.Palette []color.Color
p := color.Palette{
  color.RGBA{R: 239 G: 242 B: 245 A: 255},  // 0xEFF2F5
  color.RGBA{R: 172 G: 238 B: 187 A: 255},  // 0xACEEBB
  color.RGBA{R: 74  G: 194 B: 107 A: 255},  // 0x4AC26B
  color.RGBA{R: 45  G: 164 B: 78  A: 255},  // 0x2DA44E
  color.RGBA{R: 17  G: 99  B: 41  A: 255},  // 0x116329
}
```

</v-click>

<img v-click src="/images/actual-gh-contributions.png" class="absolute top-60 right-15" style="width: 40%; height: auto;"/>
<arrow v-click="5" x1="360" y1="400" x2="850" y2="330" color="#F00" width="2" arrowSize="1" />

---

# 🌍 A real world example

Github contribution table

<img src="/images/actual-gh-contributions.png" class="absolute top-30 left-30" style="width: 70%; height: auto;"/>
<img v-click src="/images/generated-gh-contributions.png" class="absolute bottom-30 left-40" style="width: 65%; height: auto;"/>

---
layout: lblue-fact
---

Fun fact

---
layout: fact
---

A less known version of the Monalisa is a paint by number painting made in Go

<img src="/images/leonardo-da-vinci.jpg" class="absolute top-5 left-15" style="width: 15%; height: auto;"/>
<img src="/images/Mona_Lisa,_by_Leonardo_da_Vinci.jpg" class="absolute bottom-5 right-15" style="width: 10%; height: auto;"/>
<img v-click src="/images/monalisaPaintByNumber.png" class="absolute top-10 left-80" style="width: 38%; height: auto;"/>

---
layout: center
---

# 🗃️ Layering Images and Text

Moving away from setting colors to pixels

So far we have used a lot:

```go
func (p *RGBA) Set(x, y int, c color.Color) // For RGBA
```

We are now moving to:

```go
func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)
```

---

# 🗃️ Layering Images and Text

A basic composition of two images

<v-click>
1. Image 1 will be a green rectangle.

```go
r := image.Rect(0, 0, 800, 600) 
img := image.NewRGBA(r)
for x := range r.Max.X {
  for y := range r.Max.Y {
    img.Set(x, y, color.RGBA{G: 150, A: 255})
  }
}
```
</v-click>

<v-click>
2. Image 2 will be a white rectangle

```go
img := image.NewRGBA(r)
```

</v-click>

<v-click>
3. We set the pixels of the image to the color we want

```go
```

</v-click>

<v-click>
4. We encode the image into a file with a specific format

```go
f, _ := os.Create("green.png")
png.Encode(f, img)
```

---

## 🖼️ Layering Images and Text

- Composing visuals with layers
- Adding text overlays
- Creating banners and posters


---
layout: lblue-fact
---

Fun fact

---
layout: fact
---

https://en.m.wikipedia.org/wiki/Colored_Mona_Lisa

<img src="/images/leonardo-da-vinci.jpg" class="absolute top-5 left-15" style="width: 15%; height: auto;"/>
<img src="/images/Mona_Lisa,_by_Leonardo_da_Vinci.jpg" class="absolute bottom-5 right-15" style="width: 10%; height: auto;"/>
<img v-click src="/images/monalisaPaintByNumber.png" class="absolute top-10 left-80" style="width: 38%; height: auto;"/>

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

- It shows that Go can hold its ground when working with images
  - Go is not the first language mentioned for image creation, so of course I had to do that (wink to the way Ron Evans says, Go is not for that, so of course I had to do it in Go)
- 🎨 Why Color in Code?
  - Programming isn't just logic—it's also art
  - Go is fast, simple, and surprisingly good at graphics
  - We'll use (almost) only the standard library!

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