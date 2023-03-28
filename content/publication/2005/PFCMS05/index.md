---
title: 'Exploiting the scanning sequence for automatic registration of large sets of range maps'
authors:
  - Paolo Pingi
  - Andrea Fasano
  - Paolo Cignoni
  - Claudio Montani
  - Roberto Scopigno
date: '2005-01-01T00:00:00Z'
publication_types: ['1']
publication: '*Computer Graphics Forum*'
featured: false

url_pdf: http://vcg.isti.cnr.it/Publications/2005/PFCMS05/eg05_register_final_uploaded.pdf
image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
abstract: 'Range map registration is still the most time consuming phase in the processing of 3D scanning data. This is because real scanning sets are composed of hundreds of range maps and their registration is still partially manual. We propose a new method to manage complex scan sets acquired by following a regular scanner pose pattern. Our goal is to define an initial adjacency graph by coarsely aligning couples of range maps that we know are partially overlapping thanks to the adopted scanning strategy. For a pair of partially overlapping range maps, our iterative solution locates pairs of correspondent vertices through the computation of a regular n&times;n kernel which takes into account vertex normals and is defined in the 2D space of the range map (represented in implicit 2D format rather than as a triangle mesh in 3D space). The shape-characterization kernel and the metrics defined give a sufficiently accurate shape matching, which has been proven to fit well the requirements of automatic registration. This initial set of adjacency arcs can then be augmented by the automatic identification of the other significant arcs, by adopting a criterion based on approximate range map overlap computation. With respect to the solutions present in literature, the simplifications and assumptions adopted make our solution specifically oriented to complex 3D scanning campaigns (hundreds of range maps). The proposed method can coarsely register range maps in parallel with the acquisition activity and this is a valuable help in assessing on site the completeness of the sampling of large objects.'
---
{{< figure src="http://vcg.isti.cnr.it/Publications/2005/PFCMS05/pic_arki.png" >}}
{{< figure src="http://vcg.isti.cnr.it/Publications/2005/PFCMS05/pic_sepolcro2.png" >}}
{{< figure src="http://vcg.isti.cnr.it/Publications/2005/PFCMS05/pic_spyral_stripe_small2.png" >}}
