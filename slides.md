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

<img src="/images/actual-cave-with-sand.png" class="absolute top-25 left-20" style="width: 45%; height: auto;"/>
<img v-click="+1" src="/images/forest.png" class="absolute bottom-5 left-10" style="width: 23%; height: auto;"/>
<img v-click="+2" src="/images/bgGradient.png" class="absolute bottom-25 left-85" style="width: 20%; height: auto;"/>
<img v-click="+3" src="/images/LunchBreak.gif" class="absolute bottom-8 left-67" style="width: 30%; height: auto;"/>
<img v-click="+4" src="/images/game_of_life.gif" class="absolute bottom-10 right-5" style="width: 40%; height: auto;"/>

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

# 🖼️ Our First Image

The basics: step by step

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

</v-click>

<img v-click="1" src="/images/bounds.png" class="absolute top-18 right-10" style="width: 28%; height: auto;"/>
<img v-click="5" src="/images/green.png" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>

---

# 🖼️ Our First Image

Full code

```go{all|11|12|13-17|18-19|all}
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

<img v-click=5 src="/images/green.png" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>

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
var b, g uint8
switch (x/32 + y/32) % 2 {
case 0:
  b = 150
default:
  g = 150
}
img.Set(x, y, color.RGBA{B: b, G: g, A: 255})
```

```go
b := float64(x) / float64(r.Max.X) * 255
g := float64(y) / float64(r.Max.Y) * 255
img.Set(x, y, color.RGBA{B: uint8(b), G: uint8(g), A: 255})
```

```go
func paintWith(/***/) color.Color {
  /***/
}

img.Set(x, y, paintWith(/***/))
```
````

<v-click at=4>

````md magic-move
```go
func paintWith(x, y int, r image.Rectangle) color.Color {
  b := float64(x) / float64(r.Max.X) * 255
  g := float64(y) / float64(r.Max.Y) * 255
  return color.RGBA{B: uint8(b), G: uint8(g), A: 255}
}
```

```go
func paintWith(x, y int) color.Color {
  r := (1 + math.Cos(float64(x)/10)) * 255
  g := (1 + math.Sin(float64(y)/10)) * 255
  return color.RGBA{
    R: uint8(r),
    G: uint8(g),
    A: 255,
  }
}
```

```go
func paintWith(t time.Time) color.Color {
  nano := uint32(t.UnixNano())
  r := uint8((nano >> 16) & 0xFF)
  g := uint8((nano >> 8) & 0xFF)
  b := uint8(nano & 0xFF)
  return color.RGBA{
    R: uint8(r),
    G: uint8(g),
    B: uint8(b),
    A: 255,
  }
}
```
````
</v-click>

<img src="/images/green.png" class="absolute top-18 right-10" style="width: 28%; height: auto;"/>
<img v-click="+1" src="/images/blue.png" class="absolute top-35 right-20" style="width: 30%; height: auto;"/>
<img v-click="+2" src="/images/bgCheckerboard.png" class="absolute top-50 right-35" style="width: 30%; height: auto;"/>
<img v-click="+3" src="/images/bgGradient.png" class="absolute top-65 right-50" style="width: 28%; height: auto;"/>
<!-- <img v-click="+4" v-click.hide src="/images/bgGradient.png" class="absolute top-50 right-25" style="width: 50%; height: auto;"/> -->
<img v-click="+5" src="/images/plaid.png" class="absolute top-50 right-25" style="width: 40%; height: auto;"/>
<img v-click="+6" src="/images/timeflow.png" class="absolute top-50 right-25" style="width: 45%; height: auto;"/>
<img v-click="+7" src="/images/timeflow.gif" class="absolute top-21 right-5" style="width: 50%; height: auto;"/>

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

We decide what color to set for each pixels

<!-- And this is the basics of creating an image, now setting pixels is fun and imaginative, but we can use some tools to improve our outcome -->

---

# 🧮 From Numbers and Maps

Images from input

