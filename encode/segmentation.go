package encode

type Segment struct {
	Mode QRMode
	Data string
}

type DPState struct {
	Cost      float64
	PrevMode  QRMode
	PrevIndex int
	RunLength int
}

func greedySegmentMessage(msg string) []Segment {
	var segments []Segment

	currentMode := detectBestMode(string(msg[0]))
	start := 0

	for i := 1; i < len(msg); i++ {
		charMode := detectBestMode(string(msg[i]))

		if charMode != currentMode {
			segments = append(segments, Segment{
				Mode: currentMode,
				Data: msg[start:i],
			})
			start = i
			currentMode = charMode
		}
	}

	segments = append(segments, Segment{
		Mode: currentMode,
		Data: msg[start:],
	})

	return segments
}

func segmentMessage(msg string) []Segment {
	n := len(msg)

	dp := make([]map[QRMode]DPState, n+1)

	for i := range dp {
		dp[i] = make(map[QRMode]DPState)
	}

	dp[0][Numeric] = DPState{Cost: 0, RunLength: 0}
	dp[0][Alphanumeric] = DPState{Cost: 0, RunLength: 0}
	dp[0][Byte] = DPState{Cost: 0, RunLength: 0}

	// DP Transition
	for i := 0; i < n; i++ {
		c := msg[i]

		for mode, state := range dp[i] {

			// try all possible next modes
			for _, newMode := range []QRMode{Numeric, Alphanumeric, Byte} {

				if !canEncode(newMode, c) {
					continue
				}

				cost := state.Cost
				newRunLength := 1

				if mode == newMode {
					newRunLength = state.RunLength + 1
				} else {
					// switching modes
					cost += float64(4 + charCountBits(newMode))
				}

				// add encoding cost
				cost += float64(incrementalCost(newMode, newRunLength))

				nextState, exists := dp[i+1][newMode]

				if !exists || cost < nextState.Cost {
					dp[i+1][newMode] = DPState{
						Cost:      cost,
						PrevMode:  mode,
						PrevIndex: i,
						RunLength: newRunLength,
					}
				}
			}
		}
	}

	// Find best ending mode
	bestMode := Byte
	bestCost := 1e18

	for mode, state := range dp[n] {
		if state.Cost < bestCost {
			bestCost = state.Cost
			bestMode = mode
		}
	}

	// Reconstruct path
	var modes []QRMode

	i := n
	mode := bestMode

	for i > 0 {
		state := dp[i][mode]
		modes = append([]QRMode{mode}, modes...)
		mode = state.PrevMode
		i = state.PrevIndex
	}

	// Convert modes -> segments
	var segments []Segment

	start := 0
	currentMode := modes[0]

	for i := 1; i < len(modes); i++ {
		if modes[i] != currentMode {
			segments = append(segments, Segment{Mode: currentMode, Data: msg[start:i]})
			start = i
			currentMode = modes[i]
		}
	}

	segments = append(segments, Segment{Mode: currentMode, Data: msg[start:]})

	return segments
}
