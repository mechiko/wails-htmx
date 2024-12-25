package lite

import (
	"fmt"

	"firstwails/entity"
)

func (r *dbReport) GetAllVCodesMap() (m entity.VidapMap, err error) {
	defer func() {
		// обязательно, иначе лочится рессурс репозитория для указателя на db в DbService
		r.сloseDB()
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()
	m = make(entity.VidapMap)
	rl := &entity.VidapList{}
	query := `select * from vid_ap order by vcode;`

	err = r.db.Select(&rl.Items, query)
	if err != nil {
		return m, fmt.Errorf("report:GetAllVCodes %w", err)
	}
	rl.Total = int64(len(rl.Items))

	for _, v := range rl.Items {
		m[v.Vcode] = v.Name
	}
	return m, nil
}
