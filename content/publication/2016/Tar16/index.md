---
title: 'Volume-encoded UV-maps'
authors:
  - Marco Tarini
date: '2016-01-01T00:00:00Z'
publication_types: ['1']
publication: '*ACM Trans. on Graphics - Siggraph 2016*'
featured: false

url_pdf: http://vcg.isti.cnr.it/Publications/2016/Tar16/volume-encoded_UV-maps.pdf
image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
abstract: 'UV-maps are required in order to apply a 2D texture over a 3D model.Conventional UV-maps are defined by an assignment of uv positions to mesh vertices. We present an alternative representation, volume-encoded UV-maps, in which each point on the surface is mapped to a uv position which is solely a function of its 3D position. This function is tailored for a target surface: its restriction to the surface is a parametrization exhibiting high quality, e.g. in terms of angle and area preservation; and, near the surface, it is almost constant for small orthogonal displacements.The representation is applicable to a wide range of shapes and UV-maps, and unlocks several key advantages: it removes the need to duplicate vertices in the mesh to encode cuts in the map; it makes the UV-map representation independent from the meshing of the surface; the same texture, and even the same UV-map, can be shared by multiple geometrically similar models (e.g.alllevelsofa LoDpyramid); UV-maps can be applied to representations other than polygonal meshes, like point clouds or set of registered range-maps.Our schema is cheap on GPU computational and memory resources, requiring only a single, cache-coherent indirection to a small volumetric texture per fragment.We also provide an algorithm to construct a volume-encoded UV-map given a target surface.           	 Project page (with demo, sources, etc).'
url_pdf: http://vcg.isti.cnr.it/Publications/2016/Tar16/volume-encoded_UV-maps_additional.pdf
---
{{< figure src='http://vcg.isti.cnr.it/Publications/2016/Tar16/composed3.png' >}}
{{< figure src='http://vcg.isti.cnr.it/Publications/2016/Tar16/combined2.png' >}}
[ Project page (with demo, sources, etc). ](http://vcg.isti.cnr.it/volume-encoded-uv-maps/)

