package myhub

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/zikwall/myhub/internal/application/domain"
)

func (s *DatabaseTestSuite) TestRepository_AddChannel() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	repo := New(s.container.Pool)

	tests := []struct {
		name    string
		channel domain.Channel
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "successfully add channel",
			channel: domain.Channel{
				ID:          uuid.MustParse("c59cbd01-8d78-4b91-92b5-c08852973066"),
				Name:        "test_1",
				Description: "test_desc",
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.wantErr(
				t,
				repo.AddChannel(ctx, tt.channel),
				fmt.Sprintf("AddChannel(%v)", tt.channel),
			)
		})
	}
}
