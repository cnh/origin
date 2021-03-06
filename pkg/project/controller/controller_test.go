package controller

import (
	"testing"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	kclient "github.com/GoogleCloudPlatform/kubernetes/pkg/client"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"

	osclient "github.com/openshift/origin/pkg/client"
)

func TestSyncNamespaceThatIsTerminating(t *testing.T) {
	mockKubeClient := &kclient.Fake{}
	mockOriginClient := &osclient.Fake{}
	nm := NamespaceController{
		KubeClient: mockKubeClient,
		Client:     mockOriginClient,
	}
	//now := util.Now()
	testNamespace := &api.Namespace{
		ObjectMeta: api.ObjectMeta{
			Name:            "test",
			ResourceVersion: "1",
			//			DeletionTimestamp: &now,
		},
		//		Spec: api.NamespaceSpec{
		//			Finalizers: []api.FinalizerName{"kubernetes"},
		//		},
		Status: api.NamespaceStatus{
			Phase: api.NamespaceTerminating,
		},
	}
	err := nm.Handle(testNamespace)
	if err != nil {
		t.Errorf("Unexpected error when handling namespace %v", err)
	}

	// TODO: we will expect a finalize namespace call after rebase
	expectedActionSet := util.NewStringSet(
		"list-buildconfig",
		"list-policies",
		"list-imagerepositries",
		"list-policyBindings",
		"list-roleBinding",
		"list-role",
		"list-routes",
		"list-builds",
		"list-deploymentconfig",
		"list-deployment")
	actionSet := util.NewStringSet()
	for i := range mockKubeClient.Actions {
		actionSet.Insert(mockKubeClient.Actions[i].Action)
	}
	for i := range mockOriginClient.Actions {
		actionSet.Insert(mockOriginClient.Actions[i].Action)
	}
	if !actionSet.HasAll(expectedActionSet.List()...) {
		t.Errorf("Expected actions: %v, but got: %v", expectedActionSet, actionSet)
	}
}

func TestSyncNamespaceThatIsActive(t *testing.T) {
	mockKubeClient := &kclient.Fake{}
	mockOriginClient := &osclient.Fake{}
	nm := NamespaceController{
		KubeClient: mockKubeClient,
		Client:     mockOriginClient,
	}
	//now := util.Now()
	testNamespace := &api.Namespace{
		ObjectMeta: api.ObjectMeta{
			Name:            "test",
			ResourceVersion: "1",
			//      DeletionTimestamp: &now,
		},
		//    Spec: api.NamespaceSpec{
		//      Finalizers: []api.FinalizerName{"kubernetes"},
		//    },
		Status: api.NamespaceStatus{
			Phase: api.NamespaceActive,
		},
	}
	err := nm.Handle(testNamespace)
	if err != nil {
		t.Errorf("Unexpected error when handling namespace %v", err)
	}
	actionSet := util.NewStringSet()
	for i := range mockKubeClient.Actions {
		actionSet.Insert(mockKubeClient.Actions[i].Action)
	}
	for i := range mockOriginClient.Actions {
		actionSet.Insert(mockOriginClient.Actions[i].Action)
	}
	if len(actionSet) != 0 {
		t.Errorf("Expected no action from controller, but got: %v", actionSet)
	}
}
