---
title: 'Mixing tone mapping operators on the GPU by differential zone mapping based on psychophysical experiments'
authors:
  - Francesco Banterle
  - Alessandro Artusi
  - Elena Sikudova
  - Patrick Ledda
  - Thomas Edward William Bashford-Rogers
  - Alan Chalmers
  - Marina Bloj
date: '2016-01-01T00:00:00Z'
publication_types: ['1']
publication: '*Signal Processing: Image Communication*'
featured: false

url_pdf: http://vcg.isti.cnr.it/Publications/2016/BASLBCB16/main_personal.pdf
image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
abstract: 'Abstract In this paper, we present a new technique for displaying High Dynamic Range (HDR) images on Low Dynamic Range (LDR) displays in an efficient way on the GPU. The described process has three stages. First, the input image is segmented into luminance zones. Second, the tone mapping operator (TMO) that performs better in each zone is automatically selected. Finally, the resulting tone mapping (TM) outputs for each zone are merged, generating the final LDR output image. To establish the TMO that performs better in each luminance zone we conducted a preliminary psychophysical experiment using a set of HDR images and six different TMOs. We validated our composite technique on several (new) HDR images and conducted a further psychophysical experiment, using an HDR display as the reference, that establishes the advantages of our hybrid three-stage approach over a traditional individual TMO. Finally, we present a GPU version, which is perceptually equal to the standard version but with much improved computational performance.     Link to the official publication.          The source code of this method is part of the HDR Toolbox and it is implemented in the function BanterleTMO.m.'
---
[Link](https://www.sciencedirect.com/science/article/pii/S0923596516301308)

[The source code](https://github.com/banterle/HDR_Toolbox)

