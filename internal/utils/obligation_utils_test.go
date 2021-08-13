package utils

import (
	"testing"

	"github.com/ozonva/ova-obligation-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestReplaceIdxToID(t *testing.T) {
	type test struct {
		name string
		arg  []entity.Obligation
		want map[uint]entity.Obligation
	}

	tests := []test{
		{
			name: "With values",
			arg: []entity.Obligation{
				{ID: 1, Title: "test1", Description: "test1"},
				{ID: 2, Title: "test2", Description: "test2"},
			},
			want: map[uint]entity.Obligation{
				1: {ID: 1, Title: "test1", Description: "test1"},
				2: {ID: 2, Title: "test2", Description: "test2"},
			},
		},
		{
			name: "With empty obligation list",
			arg:  make([]entity.Obligation, 0),
			want: make(map[uint]entity.Obligation),
		},
		{
			name: "If obligation list is nill",
			arg:  nil,
			want: make(map[uint]entity.Obligation),
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, ReplaceIdxToID(tt.arg), tt.name)
	}
}

func TestChunkObligations(t *testing.T) {
	type args struct {
		obligations []entity.Obligation
		size        int
	}

	type test struct {
		name string
		args args
		want [][]entity.Obligation
	}

	tests := []test{
		{
			name: "With values and without a remainder",
			args: args{
				obligations: []entity.Obligation{
					{ID: 1, Title: "test1", Description: "test1"},
					{ID: 2, Title: "test2", Description: "test2"},
					{ID: 3, Title: "test3", Description: "test3"},
					{ID: 4, Title: "test4", Description: "test4"},
					{ID: 5, Title: "test5", Description: "test5"},
					{ID: 6, Title: "test6", Description: "test6"},
				},
				size: 2,
			},
			want: [][]entity.Obligation{
				{
					{ID: 1, Title: "test1", Description: "test1"},
					{ID: 2, Title: "test2", Description: "test2"},
				},
				{
					{ID: 3, Title: "test3", Description: "test3"},
					{ID: 4, Title: "test4", Description: "test4"},
				},
				{
					{ID: 5, Title: "test5", Description: "test5"},
					{ID: 6, Title: "test6", Description: "test6"},
				},
			},
		},
		{
			name: "With values and with a remainder",
			args: args{
				obligations: []entity.Obligation{
					{ID: 1, Title: "test1", Description: "test1"},
					{ID: 2, Title: "test2", Description: "test2"},
					{ID: 3, Title: "test3", Description: "test3"},
					{ID: 4, Title: "test4", Description: "test4"},
					{ID: 5, Title: "test5", Description: "test5"},
					{ID: 6, Title: "test6", Description: "test6"},
					{ID: 7, Title: "test7", Description: "test7"},
				},
				size: 2,
			},
			want: [][]entity.Obligation{
				{
					{ID: 1, Title: "test1", Description: "test1"},
					{ID: 2, Title: "test2", Description: "test2"},
				},
				{
					{ID: 3, Title: "test3", Description: "test3"},
					{ID: 4, Title: "test4", Description: "test4"},
				},
				{
					{ID: 5, Title: "test5", Description: "test5"},
					{ID: 6, Title: "test6", Description: "test6"},
				},
				{
					{ID: 7, Title: "test7", Description: "test7"},
				},
			},
		},
		{
			name: "With values and if size more than obligation list length",
			args: args{
				obligations: []entity.Obligation{
					{ID: 1, Title: "test1", Description: "test1"},
					{ID: 2, Title: "test2", Description: "test2"},
				},
				size: 4,
			},
			want: [][]entity.Obligation{
				{
					{ID: 1, Title: "test1", Description: "test1"},
					{ID: 2, Title: "test2", Description: "test2"},
				},
			},
		},
		{
			name: "Without values",
			args: args{
				obligations: make([]entity.Obligation, 0),
				size:        4,
			},
			want: make([][]entity.Obligation, 0),
		},
		{
			name: "With nil obligations list",
			args: args{
				obligations: nil,
				size:        2,
			},
			want: make([][]entity.Obligation, 0),
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, ChunkObligations(tt.args.obligations, tt.args.size), tt.name)
	}
}
