package workflows

import (
	operators2 "backend/module/vulnscan/nuclei/v2/pkg/operators"
	output2 "backend/module/vulnscan/nuclei/v2/pkg/output"
	progress2 "backend/module/vulnscan/nuclei/v2/pkg/progress"
	protocols2 "backend/module/vulnscan/nuclei/v2/pkg/protocols"
	types2 "backend/module/vulnscan/nuclei/v2/pkg/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWorkflowsSimple(t *testing.T) {
	progressBar, _ := progress2.NewStatsTicker(0, false, false, false, 0)

	workflow := &Workflow{Options: &protocols2.ExecuterOptions{Options: &types2.Options{TemplateThreads: 10}}, Workflows: []*WorkflowTemplate{
		{Executers: []*ProtocolExecuterPair{{
			Executer: &mockExecuter{result: true}, Options: &protocols2.ExecuterOptions{Progress: progressBar}},
		}},
	}}

	matched := workflow.RunWorkflow("https://test.com")
	require.True(t, matched, "could not get correct match value")
}

func TestWorkflowsSimpleMultiple(t *testing.T) {
	progressBar, _ := progress2.NewStatsTicker(0, false, false, false, 0)

	var firstInput, secondInput string
	workflow := &Workflow{Options: &protocols2.ExecuterOptions{Options: &types2.Options{TemplateThreads: 10}}, Workflows: []*WorkflowTemplate{
		{Executers: []*ProtocolExecuterPair{{
			Executer: &mockExecuter{result: true, executeHook: func(input string) {
				firstInput = input
			}}, Options: &protocols2.ExecuterOptions{Progress: progressBar}},
		}},
		{Executers: []*ProtocolExecuterPair{{
			Executer: &mockExecuter{result: true, executeHook: func(input string) {
				secondInput = input
			}}, Options: &protocols2.ExecuterOptions{Progress: progressBar}},
		}},
	}}

	matched := workflow.RunWorkflow("https://test.com")
	require.True(t, matched, "could not get correct match value")

	require.Equal(t, "https://test.com", firstInput, "could not get correct first input")
	require.Equal(t, "https://test.com", secondInput, "could not get correct second input")
}

func TestWorkflowsSubtemplates(t *testing.T) {
	progressBar, _ := progress2.NewStatsTicker(0, false, false, false, 0)

	var firstInput, secondInput string
	workflow := &Workflow{Options: &protocols2.ExecuterOptions{Options: &types2.Options{TemplateThreads: 10}}, Workflows: []*WorkflowTemplate{
		{Executers: []*ProtocolExecuterPair{{
			Executer: &mockExecuter{result: true, executeHook: func(input string) {
				firstInput = input
			}, outputs: []*output2.InternalWrappedEvent{
				{OperatorsResult: &operators2.Result{}, Results: []*output2.ResultEvent{{}}},
			}}, Options: &protocols2.ExecuterOptions{Progress: progressBar}},
		}, Subtemplates: []*WorkflowTemplate{{Executers: []*ProtocolExecuterPair{{
			Executer: &mockExecuter{result: true, executeHook: func(input string) {
				secondInput = input
			}}, Options: &protocols2.ExecuterOptions{Progress: progressBar}},
		}}}},
	}}

	matched := workflow.RunWorkflow("https://test.com")
	require.True(t, matched, "could not get correct match value")

	require.Equal(t, "https://test.com", firstInput, "could not get correct first input")
	require.Equal(t, "https://test.com", secondInput, "could not get correct second input")
}

func TestWorkflowsSubtemplatesNoMatch(t *testing.T) {
	progressBar, _ := progress2.NewStatsTicker(0, false, false, false, 0)

	var firstInput, secondInput string
	workflow := &Workflow{Options: &protocols2.ExecuterOptions{Options: &types2.Options{TemplateThreads: 10}}, Workflows: []*WorkflowTemplate{
		{Executers: []*ProtocolExecuterPair{{
			Executer: &mockExecuter{result: false, executeHook: func(input string) {
				firstInput = input
			}}, Options: &protocols2.ExecuterOptions{Progress: progressBar}},
		}, Subtemplates: []*WorkflowTemplate{{Executers: []*ProtocolExecuterPair{{
			Executer: &mockExecuter{result: true, executeHook: func(input string) {
				secondInput = input
			}}, Options: &protocols2.ExecuterOptions{Progress: progressBar}},
		}}}},
	}}

	matched := workflow.RunWorkflow("https://test.com")
	require.False(t, matched, "could not get correct match value")

	require.Equal(t, "https://test.com", firstInput, "could not get correct first input")
	require.Equal(t, "", secondInput, "could not get correct second input")
}

