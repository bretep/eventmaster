package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ContextLogic/eventmaster/eventmaster"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

func NewHTTPServer(tlsConfig *tls.Config, store *EventStore) *http.Server {
	r := httprouter.New()
	h := httpHandler{
		store: store,
	}

	// API endpoints
	r.POST("/v1/event", wrapHandler(h.handleAddEvent))
	r.GET("/v1/event", wrapHandler(h.handleGetEvent))
	r.POST("/v1/topic", wrapHandler(h.handleAddTopic))
	r.PUT("/v1/topic/:name", wrapHandler(h.handleUpdateTopic))
	r.GET("/v1/topic", wrapHandler(h.handleGetTopic))
	r.DELETE("/v1/topic/:name", wrapHandler(h.handleDeleteTopic))
	r.POST("/v1/dc", wrapHandler(h.handleAddDc))
	r.PUT("/v1/dc/:name", wrapHandler(h.handleUpdateDc))
	r.GET("/v1/dc", wrapHandler(h.handleGetDc))

	r.GET("/v1/health", wrapHandler(h.handleHealthCheck))

	// GitHub webhook endpoint
	r.POST("/v1/github_event", wrapHandler(h.handleGitHubEvent))

	// UI endpoints
	r.GET("/", h.HandleMainPage)
	r.GET("/add_event", h.HandleCreatePage)
	r.GET("/topic", h.HandleTopicPage)
	r.GET("/dc", h.HandleDcPage)

	// JS file endpoints
	r.ServeFiles("/js/*filepath", http.Dir("ui/js"))
	r.ServeFiles("/bootstrap/*filepath", http.Dir("ui/bootstrap"))
	r.ServeFiles("/css/*filepath", http.Dir("ui/css"))

	return &http.Server{
		Handler:   r,
		TLSConfig: tlsConfig,
	}
}

func wrapHandler(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()
		defer func() {
			httpReqLatencies.WithLabelValues(r.URL.Path).Observe(trackTime(start))
		}()
		httpReqCounter.WithLabelValues(r.URL.Path).Inc()
		h(w, r, ps)
	}
}

type httpHandler struct {
	store *EventStore
}

func (h *httpHandler) sendError(w http.ResponseWriter, code int, err error, message string, path string) {
	httpRespCounter.WithLabelValues(path, fmt.Sprintf("%d", code)).Inc()
	errMsg := fmt.Sprintf("%s: %s", message, err.Error())
	fmt.Println(errMsg)
	w.WriteHeader(code)
	w.Write([]byte(errMsg))
}

func (h *httpHandler) sendResp(w http.ResponseWriter, key string, val string, path string) {
	var response []byte
	if key == "" {
		response = []byte(val)
	} else {
		resp := make(map[string]string)
		resp[key] = val
		var err error
		response, err = json.Marshal(resp)
		if err != nil {
			h.sendError(w, http.StatusInternalServerError, err, "Error marshalling response to JSON", path)
			return
		}
	}
	httpRespCounter.WithLabelValues(path, "200").Inc()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *httpHandler) handleAddEvent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var evt UnaddedEvent

	body, err := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &evt); err != nil {
		h.sendError(w, http.StatusBadRequest, err, "Error decoding JSON event", r.URL.Path)
		return
	}

	id, err := h.store.AddEvent(&evt)
	if err != nil {
		h.sendError(w, http.StatusBadRequest, err, "Error writing event", r.URL.Path)
		return
	}
	h.sendResp(w, "event_id", id, r.URL.Path)
}

type EventResult struct {
	EventID       string                 `json:"event_id"`
	ParentEventID string                 `json:"parent_event_id"`
	EventTime     int64                  `json:"event_time"`
	Dc            string                 `json:"dc"`
	TopicName     string                 `json:"topic_name"`
	Tags          []string               `json:"tag_set"`
	Host          string                 `json:"host"`
	TargetHosts   []string               `json:"target_host_set"`
	User          string                 `json:"user"`
	Data          map[string]interface{} `json:"data"`
	ReceivedTime  int64                  `json:"received_time"`
}

type SearchResult struct {
	Results []*EventResult `json:"results"`
}

