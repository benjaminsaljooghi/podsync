package builders

import (
	"os"
	"testing"

	"github.com/mxpv/podsync/pkg/api"
	"github.com/mxpv/podsync/pkg/model"
	"github.com/stretchr/testify/require"
)

var ytKey = os.Getenv("YOUTUBE_TEST_API_KEY")

func TestQueryYTChannel(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping YT test in short mode")
	}

	builder, err := NewYouTubeBuilder(ytKey)
	require.NoError(t, err)

	channel, err := builder.listChannels(api.LinkTypeChannel, "UC2yTVSttx7lxAOAzx1opjoA")
	require.NoError(t, err)
	require.Equal(t, "UC2yTVSttx7lxAOAzx1opjoA", channel.Id)

	channel, err = builder.listChannels(api.LinkTypeUser, "fxigr1")
	require.NoError(t, err)
	require.Equal(t, "UCr_fwF-n-2_olTYd-m3n32g", channel.Id)
}

func TestBuildYTFeed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping YT test in short mode")
	}

	builder, err := NewYouTubeBuilder(ytKey)
	require.NoError(t, err)

	podcast, err := builder.Build(&model.Feed{
		Provider: api.ProviderYoutube,
		LinkType: api.LinkTypeChannel,
		ItemID:   "UCupvZG-5ko_eiXAupbDfxWw",
		PageSize: maxYoutubeResults,
	})

	require.NoError(t, err)

	require.Equal(t, "CNN", podcast.Title)
	require.NotEmpty(t, podcast.Description)

	require.Equal(t, 50, len(podcast.Items))

	for _, item := range podcast.Items {
		require.NotEmpty(t, item.Title)
		require.NotEmpty(t, item.Link)
		require.NotEmpty(t, item.IDuration)
	}
}