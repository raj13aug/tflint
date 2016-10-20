package aws

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/wata727/tflint/issue"
)

func DetectAwsInstanceInvalidType(list *ast.ObjectList, file string) []*issue.Issue {
	var issues = []*issue.Issue{}

	for _, item := range list.Filter("resource", "aws_instance").Items {
		instanceTypeToken := item.Val.(*ast.ObjectType).List.Filter("instance_type").Items[0].Val.(*ast.LiteralType).Token
		instanceTypeKey := strings.Trim(instanceTypeToken.Text, "\"")

		if !ValidInstanceType[instanceTypeKey] {
			issue := &issue.Issue{
				Type:    "WARNING",
				Message: fmt.Sprintf("%s is invalid instance type.", instanceTypeToken.Text),
				Line:    instanceTypeToken.Pos.Line,
				File:    file,
			}
			issues = append(issues, issue)
		}
	}

	return issues
}

var ValidInstanceType = map[string]bool{
	"t2.nano":     true,
	"t2.micro":    true,
	"t2.small":    true,
	"t2.medium":   true,
	"t2.large":    true,
	"m4.large":    true,
	"m4.xlarge":   true,
	"m4.2xlarge":  true,
	"m4.4xlarge":  true,
	"m4.10xlarge": true,
	"m4.16xlarge": true,
	"m3.medium":   true,
	"m3.large":    true,
	"m3.xlarge":   true,
	"m3.2xlarge":  true,
	"c4.large":    true,
	"c4.2xlarge":  true,
	"c4.4xlarge":  true,
	"c4.8xlarge":  true,
	"c3.large":    true,
	"c3.xlarge":   true,
	"c3.2xlarge":  true,
	"c3.4xlarge":  true,
	"c3.8xlarge":  true,
	"x1.32xlarge": true,
	"r3.large":    true,
	"r3.xlarge":   true,
	"r3.2xlarge":  true,
	"r3.4xlarge":  true,
	"r3.8xlarge":  true,
	"p2.xlarge":   true,
	"p2.8xlarge":  true,
	"p2.16xlarge": true,
	"g2.2xlarge":  true,
	"g2.8xlarge":  true,
	"i2.xlarge":   true,
	"i2.2xlarge":  true,
	"i2.4xlarge":  true,
	"i2.8xlarge":  true,
	"d2.xlarge":   true,
	"d2.2xlarge":  true,
	"d2.4xlarge":  true,
	"d2.8xlarge":  true,
	"t1.micro":    true,
	"m1.small":    true,
	"m1.medium":   true,
	"m1.large":    true,
	"m1.xlarge":   true,
	"c1.medium":   true,
	"c1.xlarge":   true,
	"cc2.8xlarge": true,
	"cg1.4xlarge": true,
	"m2.xlarge":   true,
	"m2.2xlarge":  true,
	"m2.4xlarge":  true,
	"cr1.8xlarge": true,
	"hi1.4xlarge": true,
	"hs1.8xlarge": true,
}