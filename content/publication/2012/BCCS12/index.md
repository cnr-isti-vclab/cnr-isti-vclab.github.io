---
title: 'A Low-Memory, Straightforward and Fast Bilateral Filter Through Subsampling in Spatial Domain'
authors:
  - Francesco Banterle
  - Massimiliano Corsini
  - Paolo Cignoni
  - Roberto Scopigno
date: '2012-01-01T00:00:00Z'
publication_types: ['paper-conference']
publication: '*Computer Graphics Forum*'
featured: false

abstract: 'abstract
 	
 	
 	
 
 
 
 
 
 
 
 function log(msg) {
 	var textarea = document.getElementById("log-area");
 	textarea.innerHTML += (msg + "\n");
 	textarea.scrollTop = textarea.scrollHeight;
 }
 
 $(document).ready(function() {
     $( "#sliderSigmaS" ).slider({
         range: "min",
         min: 1,
         max: 100,
         value: 1,
         slide: function(event, ui){
             $("#paraS").text(ui.value);
         },
         change: function(event, ui){
             $("#paraS").text(ui.value);
         }
     });
     
     $( "#sliderSigmaR" ).slider({
         range: "min",
         min: 1,
         max: 100,
         value: 1,
         slide: function(event, ui){
             $("#paraR").text(ui.value);
         },
         change: function(event, ui){
             $("#paraR").text(ui.value);
         }
     });    
 });
 
 
 SpiderGL.openNamespace();
 
 function CanvasHandler() {
 }
 
 CanvasHandler.prototype = {
 
     /////////////////////////////////////////////////////////////////////////////////////
 	onInitialize : function () {
 		var gl = this.ui.gl;
 
 		var quadPositions = new Float32Array([
 			-1.0, -1.0,
 			 1.0, -1.0,
 			-1.0,  1.0,
 			 1.0,  1.0
 		]);
 		
 		var quadTexcoords = new Float32Array([
 			0.0, 0.0,
 			1.0, 0.0,
 			0.0, 1.0,
 			1.0, 1.0
 		]);		
 		
         var quad = new SglModel(gl, {
 	        data : {
 		        vertexBuffers : {
 			        "positionBuffer" : { typedArray : quadPositions},
 			        "texcoordBuffer" : { typedArray : quadTexcoords}
 		        },
 		        indexBuffers : {
 		        }
 	        },
 	        access : {
 		        vertexStreams : {
 			        "position" : {
 				        buffer : "positionBuffer",
 				        size   : 2,
 				        type   : SGL_FLOAT32,
 				        stride : 2 * SGL_SIZEOF_FLOAT32,
 				        offset : 0
 			        },
 			        "texcoord" : {
 				        buffer : "texcoordBuffer",
 				        size   : 2,
 				        type   : SGL_FLOAT32,
 				        stride : 2 * SGL_SIZEOF_FLOAT32,
 				        offset : 0
 			        }			        
 		        },
 		        primitiveStreams : {
 			        "triangles" : {
 				        mode   : SGL_TRIANGLE_STRIP,
 				        first  : 0,
 				        count  : 4 //number of verticies
 			        }
 		        }
 	        },
 	        semantic : {
 		        bindings : {
 			        "COMMON" : {
 				        vertexStreams : {
 					        "POSITION" : [ "position" ],
 					        "TEXCOORD":  [ "texcoord" ]
 				        },
 				        primitiveStreams : {
 					        "FILL"  : [ "triangles" ]
 				        }
 			        }
 		        },
 		        chunks : {
 			        "chunk0" : {
 				        techniques : {
 					        "COMMON" : {
 						        binding : "COMMON"
 					        }
 				        }
 			        }
 		        }
 	        },
 	        logic : {
 		        parts : {
 			        "main" : {
 				        chunks : [ "chunk0" ]
 			        }
 		        }
 	        }
         });	
         this.quad = quad;
 
 		var fsQuad = new SglFragmentShader(gl,"\
    			precision highp float;                                                  \n\
                                                                                     \n\
             uniform sampler2D   sTexture;\n\
             uniform sampler2D   sTextureSamples;                                           \n\
             varying vec2        vTexcoord;                                          \n\
             uniform float scaleX, scaleY, scaleY2, sigma_r2, sigma_s2;\n\
                                                                                     \n\
                                                                                     \n\
             void main(void)                                                         \n\
             {                                                                       \n\
 	            vec3 colorRef = texture2D(sTexture, vTexcoord).xyz;                   \n\
                 vec3 color = vec3(0.0,0.0,0.0);\n\
                 float yFetch = vTexcoord.y*scaleY2;\n\
                 float weight = 0.0;\n\
                 for(int i=0;i'
image:
#    caption: 'Image'
    focal_point: ''
    preview_only: false
    share: false
url_pdf: https://vcgdata.isti.cnr.it/Publications/2012/BCCS12/j.1467-8659.2011.02078.x.pdf
---
[link to the official journal](http://onlinelibrary.wiley.com/doi/10.1111/j.1467-8659.2011.02078.x/abstract)

[Piccante](https://vcgdata.isti.cnr.it/Publicstions/2012/BCCS12/www.piccantelib.net)

[OpenGL ES 2.0 shaders source code and data](https://vcgdata.isti.cnr.it/Publicstions/2012/BCCS12/shaders_data.zip)

