# go-snr-post-parser
A library that parses rows of Bolagsverket's file data using reflection on structs with tags

### How to use
See structs.go and main_test.go

### Tests
The tests have been done by the same person that made (most of) the structs.
So any misunderstanding of the data layout will go unnoticed.

One test expect a file with example posts from Bolagsverket:
Filexempel dagliga aviseringar_20160211.zip
If it is present we use it to test post#:
100, 800, 808, 810, 811, 812, 813, 814, 816, 820, 830, 840,
880, 881, 883, 884, 885, 886, 888, 892, 893, 894, 895, 896,
930, 931, 970
 
### Todo
All the posts in Bolagsverkets document
Aviseringsposter.pdf
that are not present in structs.go

