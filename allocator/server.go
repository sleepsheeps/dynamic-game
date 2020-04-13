package allocator

import (
	agonesv1 "agones.dev/agones/pkg/apis/agones/v1"
	allocationv1 "agones.dev/agones/pkg/apis/allocation/v1"
	"agones.dev/agones/pkg/client/clientset/versioned"
	"agones.dev/agones/pkg/util/runtime"
	dallocator "dynamic-game/client/allocator"
	"errors"
	"io"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

var (
	logger    = runtime.NewLoggerWithSource("main")
	allocator *Allocator
)

func Start() (err error) {
	if allocator == nil {
		allocator = new(Allocator)
	}
	allocator.agonesClient, err = getAgonesClient()
	if err != nil {
		return
	}
	allocator.start()
	return
}

type handler func(w http.ResponseWriter, r *http.Request)

type Allocator struct {
	agonesClient *versioned.Clientset
}

func (a *Allocator) start() {
	a.registerHandle()
	if err := http.ListenAndServe(":8000", nil); err != nil {
		logger.WithError(err).Fatal("HTTP server failed to run")
	} else {
		logger.Info("HTTP server is running on port 8000")
	}
}

func (a *Allocator) registerHandle() {
	http.HandleFunc("/"+dallocator.HANDLE_HEALTHZ, a.handleHealthz)
	http.HandleFunc("/"+dallocator.HANDLE_ALLOCATE, getOnly(a.handleAllocator))
	http.HandleFunc("/"+dallocator.HANDLE_DELETE, deleteOnly(a.handleDelete))
	http.HandleFunc("/"+dallocator.HANDLE_VERSION, getOnly(a.handleVersion))
	http.HandleFunc("/"+dallocator.HANDLE_GS_VERSION, getOnly(a.handleGSVersion))
}

func (a *Allocator) handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := io.WriteString(w, "Healthy")
	if err != nil {
		logger.WithError(err).Fatal("Error writing string Healthy from /healthz")
	}
}

func (a *Allocator) handleAllocator(w http.ResponseWriter, r *http.Request) {
	fleet := r.Header.Get(dallocator.HEADER_FLEET)
	namespace := r.Header.Get(dallocator.HEADER_NAMESPACE)
	if fleet == "" || namespace == "" {
		http.Error(w, "allocate fleet or namespace is nil", http.StatusNotFound)
	}
	_, serverID, version, err := a.allocate(fleet, namespace)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("serverID", serverID)
	w.Header().Set("version", version)
}

func (a *Allocator) handleDelete(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("Name")
	fleetname := r.Header.Get("Fleet")
	namespace := r.Header.Get("NameSpace")
	if name == "" || fleetname == "" || namespace == "" {
		http.Error(w, "delete name or fleetname or namespace is nil", http.StatusNotFound)
	}
	err := a.deleteGS(name, namespace, fleetname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
}

func (a *Allocator) handleVersion(w http.ResponseWriter, r *http.Request) {
	fleetname := r.Header.Get("Fleet")
	namespace := r.Header.Get("NameSpace")
	if fleetname == "" || namespace == "" {
		http.Error(w, "get fleetname or namespace is nil", http.StatusNotFound)
	}
	version, err := a.getImageVersion(namespace, fleetname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if version == "" {
		http.Error(w, "version is nil", http.StatusInternalServerError)
	}
	w.Header().Set("version", version)
}

func (a *Allocator) handleGSVersion(w http.ResponseWriter, r *http.Request) {
	namespace := r.Header.Get("NameSpace")
	fightd := r.Header.Get("Name")
	if fightd == "" || namespace == "" {
		http.Error(w, "get fightd version  is nil", http.StatusNotFound)
	}
	version, err := a.getGSImageVersion(namespace, fightd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if version != "" {
		w.Header().Set("version", version)
	} else {
		http.Error(w, "not found", http.StatusInternalServerError)
	}
}

func (a *Allocator) allocate(fleetname string, namespace string) (state allocationv1.GameServerAllocationState,
	serverID string, version string, err error) {
	logger.WithField("namespace", namespace).Info("namespace for gsa")
	logger.WithField("fleetname", fleetname).Info("fleetname for gsa")
	state = allocationv1.GameServerAllocationUnAllocated
	// 检测有多少ready状态的server
	readyReplicas, err := a.checkReadyReplicas(fleetname, namespace)
	if err != nil {
		logger.WithError(err).Info("check ready replicas error")
		return
	}
	logger.WithField("readyReplicas", readyReplicas).Info("number of ready replicas")
	if readyReplicas < 1 {
		logger.WithField("fleetname", fleetname).Info("Insufficient ready replicas, cannot create fleet allocation")
		err = errors.New("insufficient ready replicas, cannot create fleet allocation")
		return
	}
	allocationInterface := a.agonesClient.AllocationV1().GameServerAllocations(namespace)
	gsa := &allocationv1.GameServerAllocation{
		Spec: allocationv1.GameServerAllocationSpec{
			Required: metav1.LabelSelector{MatchLabels: map[string]string{agonesv1.FleetNameLabel: fleetname}},
		}}
	gsa, err = allocationInterface.Create(gsa)
	if err != nil {
		logger.WithError(err).Info("Failed to create allocation")
		return
	}
	logger.Info("New GameServer allocated: ", gsa.Status.State, "the name is ", gsa.ObjectMeta.Name)
	serverID = gsa.ObjectMeta.Name
	state = gsa.Status.State
	version, err = a.getImageVersion(namespace, fleetname)
	return
}

func (a *Allocator) deleteGS(name string, namespace string, fleetname string) error {
	logger.WithField("namespace", namespace).Info("namespace for gsa")
	logger.WithField("fleetname", fleetname).Info("fleetname for gsa")
	gsInterface := a.agonesClient.AgonesV1().GameServers(namespace)
	gs, err := gsInterface.Get(name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	gs.Status.State = agonesv1.GameServerStateReady
	_, err = gsInterface.Update(gs)
	if err != nil {
		return err
	}
	return err
}

// 获取fleet的镜像version
func (a *Allocator) getImageVersion(namespace string, fleetname string) (string, error) {
	fleetInterface := a.agonesClient.AgonesV1().Fleets(namespace)
	fleet, err := fleetInterface.Get(fleetname, metav1.GetOptions{})
	if err != nil {
		logger.WithError(err).Info("Get fleet failed")
		return "", err
	}
	containers := fleet.Spec.Template.Spec.Template.Spec.Containers
	if len(containers) > 0 {
		return containers[0].Image, nil
	}
	return "", errors.New("no enough containers")
}

func (a *Allocator) getGSImageVersion(namespace, name string) (string, error) {
	gsInterface := a.agonesClient.AgonesV1().GameServers(namespace)
	gs, err := gsInterface.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return gs.Spec.Template.Spec.Containers[0].Image, nil
}

func (a *Allocator) checkReadyReplicas(fleetname string, namespace string) (int32, error) {
	fleetInterface := a.agonesClient.AgonesV1().Fleets(namespace)
	fleet, err := fleetInterface.Get(fleetname, metav1.GetOptions{})
	if err != nil {
		return 0, err
	}
	return fleet.Status.ReadyReplicas, nil
}
