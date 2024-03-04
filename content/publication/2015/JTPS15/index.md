---
title: 'Instant Field-aligned Meshes'
authors:
  - Olga Sorkine-Hornung
  - Daniele Panozzo
  - Marco Tarini
  - Wenzel Jakob
date: '2015-01-01T00:00:00Z'
publication_types: ['paper-conference']
publication: '*ACM Trans. on Graphics*'
featured: false

url_pdf: https://vcgdata.isti.cnr.it/Publications/2015/JTPS15/instant-meshes-SA-2015-jakob-et-al-compressed.pdf
image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
abstract: 'We present a novel approach to remesh a surface into an isotropic triangular or quad-dominant mesh using a unified local smoothing operator that optimizes both the edge orientations and vertex positions in the output mesh. Our algorithm produces meshes with high isotropy while naturally aligning and snapping edges to sharp features. The method is simple to implement and parallelize, and it can process a variety of input surface representations, such as point clouds, range scans and triangle meshes. Our full pipeline executes instantly (less than a second) on meshes with hundreds of thousands of faces, enabling new types of interactive workflows. Since our algorithm avoids any global optimization, and its key steps scale linearly with input size, we are able to process extremely large meshes and point clouds, with sizes exceeding several hundred million elements. To demonstrate the robustness and effectiveness of our method, we apply it to hundreds of models of varying complexity and provide our cross-platform reference implementation in the supplemental material.  Code and executables on gitHub!'
url_video: https://vcgdata.isti.cnr.it/Publications/2015/JTPS15/instant-field-aligned-meshes.mp4
---
{{< figure src="https://vcgdata.isti.cnr.it/Publications/2015/JTPS15/posbrush.jpg" >}}
{{< figure src="https://vcgdata.isti.cnr.it/Publications/2015/JTPS15/bunny_varyingscale.jpg" >}}
{{< figure src="https://vcgdata.isti.cnr.it/Publications/2015/JTPS15/database.jpg" >}}
{{< figure src="https://vcgdata.isti.cnr.it/Publications/2015/JTPS15/box_ext_int.jpg" >}}
{{< figure src="https://vcgdata.isti.cnr.it/Publications/2015/JTPS15/stmatthew.jpg" >}}
{{< figure src="https://vcgdata.isti.cnr.it/Publications/2015/JTPS15/box_ext_int_pos.jpg" >}}
{{< figure src="https://vcgdata.isti.cnr.it/Publications/2015/JTPS15/edit_orient_stacked.jpg" >}}
{{< figure src="https://vcgdata.isti.cnr.it/Publications/2015/JTPS15/pos_sing.png" >}}
{{< figure src="https://vcgdata.isti.cnr.it/Publications/2015/JTPS15/a02_representative_image.jpg" >}}
{{< figure src="https://vcgdata.isti.cnr.it/Publications/2015/JTPS15/res_3dscan.jpg" >}}
{{< figure src="https://vcgdata.isti.cnr.it/Publications/2015/JTPS15/a01_representative.png" >}}
[Code and executables on gitHub!](https://github.com/wjakob/instant-meshes)