func (h *httpHandler) handleGetEvent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var q eventmaster.Query

	// read from request body first - if there's an error, read from query params
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&q); err != nil {
		query := r.URL.Query()
		if val, ok := query["event_id"]; ok {
			if len(val) > 0 {
				q.EventId = val[0]
			}
		}
		q.ParentEventId = query["parent_event_id"]
		q.Dc = query["dc"]
		q.Host = query["host"]
		q.TargetHostSet = query["target_host_set"]
		q.User = query["user"]
		q.TagSet = query["tag_set"]
		q.TopicName = query["topic_name"]
		q.SortField = query["sort_field"]
		for _, elem := range query["sort_ascending"] {
			if strings.ToLower(elem) == "true" {
				q.SortAscending = append(q.SortAscending, true)
			} else if strings.ToLower(elem) == "false" {
				q.SortAscending = append(q.SortAscending, false)
			}
		}
		if len(q.SortField) != len(q.SortAscending) {
			h.sendError(w, http.StatusBadRequest, errors.New("sort_field and sort_ascending don't match"), "Error", r.URL.Path)
			return
		}
		if len(query["data"]) > 0 {
			q.Data = query["data"][0]
		}
		if startEventTime := query.Get("start_event_time"); startEventTime != "" {
			q.StartEventTime, _ = strconv.ParseInt(startEventTime, 10, 64)
		}
		if endEventTime := query.Get("end_event_time"); endEventTime != "" {
			q.EndEventTime, _ = strconv.ParseInt(endEventTime, 10, 64)
		}
		if startReceivedTime := query.Get("start_received_time"); startReceivedTime != "" {
			q.StartReceivedTime, _ = strconv.ParseInt(startReceivedTime, 10, 64)
		}
		if endReceivedTime := query.Get("end_received_time"); endReceivedTime != "" {
			q.EndReceivedTime, _ = strconv.ParseInt(endReceivedTime, 10, 64)
		}
		if start := query.Get("start"); start != "" {
			startIndex, _ := strconv.ParseInt(start, 10, 32)
			q.Start = int32(startIndex)
		}
		if limit := query.Get("limit"); limit != "" {
			resultSize, _ := strconv.ParseInt(limit, 10, 32)
			q.Limit = int32(resultSize)
		}
	}

	events, err := h.store.Find(&q)
	if err != nil {
		h.sendError(w, http.StatusInternalServerError, err, "Error executing query", r.URL.Path)
		return
	}
	var results []*EventResult
	for _, ev := range events {
		results = append(results, &EventResult{
			EventID:       ev.EventID,
			ParentEventID: ev.ParentEventID,
			EventTime:     ev.EventTime,
			Dc:            h.store.getDcName(ev.DcID),
			TopicName:     h.store.getTopicName(ev.TopicID),
			Tags:          ev.Tags,
			Host:          ev.Host,
			TargetHosts:   ev.TargetHosts,
			User:          ev.User,
			Data:          ev.Data,
		})
	}
	sr := SearchResult{
		Results: results,
	}
	jsonSr, err := json.Marshal(sr)
	if err != nil {
		h.sendError(w, http.StatusInternalServerError, err, "Error marshalling results into JSON", r.URL.Path)
		return
	}
	h.sendResp(w, "", string(jsonSr), r.URL.Path)
}

func (h *httpHandler) handleAddTopic(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	td := TopicData{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&td); err != nil {
		h.sendError(w, http.StatusBadRequest, err, "Error decoding JSON event", r.URL.Path)
		return
	}

	if td.Name == "" {
		h.sendError(w, http.StatusBadRequest, errors.New("Must include topic_name in request"), "Error adding topic", r.URL.Path)
		return
	}

	id, err := h.store.AddTopic(td)
	if err != nil {
		h.sendError(w, http.StatusBadRequest, err, "Error adding topic", r.URL.Path)
		return
	}
	h.sendResp(w, "topic_id", id, r.URL.Path)
}

func (h *httpHandler) handleUpdateTopic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var td TopicData
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.sendError(w, http.StatusBadRequest, err, "Error reading request body", r.URL.Path)
		return
	}
	err = json.Unmarshal(reqBody, &td)
	if err != nil {
		h.sendError(w, http.StatusBadRequest, err, "Error JSON decoding body of request", r.URL.Path)
		return
	}

	topicName := ps.ByName("name")
	if topicName == "" {
		h.sendError(w, http.StatusBadRequest, errors.New("Must provide topic name in URL"), "Error updating topic, no topic name provided", r.URL.Path)
		return
	}
	id, err := h.store.UpdateTopic(topicName, td)
	if err != nil {
		h.sendError(w, http.StatusBadRequest, err, "Error updating topic", r.URL.Path)
		return
	}
	h.sendResp(w, "topic_id", id, r.URL.Path)
}

