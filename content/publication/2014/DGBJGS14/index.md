---
title: 'eXploreMaps: Efficient Construction and Ubiquitous Exploration of Panoramic View Graphcs of Complex 3D Environments'
authors:
  - Marco Di Benedetto
  - Fabio Ganovelli
  - marcos balsa Rodriguez
  - Alberto Jaspe Villanueva
  - Enrico Gobbetti
  - Roberto Scopigno
date: '2014-01-01T00:00:00Z'
publication_types: ['1']
publication: '*Computer Graphics Forum, 33(2), 2014. Proc. Eurographics 2014, To appear*'
featured: false

image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
url_pdf: http://vcg.isti.cnr.it/Publications/2014/DGBJGS14/eg2014-exploremaps.pdf
abstract: 'http://www.exploremaps.org We introduce a novel efficient technique for automatically transforming a generic renderable 3D scene into a simple graph representation named ExploreMaps, where nodes are nicely placed point of views, called probes, and arcs are smooth paths between neighboring probes. Each probe is associated with a panoramic image enriched with preferred viewing orientations, and each path with a panoramic video. Our GPU-accelerated unattended construction pipeline distributes probes so as to guarantee coverage of the scene while accounting for perceptual criteria before finding smooth, good looking paths between neighboring probes. Images and videos are precomputed at construction time with off-line photorealistic rendering engines, providing a convincing 3D visualization beyond the limits of current real-time graphics techniques. At run-time, the graph is exploited both for creating automatic scene indexes and movie previews of complex scenes and for supporting interactive exploration through a low-DOF assisted navigation interface and the visual indexing of the scene provided by the selected viewpoints. Due to negligible CPU overhead and very limited use of GPU functionality, real-time performance is achieved on emerging web-based environments based on WebGL even on low-powered mobile devices.'
---
{{< figure src='http://vcg.isti.cnr.it/Publications/2014/DGBJGS14/result-medieval-graph.png' >}}
{{< figure src='http://vcg.isti.cnr.it/Publications/2014/DGBJGS14/webgl-gui-diagram.jpg' >}}
[http://www.exploremaps.org](http://www.exploremaps.org)

<iframe width="580" height="435" src="http://www.youtube.com/embed/yW3uVEzd-Pg" frameborder="0" frameborder="0" allowfullscreen>

