package bot

import "testing"

func TestInitialState(t *testing.T) {
	cases := []struct {
		Title            string
		message          string
		expectedResponse string
		mockBotResponse  string
	}{
		{
			Title:            "Sending /quote command",
			message:          "/quote",
			expectedResponse: "entered quote state",
			mockBotResponse:  "entered quote state",
		},
		{
			Title:            "Sending unknown command",
			message:          "hello",
			expectedResponse: welcomeMessage,
		},
	}

	for _, test := range cases {
		mockBot := &mockStockerBot{
			mockEnterQuoteStateResponse: test.mockBotResponse,
		}

		initialState := &InitialState{
			stockerBot: mockBot,
		}

		got := initialState.buildResponse(test.message)
		want := test.expectedResponse

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}

type mockStockerBot struct {
	mockEnterQuoteStateResponse   string
	mockEnterInitialStateResponse string
}

func (m mockStockerBot) enterQuoteState() (response string) {
	return m.mockEnterQuoteStateResponse
}

func (m mockStockerBot) enterInitialState() (response string) {
	return m.mockEnterInitialStateResponse
}
