package utils

import "github.com/ozonva/ova-obligation-api/internal/entity"

func ChunkObligations(obligations []entity.Obligation, chunkSize int) [][]entity.Obligation {
	chunks := make([][]entity.Obligation, 0)
	sliceLen := len(obligations)
	for i := 0; i < sliceLen; i += chunkSize {
		end := i + chunkSize

		if end > sliceLen {
			end = sliceLen
		}

		chunks = append(chunks, obligations[i:end])
	}

	return chunks
}

func ChunkObligationPointers(obligations []*entity.Obligation, chunkSize int) [][]*entity.Obligation {
	chunks := make([][]*entity.Obligation, 0)
	sliceLen := len(obligations)
	for i := 0; i < sliceLen; i += chunkSize {
		end := i + chunkSize

		if end > sliceLen {
			end = sliceLen
		}

		chunks = append(chunks, obligations[i:end])
	}

	return chunks
}

func ReplaceIdxToID(obligations []entity.Obligation) map[uint]entity.Obligation {
	keyToEntity := make(map[uint]entity.Obligation)
	for _, v := range obligations {
		keyToEntity[v.ID] = v
	}

	return keyToEntity
}
