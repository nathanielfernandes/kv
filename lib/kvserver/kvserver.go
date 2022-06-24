package kvserver

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nathanielfernandes/kv/lib/kv"
	rl "github.com/nathanielfernandes/rl"
)

const MAX_K_LENGTH = 256
const MAX_V_LENGTH = 512

type KVServer struct {
	kv *kv.KV

	set_rlm *rl.RatelimitManager
	get_rlm *rl.RatelimitManager
}

func trunicate(s string, maxlen int) string {
	if len(s) > maxlen {
		return s[:maxlen]
	}
	return s
}

func NewKVServer() *KVServer {
	return &KVServer{
		kv: kv.NewKV(),

		set_rlm: rl.NewRatelimitManager(10_000, 3600_000),
		get_rlm: rl.NewRatelimitManager(100, 6000),
	}
}

func (kvs *KVServer) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cors(w)

	id := getID(r)
	if kvs.get_rlm.IsRatelimited(id) {
		http.Error(w, "ratelimited", http.StatusTooManyRequests)
		return
	}

	key := ps.ByName("key")

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(kvs.kv.Get(key, "")))
}

func (kvs *KVServer) Set(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cors(w)

	id := getID(r)
	if kvs.set_rlm.IsRatelimited(id) {
		http.Error(w, "ratelimited", http.StatusTooManyRequests)
		return
	}

	key := ps.ByName("key")
	value := ps.ByName("value")

	kvs.kv.Set(trunicate(key, MAX_K_LENGTH), trunicate(value, MAX_V_LENGTH))
	w.WriteHeader(http.StatusOK)
}

func (kvs *KVServer) RedirectTo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := getID(r)
	if kvs.get_rlm.IsRatelimited(id) {
		http.Error(w, "ratelimited", http.StatusTooManyRequests)
		return
	}

	key := ps.ByName("key")
	url := kvs.kv.Get(key, "https://nathanielfernandes.ca/")

	http.Redirect(w, r, url, http.StatusFound)
}
