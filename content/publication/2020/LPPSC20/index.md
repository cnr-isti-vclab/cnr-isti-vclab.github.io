---
title: 'LoopyCuts: Practical Feature-Preserving Block Decomposition'
authors:
  - Marco Livesu
  - Nico Pietroni
  - Enrico Puppo
  - Alla Sheffer
  - Paolo Cignoni
date: '2020-01-01T00:00:00Z'
publication_types: ['1']
publication: '*ACM Trans. on Graphics - Siggraph 2020*'
featured: false

abstract: 'We present a new fully automatic block-decomposition algorithm for feature-preserving, strongly hex-dominant meshing, that yields results with a drastically larger percentage of hex elements than prior art. Our method is guided by a surface field that conforms to both surface curvature and feature lines, and exploits an ordered set of cutting loops that evenly cover the input surface, defining an arrangement of loops suitable for hex-element generation. We decompose the solid into coarse blocks by iteratively cutting it with surfaces bounded by these loops. The vast majority of the obtained blocks can be turned into hexahedral cells via simple midpoint subdivision.    GitHub Code'
image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
url_pdf: http://vcg.isti.cnr.it/Publications/2020/LPPSC20/LPPSC20.pdf
---
{{< figure src='http://vcg.isti.cnr.it/Publications/2020/LPPSC20/organic_hexa.png' >}}
{{< figure src='http://vcg.isti.cnr.it/Publications/2020/LPPSC20/cad_hexa.png' >}}
{{< figure src='http://vcg.isti.cnr.it/Publications/2020/LPPSC20/0_LPPSC20.jpg' >}}
{{< figure src='http://vcg.isti.cnr.it/Publications/2020/LPPSC20/teaser.png' >}}
{{< figure src='http://vcg.isti.cnr.it/Publications/2020/LPPSC20/cad_non_hexa.png' >}}
[GitHub Code](https://github.com/mlivesu/LoopyCuts)

<iframe width="580" height="435" src="https://www.youtube.com/embed/n-rWtLi3LSU" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" frameborder="0" allowfullscreen>

