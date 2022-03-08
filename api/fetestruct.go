package api

import (
  "go.mongodb.org/mongo-driver/bson/primitive"
  "time"
)

type Insertable interface {
  SetDateAdded()
  SetDateUpdated()
}

type FeatureDefinition struct {
  // stores definition and metadata
}

type Tag struct {
  // defines tags for labeling features, environments
}

type FeatureEnvironment struct {
  // stores environmental expectations for a feature
  // defition (python version, packages with version)
}

type DefinitionTest struct {
  // stores definition test information
}

type DefinitionSubTest struct {
  // each definition test has one or more subtests, which describe the
  // specific test parameters
}

type FeatureDefinitionRequestEventLog struct {
  // defines requests, the feature requested
  // and retrieval/test success
}

func (i *  Insertable) SetDateAdded() {
  // set the record instertion date
}

func (I *Insertable) SetDateUpdated() {
  // set the record update date
}
