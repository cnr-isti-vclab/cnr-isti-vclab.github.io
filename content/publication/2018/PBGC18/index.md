---
title: 'Scalable non-rigid registration for multi-view stereo data'
authors:
  - Gianpaolo Palma
  - Tamy Boubekeur
  - Fabio Ganovelli
  - Paolo Cignoni
date: '2018-06-01T00:00:00Z'
publication_types: ['article-journal']
publication: '*ISPRS Journal of Photogrammetry and Remote Sensing*'
featured: false
doi: '10.1016/j.isprsjprs.2018.06.012'

url_pdf: https://vcgdata.isti.cnr.it/Publications/2018/PBGC18/paper.pdf
image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
abstract: 'We propose a new non-rigid registration method for large 3D meshes from Multi-View Stereo (MVS) reconstruction characterized by low-frequency shape deformations induced by several factors, such as low sensor quality and irregular sampling object coverage. Starting from a reference model to which we want to align a new 3D mesh, our method starts by decomposing it in patches using a Lloyd clustering before running an ICP local registration for each patch. Then, we improve the alignment using few geometric constraints and finally, we build a global deformation function that blends the estimated per-patch transformations. This function is structured on top of a deformation graph derived from the dual graph of the clustering. Our algorithm is iterated until convergence, increasing progressively the number of patches in the clustering to capture smaller deformations. The method comes with a scalable multicore implementation that enables, for the first time, the alignment of meshes made of tens of millions of triangles in a few minutes. We report extensive experiments of our algorithm on several dense Multi-View Stereo models, using a 3D scan or another MVS reconstruction as reference. Beyond MVS data, we also applied our algorithm to different scenarios, exhibiting more complex and larger deformations, such as 3D motion capture dataset or 3D scans of dynamic objects. The good alignment results obtained for both datasets highlights the efficiency and the flexibility of our approach.'

---

{{< youtube "99Sq92JfKkM" >}}