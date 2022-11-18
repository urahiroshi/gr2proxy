package internal

import "encoding/json"

type Header map[string][]string
type Trailer map[string][]string

type Response struct {
	Body       string  `json:"Body"`
	StatusCode int     `json:"StatusCode"`
	Header     Header  `json:"Header"`
	Trailer    Trailer `json:"Trailer"`
}

func NewResponse() *Response {
	header := make(Header)
	trailer := make(Trailer)
	return &Response{Header: header, Trailer: trailer}
}

// FixtureMap gets Response object by such a keys:
// response := fixtureMap[serviceName][methodName][requestHttpHeader][requestBodyAsJsonString]
//
// It assumes this json output
// {
// 	"helloworld.Greeter": {
// 		"SayGoodbye": {
//			"*": {
//				"{\"hoge\":\"hogehoge\",\"fuga\":\"fugafuga\"}": {
//					"Body": "{\"value\":\"hoge\"}",
//					"StatusCode": 200,
//					"Header": {
//						"Content-Type": ["application/grpc"]
//					},
//					"Trailer": {
//						"Grpc-Message": [""],
//						"Grpc-Status": ["0"]
//					}
//				}
//			}
// 		}
// 	}
// }
type FixtureMap map[string]map[string]map[string]map[string]*Response

func NewFixtureMapFromJson(fixtureJson []byte) (FixtureMap, error) {
	var fixtureMap FixtureMap
	err := json.Unmarshal(fixtureJson, &fixtureMap)
	if err != nil {
		return nil, err
	}
	return fixtureMap, nil
}

func (m FixtureMap) SaveResponse(
	serviceName string,
	methodName string,
	requestHttpHeader string,
	requestBody string,
	response *Response,
) {
	if m[serviceName] == nil {
		m[serviceName] = make(map[string]map[string]map[string]*Response)
	}
	if m[serviceName][methodName] == nil {
		m[serviceName][methodName] = make(map[string]map[string]*Response)
	}
	if m[serviceName][methodName][requestHttpHeader] == nil {
		m[serviceName][methodName][requestHttpHeader] = make(map[string]*Response)
	}
	m[serviceName][methodName][requestHttpHeader][requestBody] = response
}

func (m FixtureMap) GetResponse(serviceName, methodName, requestHttpHeader, requestBody string) *Response {
	return m[serviceName][methodName][requestHttpHeader][requestBody]
}

func (m FixtureMap) ToJson() ([]byte, error) {
	return json.Marshal(m)
}
