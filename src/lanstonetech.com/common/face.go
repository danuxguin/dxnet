package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Faceslice struct {
	Face []struct {
		Attribute struct {
			Age struct {
				Range float64
				Value float64
			}
			Gender struct {
				Confidence float64
				Value      string
			}
			Race struct {
				Confidence float64
				Vaule      string
			}
		}
		Face_id  string
		Position struct {
			Center struct {
				X float64
				Y float64
			}

			Eye_left struct {
				X float64
				Y float64
			}

			Eye_right struct {
				X float64
				Y float64
			}

			Mouth_left struct {
				X float64
				Y float64
			}

			Mouth_right struct {
				X float64
				Y float64
			}
			Nose struct {
				X float64
				Y float64
			}

			Height float64
			Width  float64
		}

		Tag string
	}
	Img_height int
	Img_id     string
	Img_width  int
	Session_id string
	url        string
}

func FaceDecodeDetect(data []byte) Faceslice {
	var f Faceslice
	json.Unmarshal(data, &f)
	return f
}

func get(url string) (b []byte, err error) {
	res, e := http.Get(url)
	if e != nil {
		err = e
		return
	}
	data, e := ioutil.ReadAll(res.Body)
	if e != nil {
		err = e
		return
	}
	res.Body.Close()
	return data, nil
}

/*
const (
	apiurl    = "https://apicn.faceplusplus.com"
	apikey    = "353412255ae223d99559777df0b1dd3d"
	apisecret = "YuUsRLsUajXC2gXH4Hnz31v8J0LbPFwK"
)
	//url := apiurl + "/v2/detection/detect?url=" + picurl + "&api_secret=" + apisecret + "&api_key=" + apikey
*/

func FaceDetectionDetect(url string) (bool, error) {

	tmp, err := get(url)
	if len(FaceDecodeDetect(tmp).Face) == 0 {
		return false, err
	}

	return true, nil
}
