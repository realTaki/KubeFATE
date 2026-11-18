package k8sclient

import (
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
)

func TestToIngressPorts(t *testing.T) {
	ingresses := []networkingv1.IngressLoadBalancerIngress{
		{Ports: []networkingv1.IngressPortStatus{
			{Port: 80, Protocol: corev1.ProtocolTCP},
			{Port: 53, Protocol: corev1.ProtocolUDP},
		}},
		{Ports: []networkingv1.IngressPortStatus{
			{Port: 443, Protocol: corev1.ProtocolTCP},
		}},
	}
	want := []string{"80/TCP,53/UDP", "443/TCP"}
	if got := ToIngressPorts(ingresses); !reflect.DeepEqual(got, want) {
		t.Fatalf("ToIngressPorts() = %v, want %v", got, want)
	}
}
