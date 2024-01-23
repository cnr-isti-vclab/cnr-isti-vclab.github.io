---
title: 'Soft Transparency for Point Cloud Rendering'
authors:
  - Patrick Seemann
  - Gianpaolo Palma
  - Matteo Dellepiane
  - Paolo Cignoni
  - Michael Goesele
date: '2018-01-01T00:00:00Z'
publication_types: ['1']
publication: '*Eurographics Symposium on Rendering - Experimental Ideas & Implementations*'
featured: false

url_pdf: https://vcgdata.isti.cnr.it/Publications/2018/SPDCG18/soft_transparency.pdf
abstract: 'We propose a novel rendering framework for visualizing point data with complex structures and/or different quality of data. The point cloud can be characterized by setting a per-point scalar field associated to the aspect that differentiates the parts of the dataset (i.e. uncertainty given by local normal variation). Our rendering method uses the scalar field to render points as solid splats or semi-transparent spheres with non-uniform density to produce the final image. To that end, we derive a base model for integrating density in (intersecting) spheres for both the uniform and non-uniform setting and introduce a simple and fast approximation which yields interactive rendering speeds for millions of points. Because our method only relies on the basic OpenGL rasterization pipeline, rendering properties can be adjusted in real-time by user. The method has been tested on several datasets with different characteristics, and user studies show that a clearer understanding of the scene is possible in comparison with point splatting techniques and basic transparency rendering.'
url_video: https://vcgdata.isti.cnr.it/Publications/2018/SPDCG18/video_final.mp4
image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
---
<iframe width="580" height="435" src="https://www.youtube.com/embed/hWz4ReLG36U?rel=0" frameborder="0" allow="autoplay; encrypted-media" frameborder="0" allowfullscreen>

