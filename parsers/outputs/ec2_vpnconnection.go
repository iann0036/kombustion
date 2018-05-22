package outputs
import (
	"github.com/KablamoOSS/kombustion/types"
)

func ParseEC2VPNConnection(name string, data string) (cf types.ValueMap, err error) {
	
	cf = types.ValueMap{
		name: types.ValueMap{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-EC2VPNConnection-" + name,
				},
			},
		},
	}

	

	return
}