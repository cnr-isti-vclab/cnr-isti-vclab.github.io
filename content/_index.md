---
# Leave the homepage title empty to use the site title
title:
type: landing

sections:
  - block: markdown
    content:
      title: 
        Visual Computing Lab
      text:      
        The Visual Computing Lab is a research group of the [CNR-ISTI](https://www.isti.cnr.it/) (Italian National Research Council - Institute for Information Science and Technology) in Pisa, Italy. The lab is headed by [Paolo Cignoni]({{< relref "/authors/paolo-cignoni" >}}), and is located in the CNR-ISTI [building](https://www.isti.cnr.it/en/about/reach-us) in the CNR Research Area in Pisa, Italy.<br/>
        {{< gallery album="gallery" >}}
    design:
      css_class: universal-wrapper
   
  
  - block: collection
    content:
      title: Latest News
      subtitle:
      text:
      count: 5
      filters:
        author: ''
        category: ''
        exclude_featured: false
        publication_type: ''
        tag: ''
      offset: 0
      order: desc
      page_type: post
    design:
      view: showcase 
      columns: '1'
      css_class: universal-wrapper
---