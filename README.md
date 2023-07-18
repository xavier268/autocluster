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


List of files to process :

  id    full path name
   0    C:\Users\xavie\Desktop\autocluster\2023.findings-acl.426.pdf
   1    C:\Users\xavie\Desktop\autocluster\LICENSE
   2    C:\Users\xavie\Desktop\autocluster\README.md
   3    C:\Users\xavie\Desktop\autocluster\cluster\cluster.go
   4    C:\Users\xavie\Desktop\autocluster\cluster\cluster_test.go
   5    C:\Users\xavie\Desktop\autocluster\cluster\linkage.go
   6    C:\Users\xavie\Desktop\autocluster\cluster\medoid.go
   7    C:\Users\xavie\Desktop\autocluster\cluster\version.go
   8    C:\Users\xavie\Desktop\autocluster\cmd\clusterize\main.go
   9    C:\Users\xavie\Desktop\autocluster\distance\cache.go
  10    C:\Users\xavie\Desktop\autocluster\distance\cache_test.go
  11    C:\Users\xavie\Desktop\autocluster\distance\euclid.go
  12    C:\Users\xavie\Desktop\autocluster\distance\file.go
  13    C:\Users\xavie\Desktop\autocluster\distance\file_test.go
  14    C:\Users\xavie\Desktop\autocluster\distance\matrix.go
  15    C:\Users\xavie\Desktop\autocluster\distance\matrix_test.go
  16    C:\Users\xavie\Desktop\autocluster\distance\string.go
  17    C:\Users\xavie\Desktop\autocluster\distance\version.go
  18    C:\Users\xavie\Desktop\autocluster\go.mod
  19    C:\Users\xavie\Desktop\autocluster\go.sum
  20    C:\Users\xavie\Desktop\autocluster\testFiles\test.docx
  21    C:\Users\xavie\Desktop\autocluster\testFiles\test.html
  22    C:\Users\xavie\Desktop\autocluster\testFiles\test.html.zip
  23    C:\Users\xavie\Desktop\autocluster\testFiles\test.xlsx

cache loaded from C:\Users\xavie\AppData\Local\Temp\fileDistance.cache (209463 values)

Computing distance matrix : 24/24
cache saved to C:\Users\xavie\AppData\Local\Temp\fileDistance.cache ( 209486 values)

Distance matrix (truncated after 10 values) :
               0               1               2               3               4               5               6               7               8               9              10        [...23]
    0   0.000000        1.000499        1.000129        1.001316        1.000567        1.000154        1.000392        0.999920        1.000755        1.000869        1.000105
    1   1.000499        0.000000        0.915625        0.990047        0.964246        0.939063        0.964063        0.951562        0.958254        0.978317        0.928125
    2   1.000129        0.915625        0.000000        0.951283        0.903911        0.818859        0.896610        0.863971        0.917457        0.944493        0.818452
    3   1.001316        0.990047        0.951283        0.000000        0.897328        0.920901        0.887376        0.978523        0.893138        0.915663        0.948140
    4   1.000567        0.964246        0.903911        0.897328        0.000000        0.877095        0.850279        0.954190        0.809298        0.913270        0.871508
    5   1.000154        0.939063        0.818859        0.920901        0.877095        0.000000        0.781356        0.895782        0.896584        0.919341        0.784119
    6   1.000392        0.964063        0.896610        0.887376        0.850279        0.781356        0.000000        0.933898        0.880455        0.908933        0.825424
    7   0.999920        0.951562        0.863971        0.978523        0.954190        0.895782        0.933898        0.000000        0.957306        0.967042        0.883929
    8   1.000755        0.958254        0.917457        0.893138        0.809298        0.896584        0.880455        0.957306        0.000000        0.854293        0.858634
    9   1.000869        0.978317        0.944493        0.915663        0.913270        0.919341        0.908933        0.967042        0.854293        0.000000        0.874241
   10   1.000105        0.928125        0.818452        0.948140        0.871508        0.784119        0.825424        0.883929        0.858634        0.874241        0.000000
[...23]


Computing clusters 24/24

Annotated dendrogramme of clusters :