func TestWorkflowsSubtemplatesWithMatcher(t *testing.T) {
	progressBar, _ := progress2.NewStatsTicker(0, false, false, false, 0)

	var firstInput, secondInput string
	workflow := &Workflow{Options: &protocols2.ExecuterOptions{Options: &types2.Options{TemplateThreads: 10}}, Workflows: []*WorkflowTemplate{
		{Executers: []*ProtocolExecuterPair{{
			Executer: &mockExecuter{result: true, executeHook: func(input string) {
				firstInput = input
			}, outputs: []*output2.InternalWrappedEvent{
				{OperatorsResult: &operators2.Result{
					Matches:  map[string]struct{}{"tomcat": {}},
					Extracts: map[string][]string{},
				}},
			}}, Options: &protocols2.ExecuterOptions{Progress: progressBar}},
		}, Matchers: []*Matcher{{Name: "tomcat", Subtemplates: []*WorkflowTemplate{{Executers: []*ProtocolExecuterPair{{
			Executer: &mockExecuter{result: true, executeHook: func(input string) {
				secondInput = input
			}}, Options: &protocols2.ExecuterOptions{Progress: progressBar}},
		}}}}}},
	}}

	matched := workflow.RunWorkflow("https://test.com")
	require.True(t, matched, "could not get correct match value")

	require.Equal(t, "https://test.com", firstInput, "could not get correct first input")
	require.Equal(t, "https://test.com", secondInput, "could not get correct second input")
}

func TestWorkflowsSubtemplatesWithMatcherNoMatch(t *testing.T) {
	progressBar, _ := progress2.NewStatsTicker(0, false, false, false, 0)

	var firstInput, secondInput string
	workflow := &Workflow{Options: &protocols2.ExecuterOptions{Options: &types2.Options{TemplateThreads: 10}}, Workflows: []*WorkflowTemplate{
		{Executers: []*ProtocolExecuterPair{{
			Executer: &mockExecuter{result: true, executeHook: func(input string) {
				firstInput = input
			}, outputs: []*output2.InternalWrappedEvent{
				{OperatorsResult: &operators2.Result{
					Matches:  map[string]struct{}{"tomcat": {}},
					Extracts: map[string][]string{},
				}},
			}}, Options: &protocols2.ExecuterOptions{Progress: progressBar}},
		}, Matchers: []*Matcher{{Name: "apache", Subtemplates: []*WorkflowTemplate{{Executers: []*ProtocolExecuterPair{{
			Executer: &mockExecuter{result: true, executeHook: func(input string) {
				secondInput = input
			}}, Options: &protocols2.ExecuterOptions{Progress: progressBar}},
		}}}}}},
	}}

	matched := workflow.RunWorkflow("https://test.com")
	require.False(t, matched, "could not get correct match value")

	require.Equal(t, "https://test.com", firstInput, "could not get correct first input")
	require.Equal(t, "", secondInput, "could not get correct second input")
}

type mockExecuter struct {
	result      bool
	executeHook func(input string)
	outputs     []*output2.InternalWrappedEvent
}

// Compile compiles the execution generators preparing any requests possible.
func (m *mockExecuter) Compile() error {
	return nil
}

// Requests returns the total number of requests the rule will perform
func (m *mockExecuter) Requests() int {
	return 1
}

// Execute executes the protocol group and  returns true or false if results were found.
func (m *mockExecuter) Execute(input string) (bool, error) {
	if m.executeHook != nil {
		m.executeHook(input)
	}
	return m.result, nil
}

// ExecuteWithResults executes the protocol requests and returns results instead of writing them.
func (m *mockExecuter) ExecuteWithResults(input string, callback protocols2.OutputEventCallback) error {
	if m.executeHook != nil {
		m.executeHook(input)
	}
	for _, output := range m.outputs {
		callback(output)
	}
	return nil
}
