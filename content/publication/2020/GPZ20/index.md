---
title: 'Reinforcement of General Shell Structures'
authors:
  - Francisca  Gil-Ureta
  - Nico Pietroni
  - Denis Zorin
date: '2020-01-01T00:00:00Z'
publication_types: ['paper-conference']
publication: '*ACM Trans. on Graphics - Presented at Siggraph 2020*'
featured: false

image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
abstract: 'We introduce an efficient method for designing shell reinforcements of minimal weight. Inspired by classical Michell trusses, we create a reinforcement layout whose members are aligned with optimal stress directions, then optimize their shape minimizing the volume while keeping stresses bounded. We exploit two predominant techniques for reinforcing shells: adding ribs aligned with stress directions and using thicker walls on regions of high stress. Most previous work can generate either only ribs or only variablethickness walls. However, in the general case, neither approach by itself will provide optimal solutions. By using a more precise volume model, our method is capable of producing optimized structures with the full range of qualitative behaviors: from ribs to walls, and smoothly transitioning in between. Our method includes new algorithms for determining the layout of reinforcement structure elements, and an efficient algorithm to optimize their shape, minimizing a non-linear non-convex functional at a fraction of the cost and with better optimality compared to standard solvers. We demonstrate the optimization results for a variety of shapes, and the improvements it yields in the strength of 3D-printed objects.'
url_pdf: https://vcgdata.isti.cnr.it/Publications/2020/GPZ20/fgil_shellopt_tex.pdf
---
{{< figure src="https://vcgdata.isti.cnr.it/Publications/2020/GPZ20/0_thumb_reinforced.png" >}}
