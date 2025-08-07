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
---

# Let's Put Some Color to Our Programming with Go
Unleashing creativity with pixels, palettes, and Go!

---

# ⚠️ Disclaimer

This talk contains:
- Images 🖼️
- GIFs 🎞️
- Pixels 🧮
- And a lot of Go 🐹

All visuals are created with Go (yes, even the silly ones).
Code will be shared — creativity is encouraged!

---

# 🧠 What do these have in common?

<!-- Show a sequence of 3 images or gifs created with Go -->

They were all made with Go — and mostly with the standard library!

---

# 🎨 Challenge Accepted

I tried to create **all** the images in this presentation using Go.

Why? Because:
- It’s fun
- It’s weirdly satisfying
- And it shows what Go can do beyond servers and CLIs


---

## 👋 Welcome!
- Who I am
- Why I love Go
- What this talk is about: creativity, color, and fun!

---

## 🎨 Why Color in Code?
- Programming isn't just logic—it's also art
- Go is fast, simple, and surprisingly good at graphics
- We'll use (almost) only the standard library!

---

## 🌈 Creating Palettes and Images
- How to define and use colors in Go
- Building palettes
- Drawing basic shapes and rectangles
---

## 🧮 From Numbers to Maps
- Visualizing matrices as forests, terrain, or heatmaps
- Mapping values to colors
- Example: turning a grid of numbers into a landscape

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
layout: end
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