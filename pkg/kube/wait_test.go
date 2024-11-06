package kube

import (
	"context"
	"testing"
	"time"

	"k8s.io/client-go/kubernetes/fake"

	_ "k8s.io/cli-runtime/pkg/resource"
)

func Test_waiter_waitForDeletedResources(t *testing.T) {
	type args struct {
		ctx      context.Context
		resource ResourceList
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "wait for deleted resources",
			args: args{
				ctx:      context.TODO(),
				resource: ResourceList{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &waiter{
				c: ReadyChecker{
					client:        fake.NewSimpleClientset(),
					log:           func(string, ...interface{}) {},
					checkJobs:     false,
					pausedAsReady: false,
				},
				timeout: time.Duration(10000),
				log:     func(string, ...interface{}) {},
			}
			got := w.waitForDeletedResources(tt.args.resource)
			if (got == nil) != tt.want {
				t.Errorf("waiter.waitForDeletedResources() got = %v, want %v", got, tt.want)
				return
			}

		})
	}
}
