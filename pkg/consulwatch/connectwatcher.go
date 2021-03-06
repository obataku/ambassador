package consulwatch

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"

	"github.com/datawire/ambassador/pkg/dlog"
)

type ConnectLeafWatcher struct {
	consul *api.Client
	plan   *watch.Plan
}

func NewConnectLeafWatcher(consul *api.Client, service string) (*ConnectLeafWatcher, error) {
	if service == "" {
		err := errors.New("service name is empty")
		return nil, err
	}

	watcher := &ConnectLeafWatcher{consul: consul}

	plan, err := watch.Parse(map[string]interface{}{"type": "connect_leaf", "service": service})
	if err != nil {
		return nil, err
	}

	watcher.plan = plan

	return watcher, nil
}

func (w *ConnectLeafWatcher) Watch(handler func(*Certificate, error)) {
	w.plan.HybridHandler = func(val watch.BlockingParamVal, raw interface{}) {
		if raw == nil {
			handler(nil, fmt.Errorf("unexpected empty/nil response from consul"))
			return
		}

		v, ok := raw.(*api.LeafCert)
		if !ok {
			handler(nil, fmt.Errorf("unexpected raw type. expected: %T, was: %T", &api.LeafCert{}, raw))
			return
		}

		certificate := &Certificate{
			PEM:           v.CertPEM,
			PrivateKeyPEM: v.PrivateKeyPEM,
			ValidBefore:   v.ValidBefore,
			ValidAfter:    v.ValidAfter,
			SerialNumber:  v.SerialNumber,
			Service:       v.Service,
			ServiceURI:    v.ServiceURI,
		}

		handler(certificate, nil)
	}
}

func (w *ConnectLeafWatcher) Start(ctx context.Context) error {
	return w.plan.RunWithClientAndLogger(w.consul, dlog.StdLogger(ctx, dlog.LogLevelInfo))
}

func (w *ConnectLeafWatcher) Stop() {
	w.plan.Stop()
}

// ConnectCARootsWatcher watches the Consul Connect CA roots endpoint for changes and invokes a a handler function
// whenever it changes.
type ConnectCARootsWatcher struct {
	consul *api.Client
	plan   *watch.Plan
}

func NewConnectCARootsWatcher(consul *api.Client) (*ConnectCARootsWatcher, error) {
	watcher := &ConnectCARootsWatcher{consul: consul}

	plan, err := watch.Parse(map[string]interface{}{"type": "connect_roots"})
	if err != nil {
		return nil, err
	}

	watcher.plan = plan

	return watcher, nil
}

func (w *ConnectCARootsWatcher) Watch(handler func(*CARoots, error)) {
	w.plan.HybridHandler = func(val watch.BlockingParamVal, raw interface{}) {
		if raw == nil {
			handler(nil, fmt.Errorf("unexpected empty/nil response from consul"))
			return
		}

		v, ok := raw.(*api.CARootList)
		if !ok {
			handler(nil, fmt.Errorf("unexpected raw type. expected: %T, was: %T", &api.CARootList{}, raw))
			return
		}

		rootsMap := make(map[string]CARoot)
		for _, root := range v.Roots {
			rootsMap[root.ID] = CARoot{
				ID:     root.ID,
				Name:   root.Name,
				PEM:    root.RootCertPEM,
				Active: root.Active,
			}
		}

		roots := &CARoots{
			ActiveRootID: v.ActiveRootID,
			TrustDomain:  v.TrustDomain,
			Roots:        rootsMap,
		}

		handler(roots, nil)
	}
}

func (w *ConnectCARootsWatcher) Start(ctx context.Context) error {
	return w.plan.RunWithClientAndLogger(w.consul, dlog.StdLogger(ctx, dlog.LogLevelInfo))
}

func (w *ConnectCARootsWatcher) Stop() {
	w.plan.Stop()
}
