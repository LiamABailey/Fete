package api

import (
  "go.mongodb.org/mongo-driver/bson/primitive"
  "time"
)

type Insertable interface {
  SetDateAdded()
  SetDateUpdated()
}

type DefinitionSubTest struct {
  // each definition test has one or more subtests, which describe the
  // specific test parameters
}

func (dst *  DefinitionSubTest) SetDateAdded() {
  // set the record instertion date
}

func (dst  *DefinitionSubTest) SetDateUpdated() {
  // set the record update date
}
type DefinitionTest struct {
  // stores definition test information
}

func (dt *  DefinitionTest) SetDateAdded() {
  // set the record instertion date
}

func (dt *DefinitionTest) SetDateUpdated() {
  // set the record update date
}

type FeatureDefinition struct {
  ID                    bson.Objectid   `json:"id" bson:"_id"`
  FeatureID             int             `json:"feature_id" bson:"id"`
  VersionID             int             `json:"version_id" bson:"version_id"`
  PackageId             bson.Objectid   `json:"package_id" bson:"package_id"`
  FeatureEnvironmentID  bson.ObjectId   `json:"feature_environment_id" bson:"feature_environment_id"`
  FeatureName           string          `json:"feature_name" bson:"feature_name"`
  VersionName           string          `json:"version_name" bson:"version_name"`
  InputArgs             []Variable      `json:"input_args" bson:"input_args"`
  RetrunValues          []Variable      `json:"return_values" bson:"return_values"`
  DateAdded             time.Time       `json:"date_added" bson:"date_added"`
  DateUpdated           time.Time       `json:"date_updated" bson:"date_updated"`
}

type FeatureDefinitionRequestEventLog struct {
  // defines requests, the feature requested
  // and retrieval/test success
}

func (l *  FeatureDefinitionRequestEventLog) SetDateAdded() {
  // set the record instertion date
}

func (l *FeatureDefinitionRequestEventLog) SetDateUpdated() {
  // set the record update date
}

type FeatureEnvironment struct {
  // stores environmental expectations for a feature
  // defition (python version, packages with version)
}

func (fe *  FeatureEnvironment) SetDateAdded() {
  // set the record instertion date
}

func (fe *FeatureEnvironment) SetDateUpdated() {
  // set the record update date
}

func (fd *  FeatureDefinition) SetDateAdded() {
  // set the record instertion date
}

func (fd *FeatureDefinition) SetDateUpdated() {
  // set the record update date
}

type Tag struct {
  // defines tags for labeling features, environments
}

func (t *  Tag) SetDateAdded() {
  // set the record instertion date
}

func (t *Tag) SetDateUpdated() {
  // set the record update date
}

type Variable struct {
  // defines variables used to define name & type for function inputs
  Name    string  `json:"name" bson:"name"`
  Type    string  `json:"type" bson:"type"`
  Default string  `json:"default" bson:"default"`
}
