Benchmark
=========

![](benchmark_result.png)


```
go test -bench=.
PASS
BenchmarkDecode_1_webp_a_chai2010_webp-4                                       	     500	   3256979 ns/op
BenchmarkDecode_1_webp_a_chai2010_webp_cbuf-4                                  	     500	   3036237 ns/op
BenchmarkDecode_1_webp_a_x_image_webp-4                                        	     200	   8059742 ns/op
BenchmarkDecode_1_webp_ll_chai2010_webp-4                                      	     500	   3733524 ns/op
BenchmarkDecode_1_webp_ll_chai2010_webp_cbuf-4                                 	     500	   3741486 ns/op
BenchmarkDecode_1_webp_ll_x_image_webp-4                                       	     200	   8406712 ns/op
BenchmarkDecode_2_webp_a_chai2010_webp-4                                       	     500	   3280379 ns/op
BenchmarkDecode_2_webp_a_chai2010_webp_cbuf-4                                  	     500	   3108146 ns/op
BenchmarkDecode_2_webp_a_x_image_webp-4                                        	     200	   6959859 ns/op
BenchmarkDecode_2_webp_ll_chai2010_webp-4                                      	     500	   2841670 ns/op
BenchmarkDecode_2_webp_ll_chai2010_webp_cbuf-4                                 	     500	   2855405 ns/op
BenchmarkDecode_2_webp_ll_x_image_webp-4                                       	     200	   6921478 ns/op
BenchmarkDecode_3_webp_a_chai2010_webp-4                                       	     100	  11406906 ns/op
BenchmarkDecode_3_webp_a_chai2010_webp_cbuf-4                                  	     100	  10064477 ns/op
BenchmarkDecode_3_webp_a_x_image_webp-4                                        	      50	  23756447 ns/op
BenchmarkDecode_3_webp_ll_chai2010_webp-4                                      	     100	  11229056 ns/op
BenchmarkDecode_3_webp_ll_chai2010_webp_cbuf-4                                 	     100	  10147847 ns/op
BenchmarkDecode_3_webp_ll_x_image_webp-4                                       	      50	  22879775 ns/op
BenchmarkDecode_4_webp_a_chai2010_webp-4                                       	     500	   2236276 ns/op
BenchmarkDecode_4_webp_a_chai2010_webp_cbuf-4                                  	     500	   2152100 ns/op
BenchmarkDecode_4_webp_a_x_image_webp-4                                        	     200	   5490799 ns/op
BenchmarkDecode_4_webp_ll_chai2010_webp-4                                      	    1000	   2233461 ns/op
BenchmarkDecode_4_webp_ll_chai2010_webp_cbuf-4                                 	    1000	   2116654 ns/op
BenchmarkDecode_4_webp_ll_x_image_webp-4                                       	     300	   5253699 ns/op
BenchmarkDecode_5_webp_a_chai2010_webp-4                                       	     300	   5390873 ns/op
BenchmarkDecode_5_webp_a_chai2010_webp_cbuf-4                                  	     300	   5304526 ns/op
BenchmarkDecode_5_webp_a_x_image_webp-4                                        	     100	  13717136 ns/op
BenchmarkDecode_5_webp_ll_chai2010_webp-4                                      	     500	   3615680 ns/op
BenchmarkDecode_5_webp_ll_chai2010_webp_cbuf-4                                 	     500	   3450953 ns/op
BenchmarkDecode_5_webp_ll_x_image_webp-4                                       	     200	   7962092 ns/op
BenchmarkDecode_blue_purple_pink_large_lossless_chai2010_webp-4                	     200	   8242255 ns/op
BenchmarkDecode_blue_purple_pink_large_lossless_chai2010_webp_cbuf-4           	     200	   7394029 ns/op
BenchmarkDecode_blue_purple_pink_large_lossless_x_image_webp-4                 	     100	  19587284 ns/op
BenchmarkDecode_blue_purple_pink_large_no_filter_lossy_chai2010_webp-4         	     500	   3484234 ns/op
BenchmarkDecode_blue_purple_pink_large_no_filter_lossy_chai2010_webp_cbuf-4    	     500	   3171442 ns/op
BenchmarkDecode_blue_purple_pink_large_no_filter_lossy_x_image_webp-4          	     200	   8381170 ns/op
BenchmarkDecode_blue_purple_pink_large_normal_filter_lossy_chai2010_webp-4     	     300	   3830255 ns/op
BenchmarkDecode_blue_purple_pink_large_normal_filter_lossy_chai2010_webp_cbuf-4	     500	   3526849 ns/op
BenchmarkDecode_blue_purple_pink_large_normal_filter_lossy_x_image_webp-4      	     100	  13233320 ns/op
BenchmarkDecode_blue_purple_pink_large_simple_filter_lossy_chai2010_webp-4     	     500	   3586785 ns/op
BenchmarkDecode_blue_purple_pink_large_simple_filter_lossy_chai2010_webp_cbuf-4	     500	   3317595 ns/op
BenchmarkDecode_blue_purple_pink_large_simple_filter_lossy_x_image_webp-4      	     100	  10658441 ns/op
BenchmarkDecode_blue_purple_pink_lossless_chai2010_webp-4                      	    2000	    751205 ns/op
BenchmarkDecode_blue_purple_pink_lossless_chai2010_webp_cbuf-4                 	    2000	    711052 ns/op
BenchmarkDecode_blue_purple_pink_lossless_x_image_webp-4                       	    1000	   1813983 ns/op
BenchmarkDecode_blue_purple_pink_lossy_chai2010_webp-4                         	    3000	    343784 ns/op
BenchmarkDecode_blue_purple_pink_lossy_chai2010_webp_cbuf-4                    	    5000	    331602 ns/op
BenchmarkDecode_blue_purple_pink_lossy_x_image_webp-4                          	    2000	    848802 ns/op
BenchmarkDecode_gopher_doc_1bpp_lossless_chai2010_webp-4                       	   30000	     55913 ns/op
BenchmarkDecode_gopher_doc_1bpp_lossless_chai2010_webp_cbuf-4                  	   30000	     44023 ns/op
BenchmarkDecode_gopher_doc_1bpp_lossless_x_image_webp-4                        	   10000	    122832 ns/op
BenchmarkDecode_gopher_doc_2bpp_lossless_chai2010_webp-4                       	   20000	     68660 ns/op
BenchmarkDecode_gopher_doc_2bpp_lossless_chai2010_webp_cbuf-4                  	   30000	     55766 ns/op
BenchmarkDecode_gopher_doc_2bpp_lossless_x_image_webp-4                        	   10000	    157441 ns/op
BenchmarkDecode_gopher_doc_4bpp_lossless_chai2010_webp-4                       	   20000	     93446 ns/op
BenchmarkDecode_gopher_doc_4bpp_lossless_chai2010_webp_cbuf-4                  	   20000	     77913 ns/op
BenchmarkDecode_gopher_doc_4bpp_lossless_x_image_webp-4                        	    5000	    218572 ns/op
BenchmarkDecode_gopher_doc_8bpp_lossless_chai2010_webp-4                       	   10000	    160015 ns/op
BenchmarkDecode_gopher_doc_8bpp_lossless_chai2010_webp_cbuf-4                  	   10000	    144743 ns/op
BenchmarkDecode_gopher_doc_8bpp_lossless_x_image_webp-4                        	    5000	    324870 ns/op
BenchmarkDecode_tux_lossless_chai2010_webp-4                                   	     500	   2761451 ns/op
BenchmarkDecode_tux_lossless_chai2010_webp_cbuf-4                              	     500	   2573258 ns/op
BenchmarkDecode_tux_lossless_x_image_webp-4                                    	     200	   6837351 ns/op
BenchmarkDecode_video_001_lossy_chai2010_webp-4                                	    3000	    441210 ns/op
BenchmarkDecode_video_001_lossy_chai2010_webp_cbuf-4                           	    3000	    463008 ns/op
BenchmarkDecode_video_001_lossy_x_image_webp-4                                 	    1000	   1066743 ns/op
BenchmarkDecode_video_001_chai2010_webp-4                                      	    3000	    513669 ns/op
BenchmarkDecode_video_001_chai2010_webp_cbuf-4                                 	    3000	    460857 ns/op
BenchmarkDecode_video_001_x_image_webp-4                                       	    2000	   1149123 ns/op
BenchmarkDecode_yellow_rose_lossless_chai2010_webp-4                           	     500	   3916892 ns/op
BenchmarkDecode_yellow_rose_lossless_chai2010_webp_cbuf-4                      	     500	   3702229 ns/op
BenchmarkDecode_yellow_rose_lossless_x_image_webp-4                            	     100	  10895141 ns/op
BenchmarkDecode_yellow_rose_lossy_with_alpha_chai2010_webp-4                   	     500	   2340825 ns/op
BenchmarkDecode_yellow_rose_lossy_with_alpha_chai2010_webp_cbuf-4              	     500	   2531316 ns/op
BenchmarkDecode_yellow_rose_lossy_with_alpha_x_image_webp-4                    	     200	   7444750 ns/op
BenchmarkDecode_yellow_rose_lossy_chai2010_webp-4                              	     500	   2212347 ns/op
BenchmarkDecode_yellow_rose_lossy_chai2010_webp_cbuf-4                         	    1000	   2001677 ns/op
BenchmarkDecode_yellow_rose_lossy_x_image_webp-4                               	     300	   5464172 ns/op
ok  	github.com/chai2010/webp/bench	143.753s
```
