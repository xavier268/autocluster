# autocluster

[![Go Reference](https://pkg.go.dev/badge/github.com/xavier268/autocluster.svg)](https://pkg.go.dev/github.com/xavier268/autocluster) [![Go Report Card](https://goreportcard.com/badge/github.com/xavier268/autocluster)](https://goreportcard.com/report/github.com/xavier268/autocluster)

Automatic, hierarchical clustering of any files. Pure go.

## What is this project for ...

Imagine you have a bunch of files in a folder (within sub-folders, of course) and you want to put some order here. You want an easy way to know what file is *similar* to another one without going through all of them ?

This project is for you !

## Example

Let's see how the current github project folder would do ...

```
> go run ./... .

```

Update in progress  **to do**

## It's all about distance ...

You may want to have a quick look at the paper attached below.

The goal is to construct the distance matrix, by measuring a *distance* from one file to another.
First, because we care only about content, and not format, we do everything we can to extract text content from the file, to make sure we compare actual content and not format. For instance, in the test folder, you will find a couple of files with almost identical content : a word file, an excel file, and html file, a zipped file, ... A naïve comparison of these files would fail, and they would be put in far away clsuters, because of the significant difference in format. Because we first extract text before computing the distance, the results are much better. You can see in the clustering above that these 4 test files are aggregated together naturally, showing reasonnably small distance between them - these are files 20 to 23 above.

Okay, we get the content that matters from each file, then what ?

Then, we use the **gzip distance** as described in that paper : https://github.com/xavier268/autocluster/blob/main/2023.findings-acl.426.pdf . The idea is that if two *contents* are similar, then gzip can compress them both more efficiently together than separately. Simple, no ? It turns out to be surprisingly performant !

And now, we have a matrix of distance between actual files contents.

We need to define distance between *clusters*, called a *linkage distance*. There are many different way to do that (see https://en.wikipedia.org/wiki/Hierarchical_clustering , for commonly used linkage ). Various choices are available as an option on the cli.

And that's it ! We can now navigate our data, look for natural grouping of our files, without actually having to review the whole stuff ...

And because all computations are cached, it may take some time the first time, but will be much faster we you do it again.

## Some more details and potential improvements 

* Its a pure go project - no CGo, not external librairies, only standard Go libs ! And no licensing hassle ...
* Caching has been added to avoid recalculation of file-to-file distance. We use a map from a pair of digest for file contents to a float64 distance value. Space is very reasonnable, currently, my cache contains 210k values and weight less than 21MB.
* Currently, optimized file format are excel, word, any textual format, zip, zlib, gzip ... probably other as we move forward.
    * pdf is a challenge, though, at this stage. *Unidoc/Unipdf* requires a subscription and a licence, and I ruled that out. Moreover, often pdf would contain scanned image, and will in any case required an OCR pass before content can be effectively processed. 
    Do you want to try and address this challenge ?

* Project is structured in 2 packages :
    * *distance* is only taking care of the computation of distance between single elements. This package does all the file distance processing, text extraction, etc ... A *distance* must simply follow a basic interface, accepting two int ids, and returning a float64.
    * *cluster* does the actual grouping and processing, when passed a *distance*. It defines how a *linkage distance* is used to measure *distance* between *clusters*.
* not much optimization was done beyond the obvious for performance. Clarity was prefered over perfomance. There might be room to improve here also, but I did not feel the need ...

See detailled documentation on https://pkg.go.dev/