(cluster content ....)          ( level / link distance )
+---[0 12 9 3 4 8 5 6 16 14 18 19 7 17 2 11 15 10 13 1 23 20 21 22]     (8 / 1.032926)
   +---[0 12 9 3 4 8 5 6 16 14 18 19 7 17 2 11 15 10 13 1]      (7 / 1.001316)
   |  +---[0] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -   C:\Users\xavie\Desktop\autocluster\2023.findings-acl.426.pdf
   |  +---[12 9 3 4 8 5 6 16 14 18 19 7 17 2 11 15 10 13 1]     (6 / 0.990047)
   |     +---[12 9 3 4 8 5 6 16 14]     (4 / 0.939560)
   |     |  +---[12 9]  (1 / 0.860806)
   |     |  |  +---[12] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  C:\Users\xavie\Desktop\autocluster\distance\file.go
   |     |  |  +---[9] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -   C:\Users\xavie\Desktop\autocluster\distance\cache.go
   |     |  |
   |     |  +---[3 4 8 5 6 16 14]       (3 / 0.925092)
   |     |     +---[3 4 8]      (2 / 0.897328)
   |     |     |  +---[3] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -   C:\Users\xavie\Desktop\autocluster\cluster\cluster.go
   |     |     |  +---[4 8]     (1 / 0.809298)
   |     |     |     +---[4] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -   C:\Users\xavie\Desktop\autocluster\cluster\cluster_test.go
   |     |     |     +---[8] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -   C:\Users\xavie\Desktop\autocluster\cmd\clusterize\main.go
   |     |     |
   |     |     |
   |     |     +---[5 6 16 14]  (2 / 0.859249)
   |     |        +---[5 6]     (1 / 0.781356)
   |     |        |  +---[5] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -   C:\Users\xavie\Desktop\autocluster\cluster\linkage.go
   |     |        |  +---[6] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -   C:\Users\xavie\Desktop\autocluster\cluster\medoid.go
   |     |        |
   |     |        +---[16 14]   (1 / 0.809651)
   |     |           +---[16] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  C:\Users\xavie\Desktop\autocluster\distance\string.go
   |     |           +---[14] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  C:\Users\xavie\Desktop\autocluster\distance\matrix.go
   |     |
   |     |
   |     |
   |     |
   |     +---[18 19 7 17 2 11 15 10 13 1]       (5 / 0.956250)
   |        +---[18 19 7 17 2 11 15 10 13]      (4 / 0.897196)
   |        |  +---[18 19 7 17] (2 / 0.785714)
   |        |  |  +---[18 19]   (1 / 0.662338)
   |        |  |  |  +---[18] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  C:\Users\xavie\Desktop\autocluster\go.mod
   |        |  |  |  +---[19] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  C:\Users\xavie\Desktop\autocluster\go.sum
   |        |  |  |
   |        |  |  +---[7 17]    (1 / 0.292308)
   |        |  |     +---[7] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -   C:\Users\xavie\Desktop\autocluster\cluster\version.go
   |        |  |     +---[17] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  C:\Users\xavie\Desktop\autocluster\distance\version.go
   |        |  |
   |        |  |
   |        |  +---[2 11 15 10 13]      (3 / 0.830816)
   |        |     +---[2] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -   C:\Users\xavie\Desktop\autocluster\README.md
   |        |     +---[11 15 10 13]     (2 / 0.779762)
   |        |        +---[11 15]        (1 / 0.652568)
   |        |        |  +---[11] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  C:\Users\xavie\Desktop\autocluster\distance\euclid.go
   |        |        |  +---[15] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  C:\Users\xavie\Desktop\autocluster\distance\matrix_test.go
   |        |        |
   |        |        +---[10 13]        (1 / 0.601190)
   |        |           +---[10] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  C:\Users\xavie\Desktop\autocluster\distance\cache_test.go
   |        |           +---[13] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  C:\Users\xavie\Desktop\autocluster\distance\file_test.go
   |        |
   |        |
   |        |
   |        |
   |        +---[1] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -   C:\Users\xavie\Desktop\autocluster\LICENSE
   |
   |
   |
   +---[23 20 21 22]    (3 / 0.749830)
      +---[23] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  C:\Users\xavie\Desktop\autocluster\testFiles\test.xlsx
      +---[20 21 22]    (2 / 0.077442)
         +---[20] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  C:\Users\xavie\Desktop\autocluster\testFiles\test.docx
         +---[21 22]    (1 / 0.000000)
            +---[21] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  C:\Users\xavie\Desktop\autocluster\testFiles\test.html
            +---[22] -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  C:\Users\xavie\Desktop\autocluster\testFiles\test.html.zip






