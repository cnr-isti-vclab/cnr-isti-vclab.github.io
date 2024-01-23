---
title: 'Mill and fold: Shape simplification for fabrication'
authors:
  - Alessandro Muntoni
  - Stefano Nuvoli
  - Andreas Scalas
  - Alessandro Tola
  - Luigi Malomo
  - Riccardo Scateni
date: '2019-01-01T00:00:00Z'
publication_types: ['1']
publication: '*Computer & Graphics*'
featured: false

image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
url_pdf: https://vcgdata.isti.cnr.it/Publications/2019/MNSTMS19/mill_and_fold.pdf
abstract: 'We introduce a pipeline for simplifying digital 3D shapes and fabricate them using 2D polygonal flat parts. Our method generates shapes that, once unfolded, can be fabricated with CNC milling machines using special tools called V-Grooves. These tools create V-shaped furrows at given angles depending on the shape of the used tool. Milling the edges of each flat facet simplifies the manual assembly, which consists only in folding adjacent facets at a constrained angle. Our method generates simplified shapes where every dihedral angle between adjacent facets belongs to a restricted set, thus making the assembly process quicker and more straightforward. Firstly, our method automatically computes a simplified version of the input model, using the marching cubes algorithm on the original mesh and iteratively performing local changes on the resulting triangle mesh. The user can then perform an additional manual simplification to remove unwanted facets. Finally, an unfolding algorithm, which takes into account the thickness of the material, flattens the polygonal facets onto the 2D plane, so that a CNC milling machine can fabricate it from a sheet of rigid material.'
---
