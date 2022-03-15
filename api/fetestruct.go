package api

import (
  "go.mongodb.org/mongo-driver/bson/primitive"
  "time"
)

type Insertable interface {
  SetDateAdded()
  SetDateUpdated()
}

// defines test information
type DefinitionSubTest struct {
  ID                    bson.ObjectId       `json:"id" bson:"_id"`
  DefinitionTestIDs     []bson.ObjectId     `json:"DefinitionTest" bson:"related_feature_definitions"`
  SubTestName           string              `json:"subtest_name" bson:"subtest_name"`
  Inputs                []Values            `json:"inputs" bson:"inputs"`
  Expectations          []EvalDefiniton     `json:"expectations" bson:"expected_results_with_tests"`
  DateAdded             time.Time           `json:"date_added" bson:"date_added"`
  DateUpdated           time.Time           `json:"date_updated" bson:"date_updated"`
  Tags                  []Tag               `json:"tags" bson:"tags"`
}

func (dst *  DefinitionSubTest) SetDateAdded() {
  // set the record instertion date
}

func (dst  *DefinitionSubTest) SetDateUpdated() {
  // set the record update date
}

// type defining linkages between subtests and features
type DefinitionTest struct {
  ID                    bson.ObjectId       `json:"id" bson:"_id"`
  FeatureDefinitionIDs  []bson.ObjectId     `json:"related_feature_definitons" bson:"related_feature_definitions"`
  DefintionSubTestIDs   []DefinitionSubTest `json:"subtests" bson:"subtests"`
  TestName              string              `json:"test_name" bson:"test_name"`
  TestDescription       string              `json:"test_description" bson:"test_description"`
  DateAdded             time.Time           `json:"date_added" bson:"date_added"`
  DateUpdated           time.Time           `json:"date_updated" bson:"date_updated"`
  Tags                  []Tag               `json:"tags" bson:"tags"`
}

func (dt *  DefinitionTest) SetDateAdded() {
  // set the record instertion date
}

func (dt *DefinitionTest) SetDateUpdated() {
  // set the record update date
}

type EvalDefinition struct {
  ExpectedOutput    Value   `json:"expected_output" bson:"expected_output"`
  AssertType        string  `json:"assertion_type" bson:"assertion_type"`
  Atol              float   `json:"atol" bson:"absolute_tolerance"`
  Rtol              float   `json:"rtol" bson:"relative_tolerance"`
  Ordered           bool    `json:"ordered" bson:"ordered"`
}

type FeatureDefinition struct {
  FeatureDefinitionID   bson.Objectid   `json:"id" bson:"_id"`
  FeatureID             int             `json:"feature_id" bson:"id"`
  VersionID             int             `json:"version_id" bson:"version_id"`
  PackageID             bson.Objectid   `json:"package_id" bson:"package_id"`
  FeatureEnvironmentID  bson.ObjectId   `json:"feature_environment_id" bson:"feature_environment_id"`
  FeatureName           string          `json:"feature_name" bson:"feature_name"`
  VersionName           string          `json:"version_name" bson:"version_name"`
  InputArgs             []Variable      `json:"input_args" bson:"input_args"`
  ReturnValues          []Variable      `json:"return_values" bson:"return_values"`
  DateAdded             time.Time       `json:"date_added" bson:"date_added"`
  DateUpdated           time.Time       `json:"date_updated" bson:"date_updated"`
  Tags                  []Tag           `json:"tags" bson:"tags"`
}

func (fd *  FeatureDefinition) SetDateAdded() {
  // set the record instertion date
}

func (fd *FeatureDefinition) SetDateUpdated() {
  // set the record update date
}

type FeatureEnvironment struct {
  ID                    bson.ObjectId   `json:"id" bson:"_id"`
  PythonVersion         []string        `json:"python_version" bson:"python_version"`
  RequiredPackages      []Packages      `json:"required_packages" bson:"required_packages"`
  DateAdded             time.Time       `json:"date_added" bson:"date_added"`
  DateUpdated           time.Time       `json:"date_updated" bson:"date_updated"`
  Tags                  []Tag           `json:"tags" bson:"tags"`
}

func (fe *  FeatureEnvironment) SetDateAdded() {
  // set the record instertion date
}

func (fe *FeatureEnvironment) SetDateUpdated() {
  // set the record update date
}

type Packages struct {
  PackageName       string    `json:"name" bson:"name"`
  Version           string    `json:"version" bson:"version"`
  ImportStatement   string    `json:"import_statement" bson:"import_statement"`
  LocalWHLName      string     `json:"local_whl_name" bson:"local_whl_name"`
}

type Tag struct {
  ID             bson.ObjectId   `json:"id" bson:"_id"`
  Name           string          `json:"name" bson:"name"`
  Description    string          `json:"description" bson:"description"`
  DateAdded      time.Time       `json:"date_added" bson:"date_added"`
  DateUpdated    time.Time       `json:"date_updated" bson:"date_updated"`
}

func (t *  Tag) SetDateAdded() {
  // set the record instertion date
}

func (t *Tag) SetDateUpdated() {
  // set the record update date
}

type FeatureDefinitionRequestEventLog struct {
  ID                    bson.ObjectId   `json:"id" bson:"_id"`
  Requestor             string          `json:"requestor" bson:"requestor"`
  FeatureDefinitionID   string          `json:"feature_definition" bson:"feature_definition"`
  RequireTestPass       bool            `json:"test_pass_required" bson:"test_pass_required"`
  DefinitionReturned    bool            `json:"definition_returned" bson:"definition_returned"`
  ReturnStatusMessage   string          `json:"return_status_message" bson:"return_status_message"`
  RequestedAt           time.Time       `json:"requested_at" bson:"requested_at"`
}

type FeatureDefinitionTestEventLog struct {
  ID                    bson.ObjectId   `json:"id" bson:"_id"`
  Requestor             string          `json:"requestor" bson:"requestor"`
  FeatureDefinitionID   string          `json:"feature_definition" bson:"feature_definition"`
  DefinitionTest        bson.ObjectId   `json:"definition_test" bson:"definition_test"`
  SubTestResults        []SubTestResult `json:"subtest_results" bson:"subtest_results"`
  TestPassed            bool            `json:"test_passed" bson:"test_passed"`
  RequestedAt           time.Time       `json:"requested_at" bson:"requested_at"`
}

type SubTestResult struct {
  ID                    bson.ObjectId   `json:"id" bson:"_id"`
  DefinitionSubTest     bson.ObjectId   `json:"definition_subtest" bson:"definition_subtest"`
  FeatureDefinitionID   string          `json:"feature_definition" bson:"feature_definition"`
  TestPassed            bool            `json:"test_passed" bson:"test_passed"`
  Message               string          `json:"test_message" bson:"test_message"`
  TestExecutedAt        time.Time       `json:"test_executed_at" bson:"test_executed_at"`
}

type Value struct {
  Value   string  `json:"value" bson:"value"`
  Type    string  `json:"type" bson:"type"`
}

type Variable struct {
  // defines variables used to define name & type for function inputs
  Name    string  `json:"name" bson:"name"`
  Type    string  `json:"type" bson:"type"`
  Default string  `json:"default" bson:"default"`
}