func (h *httpHandler) handleGetTopic(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	topics, err := h.store.GetTopics()
	if err != nil {
		h.sendError(w, http.StatusInternalServerError, err, "Error getting topics from store", r.URL.Path)
		return
	}

	topicSet := make(map[string][]TopicData)
	topicSet["results"] = topics
	str, err := json.Marshal(topicSet)
	if err != nil {
		h.sendError(w, http.StatusInternalServerError, err, "Error marshalling response to JSON", r.URL.Path)
		return
	}
	h.sendResp(w, "", string(str), r.URL.Path)
}

func (h *httpHandler) handleDeleteTopic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	topicName := ps.ByName("name")
	if topicName == "" {
		h.sendError(w, http.StatusBadRequest, errors.New("Must provide topic name in URL"), "Error deleting topic, no topic name provided", r.URL.Path)
		return
	}
	err := h.store.DeleteTopic(&eventmaster.DeleteTopicRequest{
		TopicName: topicName,
	})
	if err != nil {
		h.sendError(w, http.StatusInternalServerError, err, "Error deleting topic from store", r.URL.Path)
		return
	}
	h.sendResp(w, "", "", r.URL.Path)
}

func (h *httpHandler) handleAddDc(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var dd DcData
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.sendError(w, http.StatusBadRequest, err, "Error reading request body", r.URL.Path)
		return
	}
	err = json.Unmarshal(reqBody, &dd)
	if err != nil {
		h.sendError(w, http.StatusBadRequest, err, "Error JSON decoding body of request", r.URL.Path)
		return
	}
	id, err := h.store.AddDc(&eventmaster.Dc{
		DcName: dd.Name,
	})
	if err != nil {
		h.sendError(w, http.StatusBadRequest, err, "Error adding dc", r.URL.Path)
		return
	}
	h.sendResp(w, "dc_id", id, r.URL.Path)
}

func (h *httpHandler) handleUpdateDc(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var dd DcData
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.sendError(w, http.StatusBadRequest, err, "Error reading request body", r.URL.Path)
		return
	}
	err = json.Unmarshal(reqBody, &dd)
	if err != nil {
		h.sendError(w, http.StatusBadRequest, err, "Error JSON decoding body of request", r.URL.Path)
		return
	}
	dcName := ps.ByName("name")
	if dcName == "" {
		h.sendError(w, http.StatusBadRequest, err, "Error updating topic, no topic name provided", r.URL.Path)
		return
	}
	id, err := h.store.UpdateDc(&eventmaster.UpdateDcRequest{
		OldName: dcName,
		NewName: dd.Name,
	})
	if err != nil {
		h.sendError(w, http.StatusBadRequest, err, "Error updating dc", r.URL.Path)
		return
	}
	h.sendResp(w, "dc_id", id, r.URL.Path)
}

func (h *httpHandler) handleGetDc(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	dcSet := make(map[string][]DcData)
	dcs, err := h.store.GetDcs()
	if err != nil {
		h.sendError(w, http.StatusInternalServerError, err, "Error getting dcs from store", r.URL.Path)
		return
	}
	dcSet["results"] = dcs
	str, err := json.Marshal(dcSet)
	if err != nil {
		h.sendError(w, http.StatusInternalServerError, err, "Error marshalling response to JSON", r.URL.Path)
		return
	}
	h.sendResp(w, "", string(str), r.URL.Path)
}

func (h *httpHandler) handleGitHubEvent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var info map[string]interface{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.sendError(w, http.StatusBadRequest, err, "Error reading request body", r.URL.Path)
		return
	}
	if err := json.Unmarshal(reqBody, &info); err != nil {
		h.sendError(w, http.StatusBadRequest, err, "Error JSON decoding body of request", r.URL.Path)
		return
	}

	id, err := h.store.AddEvent(&UnaddedEvent{
		Dc:        "github",
		Host:      "github",
		TopicName: "github",
		Data:      info,
	})
	if err != nil {
		h.sendError(w, http.StatusInternalServerError, err, "Error adding event to store", r.URL.Path)
		return
	}
	h.sendResp(w, "event_id", id, r.URL.Path)
}

func (h *httpHandler) handleHealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// TODO: make this more useful
	h.sendResp(w, "", "", r.URL.Path)
}
