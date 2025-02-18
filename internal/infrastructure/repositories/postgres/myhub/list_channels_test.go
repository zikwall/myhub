//go:build integration

package myhub

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/zikwall/myhub/internal/application/domain"
)

func (s *DatabaseTestSuite) TestRepository_ListChannels() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	repo := New(s.container.Pool)

	tests := []struct {
		name      string
		want      []domain.Channel
		prepareFn func([]domain.Channel)
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "successfully list channels",
			want: []domain.Channel{
				{
					ID:          uuid.MustParse("c59cbd01-8d78-4b91-92b5-c08852973066"),
					Name:        "test_1",
					Description: "test_desc",
				},
				{
					ID:          uuid.MustParse("c59cbd01-8d78-4b91-92b5-c08852973065"),
					Name:        "test_2",
					Description: "test_desc_2",
				},
			},
			prepareFn: func(channels []domain.Channel) {
				for _, channel := range channels {
					require.NoError(s.T(), repo.AddChannel(ctx, channel))
				}
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.prepareFn(tt.want)

			got, err := repo.ListChannels(ctx)
			tt.wantErr(t, err, fmt.Sprintf("ListChannels()"))

			require.Equal(t, tt.want, got)
		})
	}
}
