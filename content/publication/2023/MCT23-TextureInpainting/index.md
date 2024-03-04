---
title: 'Texture Inpainting for Photogrammetric Models'
authors:
  - Andrea Maggiordomo
  - Paolo Cignoni
  - Marco Tarini
date: '2023-02-28'
publication_types: ['article-journal']
publication: '*Computer Graphics Forum*'
featured: false
doi: ' https://doi.org/10.1111/cgf.14735'

url_pdf: https://vcgdata.isti.cnr.it/Publications/2023/MCT23-TextureInpainting/2023TextureInpainting.pdf
image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
tags:
  - Geometry Processing
  - Texture Inpainting
  - 3D Scanning
  - Visual AI
abstract: 'We devise a technique designed to remove the texturing artefacts that are typical of 3D models representing real-world objects, acquired by photogrammetric techniques. Our technique leverages the recent advancements in inpainting of natural colour images, adapting them to the specific context. A neural network, modified and trained for our purposes, replaces the texture areas containing the defects, substituting them with new plausible patches of texels, reconstructed from the surrounding surface texture. We train and apply the network model on locally reparametrized texture patches, so to provide an input that simplifies the learning process, because it avoids any texture seams, unused texture areas, background, depth jumps and so on. We automatically extract appropriate training data from real-world datasets. We show two applications of the resulting method: one, as a fully automatic tool, addressing all problems that can be detected by analysing the UV-map of the input model; and another, as an interactive semi-automatic tool, presented to the user as a 3D ‘fixing’ brush that has the effect of removing artefacts from any zone the users paints on. We demonstrate our method on a variety of real-world inputs and provide a reference usable implementation. '
---
