package lite

import (
	"context"
	_ "embed"
	"fmt"

	"firstwails/repo/lite/liteboil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (c *dbLite) Version() (s string, err error) {
	defer func() {
		// обязательно, иначе лочится рессурс репозитория для указателя на db в DbService
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
		c.сloseDB()
	}()

	ctx := context.Background()
	if out, err := liteboil.Dboptions(qm.Where("name=?", "version")).One(ctx, c.db); err != nil {
		return "", fmt.Errorf("%w", err)
	} else {
		return out.Value, nil
	}
}
