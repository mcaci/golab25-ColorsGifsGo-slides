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
      <a href="https://www.linkedin.com/in/michele-caci-47770132/" target="_blank" alt="Michele's Linkedin" title="Michele's Linkedin"
        class="text-xl slidev-icon-btn opacity-50 !border-none !hover:text-white">
        <carbon-logo-linkedin />
      </a>
    </div>
</div>

<!-- Hi everyone and thank you for joining.
Today I want to offer you a break from all the seriousness of the day
and deep dive into a relaxing and fun moment where you'll get the tools to create some beautiful art with Go
 -->

---
transition: fade-out
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
Monday October 6th, 2025 (45 min @15:15)

---

# Can you guess what do these have in common?

<img src="/images/actual-cave-with-sand.png" class="absolute top-25 left-20" style="width: 45%; height: auto;"/>
<img src="/images/bgGradient.png" class="absolute bottom-5 left-40" style="width: 30%; height: auto;"/>
<img src="/images/game_of_life.gif" class="absolute bottom-10 right-5" style="width: 40%; height: auto;"/>

<!-- Having an idea of the topic of the day,
I can guess that you can guess what do these images and gifs have in common-->

---
layout: lblue-fact
---

They are all made in Go

<!-- And the answer is straightforward, they are all made in Go -->

---
layout: fact
---

# ⚠️ Disclaimer

<v-click>
Most of the images you are about to see are made in Go

