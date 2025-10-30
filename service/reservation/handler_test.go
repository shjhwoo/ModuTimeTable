package reservation

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAvailableTimeSlotsByRoom(t *testing.T) {

	path := "/room/timeSlots?roomId=1"

	req, err := http.NewRequestWithContext(ctx, "GET", path, nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	t.Logf("Response Body: %s", w.Body.String())
}
