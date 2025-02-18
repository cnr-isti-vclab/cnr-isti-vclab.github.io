---
title: 'Animation-Aware Quadrangulation'
authors:
  - Giorgio Marcias
  - Nico Pietroni
  - Daniele Panozzo
  - Enrico Puppo
  - Olga Sorkine
date: '2013-01-01T00:00:00Z'
publication_types: ['paper-conference']
publication: '*Computer Graphics Forum SGP 2013*'
featured: false

image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
abstract: 'Geometric meshes that model animated characters must be designed while taking into account the deformations that the shape will undergo during animation. We analyze an input sequence of meshes with point-to-point correspondence, and we automatically produce a quadrangular mesh that fits well the input animation. We first analyze the local deformation that the surface undergoes at each point, and we initialize a cross field that remains as aligned as possible to the principal directions of deformation throughout the sequence. We then smooth this cross field based on an energy that uses a weighted combination of the initial field and the local amount of stretch. Finally, we compute a field-aligned quadrangulation with an off-the-shelf method. Our technique is fast and very simple to implement, and it significantly improves the quality of the output quad mesh and its suitability for character animation, compared to creating the quad mesh based on a single pose. We present experimental results and comparisons with a state-of-the-art quadrangulation method, on both sequences from 3D scanning and synthetic sequences obtained by a rough animation of a triangulated model.    Talk Slides       Data       Code'
url_pdf: https://vcgdata.isti.cnr.it/Publications/2013/MPPPS13/AnimationAwareQuadrFinal.pdf
---
{{< figure src="https://vcgdata.isti.cnr.it/Publications/2013/MPPPS13/ethfaceOur1.png" >}}
[ Talk Slides ](https://vcgdata.isti.cnr.it/Publicstions/2013/MPPPS13/AnimationAwareQuad.pptx)

[ Data ](https://vcgdata.isti.cnr.it/Publicstions/2013/MPPPS13/Animation-Aware_Quadrangulation_data.zip)

[ Code ](https://vcgdata.isti.cnr.it/Publicstions/2013/MPPPS13/Animation-Aware_Quadrangulation_code.zip)

{{< youtube "d4ZalZ48vXE" >}}