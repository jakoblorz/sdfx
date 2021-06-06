package render

import (
	"github.com/jakoblorz/sdfx/sdf"
)

/***********************
 * Hitable
 ************************/
type HitRecord struct {
	t        float64  // which t generated the hit
	p        sdf.V3   // which point when hit
	normal   sdf.V3   // normal at that point
	material Material // the material associated to this record
}

func RecordHit(t float64, p, normal sdf.V3, material Material) *HitRecord {
	return &HitRecord{
		t:        t,
		p:        p,
		normal:   normal,
		material: material,
	}
}

func (h *HitRecord) T() float64 {
	return h.t
}

func (h *HitRecord) P() sdf.V3 {
	return h.p
}

func (h *HitRecord) Normal() sdf.V3 {
	return h.normal
}

func (h *HitRecord) Material() Material {
	return h.material
}

// Hitable defines the interface of objects that can be hit by a ray
type Hitable interface {
	Hit(r *Ray3, tMin float64, tMax float64) (bool, *HitRecord)
}

// HitableList defines a simple list of hitable
type HitableList []Hitable

// hit defines the method for a list of hitables: will return the one closest
func (hl HitableList) Hit(r *Ray3, tMin float64, tMax float64) (bool, *HitRecord) {
	var res *HitRecord
	hitAnything := false

	closestSoFar := tMax

	for _, h := range hl {
		if hit, hr := h.Hit(r, tMin, closestSoFar); hit {
			hitAnything = true

			if hr.t < closestSoFar {
				res = hr
				closestSoFar = hr.t
			}
		}
	}

	return hitAnything, res
}
