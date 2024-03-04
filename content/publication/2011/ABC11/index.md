---
title: 'A Survey of Specularity Removal Methods'
authors:
  - Alessandro Artusi
  - Francesco Banterle
  - Dmitry  Chetverikov
date: '2011-01-01T00:00:00Z'
publication_types: ['paper-conference']
publication: '*Computer Graphics Forum*'
featured: false

abstract: 'abstract
 	
 	
 	
 
 
 
 
 
 AbstractIn the real world, materials often show both diffuse and specular reflections. To separate these two reflection components may help applications which need consistent object surface appearance.
 In fact, many algorithms used in numerous tasks of computer vision, computer graphics and image processing,  work under the assumption of perfect Lambertian surfaces (or perfect diffuse reflection). They consider specular pixels (or highlights) as outliers or noise.
 Moreover, the two reflection components can be processed separately and afterwards recombined to produce particular visual effects [MZK*06] (dichromatic editing  Figure.1).
 
 
 
 
 
 
 
 Figure 2: Example of an image where the presence of highlights generates the loss of details and colour information. Details and colours are completely washed out in the highlights region.
 
 
 
 
 Several specularity removal techniques are available in literature; they differ in the information they use and in how it is used. Table 1 summarises how these techniques may be classified and in which paragraph of this survey they are explained. The notation used in Table 1 for the categories is the following: CSA, colour space analysis; NA, neighbourhood analysis; POL, polarization; IS, image sequences; MFI, multiple-flash images.
 The notation used in Table 1 for the techniques is the following: CRM, using colour reflection model [KSK87, KSK88]; TDA, 2D diagram approach [ST95b, SK00]; BM, Bajcsy et al. method [BLL96]; USFI, use of specular-free images [TI05b, YCK06, SC09]; PDE, PDE approach [MZK*06]; IT, in-panting technique [TLQ06]; CC, use of colour information and classifier [TFA05]; TS, separation of highlight reflections on textured surfaces [TLQ06]; FR, Fresnel methods [Ang07]; HM, histogram methods [CGS06]; HFLF, high-low frequency separation [LPD07, NGR06]; MBS, multi-baseline-stereo [LB92, LYK*03, LS01]; MII, deriving intrinsic images from image sequences with illumination changes [Wei01]; CP, colour and polarisation methods [NFB97, KLHS02, MPC*07, USGG04]; MF, multi-flash methods [FRTT04, ARNL05].
 
 
     
       Table 1: The
         classification of highlights removal methods. These methods are
         classified based on the number of image used as input an the type of
         information used. L and G respectively stand for global and local
         approach. The number, in the first column, indicates the paragraph in
         the paper where the technique appears in the survey. 
       
       
         
           
             
         
         
           
           
           
           
           Single
               Image
           
           Multiple
               Images
           
         
         
           
           
           
           
           
             
           
             
         
         
           Tech. - Sec.
           
           Type
           
           CSA
           
           NA
           
           IS
           
           MFI
           
           POL
           
         
         
           
             
         
         
           CRM - 3.1.1
           
           G
           
           x
           
           -
           -
           -
           -
         
         
           TDA - 3.1.2
           
           L
           
           x
           
           -
           -
           -
           -
         
         
           BM - 3.1.3
           
           G
           
           x
           
           -
           -
           -
           -
         
         
           USFI - 3.2.1
           
           L
           
           -
           x
           
           -
           -
           -
         
         
           PDE - 3.2.2
           
           L
           -
           x
           
           -
           -
           -
         
         
           IT - 3.2.3
           
           L
           -
           x
           
           -
           -
           -
         
         
           TS - 3.2.4
           
           L
           -
           x
           
           -
           -
           -
         
         
           CC - 3.2.5
           
           L
           -
           x
           
           -
           -
           -
         
         
           FR - 3.2.6
           
           G
           
           -
           x
           
           -
           -
           -
         
         
           HM - 4
           
           G
           
           -
           -
           x
           
           -
           -
         
         
           HFLF - 4
           
           G
           -
           -
           x
           
           -
           -
         
         
           MBS - 4.1
           
           G
           -
           -
           x
           
           -
           -
         
         
           MII - 4.2
           
           L
           
           -
           -
           x
           
           -
           -
         
         
           CP - 4.3
           
           G
           
           -
           -
           -
           -
           
           x
           
         
         
           MF - 4.4
           
           L
           
           -
           -
           -
           x
           
           -
         
         
           
             
         
       
     
     
     
 
 Discussion
 Several factors may influence which approach is superior to another one, such as the number of images
 to be captured, automatic operation vs. manual help, light constraints and
 the reflectance model used, merits of quality, and the hardware used during the acquisition phase. Table
 2 provides a comparative summary of the specular removal methods discussed in this survey. 
 
 
 
     
       Table 2: Comparison of
         the highlight removal techniques by their characteristics. The methods
         are grouped into the single-image category (top) and the multiple-image
         category (bottom). User Interaction:
         MS manual segmentation, A automatic, NS no segmentation, Light
           Requirement: IC illuminant compensation, DM dichromatic
         reflectance model, and FM flash model. Hardware:
         S single camera, M multiple camera, * polarised filters, FS flash system
       
       
       
         
           
             
         
         
           Technique
           
           Images
           
           User
               Interaction
           
           Light
               Requirement
           
           Hardware
           
         
         
           
             
           
         
         
           Colour Space [KSK87, KSK88, ST95b, SK00, BLL96]
           
           1
           
           MS
           
           IC-DM
           
           S
           
         
         
           Specular-Free Image [TI05b, YCK06, SC09]
           
           1
           
           A
           
           IC-DM
           S
         
         
           Inpainting [TLQ06]
           
           1
           
           MS
           
           IC-DM
           S
         
         
           PDE [MZK*06]
           
           1
           
           -
           
           IC-DM
           S
         
         
           Textured Surfaces [TLQ06]
           
           1
           
           MS
           
           IC-DM
           S
         
         
           Colour Classifier [TFA05]
           
           1
           
           -
           
           -
           
           S
         
         
           Fresnel Coefficient [Ang07]
           
           1
           
           MS
           
           No IC-DM
           S
         
         
           Multi-Baseline Stereo [LB92, LYK*03, LS01]
           
           50-70
           
           NS
           DM
           
           M
           
         
         
           Deriving Intrinsic Images [Wei01]
           
           40-70
           
           NS
           Lighting
             changes
           
           S
           
         
         
           Polarization [NFB97, KLHS02, MPC*07, USGG04]
           
           6-10
           
           NS
           DM
           
           S*
           
         
         
           Histogram [CGS06]
           
           200
           
           NS
           -
           
           S
           
         
         
           High-Low Frequency Separation [LPD07, NGR06]
           
           4, 32
           
           NS
           
           -
           
           S
           
         
         
           Multi-flash [FRTT04, ARNL05]
           
           4-8,2
           
           NS
           
           FM
           
           FS
           
         
         
           
             
         
       
     
     
 
 
 Finally we have performed several experiment, to compare different specularity removal techniques, on various input images starting from single-colour, multi-colour and textured surfaces, increasing the texture complexity.
 
 
 ReferencesFor the references please see the paper.'
image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
url_pdf: https://vcgdata.isti.cnr.it/Publications/2011/ABC11/j.1467-8659.2011.01971.x.pdf
---
