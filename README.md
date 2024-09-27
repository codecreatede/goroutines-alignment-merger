# goroutines-alignment-merger

- large scale alignment merger for phylogenomics.
- a large phylogenomics datasets merged in less than few minutes for the genome scale alignment.
- check the alignment-go package which has all alignment utilities.

```
# run this on the fasta if your alignment fasta are NCBI formatted to make it linearized.

awk '/^>/ {printf("\n%s\n",$0);next; } { printf("%s",$0);}  \
                         END {printf("\n");}' fastafile > alignment.go.fasta

go run main.go
  -alignmentfile string
        file (default "path to the alignment file")
  -title string
        gene name (default "name of the header of the gene")
[gauravsablok@fedora]~/Desktop/codecreatede/goroutines-alignment-clipper% \
      go run main.go -alignmentfile ./sample-files/alignment.go.fasta -title merged
> merged
ATG.....

a large phylogenomics datasets merged in less than few minutes for the genome scale alignment. 

```
Gaurav Sablok
