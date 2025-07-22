---
# You can also start simply with 'default'
theme: seriph
# random image from a curated Unsplash collection by Anthony
# like them? see https://unsplash.com/collections/94734566/slidev
background: https://cover.sli.dev
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
transition: fade-out
backgroundSize: 80%
---

# Disclaimer

<!-- Import go fact layouts -->

This talk will contains images and gifs, lots of them

Or

What do this, this and this have in common? They have all been created with Go
<!-- Show sequence of images and gifs meanwhile -->

Also try to create all of the images of the slides with Go and share the code for them

---
transition: fade-out
backgroundSize: 80%
---

# About myself

Change a bit the layout of the slide, use more images and gifs and maybe a photo of last year's golab

<!--  Old description:

I’m a Gopher since 2018
- Practised via the [Tour of Go](https://go.dev/tour/) and [exercism](https://exercism.org/)

I work in Amadeus where I develop and operate Splunk for logging in the cloud

I speak at conferences, about Go and cloud topics

I previously worked in Java, Scala and C++

Besides programming I enjoy swimming, cooking and learning languages
- GOのワークショップへようこそ！ -->

---
transition: fade-out
layout: fact
class: 'text-white bg-#00ADD8 font-size-10'
---

Go is a simple language

---
transition: fade-out
backgroundSize: 80%
---
<!-- 
layout: image-right
image: /images/Gophers1.jpeg 
-->

# Let's Go colors

1 slide to present colors in the standard library.

Can I use colors here? Too early maybe.

Still this is one of the basic components (try not to sound too serious, let's bring the fun with Go!)

---
transition: fade-out
---

# Let's Go images

Everything starts from our rectangle friend

Can we encode a more or less nice Gopher? How about turning it to a svg? (svg encoding is not stdlib)

My two ways to go for images

#1 -> Setting pixels (from file or whatever you want)
#2 -> Composing existing images

References to the matrix

Examples from advent of code?
<v-click>

````md magic-move {lines: true}
```go {all|2|2,6}
for i := range img {
  for j := range img[i] {
    mat[i][j] = color.RGBA{}
  }
}
```

```go {2|all}
// no time now: TBA
```
````
</v-click>

---
transition: fade-out
layout: lblue-fact
---

We keep a transition here
<!-- import lblue-fact for this to work -->

---
transition: fade-out
---

# How about some ascii art?

This section needs an external library, I said almost all stdlib but that's not so bad after all (frowny gopher here? could this talk be a conversation with a gopher?)

Here there can be some live coding (but also in general, to be seen)

It might be funny to show the transition between what the first try looks like and how to properly draw ascii art

---
transition: fade-out
layout: lblue-fact
---

Gem #2: A human cares about the error message
<v-click>
<div class="font-size-8">a program cares about what kind of error it is</div>
</v-click>

We keep a transition here
<!-- import lblue-fact for this to work -->

---
transition: fade-out
---

# Give me gifs

Quick explanation of gifs (frames + rate and the encoder) then show a selection of them to show what it means

---
transition: fade-out
layout: image-right
image: /images/Gophers10.jpeg
backgroundSize: 80%
---

# Conclusions

Simple: you didn't know? Now you know (jk, TBA)

<!-- ---
KEEP A FACT OR NOT?

layout: fact
transition: fade-out
class: "font-size-7.8"
---

And making our code simpler is how we step up our Go game! -->

---
layout: lblue-end
transition: fade-out
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