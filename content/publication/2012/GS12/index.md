---
title: 'OCME: out-of-core Mesh Editing Made Practical'
authors:
  - Fabio Ganovelli
  - Roberto Scopigno
date: '2012-01-01T00:00:00Z'
publication_types: ['paper-conference']
publication: '*IEEE Computer Graphics and Applications*'
featured: false

image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
url_pdf: https://vcgdata.isti.cnr.it/Publications/2012/GS12/ocme_preprint.pdf
abstract: 'OCME (Out-of-Core Mesh Editing) is a novel data-structure  and related algorithms for out-of-core editing of large meshes. OCME uses a hashed multigrid where the triangles are inserted on the base of their size and position. This choice allows a rapid access and, on average, a constant construction time per triangle. Unlike previous approaches, no explicit hierarchy is maintained and therefore insertion/modification/deletion of data does not require costly refitting procedures. OCME stores attributes locally, for example it allows to assign vertex color only to a small subparts of the dataset, and naturally handles multiple-scale datasets.'
---
{{< figure src="https://vcgdata.isti.cnr.it/Publications/2012/GS12/panels_hierarchy.png" >}}
<iframe width="580" height="435" src="http://www.youtube.com/embed/JrpwGsbUfiw" frameborder="0" frameborder="0" allowfullscreen>

