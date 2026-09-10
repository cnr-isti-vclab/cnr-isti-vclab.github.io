---
# delete the following three lines if you want that your page appears:
# Here is needed to avoid that the example paper appears in the list of publications
# _build:
#   render: always
#   list: never

title: "ArchiClip: Learning joint text–geometry representations for 3D architectural freeform surfaces"
authors:
- Andrea Favilli
- Francesco Laccone
- Luigi Malomo
- Nicola Messina
- Fabio Carrara
- Paolo Cignoni
- Daniela Giorgi
# author_notes:
# - "Equal contribution"
# - "Equal contribution"
# Note use this date format
date: "2026-08-31T00:00:00Z"
doi: "10.1016/j.cad.2026.104156"

# Schedule page publish date (NOT publication's date).
publishDate: "2026-08-31T00:00:00Z"

# Publication type.
# Accepts a single type but formatted as a YAML list (for Hugo requirements).
# Enter a publication type from the CSL standard.
publication_types: ["article-journal"]

# Publication name and optional abbreviated publication name.
publication: "*Computer-Aided Design*"
publication_short: ""

abstract: We present ArchiClip, a method for learning multi‑modal joint representations of text and freeform architectural surfaces. Building on the widely adopted multi‑modal contrastive learning paradigm, our approach focuses specifically on distilling knowledge from the architectural domain and enabling robust understanding of freeform 3D geometry. To support this, we introduce a strategy for constructing ArchiShape, a curated 3D dataset enriched with domain‑specific, rich textual descriptions. ArchiShape is generated automatically by combining procedural modeling and geometric shape descriptors with pre‑trained large language models, eliminating the need for manual annotation. We then design a bi‑modal pre‑training framework that learns aligned embeddings for both textual descriptions and architectural freeform surfaces. The framework incorporates a 3D backbone network tailored to architectural geometry and a custom batch‑sampling scheme to ensure efficient training. We evaluate ArchiClip on cross‑modal 3D retrieval of architectural freeforms, demonstrating its ability to encode rich geometric and domain‑specific concepts (e.g., curvature, spatial organization), and highlighting the benefits of domain‑aware multi‑modal representation learning.

# Summary. An optional shortened abstract, it appears in the list of publications.
# summary: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis posuere tellus ac convallis placerat. Proin tincidunt magna sed ex sollicitudin condimentum.
# Add keywords here (example keywords below)
tags:
#- Digital Heritage 
- Architectural Geometry
#- Geometry Processing 
#- Digital Fabrication
#- Environment Monitoring
- Visual AI
#- XR and Advanced UI

# Additional Tags to be considered: 
#- 3D Reconstruction
#- 3D Printing
#- 3D Scanning
#- Texture Mapping
#- Coral Reef
#- TagLab

featured: false

# links:
# - name: ""
#   url: ""
url_pdf: https://vcgdata.isti.cnr.it/Publications/2026/ArchiClip/Favilli2026_ArchiClip.pdf
url_code: 'https://github.com/cnr-isti-vclab/ArchiClip'
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
  focal_point: ""
  preview_only: false

# Associated Projects web Page(optional).
# Not the funding project
projects: []

# Slides (optional).
#   Associate this publication with Markdown slides.
#   Simply enter your slide deck's filename without extension.
#   E.g. `slides: "example"` references `content/slides/example/index.md`.
#   Otherwise, set `slides: ""`.
slides: ''
---

