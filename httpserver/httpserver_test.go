package httpserver

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func checkStatusCode(got, expect int, t *testing.T) {
    if got != expect {
        t.Errorf("Wrong status code, expect %d, got %d", expect, got)
    }
}

func checkStrEqual(key, got, expect string, t *testing.T) {
    if got != expect {
        t.Errorf("Wrong %s, expect %s, got %s", key, expect, got)
    }
}

func TestRootHandler(t *testing.T) {
    r, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatalf("rootHandler has error, message: %s", err)
    }

    w := httptest.NewRecorder()
    srv := http.HandlerFunc(rootHandler)
    srv.ServeHTTP(w, r)
    checkStatusCode(w.Code, http.StatusOK, t)
    checkStrEqual("response Body", w.Body.String(), welcomeMsg, t)
}
