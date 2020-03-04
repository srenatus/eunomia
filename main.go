package main

import (
	"flag"
	"fmt"
	"log"

	iamutil "github.com/mhausenblas/eunomia/pkg/iamutil"
)

func main() {
	iamfile := flag.String("iam", "", "The IAM policy file you want to use as an input")
	rbacfile := flag.String("rbac", "", "The RBAC role file you want to use as an input")
	flag.Parse()
	if *iamfile == "" && *rbacfile == "" {
		log.Fatalln("Need at least either an IAM policy file or RBAC role file to continue, sorry :(")
	}
	regocontent := ""
	if *iamfile != "" {
		log.Printf("Parsing IAM file %v", *iamfile)
		iamcontent, err := iamutil.DumpJSON(*iamfile)
		if err != nil {
			log.Fatalln(fmt.Sprintf("Can't parse IAM policy file: %v", err))
		}
		regocontent = iamcontent
	}
	if *rbacfile != "" {
		log.Printf("Parsing RBAC file %v", *rbacfile)

	}
	fmt.Println(regocontent)
}
