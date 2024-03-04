---
title: 'Texture Defragmentation for Photo-Reconstructed 3D Models'
authors:
  - Andrea Maggiordomo
  - Paolo Cignoni
  - Marco Tarini
date: '2021-01-01T00:00:00Z'
publication_types: ['paper-conference']
publication: '*Computer Graphics Forum*'
featured: false

image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
url_pdf: http://vcg.isti.cnr.it/Publications/2021/MCT21/2021 CGF Texture Defragmentation.pdf
tags:
  - Geometry Processing
  - 3D Scanning

abstract: 'We propose a method to improve an existing parametrization (UV-map layout) of a textured 3D model, targeted explicitly at alleviating typical defects afflicting models generated with automatic photo-reconstruction tools from real-world objects. This class of 3D data is becoming increasingly important thanks to the growing popularity of reliable, ready-to-use photogrammetry software packages. The resulting textured models are richly detailed, but their underlying parametrization typically falls short of many practical requirements, particularly exhibiting excessive fragmentation and consequent problems. Producing a completely new UV-map, with standard parametrization techniques, and then resampling a new texture image, is often neither practical nor desirable for at least two reasons: first, these models have characteristics (such as inconsistencies, high resolution) that make them unfit for automatic or manual parametrization; second, the required resampling leads to unnecessary signal degradation because this process is unaware of the original texel densities. In contrast, our method improves the existing UV-map instead of replacing it, balancing the reduction of the map fragmentation with signal degradation due to resampling, while also avoiding oversampling of the original signal. The proposed approach is fully automatic and extensively tested on a large benchmark of photo-reconstructed models; quantitative evaluation evidences a drastic and consistent improvement of the mappings.   https://doi.org/10.1111/cgf.142615    Reference implementation available at https://github.com/maggio-a/texture-defrag'
---
{{< figure src="http://vcg.isti.cnr.it/Publications/2021/MCT21/texture_defrag_6.png" >}}
{{< figure src="http://vcg.isti.cnr.it/Publications/2021/MCT21/texture_defrag_2.png" >}}
{{< figure src="http://vcg.isti.cnr.it/Publications/2021/MCT21/texture_defrag_4.png" >}}
{{< figure src="http://vcg.isti.cnr.it/Publications/2021/MCT21/texture_defrag_5.png" >}}
{{< figure src="http://vcg.isti.cnr.it/Publications/2021/MCT21/texture_defrag_3.png" >}}
{{< figure src="http://vcg.isti.cnr.it/Publications/2021/MCT21/texture_defrag_1.png" >}}
[https://doi.org/10.1111/cgf.142615](https://doi.org/10.1111/cgf.142615)

[https://github.com/maggio-a/texture-defrag](https://github.com/maggio-a/texture-defrag)

