package v1beta1

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("SuperStream spec", func() {
	var (
		namespace = "default"
		ctx       = context.Background()
	)

	It("creates a superstream with default settings", func() {
		expectedSpec := SuperStreamSpec{
			Name:       "test-super-stream",
			Partitions: 3,
			RabbitmqClusterReference: RabbitmqClusterReference{
				Name: "some-cluster",
			},
		}

		superStream := SuperStream{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-super-stream",
				Namespace: namespace,
			},
			Spec: SuperStreamSpec{
				Name: "test-super-stream",
				RabbitmqClusterReference: RabbitmqClusterReference{
					Name: "some-cluster",
				},
			},
		}
		Expect(k8sClient.Create(ctx, &superStream)).To(Succeed())
		fetchedSuperStream := &SuperStream{}
		Expect(k8sClient.Get(ctx, types.NamespacedName{
			Name:      superStream.Name,
			Namespace: superStream.Namespace,
		}, fetchedSuperStream)).To(Succeed())
		Expect(fetchedSuperStream.Spec).To(Equal(expectedSpec))
	})

	It("creates a superstream with specified settings", func() {
		expectedSpec := SuperStreamSpec{
			Name:       "test-super-stream2",
			Partitions: 5,
			RabbitmqClusterReference: RabbitmqClusterReference{
				Name: "some-cluster",
			},
		}

		superStream := SuperStream{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-super-stream2",
				Namespace: namespace,
			},
			Spec: SuperStreamSpec{
				Name:       "test-super-stream2",
				Partitions: 5,
				RabbitmqClusterReference: RabbitmqClusterReference{
					Name: "some-cluster",
				},
			},
		}
		Expect(k8sClient.Create(ctx, &superStream)).To(Succeed())
		fetchedSuperStream := &SuperStream{}
		Expect(k8sClient.Get(ctx, types.NamespacedName{
			Name:      superStream.Name,
			Namespace: superStream.Namespace,
		}, fetchedSuperStream)).To(Succeed())
		Expect(fetchedSuperStream.Spec).To(Equal(expectedSpec))
	})

})