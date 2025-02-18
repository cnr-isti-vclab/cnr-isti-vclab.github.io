---
title: 'High Dynamic Range Point Clouds for Real-Time Relighting'
authors:
  - Manuele Sabbadin
  - Gianpaolo Palma
  - Francesco Banterle
  - Tamy Boubekeur
  - Paolo Cignoni
date: '2019-11-01T00:00:00Z'
publication_types: ['article-journal']
publication: '*Computer Graphics Forum*'
featured: false
doi: '10.1111/cgf.13857'

links:
  - name: Additional material
    url: https://vcgdata.isti.cnr.it/Publications/2019/SPBBC19/RealTime_PBGI__Additional.pdf
url_pdf: https://vcgdata.isti.cnr.it/Publications/2019/SPBBC19/RealTime_PBGI.pdf
url_slides: https://vcgdata.isti.cnr.it/Publications/2019/SPBBC19/hdr_pbgi.pptx

image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
abstract: 'Acquired 3D point clouds make possible quick modeling of virtual scenes from the real world.With modern 3D capture pipelines, each point sample often comes with additional attributes such as normal vector and color response. Although rendering and processing such data has been extensively studied, little attention has been devoted using the light transport hidden in the recorded per-sample color response to relight virtual objects in visual effects (VFX) look-dev or augmented reality (AR) scenarios. Typically, standard relighting environment exploits global environment maps together with a collection of local light probes to reflect the light mood of the real scene on the virtual object. We propose instead a unified spatial approximation of the radiance and visibility relationships present in the scene, in the form of a colored point cloud. To do so, our method relies on two core components: High Dynamic Range (HDR) expansion and real-time Point-Based Global Illumination (PBGI). First, since an acquired color point cloud typically comes in Low Dynamic Range (LDR) format, we boost it using a single HDR photo exemplar of the captured scene that can cover part of it. We perform this expansion efficiently by first expanding the dynamic range of a set of renderings of the point cloud and then projecting these renderings on the original cloud. At this stage, we propagate the expansion to the regions not covered by the renderings or with low-quality dynamic range by solving a Poisson system. Then, at rendering time, we use the resulting HDR point cloud to relight virtual objects, providing a diffuse model of the indirect illumination propagated by the environment. To do so, we design a PBGI algorithm that exploits the GPU''s geometry shader stage as well as a new mipmapping operator, tailored for G-buffers, to achieve real-time performances. As a result, our method can effectively relight virtual objects exhibiting diffuse and glossy physically-based materials in real time. Furthermore, it accounts for the spatial embedding of the object within the 3D environment. We evaluate our approach on manufactured scenes to assess the error introduced at every step from the perfect ground truth. We also report experiments with real captured data, covering a range of capture technologies, from active scanning to multiview stereo reconstruction.'
tags:
- Rendering and Advanced UI
---
{{< figure class="pubFig" src="https://vcgdata.isti.cnr.it/Publications/2019/SPBBC19/office.png" >}}
{{< figure class="pubFig" src="https://vcgdata.isti.cnr.it/Publications/2019/SPBBC19/desk.png" >}}
{{< figure class="pubFig" src="https://vcgdata.isti.cnr.it/Publications/2019/SPBBC19/atrium.png" >}}

{{< youtube "szBWhY1b4x8" >}}