```
100021112110202312022010330204312040000111012143445142221414240220240442332040010320133120230011020
111110201332323210211143123214321343332124211413115514155115511033421001222101204330001300333011010
222100123300231230122203432310224441551434231352112532354252244334410042212441233243220102033110020
212210311310102334321243104133012535144151532555155341325352512544453402400411340202133231000102020
112002102322001003304233302040124544411225523533534235113353522233535550432232202401221022110311110
112102131331201432320312233124434232544144233241123334112232531521542551332434224211234133132330300
201212231001314001232020011521543445152251413225334151155332443144315433444321403022301110231133320
222200212012234021440323522251222114255435321245463652634454445441313331334221403314121340322201323
011101313030124034011053132411352313135432265223564644563435541314314453511355132441242220212301210
330120320004414224001225342325133541125622665563263444434232635442123423514224513321324324321221201
232201002203142013213541415521331136252365342544552365442625322555515234315251334143110114433011330
231032032133203330135134524113312353466322543643543364665636365344645123214255433423420443411102113
323210134414131324325554245224653435365453364456243243426526432462344551442433224554110342343103002
232211304402432454533435252145662665532355636326423563423433555236565455345421453313301243030201033
301334404444113432444154526622565534622453452465656373363346236352264432242255511441531301201430310
021031402014322514331554342423522244643245347463537353546733344353665264632651255422445201241324231
311341433423143525332424243226653324554455444337664375746566644363342254524365544454512533412221003
100141044421132331524146432626626222746377344775337675474374667736346526345525531112515524220332342
103100424013355243423532363653563355763646667345657754366374646535626642326366655331552511413232033
010411402325242212332662662566365363733444434344454736635435565544436445432562223352255125432214413
321313023231141232536442263534773775537454346454663656656365364557334363362446463531553445101214414
240014233445325423455462553425364653653645333436556455474746633767756644256625645261111551134401102
141243103121441452453322334274356634453674654475774776665766755335473637655355324552442441555404422
322032122454423446326624652657555775743558467874777864887767556333635645342466232542452344223021030
341012035442545266666542455563533577735885866784868684486868864657367574557543634634445432352101333
020332133545225646653645453347737756754567886778577468644476777475537454454734354466644222142343014
443312425144123235564325363434377486685588484488457888748854756668646745465366635445226143535350403
401011343544455465645234773757567857786858744667847887678865684557657673336765353252254124535144301
231023111541232646654364676655554478448887447788656586858686874555544756355355425663634445444231041
440414511232163335524763454743366676656778445597669969755446578885785477433565366464462411314121213
423133115311242232226334356676577546867445799868866958556677644564858573537455653355226445155233301
203312245131462222554746646746567788757698759759588997656957785788856475653644633464554355125443503
232515114525335654457543443555565868786576595878575876598966858547666846455373765745424444124453521
222335423146245553644554757748846867467979895568976578965955877985885478853344575356225432342255433
101542414345642523353567565774687668996586857857667878786879985798875768547745643456564266313332224
301144344462663436337536345845856858757875777768966787979577767876448646654465333646326566332352143
421524351232465434654576565767446575666755968687986969675676875569846488676754753357442332442433411
413325242333454424656334665558754657887868888788696877986559786689587455864465343775662452632451433
154244543442464674633346657856868997877985698877867876966997859957879646664844343364323322443522435
413112322322235675547345547474859589658598796866669766699978589785866454757775477664523433333434442
424333122434626456777635768578686868759696687688967977688996678795685586557475554454652235566333123
455211225265424477347444685544699887876977897876978778998988685797998647458856773634522525443114455
215254146366363747473755887445877868656789966766676898799688786567569855574684673733533243626145254
232214336435344464337365855745665767968888669969788898668797767777899554757865644635655552352235241
431445513242642634373387444466656585686776979787998778778876688797776987575785547434636436252135425
341254135453625764555454784558697758996768789898879887979776988667895866467544537467352465554231341
353511356522232754333665846875756885597979987979797998788867898755966966747777863335465653622312521
342131425455252547374557454676889585769899979877988878787678998666857676478654857544553344324545533
423514554622265364535758656757779586677868668897789789888887987885766775484647466733766544444641311
443223335562467674644684785459675797877798787899879888997799677698799557567657744564575455564221324
434334525556625335565445744449679766879887898799888779777969797666677897876666577353534665444633525
254452226436653735365676845666667659869878998777989789787868689999995856864675465356773254455643433
541343235664626453454454858567778559888976667877797788989989978879887596787865646354356245342421323
522412226622325753744545556855766985687768979797897799987698989897779696765756544634672555453353213
412522254444564635737677845545956698889698677889799979989779898956976777758878537634642425526241551
454222136263222765655775647779568866577886976888987989877869678975585697848476455557545354635351442
135121552535465363736474668878979599679996967978778988767779668655599557568578867666446332363541241
443121425555335774663648754687768668779998779886677866797868896598977975744577757466744252426313232
413544235544623474445636785886865856978977898886979867667687979665559785856884354355373456533433444
131513455323433453365575445468776665975667767679666666676699795889878855677784375677726235226322425
434442215625364456464667564866697986985698768698867797888969976976557974477456645553454563433515241
453531215644623547465376468645845588955556796689667796966679798897589845547855353757662556342441532
332355435635654445755366546668857777995897877866676989779679698556675648665863767454324443235151225
213534313254232564576535444544764898568875856676999867867657557989768547684445736355423244324413543
124522522235244346567747584544557769579796878589766996768788685995875444645835373437362355652144245
442414332462645464457376647585848848965577665855857885698768798696575546674475573646566222314114445
112315233256652343737667637858458576968777559868597579967759677588458846648655454334245245345554224
130224422112546654465464363764457446579885795857586579789779888566478866866743355335226264212535322
243155553555342624676436455375874765489687699588798796759766675675567774646753634355346454424431253
313442313423322535325375377344687658457576699776856776787599774474677575337743334422233423431421211
234051131411522623337474757677857685464857596667557579666787555744487574544673364663643435211255240
434314134422524322443747337633658574444468845599797888867456487676744754336473656352235344411242110
004101145215335222325475546343745855484848877575546845867885845455746457433457436466565332352124123
233441434144112564624645544646557557655667845446677675588784866645475343643574552243656523444534030
121224143555135555564544776734555455674556575754675486888465845458835743367333425354526331434444031
421244422344321446364634257366433445747885577774664546565647578767333446746356442243433411415210430
134214243522231144462465353566577763536746478855876848468846687644663756364224224464635451234214111
141132203421215213235464224663666433534678865566554887458784575743373335336464662256425123321434014
213231333411351412344352524277763665543774674775545875744754576367464737663564262663112544415131113
310320312444245332353443555344534456457456364753657473447435543364646364264425336635335354223233231
314242332123434345266456226225367677737356563336765735555554546764556543652364555632435415411342334
220144410231433531546232424625336566373376474676454647665666437455766323355345563421345422444413112
331442434421142252212466563352266545334633374637735376643435675553544366564244525132151223011342202
203334211344322333542433654434655335633734475664756747766375655456225622225333343324215521200001401
133001044130025211433335626552525266367334377736466473364444543236446653253662222322415412004404203
102331031201225135443424246553365346664536353367434546473743334356454654552233145534345231113214102
212230010342131553532451526322255666663223366332674653633446456525262543246511312215123143100414322
232101044243441133154151523542336655464563325444436323646534526542245542613552531321441404211220231
230210113134231442312133523231426423645562233522662322645556325555424565225553133115020231142422303
112122132434123424454552435541116536436354343364366253452343245526436543253435122423041022130131032
303333011024122011435455241134321332555664352652523256255643224225644425233221344014220444400330023
103302231310301313234354141312323434532246526636322542364336342533331445554253314030214223423211003
110123122321424301403345325425115544551445462636362336444425435312233541552212143231043320211100213
202211312210402020033410432324431253144325523235463625332451131521414414333322211144344230221323102
210100112030114430323324002554125435321123334342551235451453442512153411112423411424123322300320102
202220203220032000004314434423331253245534454123425114423553133414151232240304414101402311330002020
012211131001201000134433142021511253314413353152553443443342425435421132024242024214322330313300110
200011111022120322143342040443334552235435211511231451452222252222113214233002413424112001211130112
202000213331311222142203242110321123234552545321425122145424451321304200332123042013120323022211012
```

