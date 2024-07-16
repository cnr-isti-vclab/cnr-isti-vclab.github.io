---

title: "Capacitive Touch Sensing on General 3D Surfaces"
authors:
- Gianpaolo Palma
- Narges Pourjafarian
- Jürgen Steimle
- Paolo Cignoni
# Note use this date format
date: "2024-07-01T00:00:00Z"
doi: "10.1145/3658185"


# Publication type.
# Accepts a single type but formatted as a YAML list (for Hugo requirements).
# Enter a publication type from the CSL standard.
publication_types: ["article-journal"]

# Publication name and optional abbreviated publication name.
publication: ACM Transactions on Graphics - [SIGGRAPH 2024 Best Paper Honorable Mention](https://blog.siggraph.org/2024/06/siggraph-2024-technical-papers-awards-best-papers-honorable-mentions-and-test-of-time.html/)
#publication_short: "ACM Trans. on Graphics"

abstract: Mutual-capacitive sensing is the most common technology for detecting multi-touch, especially on flat and simple curvature surfaces. Its extension to a more complex shape is still challenging, as a uniform distribution of sensing electrodes is required for consistent touch sensitivity across the surface. To overcome this problem, we propose a method to adapt the sensor layout of common capacitive multi-touch sensors to more complex 3D surfaces, ensuring high-resolution, robust multi-touch detection. The method automatically computes a grid of transmitter and receiver electrodes with as regular distribution as possible over a general 3D shape. It starts with the computation of a proxy geometry by quad meshing used to place the electrodes through the dual-edge graph. It then arranges electrodes on the surface to minimize the number of touch controllers required for capacitive sensing and the number of input/output pins to connect the electrodes with the controllers. We reach these objectives using a new simplification and clustering algorithm for a regular quad-patch layout. The reduced patch layout is used to optimize the routing of all the structures (surface grooves and internal pipes) needed to host all electrodes on the surface and inside the object's volume, considering the geometric constraints of the 3D shape. Finally, we print the 3D object prototype ready to be equipped with the electrodes. We analyze the performance of the proposed quad layout simplification and clustering algorithm using different quad meshing and characterize the signal quality and accuracy of the capacitive touch sensor for different non-planar geometries. The tested prototypes show precise and robust multi-touch detection with good Signal-to-Noise Ratio and spatial accuracy of about 1mm.

tags:
- Geometry Processing 
- Rendering and Advanced UI


featured: false

image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
links:
  - name: Additional material
    url: 'https://vcgdata.isti.cnr.it/Publications/2024/PPSC24-3dTouch/supplemental_material.pdf'
url_pdf: 'https://vcgdata.isti.cnr.it/Publications/2024/PPSC24-3dTouch/paper.pdf'
url_video: 'https://www.youtube.com/watch?v=m8gcV1zCJX0'
url_poster: 'https://vcgdata.isti.cnr.it/Publications/2024/PPSC24-3dTouch/poster.pdf'

---
{{< figure class="pubFig" src="https://vcgdata.isti.cnr.it/Publications/2024/PPSC24-3dTouch/papers_637s3.jpg" >}}
{{< figure class="pubFig" src="https://vcgdata.isti.cnr.it/Publications/2024/PPSC24-3dTouch/maxPlanck.jpg" >}}
{{< figure class="pubFig" src="https://vcgdata.isti.cnr.it/Publications/2024/PPSC24-3dTouch/cube.jpg" >}}
{{< youtube "m8gcV1zCJX0" >}}