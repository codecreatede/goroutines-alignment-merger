package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
 Author Gaurav Sablok
 Universitat Potsdam
 Date 2024-9-16

 a genome alignment clipper, which takes a fasta file aligned and clipped all the blocks
 with -, mismatch as invoked by the pattern. Processing large scale alignments for phylogenomics
 The biggest problem dealing with the large scale phylogenomics is that alignments are hard to filter
 and takes a load of memory for the same. A golang implementation using go routines to do everything
 faster. It only filters the blocks with the - and not mismatches.
 Check out the complete alignmentGo package for dealing with all alignment functions.
*/

func main() {
	alignment := flag.String("alignmentfile", "path to the alignment file", "file")
	title := flag.String("title", "name of the header of the gene", "gene name")

	flag.Parse()

	type alignmentID struct {
		id string
	}

	type alignmentSeq struct {
		seq string
	}

	type alignmentFilterID struct {
		id string
	}

	type alignmentFilterSeq struct {
		seq string
	}

	fOpen, err := os.Open(*alignment)
	if err != nil {
		log.Fatal(err)
	}

	// file opening and storing into structs
	fRead := bufio.NewScanner(fOpen)
	alignmentReadID := []alignmentID{}
	alignmentReadSeq := []alignmentSeq{}
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignmentReadID = append(alignmentReadID, alignmentID{
				id: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignmentReadSeq = append(alignmentReadSeq, alignmentSeq{
				seq: string(line),
			})
		}
	}
	// implementing runes and string conversion as iter.

	var seqKeep string
	for b := range alignmentReadSeq {
		for j := range alignmentReadSeq[b].seq {
			seqKeep += string(alignmentReadSeq[b].seq[j])
		}
	}

	fmt.Println(">", *title, "\n", seqKeep)
}
