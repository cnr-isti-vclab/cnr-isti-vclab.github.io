---
title: 'Reliable feature-line driven quad-remeshing'
authors:
  - Nico Pietroni
  - Stefano Nuvoli
  - Thomas Alderighi
  - Paolo Cignoni
  - Marco Tarini
date: '2021-01-01T00:00:00Z'
publication_types: ['1']
publication: '*ACM Trans. on Graphics*'
featured: false

abstract: 'We present a new algorithm for the semi-regular quadrangulation of an input surface, driven by its line features, such as sharp creases. We define a perfectly feature-aligned cross-field and a coarse layout of polygonal-shaped patches where we strictly ensure that all the feature-lines are represented as patch boundaries. To be able to consistently do so, we allow non-quadrilateral patches and T-junctions in the layout; the key is the ability to constrain the layout so that it still admits a globally consistent, T-junction-free, and pure-quad internal tessellation of its patches. This requires the insertion of additional irregular-vertices inside patches, but the regularity of the final-mesh is safeguarded by optimizing for both their number and for their reciprocal alignment. In total, our method guarantees the reproduction of feature-lines by construction, while still producing good quality, isometric, pure-quad, conforming meshes, making it an ideal candidate for CAD models. Moreover, the method is fully automatic, requiring no user intervention, and remarkably reliable, requiring little assumptions on the input mesh, as we demonstrate by batch processing the entire Thingi10K repository, with less than 0.5% of the attempted cases failing to produce a usable mesh..  DOI: 10.1145/3450626.3459941 Browsable dataset of Thingi10K remeshed objects:  www.quadmesh.cloud         Download the slides! (.pttx)'
image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
url_pdf: http://vcg.isti.cnr.it/Publications/2021/PNACT21/ReliableQuad.pdf
---
{{< figure src="http://vcg.isti.cnr.it/Publications/2021/PNACT21/00teaser.png" >}}
{{< figure src="http://vcg.isti.cnr.it/Publications/2021/PNACT21/example01.png" >}}
{{< figure src="http://vcg.isti.cnr.it/Publications/2021/PNACT21/example00.png" >}}
{{< figure src="http://vcg.isti.cnr.it/Publications/2021/PNACT21/example02.png" >}}
[ 10.1145/3450626.3459941](https://doi.org/10.1145/3450626.3459941)

[ www.quadmesh.cloud](https://www.quadmesh.cloud/Thingi10K)

[ Download the slides! (.pttx)](http://vcg.isti.cnr.it/Publicstions/2021/PNACT21/reliable-quadremesh-siggraph2021.pptx)

<iframe width="580" height="435" src="https://www.youtube.com/embed/mUzYFH7DrN0" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" frameborder="0" allowfullscreen>

