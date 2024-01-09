---
# delete the following three lines if you want that your page appears:
#_build:
#  render: always
#  list: never

title: 'Geometric deep learning for statics-aware grid shells'
authors:
  - Andrea Favilli
  - Francesco Laccone 
  - Paolo Cignoni 
  - Luigi Malomo 
  - Daniela Giorgi


date: '2024-02-01T00:00:00Z'
doi: 'https://doi.org/10.1016/j.compstruc.2023.107238'

# Schedule page publish date (NOT publication's date).
publishDate: '2017-01-01T00:00:00Z'

# Publication type.
# Legend: 0 = Uncategorized; 1 = Conference paper; 2 = Journal article;
# 3 = Preprint / Working Paper; 4 = Report; 5 = Book; 6 = Book section;
# 7 = Thesis; 8 = Patent
publication_types: ['2']

# Publication name and optional abbreviated publication name.
publication: 'Computers & Structures'
publication_short: ''

abstract: This paper introduces a novel method for shape optimization and form-finding of free-form, triangular grid shells, based on geometric deep learning. We define an architecture which consumes a 3D mesh representing the initial design of a free-form grid shell, and outputs vertex displacements to get an optimized grid shell that minimizes structural compliance, while preserving design intent. The main ingredients of the architecture are layers that produce deep vertex embeddings from geometric input features, and a differentiable loss implementing structural analysis. We evaluate the method performance on a benchmark of eighteen free-form grid shell structures characterized by various size, geometry, and tessellation. Our results demonstrate that our approach can solve the shape optimization and form finding problem for a diverse range of structures, more effectively and efficiently than existing common tools.

# Summary. An optional shortened abstract.
summary: 

tags:
  - Architectural Geometry
  - Grid Shells
  - Shape Optimization
  - Form Finding
  - evocation
  - Sun
featured: true

# links:
# - name: ""
#   url: ""
url_pdf: https://vcgdata.isti.cnr.it/Publications/2024/FLCMG24-GeometricLearningxShells/FLCMG24-GeometricLearningxShells.pdf
url_code: ''
url_dataset: 'https://github.com/cnr-isti-vclab/GeomDL4GridShell#geometric-deep-learning-for-statics-aware-grid-shells'
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

# Associated Projects (optional).
#   Associate this publication with one or more of your projects.
#   Simply enter your project's folder or file name without extension.
#   E.g. `internal-project` references `content/project/internal-project/index.md`.
#   Otherwise, set `projects: []`.
projects: [Evocation,Sun]

# Slides (optional).
#   Associate this publication with Markdown slides.
#   Simply enter your slide deck's filename without extension.
#   E.g. `slides: "example"` references `content/slides/example/index.md`.
#   Otherwise, set `slides: ""`.
slides:
---