<img v-click src="/images/forest.png" class="absolute top-35 left-55" style="width: 40%; height: auto;"/>
<img v-click src="/images/heatamp.png" class="absolute top-35 left-55" style="width: 40%; height: auto;"/>

---

# 🧮 From Numbers and Maps

Images from input

```
abaccccccccccccccaaaccccaaaaaaaaaaaaaccccccaacccccccccccccccccccccccccccccaaaaaa
abaaccaacccccccccaaaaaccccaaaaaaaaaaaaaccccaaacccccccccccccccccccccccccccccaaaaa
abaaccaaacccccccaaaaaacccaaaaaaaaaaaaaacaaaaaaaaccccccccaacccccccccccccccccccaaa
abcaaaaaaaacccccaaaaaacccaaaaaaaaaaaaaacaaaaaaaacccccccaaaacccccccccccccccccaaaa
abcaaaaaaaaccccccaaaaaccaaaaaaaaccaaaaaccaaaaaaccccccccaaaaccaaaccccccccccccaaac
abccaaaaaacccccccaaaaccaaaaaaaaaacaaaacccaaaaaacccccccccakkaaaaaacccccccccccaacc
abccaaaaaacccccccccccccaaaaaaaaaaccccccccaaaaaaccccccckkkkkkkaaacccccccccccccccc
abccaaaaaaccccccccccccccccaaaaaaaaaccccccaacaaacccccckkkkkkkkkaccccccaccaaaccccc
abccaacaaacccccaaccccccccaaacacaaaacaaccccccccccccccakkkoppkkkkicccccaaaaaaccccc
abccccccccccccaaaccccccccaacccccaaaaaaccccccccccccccjkkooppppkiicccccccaaaaccccc
abccccccccccaaaaaaaaccccccccccaaaaaaaccccccccccccccjjjooopppppiiiicccccccaaacccc
abaaacccccccaaaaaaaacccccccaacaaaaaaccccccccccccccjjjjooouuppppiiiiiicccccaacccc
abaaaccccccccaaaaaaccccccccaaaccaaaaacccccccccccjjjjjooouuuupppiiiiiiiiccccacccc
abaaaaaacccccaaaaaacccccaaaaaaaaaacaaaccccccccjjjjjjooouuuuuupppppiiiiiicccccccc
abaaaaaacccccaaaaaacccccaaaaaaaaaacccccccccccjjjjjooooouuxxuupppppqqqijjjccccccc
abaaaacccccaaaaccaaccccccaaaaaaccccccccccccciijjnooooouuuxxxuuupqqqqqqjjjdddcccc
abaaaaaccaaaaaaccacccccccaaaaaaccccccccccaaiiiinnootttuuxxxxuuvvvvvqqqjjjdddcccc
abaaaaaccaaaaaacaaaccaaccaaaaaaccccccccccaaiiinnnntttttuxxxxxvvvvvvqqqjjjdddcccc
abaaccacccaaaaacaaaaaaaccaaccaaccccccccccaaiiinnnttttxxxxxxxyyyyyvvqqqjjjdddcccc
abcccccccaaaaacccaaaaaaccccccaaaaacccccccaaiiinnntttxxxxxxxyyyyyvvvqqqjjjddccccc
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
abccccaaaaaaaaaccccccccccccaaacccccccccaaaaacccchhhmmmmmlllllllllkkkkeeeeaaccccc
abcccccccaaaaaaccccccccccaacccccccaaccaaacaacccchhhmmmmmlllgfflllkkffeeeaaaacccc
abccccccaaaaaaaccccccccccaaaaaaaaaaaaaccccaacccchhhggmmmggggffffffffffeaaaaacccc
abccaacccaacccaaaaccaccccaaaaaaaaaaaaacccccccccccgggggggggggffffffffffaacccccccc
abaaaaccaaaccccaaaaaaccccaaaaaacaaaaaaccccccccccccgggggggggaaaaccffccccccccccccc
abaaaacccccccccaaaaaaccaaaaaaaaaaaaaacccccccccccccccgggaaaaaaaacccccccccccccccca
abaaaaacccccccaaaaaaaccaaaaaaaaaaaaaacccccccccccccccccaaacccaaaaccccccccccccccaa
abaaaaacaaaaccaaaaaaaacaaaaaaaaaaaccccccccccccccccccccaaaccccaaaccccccccccaaacaa
abaaaaacaaaaccaaaaaaaaaaaaaaaaaaacccccccccccccccccccccccccccccccccccccccccaaaaaa
abaaacccaaaaccccaaaccccaaaaaaaaaaacccccccccccccccccccccccccccccccccccccccccaaaaa
```

