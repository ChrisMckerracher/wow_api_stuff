package request

import "data_dump/request/model"

type BlizzardClient interface{}

type BlizzardRequest struct {
	Region model.Region
}