Table of all clusters :

link dist.      level   cluster content .....................
1.032926        8       [0 12 9 3 4 8 5 6 16 14 18 19 7 17 2 11 15 10 13 1 23 20 21 22]
1.001316        7               [0 12 9 3 4 8 5 6 16 14 18 19 7 17 2 11 15 10 13 1]
0.000000        0                       [0]
0.990047        6                       [12 9 3 4 8 5 6 16 14 18 19 7 17 2 11 15 10 13 1]
0.939560        4                               [12 9 3 4 8 5 6 16 14]
0.860806        1                                       [12 9]
0.000000        0                                               [12]
0.000000        0                                               [9]
0.925092        3                                       [3 4 8 5 6 16 14]
0.897328        2                                               [3 4 8]
0.000000        0                                                       [3]
0.809298        1                                                       [4 8]
0.000000        0                                                               [4]
0.000000        0                                                               [8]
0.859249        2                                               [5 6 16 14]
0.781356        1                                                       [5 6]
0.000000        0                                                               [5]
0.000000        0                                                               [6]
0.809651        1                                                       [16 14]
0.000000        0                                                               [16]
0.000000        0                                                               [14]
0.956250        5                               [18 19 7 17 2 11 15 10 13 1]
0.897196        4                                       [18 19 7 17 2 11 15 10 13]
0.785714        2                                               [18 19 7 17]
0.662338        1                                                       [18 19]
0.000000        0                                                               [18]
0.000000        0                                                               [19]
0.292308        1                                                       [7 17]
0.000000        0                                                               [7]
0.000000        0                                                               [17]
0.830816        3                                               [2 11 15 10 13]
0.000000        0                                                       [2]
0.779762        2                                                       [11 15 10 13]
0.652568        1                                                               [11 15]
0.000000        0                                                                       [11]
0.000000        0                                                                       [15]
0.601190        1                                                               [10 13]
0.000000        0                                                                       [10]
0.000000        0                                                                       [13]
0.000000        0                                       [1]
0.749830        3               [23 20 21 22]
0.000000        0                       [23]
0.077442        2                       [20 21 22]
0.000000        0                               [20]
0.000000        1                               [21 22]
0.000000        0                                       [21]
0.000000        0                                       [22]



List of medoids and average internal distance per cluster :

        level   medoïd  medoïd dist                             cluster
        2       [5]     0.620183         --medoïd for-->        [5 6 16 14]
        1       [12]    0.430403         --medoïd for-->        [12 9]
        1       [16]    0.404826         --medoïd for-->        [16 14]
        4       [6]     0.773668         --medoïd for-->        [12 9 3 4 8 5 6 16 14]
        6       [10]    0.797878         --medoïd for-->        [12 9 3 4 8 5 6 16 14 18 19 7 17 2 11 15 10 13 1]
        1       [10]    0.300595         --medoïd for-->        [10 13]
        1       [4]     0.404649         --medoïd for-->        [4 8]
        3       [6]     0.730645         --medoïd for-->        [3 4 8 5 6 16 14]
        1       [21]    0.000000         --medoïd for-->        [21 22]
        3       [21]    0.206818         --medoïd for-->        [23 20 21 22]
        3       [15]    0.568920         --medoïd for-->        [2 11 15 10 13]
        5       [17]    0.703388         --medoïd for-->        [18 19 7 17 2 11 15 10 13 1]
        1       [5]     0.390678         --medoïd for-->        [5 6]
        4       [17]    0.675986         --medoïd for-->        [18 19 7 17 2 11 15 10 13]
        1       [7]     0.146154         --medoïd for-->        [7 17]
        1       [11]    0.326284         --medoïd for-->        [11 15]
        1       [18]    0.331169         --medoïd for-->        [18 19]
        2       [17]    0.440738         --medoïd for-->        [18 19 7 17]
        7       [10]    0.807989         --medoïd for-->        [0 12 9 3 4 8 5 6 16 14 18 19 7 17 2 11 15 10 13 1]
        2       [21]    0.025814         --medoïd for-->        [20 21 22]