<img v-click src="/images/hill.png" class="absolute top-35 left-35" style="width: 70%; height: auto;"/>

---

# 🧮 From Numbers and Maps

Because setting pixels is fun and imaginative, but we can use some tools to improve our images

- Look at this input:

> 525,119 -> 525,122 -> 523,122 -> 523,125 -> 529,125 -> 529,122 -> 528,122 -> 528,119

> 497,69 -> 497,73 -> 489,73 -> 489,78 -> 504,78 -> 504,73 -> 501,73 -> 501,69

> 480,38 -> 480,31 -> 480,38 -> 482,38 -> 482,35 -> 482,38 -> 484,38 -> 484,35 -> 484,38 -> 486,38 -> 486,28 -> 486,38 -> 488,38 -> 488,36 -> 488,38

> 480,38 -> 480,31 -> 480,38 -> 482,38 -> 482,35 -> 482,38 -> 484,38 -> 484,35 -> 484,38 -> 486,38 -> 486,28 -> 486,38 -> 488,38 -> 488,36 -> 488,38

> \[...\]

<v-clicks>

These are `(x, y)` coordinates in a 2D space representing walls.

From the first line:

</v-clicks>

<v-clicks depth="2">

- `525,119 -> 525,122` is a vertical wall between those two coordinates.
- `525,122 -> 523,122` is a horizontal wall between those coordinates.
</v-clicks>