[golab25](https://github.com/mcaci/golab25-ColorsGifsGo-slides/tree/main/src/go) - [adventOfCode](https://github.com/mcaci/adventOfCode) - [gopherconeu23](https://github.com/mcaci/gopherconeu23img)
</v-click>

<!--
Now, I need to give you a disclaimer here.

Most of the images you are about to see are made in Go.

And as I will share the slides later, these links are the places where you want to have a look for the Go code that generated them.

And I'm sure that for a couple of them you'll get some nice surprises.
-->

---
layout: intro
---

# 👋 Hello

Who am I?

- I'm Michele
- I use Go at work to support the logging infrastructure (Splunk) we use at Amadeus
- My hobbies include languages, board games and making silly GIFs with Go

<br/>

<img v-click src="/images/micheleRomeo.jpg" class="absolute top-15 right-15" style="width: 30%; height: auto;"/>
<arrow v-after x1="800" y1="325" x2="825" y2="250" color="#F00" width="1" arrowSize="1" />
<p v-after class="absolute bottom-50 right-46 opacity-100 transform -rotate-10" color="#F00">me</p>
<p v-click class="absolute bottom-55 right-46 opacity-100 transform -rotate-10" color="#F00">young</p>
<p v-after class="absolute top-25 right-100 opacity-100 transform -rotate-20" color="#F00">actual me</p>
<arrow v-after x1="580" y1="125" x2="640" y2="130" color="#F00" width="1" arrowSize="1" />

<img v-click src="/images/amadeus-workmark_DarkSky.png" class="absolute bottom-15 right-15" style="width: 30%; height: auto;"/>
<img v-after src="/images/splunk-logo.png" class="absolute bottom-10 right-10" style="width: 10%; height: auto;"/>
<p v-click class="absolute bottom-15 left-10 opacity-100 transform -rotate-350" color="#F00">こんにちわ！</p>
<p v-after class="absolute bottom-3 left-13 opacity-100 transform -rotate-10" color="#F00">Bom dia!  </p>
<img v-click src="/images/TTR_USA_map_graph.jpg" class="absolute bottom-2 left-65" style="width: 24%; height: auto;"/>
<img v-click src="/images/golab.gif" class="absolute bottom-8 left-40" style="width: 65%; height: auto;"/>

<!-- 
Well, I haven't presented myself yet
-->

---
layout: statement
---

<div class="font-size-10">
Why?
</div>

<!--
At this point you can legitimately ask me 
-->

---

# Why would you even want to make images/gifs in Go?

<v-clicks>

1. I love Go
2. It’s fun
3. The results can be surprising and quite satisfying

</v-clicks>

<!-- 
I'm going to offer you few simple and reasonable explanations:

- List the reasons

In addition, let me tell you this story.

Some months ago I made a search
-->

---
layout: image
image: /images/whatIsBestForImageProcessing.png
backgroundSize: 65%
---

<p v-click class="absolute top-50 right-46 opacity-100 transform -rotate-10" color="#F00" font-size="8">C\C++?</p>
<p v-click class="absolute bottom-50 left-25 opacity-100 transform -rotate-10" color="#F00" font-size="10">Python?</p>
<p v-click class="absolute bottom-20 right-46 opacity-100 transform -rotate-10" color="#F00" font-size="14">Matlab?</p>

<!-- 
"Which programming language is best for pictures". Pretty straightforward.

And when I saw the results I thought... C/C++? Phyton? Matlab?

I stopped at this sight for a moment.

Then I remembered the wise words of my almost 2-years old son, the one from the photo earlier.

This is what he often tells me: "papaaa', papaaa', no no no no no no no!"
-->

---
layout: lblue-fact
---

So of course I had to use Go!

<!-- 
Fun fact: my son may be in the room right now and may just have repeated that very same sentence.
-->

---

# 🖼️ Let's start with our first image

<v-click>

1. Define the boundary of the image as a rectangle

```go
r := image.Rect(0, 0, 1024, 768) // A 1024x768 image
```

</v-click>

<v-click>

2. Create the image variable with that boundary and a color space

```go
img := image.NewRGBA(r)
```

</v-click>

<v-click>

3. Set the color for each pixel of the image

```go
for x := range r.Max.X {
  for y := range r.Max.Y {
    img.Set(x, y, color.RGBA{G: 150, A: 255})
  }
}
```

<arrow v-after x1="800" y1="325" x2="925" y2="280" color="#F00" width="2" arrowSize="1" />
<p v-after class="absolute bottom-50 right-46 opacity-100 transform -rotate-10" color="#F00">r.Max</p>

</v-click>

<v-click>

4. Encode the image into a file with a specific format

```go
f, _ := os.Create("green.png")
png.Encode(f, img)
```

</v-click>

<img v-click="1" src="/images/bounds.png" class="absolute top-18 right-10" style="width: 28%; height: auto;"/>
<img v-click="5" src="/images/green.png" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>

<!-- 
If you want to take anything out of today, this is the moment to pay maximum attention.

Because here you'll get the 4 steps that you need to perform to start creating beautiful art with Go. Ready? Here are the steps:

Step 1.... Step 2.... Step 3.... Step 4....

With these 4 steps we get our first image: a green rectangle
-->

---

# 🖼️ Our first image's full code

```go{all|11|12|13-17|18-19|all|13-17|15}
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

<!-- 
Let's consolidate everything in one place: this is the full code that creates this green rectangle and inside the code we see the four steps together

1. Define the boundary
2. Create the image
3. Color every pixel
4. Encode to a file

Now for what's coming next you need to focus on the 3rd step: color every pixel and in particular this method call to `img.Set`, where we give the coordinate and the color we want to give to a pixel.
-->

---

# 🎨 Beyond our first image

````md magic-move
```go
img.Set(x, y, color.RGBA{G: 150, A: 255})
```

```go
img.Set(x, y, color.RGBA{B: 150, A: 255})
```

```go
switch {
  case x < 10 || y < 10 || x > r.Max.X-10 || y > r.Max.Y-10:
    img.Set(x, y, color.Black)
  case (x/128+y/128)%2 == 0:
    img.Set(x, y, color.White)
  default: // case (x/128+y/128)%2 == 1:
    img.Set(x, y, color.Black)
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
<img v-click="+2" src="/images/bwChessboard.png" class="absolute top-65 right-50" style="width: 28%; height: auto;"/>
<img v-click="+3" src="/images/bgGradient.png" class="absolute top-55 right-40" style="width: 28%; height: auto;"/>
<img v-click="+4" src="/images/plaid.png" class="absolute top-50 right-25" style="width: 40%; height: auto;"/>
<img v-click="+5" src="/images/timeflow.png" class="absolute top-50 right-25" style="width: 45%; height: auto;"/>

<!-- 
In the first image we started by setting all pixels to a medium green color

Just by updating a field of the RGBA struct, we can easily change that to amedium blue

But we can do more than monochrome images

For the low low price of some basic math we can:
- create a chess board
- or a blue and green gradient

For the even lower price of some trigonometry we can create the pattern for our next plaid.

And for those who feel like contemporary artists, how about visualizing a snapshot of time itself?
-->

---
layout: lblue-fact
---

Fun fact

<!-- Time for a fun fact! And since we are here, we stay in the context of art, which is quite appropriate since we are in Florence, a city well known for its cultural and artistic history -->

---
layout: fact
---

<v-click>

Piet Mondrian actually used Go to paint his "Composition with Red, Blue and Yellow" in 1930
</v-click>

<img src="/images/Piet_Mondriaan.jpg" class="absolute top-5 left-15" style="width: 15%; height: auto;"/>
<img src="/images/Piet_Mondriaan,_1930_-_Mondrian_Composition_II_in_Red,_Blue,_and_Yellow.jpg" class="absolute bottom-20 right-15" style="width: 10%; height: auto;"/>
<img v-click src="/images/piet-mondrian-go.png" class="absolute top-10 left-60" style="width: 50%; height: auto;"/>

<!-- 
Not bad right?
-->

---
layout: lblue-fact
---

Can we use inputs for the creation of images?

<!-- 
Now, we have decided so far which color each pixel should have.

For the next step let's address this question: can we use inputs to guide the creation of images?

...

A more guided image creation

We have been the ones deciding the color so far

Either by hardcoding it or by computing it via a function

Now we see how to use external input to drive the creation of the image
-->

---
layout: center
class: text-center
---

```
...
100021112110202312022010330204312040000111012143445142221414240220240442332040010320133120230011020
111110201332323210211143123214321343332124211413115514155115511033421001222101204330001300333011010
222100123300231230122203432310224441551434231352112532354252244334410042212441233243220102033110020
212210311310102334321243104133012535144151532555155341325352512544453402400411340202133231000102020
112002102322001003304233302040124544411225523533534235113353522233535550432232202401221022110311110
112102131331201432320312233124434232544144233241123334112232531521542551332434224211234133132330300
...
```

<br/>

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

<p v-click class="absolute top-50 left-100 opacity-100 transform -rotate-20 text-size-10" color="#F00">[][]byte</p>

<!-- 
Can we make sense of these inputs?

At a first glance these are sequences of numbers and characters...

Wait

Let me say it again, these are sequences of bytes and each line can be seen as a []byte

And put altogether they are a [][]byte
-->

---
layout: image
image: /images/welcome-to-the-matrix.jpg
---

<!-- 
Which you can absolutely interpret in this way
-->

---

# 🧮 Using Inputs and Matrices

Reading from matrices

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

<br/>

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
<arrow v-click="1" x1="380" y1="480" x2="320" y2="440" color="#F00" width="2" arrowSize="1" />

<!-- 
Or use them to follow the original 4 steps to create an image, let's apply them to this context:
1. boundaries = matrix size
2. create image = unchanged
3. set each pixel = map every cell of the matrix to a color of the image
4. don't forget to encode to file
-->

---
layout: center
class: text-center
---

```
...
100021112110202312022010330204312040000111012143445142221414240220240442332040010320133120230011020
111110201332323210211143123214321343332124211413115514155115511033421001222101204330001300333011010
222100123300231230122203432310224441551434231352112532354252244334410042212441233243220102033110020
212210311310102334321243104133012535144151532555155341325352512544453402400411340202133231000102020
112002102322001003304233302040124544411225523533534235113353522233535550432232202401221022110311110
112102131331201432320312233124434232544144233241123334112232531521542551332434224211234133132330300
...
```

<br/>

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
And this makes these apparently chaotic inputs into a forest or a hilly landscape
-->

---

# 🧮 Using Inputs and Matrices

More elaborate inputs

```
525,119 -> 525,122 -> 523,122 -> 523,125 -> 529,125 -> 529,122 -> 528,122 -> 528,119
497,69 -> 497,73 -> 489,73 -> 489,78 -> 504,78 -> 504,73 -> 501,73 -> 501,69
480,38 -> 480,31 -> 480,38 -> 482,38 -> 482,35 -> 482,38 -> 484,38 -> 484,35 -> 484,38 -> 486,38 -> 486,28 -> ...
480,38 -> 480,31 -> 480,38 -> 482,38 -> 482,35 -> 482,38 -> 484,38 -> 484,35 -> 484,38 -> 486,38 -> 486,28 -> ...
...
```

<v-click>

These are rules on 2D coordinates: $(x,y) \rightarrow (x1,y1)$
</v-click>

<!--
What happens if you have different inputs?

Let's look at these ones

These are rules on 2D coordinates and when you hear 2D coordinates

-->

---
layout: image
image: /images/welcome-to-the-matrix.jpg
---

<!-- 
... again
-->

---
layout: two-cols-header
---

# 🧮 Using Inputs and Matrices

More elaborate inputs

```
525,119 -> 525,122 -> 523,122 -> 523,125 -> 529,125 -> 529,122 -> 528,122 -> 528,119
497,69 -> 497,73 -> 489,73 -> 489,78 -> 504,78 -> 504,73 -> 501,73 -> 501,69
480,38 -> 480,31 -> 480,38 -> 482,38 -> 482,35 -> 482,38 -> 484,38 -> 484,35 -> 484,38 -> 486,38 -> 486,28 -> ...
480,38 -> 480,31 -> 480,38 -> 482,38 -> 482,35 -> 482,38 -> 484,38 -> 484,35 -> 484,38 -> 486,38 -> 486,28 -> ...
...
```

::left::

<v-click>

Rules:
- X coordinate changes?
  - it's a horizontal wall
- Y coordinate changes?
  - it's a vertical wall
</v-click>

::right::

<v-click>

We use these rules to build a `[][]byte`
</v-click>

<v-clicks>

- byte 0 is an empty space
- '*' is sand
- anything else is a wall
</v-clicks>

<!--
These are rules on 2D coordinates representing walls
 So we uses these rules to build a byte matrix with the coordinates representing this cave 
and use it as input to color our image
-->

---

# 🧮 Using Inputs and Matrices

Drawing the landscape

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

<!--
Which means that we can again map the matrix values to a specific color in the image

And it so happens that this, together with the two previous examples come from that moment of the year when I think I have free time for coding: the advent of code

The gif shows the flow of sand inside this cave envirionment. And since it's the advent of code there are two parts of the exercise. This is the first one and this is the second one.

In this moment, I want to stop and share with you this thought: creating an image out of a problem not only helps in picturing the problme itself but also have an expectation of what the solution could look like.

At this point you can say, I love this presentation, it's so much fun and instructive but, will I ever make use ot them?
-->

---
layout: lblue-fact
---

Let's see a real world example

<!-- 
Let's see a real world example
-->

---

# 🌍 GitHub contribution table

Gathering inputs

<v-click>

From the HTML of a github user's homepage

```html
<td ... id="contribution-day-component-3-6" data-level="3" ...></td>
```
</v-click>

<arrow v-click="2" x1="430" y1="240" x2="370" y2="200" color="#F00" width="2" arrowSize="1" />
<arrow v-click="3" x1="540" y1="240" x2="480" y2="200" color="#F00" width="2" arrowSize="1" />

<v-clicks>

- `id` contains the (x, y) coordinates
- `data-level` is the index of the color from a __palette__
</v-clicks>

<img src="/images/actual-gh-contributions.png" class="absolute bottom-20 left-50" style="width: 60%; height: auto;"/>

<!-- 
What if we can create the Github contribution table?
First of all we need to gather all inputs...

...and in Go, conveniently, a palette is a slice of colors 
-->
---

# 🎨 Palettes in Go

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

<img v-click src="/images/actual-gh-contributions.png" class="absolute top-40 right-15" style="width: 40%; height: auto;"/>
<arrow v-after x1="800" y1="300" x2="850" y2="250" color="#F00" width="2" arrowSize="1" />

<br/>

<v-clicks>

- $github\_input + color.Palette = contribution\_table$

<img  src="/images/generated-gh-contributions.png" class="absolute bottom-30 left-40" style="width: 65%; height: auto;"/>
</v-clicks>

---
layout: lblue-fact
---

Fun fact

<!-- Time for another fun fact -->

---
layout: fact
---

<v-click>

Did you know that a sketch of the Monalisa was made in Go?
</v-click>

<img src="/images/leonardo-da-vinci.jpg" class="absolute top-5 left-15" style="width: 15%; height: auto;"/>
<img src="/images/Mona_Lisa,_by_Leonardo_da_Vinci.jpg" class="absolute bottom-5 right-15" style="width: 10%; height: auto;"/>
<img v-click src="/images/monalisaPaintByNumber.png" class="absolute top-10 left-80" style="width: 38%; height: auto;"/>

<!-- This time we moving a bit closer, to none other than Leonardo da vinci

Did you know...?

If you ask me, I can't actually see any difference with the original -->

---
layout: lblue-fact
---

From setting pixel colors to drawing layers

<!-- So far we have used a lot:

```go
func (p *RGBA) Set(x, y int, c color.Color) // For RGBA
```

We are now moving to:

```go
func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)
``` -->

---

# 🗃️ Layering Images

A basic composition of two images

<v-click>

1. Create two images:
    - image 1 (`dst`) is a green rectangle
    - image 2 (`src`) is a smaller white rectangle

```go
func MakeLayer(r image.Rectangle, c color.Color) *image.RGBA {
  img := image.NewRGBA(r)
  for x := range r.Max.X {
    for y := range r.Max.Y {
      img.Set(x, y, c)
    }
  }
  return img
}

dst := MakeLayer(image.Rect(0, 0, 1024, 768), color.RGBA{G: 150, A: 255})
src := MakeLayer(image.Rect(0, 0, 800, 600), color.White)
```
</v-click>

<!-- We are going to see this method in 3 steps
1. Create the images
2. Draw one on top of the other
3. Encode to file
-->

---

# 🗃️ Layering Images

A basic composition of two images

2. Use `draw.Draw` to draw image 1 (`src`) onto image 2 (`dst`)

```go{all|4|5|6|7|8|all}
// dst := MakeLayer(image.Rect(0, 0, 1024, 768), color.RGBA{G: 150, A: 255})
// src := MakeLayer(image.Rect(0, 0, 800, 600), color.White)
draw.Draw(
  dst,   // the destination layer
  image.Rect(224, 168, dstR.Max.X, dstR.Max.Y),  // the area on dst where Draw can operate
  src,  // the source layer
  image.Point{224, 168}, // the starting point from which the source can be taken
  draw.Over, // the draw operation (draw src over dst)
)
```

<v-click>

3. Encode the result into a file

```go
f, _ := os.Create("white-in-green.png")
png.Encode(f, dst)
```
</v-click>

<img v-click="+1" src="/images/comp-green.png" class="absolute bottom-5 right-20" style="width: 30%; height: auto;"/>
<img v-click="+2" src="/images/comp-bounds.png" class="absolute bottom-5 right-20" style="width: 30%; height: auto;"/>
<img v-click="+4" src="/images/comp-src-point.png" class="absolute bottom-5 right-100" style="width: 20%; height: auto;"/>
<arrow v-click="+5" x1="450" y1="425" x2="665" y2="360" color="#F00" width="2" arrowSize="1" />
<arrow v-click="+5" x1="500" y1="475" x2="725" y2="400" color="#F00" width="2" arrowSize="1" />
<img v-click="+6" src="/images/comp-white-on-green.png" class="absolute bottom-5 right-20" style="width: 30%; height: auto;"/>
<arrow v-click="+6" x1="450" y1="425" x2="665" y2="360" color="#F00" width="2" arrowSize="1" />
<arrow v-click="+6" x1="500" y1="475" x2="725" y2="400" color="#F00" width="2" arrowSize="1" />

<!-- 
It makes the top left point of the src correspond to the top left point of the destination and it draws as much as it can and if it doesn't totally fill the space on the destination, what's left of the destination is left untouched
 -->
---

# 🗃️ Green and white composition's full code

```go{all|17-18|19|20-21|all}
package main
import (
  // ...
)

func MakeLayer(r image.Rectangle, c color.Color) *image.RGBA {
	img := image.NewRGBA(r)
	for x := range r.Max.X {
		for y := range r.Max.Y {
			img.Set(x, y, c)
		}
	}
	return img
}

func main() {
	dst := MakeLayer(image.Rect(0, 0, 1024, 768), color.RGBA{G: 150, A: 255})
	src := MakeLayer(image.Rect(0, 0, 800, 600), color.White)
	draw.Draw(dst, image.Rect(224, 168, dst.Rect.Max.X, dst.Rect.Max.Y), src, image.Point{224, 168}, draw.Over)
	f, _ := os.Create("comp-white-on-green.png")
	png.Encode(f, dst)
}
```

<img src="/images/comp-white-on-green.png" class="absolute top-40 right-25" style="width: 30%; height: auto;"/>

<!-- 
Again let's consolidate everything in one place: this is the full code that creates this green-white composition and inside the code we see the three steps together

1. Create the images
2. Draw one onto the other
4. Encode to a file

This technique opens up to nice applications such as
-->

---

# 🗃️ Creating a banner

<v-clicks>

1. Load a background image
2. Put a semi-transparent gray box inside
3. Add a nice gopher
4. Write some text
</v-clicks>

<v-click at=1>
````md magic-move{at:2}
```go
f, err := os.Open("golab-speakers.png")
// ...
// ...
// ...
baseImg, err := png.Decode(f)
```

```go
gray := image.NewUniform(color.RGBA{R: 150, G: 150, B: 150, A: 200})
// ...
// ...
// ...
draw.Draw(
  baseImg,
  image.Rect(110, 100, baseImg.Bounds().Dx()-110, baseImg.Bounds().Dy()-100), 
  gray, 
  image.Point{110, 100},
  draw.Over,
)
```

```go
f, err := os.Open("mcaciGopherizeMe.png")
// ...
gopherized, err := png.Decode(f)
// ...
draw.Draw(
  baseImg,
  image.Rect(130, 130, baseImg.Bounds().Dx()-130, baseImg.Bounds().Dy()-130), 
  gopherized, 
  image.Point{130, 130},
  draw.Over,
)
```

```go
import "github.com/golang/freetype"
// ...
ctx := freetype.NewContext()
// ... fill context information ...
ctx.DrawString("Colors, images and gifs:", fixed.P(1200, 400))
ctx.DrawString("bring on the fun with Go", fixed.P(1250, 550))
ctx.DrawString("by Michele Caci", fixed.P(1650, 1150))
```
````

</v-click>


<img v-click="+1" src="/images/golab-speakers.png" class="absolute top-45 right-25" style="width: 30%; height: auto;"/>
<img v-click="+2" src="/images/composition-gray.png" class="absolute top-45 right-25" style="width: 30%; height: auto;"/>
<img v-click="+3" src="/images/composition-gopher.png" class="absolute top-45 right-25" style="width: 30%; height: auto;"/>
<img v-click="+4" src="/images/composition.png" class="absolute top-45 right-25" style="width: 30%; height: auto;"/>

<!-- 
You can use anything as layers:
- images from files
- images created in Go
- and even text

And these steps are there only for an example but you can free your imagination at this point
-->

---

# 🗃️ How about some ASCII Art?

Printing to console

<v-click>

```go
package main

import figure "github.com/common-nighthawk/go-figure"

func main() {
  const figlet = "standard"
  const text = "Welcome to GoLab!"
  fig := figure.NewFigure(text, figlet, true)
  fig.Print()
}
```
</v-click>

<br/>

<v-click>

```bash
mcaci@mcaciLaptop:~/go/src/github.com/mcaci/golab25-ColorsGifsGo-slides/src/go/asciiart$ go run .
 __        __         _                                       _               ____           _               _       _
 \ \      / /   ___  | |   ___    ___    _ __ ___     ___   | |_    ___     / ___|   ___   | |       __ _  | |__   | |
  \ \ /\ / /   / _ \ | |  / __|  / _ \  | '_ ` _ \   / _ \  | __|  / _ \   | |  _   / _ \  | |      / _` | | '_ \  | |
   \ V  V /   |  __/ | | | (__  | (_) | | | | | | | |  __/  | |_  | (_) |  | |_| | | (_) | | |___  | (_| | | |_) | |_|
    \_/\_/     \___| |_|  \___|  \___/  |_| |_| |_|  \___|   \__|  \___/    \____|  \___/  |_____|  \__,_| |_.__/  (_)
```

</v-click>

<!-- We'll make use of an external library [`go-figure`](github.com/common-nighthawk/go-figure). Whose normal usage is this one -->

---

# 🗃️ How about some ASCII Art?

Drawing ASCII Art on an image

````md magic-move
```go
package main

import (
  figure "github.com/common-nighthawk/go-figure"

)

func main() {
  const figlet = "standard"
  const text = "Welcome to GoLab!"
  fig := figure.NewFigure(text, figlet, true)
  fig.Print()
}
```

```go
package main

import (
  figure "github.com/common-nighthawk/go-figure"

)

func main() {
  const figlet = "standard"
  const text = "Welcome to GoLab!"
  fig := figure.NewFigure(text, figlet, true)
  lines := fig.Slicify() // []string
}
```

```go
package main

import (
  figure "github.com/common-nighthawk/go-figure"
  "github.com/golang/freetype"
)

func main() {
  const figlet = "standard"
  const text = "Welcome to GoLab!"
  fig := figure.NewFigure(text, figlet, true)
  lines := fig.Slicify()
  // ...
  const height = 100
  ctx := freetype.NewContext()
  // ... fill context information ...
  for i := range lines {
    ctx.DrawString(lines[i], fixed.P(0, height * (i +  1)))
  }
}
```
````

<img v-click="+3" src="/images/asciiart.png" class="absolute top-55 right-15" style="width: 60%; height: auto;"/>
<img v-click="+4" src="/images/asciiart-plaid.png" class="absolute top-85 right-15" style="width: 60%; height: auto;"/>

<!-- 
If you want to use ascii art in a drawing just replace Print() with Slicify() and use the `freetype` package to write all of the strings in the retuned slice onto the background 
 -->

---
layout: lblue-fact
---

Fun fact

<!-- 
Time for another fun fact!
-->
---
layout: fact
---

<v-click>

Did you know that Pablo Picasso used Go to paint one of his famous cubist masterpiece?
</v-click>

<img src="/images/picasso-selfportrait.jpg" class="absolute top-5 left-15" style="width: 15%; height: auto;"/>
<img v-click src="/images/picasso-gopher.png" class="absolute top-10 left-80" style="width: 38%; height: auto;"/>

<!--
This is a self-portrait of Pablo Picasso, it was not made in Go but...

Did you know?

Which one do you think it is? Maybe les demoiselles d'Avignon? The Girl before a mirror? Guernica?

Well, it's none of them... it's the Gopher of course!

Just take some time to admire it, isn't it beautiful? 
 -->

---
layout: lblue-fact
---

GIFs, finally!

---

# 🎞️ What is a GIF in Go

<v-clicks>

```go
type GIF struct {
  Image []*image.Paletted // The successive images (frames).
  Delay []int             // The successive delay times, one per frame, in 100ths of a second.
  // ...
}
```

- `image.Paletted` is like `image.RBGA` constrained by a `color.Palette`
</v-clicks>

<img v-click src="/images/bgGradient.png" class="absolute bottom-10 left-25" style="width: 35%; height: auto;"/>
<img v-click src="/images/bgGradientP9.png" class="absolute bottom-10 right-25" style="width: 35%; height: auto;"/>

<!-- 
To understand the basics of a GIF in Go we can look at its type definition
-->

---

# 🎞️ How to create a GIF

2-frames GIF

1. Create two frames:
    - frame 1 (`frm1`) is a green rectangle
    - frame 2 (`frm2`) is a yellow rectangle

```go
func MakeFrame(c color.RGBA) *image.Paletted {
  r := image.Rect(0, 0, 1024, 768)
  frm := image.NewPaletted(r, palette.Plan9)
  draw.Draw(frm, r, image.NewUniform(c), image.Pt(0, 0), draw.Over)
  return frm
}

frm1 := MakeFrame(color.RGBA{G: 150, A: 255})
frm2 := MakeFrame(color.RGBA{G: 150, R: 150, A: 255})
```

<!-- 
three steps for creating a gif

1. Create the frames
2. Instantiate the gif variable
3. Encode to a file
 -->
---

# 🎞️ How to create a GIF

2-frames GIF

2. Instantiate a variable of type `gif.GIF` (`import "image/gif"`) and fill its fields

```go
g := &gif.GIF{
  Image: []*image.Paletted{frm1, frm2},
  Delay: []int{100, 100},
}
```

<v-click>

3. Encode the gif into a file

```go
f, _ := os.Create("my-first-gif.gif")
gif.EncodeAll(f, &g)
```
</v-click>

<img v-click src="/images/myFirst.gif" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>
<img v-click src="/images/thats-it.gif" class="absolute bottom-5 left-10" style="width: 55%; height: auto;"/>

---

# 🎞️ 2-frames GIF full code

```go{all}
func MakeFrame(c color.RGBA) *image.Paletted {
  r := image.Rect(0, 0, 1024, 768)
  frm := image.NewPaletted(r, palette.Plan9)
  draw.Draw(frm, r, image.NewUniform(c), image.Pt(0, 0), draw.Over)
  return frm
}

func main() {
  frm1 := MakeFrame(color.RGBA{G: 150, A: 255})
  frm2 := MakeFrame(color.RGBA{G: 150, R: 150, A: 255})
  g := &gif.GIF{
    Image: []*image.Paletted{frm1, frm2},
    Delay: []int{100, 100},
  }
  f, _ := os.Create("myFirst.gif")
  gif.EncodeAll(f, g)
}
```

<img src="/images/myFirst.gif" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>

<!-- 
Let's consolidate everything in one place: this is the full code that creates this green/yellow gif and inside the code we see the three steps together

1. Create the frames
2. Instanciate the gif variable
3. Encode to a file
 -->

---
layout: lblue-fact
---

And this is just the beginning

<img v-click src="/images/gradients.gif" class="absolute top-15 left-30" style="width: 30%; height: auto;"/>
<img v-click src="/images/timeflow.gif" class="absolute bottom-5 right-20" style="width: 30%; height: auto;"/>
<img v-click src="/images/first-cave.gif" class="absolute top-50 right-15" style="width: 40%; height: auto;"/>
<img v-click src="/images/game_of_life.gif" class="absolute top-50 left-25" style="width: 30%; height: auto;"/>
<img v-click src="/images/progressbar.gif" class="absolute bottom-10 left-80" style="width: 38%; height: auto;"/>
<img v-click src="/images/LunchBreak.gif" class="absolute top-25 right-25" style="width: 30%; height: auto;"/>

---
layout: lblue-fact
---

Fun fact

<!-- Time for the last fun fact of the day -->

---
layout: fact
---

Did you know that Picasso used Go also to create an early version of the party Gopher?

<img src="/images/picasso-selfportrait.jpg" class="absolute top-5 left-15" style="width: 15%; height: auto;"/>
<img v-click src="/images/picasso-gopher.gif" class="absolute top-10 left-80" style="width: 38%; height: auto;"/>

<!-- Here it is dancing its way out -->

---
layout: lblue-end
---

<div class="text-white font-size-10">
The end?
</div>

<!--
Now, is this the end?

I gave you the tools to start drawing images and creating GIFs with Go.

I hope you had fun while receiving them 

And I'm looking forward to see your beautiful works of art made with Go.

Discarded...
And if you are still wondering
OR
And remember, programming isn't just logic, it's also art
-->

---
layout: statement
hide: true
---

<div class="font-size-10">
Programming isn't just logic
</div>

<div v-click class="font-size-10">
it's also art!
</div>

---
layout: statement
---

<div class="font-size-10">
Could all of it have been done in other programming languages?
</div>

<img v-click src="/images/nono.gif" class="absolute top-10 left-37" style="width: 70%; height: auto;"/>

<!-- papaaaa'.... papaaaaa'... no no no no no no no!  -->

---
layout: lblue-end
---

<img src="/images/thank-you.gif" class="absolute top-10 left-5" style="width: 95%; height: auto;"/>

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
    <a href="https://www.linkedin.com/in/michele-caci-47770132/" target="_blank" alt="Michele's Linkedin" title="Michele's Linkedin"
      class="text-xl slidev-icon-btn opacity-50 !border-none !hover:text-white">
      <carbon-logo-linkedin />
    </a>
  </div>
</div>
<img src="/images/michelecaciQR.jpeg" class="absolute bottom-5 right-5 text-right" style="width: 20%; height: auto;"/>
