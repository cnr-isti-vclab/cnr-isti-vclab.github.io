---
# delete the following three lines if you want that your page appears:
# Here is needed to avoid that the example paper appears in the list of publications
#_build:
#  render: always
#  list: never

title: 'Generating HDR Video from SDR Video'
authors:
  - SaiKiran Tedla
  - Francesco Banterle
  - Trevor Canham
  - Karanpreet Raja
  - David B. Lindell
  - Kirikos N. Kutulakos
  - Jiacheng Li
  - Feiran Li
  - Daisuke Iso

date: '2026-11-10'
  
author_notes:
#  - 'Equal contribution'
#  - 'Equal contribution'
# Note use this date format
#doi: 'https://doi.org/10.1145/3648570'

# Publication type.
# Legend: 0 = Uncategorized; 1 = Conference paper; 2 = Journal article;
# 3 = Preprint / Working Paper; 4 = Report; 5 = Book; 6 = Book section;
# 7 = Thesis; 8 = Patent
publication_types: ['article-journal']

# Publication name and optional abbreviated publication name.
publication: '*Proceedings of the 19th ACM SIGGRAPH Asia Conference Papers*'
publication_short: 'SA26'

abstract: 'The high dynamic range (HDR) video ecosystem is approaching maturity, but the problem of upconverting legacy standard dynamic range (SDR) videos persists without a convincing solution. We propose a framework for HDR video synthesis from casual SDR footage by leveraging a large-scale generative video model. We introduce a Multi-Exposure Video Model (MEVM) that can predict exposure-bracketed linear SDR video sequences from a single nonlinear SDR video input. We further propose a learnable Video Merging Model (VMM) that merges the predicted exposure-bracketed video into a high-quality HDR sequence while preserving detail in both shadows and highlights. Extensive experiments, quantitative and qualitative evaluation, and a user study demonstrate that our approach enables robust HDR conversion for in-the-wild examples from casual consumer videos and even cinematic footage. Finally, our model can support HDR synthesis pipelines built upon existing SDR generative video models.'

# Add keywords here (example keywords below)
tags:
  - HDR Imaging
  - Inverse Tone Mapping
  - Reverse Tone Mapping
  - Single-Image HDR
  - HDR Videos
  - Video Models
  - Diffusion Models
  - Deep Learning
 
featured: false

# links:
# - name: ""
#   url: ""
url_pdf: ''
url_code: 'https://github.com/SonyResearch/generative_sdr_to_hdr_video'
url_dataset: ''
url_poster: ''
url_project: 'https://sdr2hdrvideo.github.io/'
url_slides: ''
url_source: ''
url_video: ''

# Featured image
# To use, add an image named `featured.jpg/png` to your page's folder.
image:
  caption: ''
  focal_point: ''
  preview_only: false
---
