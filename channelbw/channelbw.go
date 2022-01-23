package channelbw

import (
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type ChannelInfo struct {
	Epoch     uint64
	Bandwidth int
}

type ChannelInfos []ChannelInfo

func (channel ChannelInfos) Len() int {
	return len(channel)
}

func (channel ChannelInfos) Swap(i, j int) {
	channel[i], channel[j] = channel[j], channel[i]
}

func (channel ChannelInfos) Less(i, j int) bool {
	if channel[i].Epoch < channel[j].Epoch {
		return true
	}
	if channel[i].Epoch == channel[j].Epoch && channel[i].Bandwidth < channel[j].Bandwidth {
		return true
	}
	return false
}

func solve2(channels [][]uint64) int {
	channelShows := buildNormalizedShowInfo(channels)
	sort.Sort(ChannelInfos(channelShows))

	var currBW, maxBW int
	for _, timeBW := range channelShows {
		currBW += timeBW.Bandwidth
		maxBW = max(currBW, maxBW)
	}

	return maxBW
}

func buildNormalizedShowInfo(channels [][]uint64) ChannelInfos {
	var channelShows []ChannelInfo
	for i := 0; i < len(channels); i++ {
		channelShows = append(channelShows,
			ChannelInfo{channels[i][0], int(channels[i][2])})

		channelShows = append(channelShows,
			ChannelInfo{channels[i][1], -int(channels[i][2])})
	}
	return channelShows
}

func solve1(channels [][]uint64) int {
	var channelBegins, channelEnds ChannelInfos
	channelBegins, channelEnds = buildBeginEndListsForChannels(channels)
	sort.Sort(ChannelInfos(channelBegins))
	sort.Sort(ChannelInfos(channelEnds))

	var currBW, maxBW int
	var numshows = channelBegins.Len()
	var firstinWindow, curr int

	for firstinWindow < numshows && curr < numshows {
		// figures sliding windows of shows where show times overlap
		// in a sliding window of n shows, show-n starts *before* show-1 ends
		// show(n+1) starts after show(1) ends
		// if show S1's end-time = show S2's start-time, we don't consider them as
		// overlapping
		if channelEnds[firstinWindow].Epoch > channelBegins[curr].Epoch {
			currBW += channelBegins[curr].Bandwidth
			maxBW = max(currBW, maxBW)
			curr++
		} else {
			// the window slides fwd by one show, ignore the 1st show's needed bandwidth
			currBW -= channelEnds[firstinWindow].Bandwidth
			firstinWindow++
		}
	}
	return maxBW
}

func buildBeginEndListsForChannels(channels [][]uint64) (ChannelInfos, ChannelInfos) {
	var channelBegins, channelEnds ChannelInfos
	for i := 0; i < len(channels); i++ {
		channelBegins = append(channelBegins,
			ChannelInfo{channels[i][0], int(channels[i][2])})

		channelEnds = append(channelEnds,
			ChannelInfo{channels[i][1], int(channels[i][2])})
	}
	return channelBegins, channelEnds
}
