---
# delete the following three lines if you want that your page appears:
# _build:
#  render: always
#  list: never

title: 'MoReLab: A Software for User-Assisted 3D Reconstruction'
authors:
  - Arslan Siddique
  - Francesco Banterle
  - Massimiliano Corsini
  - Paolo Cignoni
  - Daniel Sommerville
  - Chris Joffe

date: '2023-07-17T00:00:00Z'
doi: 'https://doi.org/10.3390/s23146456'


publication_types: ['article-journal']

# Publication name and optional abbreviated publication name.
publication: '*Sensors, 23*(14)'

abstract: We present MoReLab, a tool for user-assisted 3D reconstruction. This reconstruction requires an understanding of the shapes of the desired objects. Our experiments demonstrate that existing Structure from Motion (SfM) software packages fail to estimate accurate 3D models in low-quality videos due to several issues such as low resolution, featureless surfaces, low lighting, etc. In such scenarios, which are common for industrial utility companies, user assistance becomes necessary to create reliable 3D models. In our system, the user first needs to add features and correspondences manually on multiple video frames. Then, classic camera calibration and bundle adjustment are applied. At this point, MoReLab provides several primitive shape tools such as rectangles, cylinders, curved cylinders, etc., to model different parts of the scene and export 3D meshes. These shapes are essential for modeling industrial equipment whose videos are typically captured by utility companies with old video cameras (low resolution, compression artifacts, etc.) and in disadvantageous lighting conditions (low lighting, torchlight attached to the video camera, etc.). We evaluate our tool on real industrial case scenarios and compare it against existing approaches. Visual comparisons and quantitative results show that MoReLab achieves superior results with regard to other user-interactive 3D modeling tools.

# Summary. An optional shortened abstract.
summary: A software for user-assisted 3D reconstruction.

tags:
  - 3D Reconstruction
featured: false

# links:
# - name: ""
url_pdf: https://www.mdpi.com/1424-8220/23/14/6456
url_code: https://github.com/cnr-isti-vclab/MoReLab

# Featured image
# To use, add an image named `featured.jpg/png` to your page's folder.
image:
  caption: 'Graphical User Interface of MoReLab'
  focal_point: ''
  preview_only: false

---

The source code for MoReLab can be found on our [Github repository](https://github.com/cnr-isti-vclab/MoReLab).