root    8       [10]    0.837872         --medoïd for-->        [0 12 9 3 4 8 5 6 16 14 18 19 7 17 2 11 15 10 13 1 23 20 21 22]
        2       [15]    0.503446         --medoïd for-->        [11 15 10 13]
        2       [8]     0.567479         --medoïd for-->        [3 4 8]


```

By default, you get **a lot** of useful information.

First, you  get a list of the files, 24 here, which the program is going to process. Nothing fancy at this point, except that git files are excluded ;-). You also notice that a unique id is associated with each file, and will be used later to save screen space.

Then, a message about a cache appears? The caching process is entirely transparent, in the temporary forlder of your computer. It significantly accelerates processing of files when you run the program later. It takes care of itself if the name or content of your file is modified, don't worry ... The size of the cache is minimal, since a digest of the file content is used to uniquely identify known content.

From there, we can compute a distance matrix (here, a 23 x 23 matrix) for all the distance between the files. It is necessary because (see below) the file distance is an expensive calculation, we want to make sure we do it only once (and we cache it).

Then, the hierarchical grouping (the **clusters**) that program selected are shown. This representation is a *dendrogramme*, showing the root cluster (the one that contains all the files) and then spliting until there are only single file clusters.
Notice how the cluster content is descripbed by an array of *id* of files. You also see the *level* and *link distance*. The *level* starts at 0 for the leaf, and increases as we move towards the root cluster. Here, the root cluster is level 8. The *link distance* represents the cost we had to accept (and minimize) we creating this cluster by merging two smaller clusters. The smaller the *link distance*, the closer the left and right components of the cluster were from each other.

Finally, we get a table of the **medoïds** for each cluster. The * medoïd* is the file that represents the best the cluster, being the closest within the cluster to all the other. You get the *level* (as above), the *id* of the *medoïd*, the average distance from the *medoïd* to the other elements of the cluster, and a reminder of the cluster content. Obviously, the *medoid* always appears in that cluster content list. Do you see how the file 21 is the *medoïd* of [ 20,21,22,23] with a very small average distance to the rest of the files 20 to 24 ? These are the test files (see below).

A lot of CLI settings are available, you can get them with :

```
> go run ./... -h

Unsupervised clustering of files
        distance version : 1.0.2
        cluster version  : 0.8.1
Usage :
  -cache string
        cache file location (default "C:\\Users\\xavie\\AppData\\Local\\Temp\\fileDistance.cache")
  -d    print dendrogramme (default true)
  -dm
        print distance matrix (truncated) (default true)
  -f    list file names (default true)
  -h    print version, usage and exit
  -link string
        select type of linkage from single, complete, upgma (default "complete")
  -med
        compute and print medoids (dendrogramme) (default true)
  -min int
        set this value to a high number to get less clusters
  -t    print tree (default true)

```

## It's all about distance ...

You may want to have a quick look at the paper attached below.

The goal is to construct the distance matrix, by measuring a *distance* from one file to another.
First, because we care only about content, and not format, we do everything we can to extract text content from the file, to make sure we compare actual content and not format. For instance, in the test folder, you will find a couple of files with almost identical content : a word file, an excel file, and html file, a zipped file, ... A naïve comparison of these files would fail, and they would be put in far away clsuters, because of the significant difference in format. Because we first extract text before computing the distance, the results are much better. You can see in the clustering above that these 4 test files are aggregated totether naturally, showing reasonnably small distance between them - these are files 20 to 23 above.

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
* not much optimization was done beyond the obvious for performance. Clarity was prefered of perfomance. There might be room to improve here also, but I did not feel the need ...

See detailled documentation on https://pkg.go.dev/