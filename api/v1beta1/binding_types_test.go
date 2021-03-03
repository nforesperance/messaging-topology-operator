package v1beta1

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("Binding spec", func() {
	var (
		namespace = "default"
		ctx       = context.Background()
	)

	It("creates a binding with default settings", func() {
		expectedSpec := BindingSpec{
			Vhost: "/",
			RabbitmqClusterReference: RabbitmqClusterReference{
				Name:      "some-cluster",
				Namespace: namespace,
			},
		}

		binding := Binding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-binding",
				Namespace: namespace,
			},
			Spec: BindingSpec{
				RabbitmqClusterReference: RabbitmqClusterReference{
					Name:      "some-cluster",
					Namespace: namespace,
				},
			},
		}
		Expect(k8sClient.Create(ctx, &binding)).To(Succeed())
		fetchedBinding := &Binding{}
		Expect(k8sClient.Get(ctx, types.NamespacedName{
			Name:      binding.Name,
			Namespace: binding.Namespace,
		}, fetchedBinding)).To(Succeed())
		Expect(fetchedBinding.Spec).To(Equal(expectedSpec))
	})

	It("creates a binding with configurations", func() {
		q := Binding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "random-q",
				Namespace: namespace,
			},
			Spec: BindingSpec{
				Vhost:           "/avhost",
				Source:          "anexchange",
				Destination:     "aqueue",
				DestinationType: "queue",
				RoutingKey:      "akey",
				Arguments: &runtime.RawExtension{
					Raw: []byte(`{"argument":"argument-value"}`),
				},
				RabbitmqClusterReference: RabbitmqClusterReference{
					Name:      "random-cluster",
					Namespace: namespace,
				},
			},
		}
		Expect(k8sClient.Create(ctx, &q)).To(Succeed())
		fetchedBinding := &Binding{}
		Expect(k8sClient.Get(ctx, types.NamespacedName{
			Name:      q.Name,
			Namespace: q.Namespace,
		}, fetchedBinding)).To(Succeed())

		Expect(fetchedBinding.Spec.Vhost).To(Equal("/avhost"))
		Expect(fetchedBinding.Spec.Source).To(Equal("anexchange"))
		Expect(fetchedBinding.Spec.Destination).To(Equal("aqueue"))
		Expect(fetchedBinding.Spec.DestinationType).To(Equal("queue"))
		Expect(fetchedBinding.Spec.RoutingKey).To(Equal("akey"))
		Expect(fetchedBinding.Spec.RabbitmqClusterReference).To(Equal(
			RabbitmqClusterReference{
				Name:      "random-cluster",
				Namespace: namespace,
			}))
		Expect(fetchedBinding.Spec.Arguments.Raw).To(Equal([]byte(`{"argument":"argument-value"}`)))
	})
})