<img v-click src="/images/cave.png" class="absolute top-45 left-35" style="width: 70%; height: auto;"/>

---

# 🧮 From Numbers and Maps

There's more

This input comes from the advent of code [2022/14](https://adventofcode.com/2022/day/14).

We have a cave and what's more, sand is coming down from it! (from coordinate `500,0`)

<img v-click src="/images/cave.png" class="absolute top-45 left-35" style="width: 70%; height: auto;"/>
<img v-click src="/images/first-cave.gif" class="absolute top-45 left-35" style="width: 70%; height: auto;"/>
<img v-click src="/images/cave-with-sand.png" class="absolute top-45 left-35" style="width: 70%; height: auto;"/>
<img v-click src="/images/actual-cave-with-sand.png" class="absolute top-45 left-35" style="width: 70%; height: auto;"/>

---

# 🧮 From Numbers and Maps

- Visualizing matrices as forests, terrain, or heatmaps
- Mapping values to colors
- Example: turning a grid of numbers into a landscape

<img v-click="[1, '+1']" src="/images/forest.png" class="absolute top-50 right-25" style="width: 30%; height: auto;"/>
<img v-click="[2, '+1']" src="/images/hill.png" class="absolute top-50 right-25" style="width: 35%; height: auto;"/>
<img v-click="3" src="/images/cave-with-sand.png" class="absolute top-50 right-25" style="width: 40%; height: auto;"/>


<!--
 
 Add joke about the matrix movie and or bitmaps 
 -->

---

# Making Art with Go

Piet Go-ndrian

Easiest way: https://www.google.com/search?q=piet+mondrian&client=ms-android-google&sca_esv=e34a8a4f92e4e401&udm=2&biw=411&bih=785&sxsrf=AE3TifMXRufYHwbrEiyrTi5GI1jaf5AZUg%3A1755465000360&ei=KEWiaPXiFa6HkdUPmJ7RoAg&gs_ssp=eJzj4tTP1TdILspLMjVg9OItyEwtUcjNz0spykzMAwBr0Qib&oq=piet+&gs_lp=EhJtb2JpbGUtZ3dzLXdpei1pbWciBXBpZXQgKgIIADIFEC4YgAQyBRAAGIAEMgUQLhiABDIFEC4YgAQyBRAAGIAESMU2UKwSWL4mcAJ4AJABAJgBgQGgAbkDqgEDNC4xuAEByAEA-AEBmAIHoALeA6gCBcICChAjGCcYyQIY6gLCAgcQIxgnGMkCwgIKEC4YgAQYQxiKBcICChAAGIAEGEMYigXCAgsQLhiABBjHARivAZgDBJIHAzYuMaAHhC-yBwM0LjG4B9YDwgcDMi03yAch&sclient=mobile-gws-wiz-img

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