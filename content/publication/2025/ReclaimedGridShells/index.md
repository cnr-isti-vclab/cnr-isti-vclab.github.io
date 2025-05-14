---
# delete the following three lines if you want that your page appears:
#_build:
#  render: always
#  list: never

title: 'Optimizing Free-Form Grid Shells with Reclaimed Elements under Inventory Constraints'
authors:
  - Andrea Favilli
  - Francesco Laccone 
  - Paolo Cignoni 
  - Luigi Malomo 
  - Daniela Giorgi


date: '2025-04-18T00:00:00Z'
doi: '10.1111/cgf.70047'

# Schedule page publish date (NOT publication's date).
publishDate: '2025-04-18T00:00:00Z'

# Publication type.
# Legend: 0 = Uncategorized; 1 = Conference paper; 2 = Journal article;
# 3 = Preprint / Working Paper; 4 = Report; 5 = Book; 6 = Book section;
# 7 = Thesis; 8 = Patent
publication_types: ['article-journal']

# Publication name and optional abbreviated publication name.
publication: '*Computer Graphics Forum (Eurographics 2025)*'
publication_short: ''

abstract: We propose a method for designing 3D architectural free-form surfaces, represented as grid shells with beams sourced from inventories of reclaimed elements from dismantled buildings. In inventory-constrained design, the reused elements must be paired with elements in the target design. Traditional solutions to this assignment problem often result in cuts and material waste or geometric distortions that affect the surface aesthetics and buildability. Our method for inventory-constrained assisted design blends the traditional assignment problem with differentiable geometry optimization to reduce cut-off waste while preserving the design intent. Additionally, we extend our approach to incorporate strain energy minimization for structural efficiency. We design differentiable losses that account for inventory, geometry, and structural constraints, and streamline them into a complete pipeline, demonstrated through several case studies. Our approach enables the reuse of existing elements for new designs, reducing the need for sourcing new materials and disposing of waste. Consequently, it can serve as an initial step towards mitigating the significant environmental impact of the construction sector.

# Summary. An optional shortened abstract.
summary: 

tags:
  - Architectural Geometry
  - Visual AI
  - Grid Shells
  - Shape Optimization
  - Form Finding
  - FAIR
featured: true

# links:
# - name: ""
#   url: ""
url_pdf: http://vcgdata.isti.cnr.it/Publications/2025/EG-ReclaimedGridShells/Optimizing%20Free-Form%20Grid%20Shells%20with%20Reclaimed%20Elements.pdf
url_code: 'https://github.com/cnr-isti-vclab/ReclaimedGridShells'
url_dataset: ''
url_poster: ''
url_project: ''
url_slides: ''
url_source: ''
url_video: 'https://youtu.be/L05gny2ikNs'

# Featured image
# To use, add an image named `featured.jpg/png` to your page's folder.
image:
  caption: ''
  focal_point: ''
  preview_only: false

# Associated Projects (optional).
#   Associate this publication with one or more of your projects.
#   Simply enter your project's folder or file name without extension.
#   E.g. `internal-project` references `content/project/internal-project/index.md`.
#   Otherwise, set `projects: []`.
projects: []

# Slides (optional).
#   Associate this publication with Markdown slides.
#   Simply enter your slide deck's filename without extension.
#   E.g. `slides: "example"` references `content/slides/example/index.md`.
#   Otherwise, set `slides: ""`.
slides:
---
{{< youtube "L05gny2ikNs" >}}
