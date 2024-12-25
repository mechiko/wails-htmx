package lite

import (
	"context"
	_ "embed"
	"fmt"

	"firstwails/repo/lite/liteboil"
)

func (c *dbLite) VcodeMap() (sm map[string]string, err error) {
	defer func() {
		// обязательно, иначе лочится рессурс репозитория для указателя на db в DbService
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
		c.сloseDB()
	}()

	ctx := context.Background()
	sm = map[string]string{}
	if out, err := liteboil.VidAps().All(ctx, c.db); err != nil {
		return sm, fmt.Errorf("%w", err)
	} else {
		for _, vid := range out {
			if _, ok := sm[vid.Vcode]; !ok {
				sm[vid.Vcode] = vid.Name
			}
		}
		return sm, nil
	}
}
