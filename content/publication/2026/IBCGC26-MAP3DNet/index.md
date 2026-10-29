---
# delete the following three lines if you want that your page appears:
# Here is needed to avoid that the example paper appears in the list of publications
#_build:
#  render: always
#  list: never

title: 'MAP3DNet: Multimodal Aesthetic Prediction for 3D Models?'
authors:
  - Samuele Intaschi
  - Francesco Banterle
  - Massimiliano Corsini
  - Daniela Giorgi
  - Paolo Cignoni
  
author_notes:
#  - 'Equal contribution'
#  - 'Equal contribution'
# Note use this date format
date: '2026-10-29T19:02:00Z'
#doi: 'https://doi.org/10.1145/3648570'


# Publication type.
# Legend: 0 = Uncategorized; 1 = Conference paper; 2 = Journal article;
# 3 = Preprint / Working Paper; 4 = Report; 5 = Book; 6 = Book section;
# 7 = Thesis; 8 = Patent
publication_types: ['article-journal']

# Publication name and optional abbreviated publication name.
publication: '*Proceedings of the 37th British Machine Vision Conference (BMVC)*'
publication_short: 'BMVC2026'

abstract: 'Predicting the aesthetic appeal of 3D models is an emerging problem at the intersection of computer vision, graphics, and human preference modeling. Unlike geometric and visual quality assessment, aesthetic prediction aims to estimate how visually appealing a 3D asset might be perceived by users. We formulate this task as a regression task, learning to predict a global aesthetic score derived from user engagement statistics collected from a 3D content sharing platform, where likes considered relative to views provide a scalable proxy for collective appreciation. We build on the ViDA3D dataset, a collection of Sketchfab models represented by multi-view renderings and artist-provided tags. To complement this textual information, we generate captions from rendered views and propose \vidanet, a lightweight multimodal architecture that combines visual features, tags, and generated captions. {\vidanet} learns to weight different viewpoints, encodes visual and textual cues through vision-language representations, and uses cross-modal attention to predict the final aesthetic score. Experiments demonstrate that our multimodal approach improves over visual-only alternatives. While visual features remain the main cue, tags and generated captions provide important complementary semantic context. {\vidanet} outperforms adapted baselines and produces rankings closely aligned with those derived from human engagement, supporting the assumption that aesthetic evaluation should be modeled as a context-aware, multimodal phenomenon.'

# Add keywords here (example keywords below)
tags:
- 3D Aesthetic
- Deep Learning
- No-reference Metrics
 
featured: false

# links:
# - name: ""
#   url: ""
url_pdf: ''
url_code: ''
url_dataset: ''
url_poster: ''
url_project: ''
url_slides: ''
url_source: ''
url_video: ''

# Featured image
# To use, add an image named `featured.jpg/png` to your page's folder.
image:
  caption: ''
  focal_point: ''
  preview_only: false

# Associated Projects web Page(optional).
# Not the funding project
projects: []

# Slides (optional).
#   Associate this publication with Markdown slides.
#   Simply enter your slide deck's filename without extension.
#   E.g. `slides: "example"` references `content/slides/example/index.md`.
#   Otherwise, set `slides: ""`.
slides:
---
<!-- Supplementary notes can be added here, including [code and math](https://wowchemy.com/docs/content/writing-markdown-latex/). -->
